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
fn search(word: &str, limit: i32, strict: bool) -> Vec<HashMap<String, String>> {
    let mut results: Vec<HashMap<String, String>> = Vec::new();

    // If strict mode is enabled, return the exact match
    if strict {
        match CACHE.get(word) {
            Some(value) => {
                for index in value {
                    results.push(DATA[*index as usize].clone());
                }
                return results;
            }
            None => return results,
        }
    }

    // Store the indexes that have already been added
    let mut already_added: Vec<i16> = Vec::new();

    // Iterate over the cache
    for (key, value) in CACHE.iter() {
        if results.len() as i32 >= limit {
            return results;
        }

        // If the word doesn't start with the same letter as the key
        match key.chars().nth(0) {
            Some(key_first_char) => match word.chars().nth(0) {
                Some(word_first_char) => {
                    if key_first_char != word_first_char {
                        continue;
                    }
                }
                None => continue,
            },
            None => continue,
        }

        // If the key is shorter than the word
        if key.len() < word.len() {
            continue;
        }

        // If the key equals the word
        if key == word {
            for index in value {
                results.push(DATA[*index as usize].clone());
            }
            return results;
        }

        // If the key doesn't contain the word
        if !key.contains(word) {
            continue;
        }

        // Iterate over the indexes
        for index in value {
            if already_added.contains(index) {
                continue;
            }
            results.push(DATA[*index as usize].clone());
            already_added.push(*index);
        }
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

    // Get the limit
    let limit = match params.get("limit") {
        Some(limit) => match limit.parse::<i32>() {
            Ok(limit) => limit,
            Err(_) => return HttpResponse::BadRequest().json(serde_json::json!({})),
        },
        None => 10,
    };

    // Get the strict mode
    let strict = match params.get("strict") {
        Some(strict) => match strict.parse::<bool>() {
            Ok(strict) => strict,
            Err(_) => return HttpResponse::BadRequest().json(serde_json::json!({})),
        },
        None => false,
    };

    // Track the start time in nanoseconds
    let start: std::time::Instant = std::time::Instant::now();

    // Query the cache
    let results: Vec<HashMap<String, String>> = search(query, limit, strict);

    // Print the elapsed time
    println!("Found {} results in {:?}", results.len(), start.elapsed());

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