use serde_json;
use lazy_static;
use std::collections::HashMap;
use actix_web::{HttpRequest, HttpResponse, HttpServer, App, web};

// Store the json data in a variable
lazy_static::lazy_static!(
    static ref DATA: Vec<HashMap<String, String>> = {
        // Read the data.json file
        let data: String = std::fs::read_to_string("../data.json").expect("Failed to read data.json");

        // Parse the data.json file
        return serde_json::from_str(&data).expect("Failed to parse data.json");
    };
);

// Remove double spaces from a string
fn remove_double_spaces(value: &str) -> String {
    let mut value = value.to_string();
    while value.contains("  ") {
        value = value.replace("  ", " ");
    }
    return value;
}

// Create a global hasjmap cache for the data using lazy_static
// The CACHE stores a string key and a vector of i16 values
lazy_static::lazy_static! {
    static ref CACHE: std::collections::HashMap<String, Vec<i16>> = {
        let mut cache = std::collections::HashMap::new();

        // Iterate over the parsed data
        for (i, item) in DATA.iter().enumerate() {
            for (_, value) in item {
                let value: String = remove_double_spaces(value);

                // Split the value by spaces
                let words: Vec<&str> = value.as_str().split(" ").collect();

                // Iterate over the words
                for word in words {
                    // Check if the word is alphanumeric
                    if !word.chars().all(char::is_alphanumeric) {
                        continue;
                    }

                    // Check if the word is already in the cache
                    if !cache.contains_key(word) {
                        cache.insert(word.to_string(), Vec::new());
                    }

                    // If the index already exists in the vector, skip it
                    if cache.get(word).unwrap().contains(&(i as i16)) {
                        continue;
                    }

                    // Push the index to the vector
                    cache.get_mut(word).unwrap().push(i as i16);
                }
            }
        }
        return cache
    };
}

// Search the cache for a word
fn search(word: &str) -> Vec<i16> {
    let mut results: Vec<i16> = Vec::new();
    // Iterate over the cache
    for (key, value) in CACHE.iter() {
        if key.contains(word) {
            results.append(&mut value.clone());
        }
    }
    return results
}

// Convert the indices to the actual data
fn indices_to_data(indices: Vec<i16>) -> Vec<HashMap<String, String>> {
    let mut results: Vec<HashMap<String, String>> = Vec::new();
    for index in indices {
        results.push(DATA[index as usize].clone());
    }
    return results
}

// The courses api endpoint
#[actix_web::get("/courses")]
async fn courses(req: HttpRequest) -> HttpResponse {
    // Get the query parameters
    let params = match web::Query::<HashMap<String, String>>::from_query(req.query_string()) {
        Ok(params) => params,
        Err(_) => return HttpResponse::BadRequest().json(serde_json::json!({})),
    };

    // Get the query
    let query = match params.get("q") {
        Some(query) => query,
        None => return HttpResponse::BadRequest().json(serde_json::json!({})),
    };

    // Track the start time in nanoseconds
    let start: std::time::Instant = std::time::Instant::now();

    // Query the cache
    let results: Vec<HashMap<String, String>> = indices_to_data(search(query));

    // Print the elapsed time
    println!("Elapsed: {:?}", start.elapsed());

    // Return the results as json
    return HttpResponse::Ok().json(results);
}

// Main Actix-Web function
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Establish a connection to http://127.0.0.1:8000/
    HttpServer::new(move || {
        App::new()
            .wrap(actix_cors::Cors::permissive())
            .service(courses)
            .wrap(actix_web::middleware::NormalizePath::trim())
    })
    .bind(("127.0.0.1", 8000))?
    .run()
    .await
}