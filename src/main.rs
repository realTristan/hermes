use serde_json;
use lazy_static;
use std::collections::HashMap;

// Store the json data in a variable
lazy_static::lazy_static!(
    static ref DATA: Vec<HashMap<String, String>> = {
        // Read the data.json file
        let data: String = std::fs::read_to_string("data.json").expect("Failed to read data.json");

        // Parse the data.json file
        return serde_json::from_str(&data).expect("Failed to parse data.json");
    };
);

// Create a global hasjmap cache for the data using lazy_static
// The CACHE stores a string key and a vector of i16 values
lazy_static::lazy_static! {
    static ref CACHE: std::collections::HashMap<String, Vec<i16>> = {
        let mut cache = std::collections::HashMap::new();

        // Iterate over the parsed data
        for (i, item) in DATA.iter().enumerate() {
            for (_, value) in item {
                // Split the value by spaces
                let words: Vec<&str> = value.as_str().split(" ").collect();

                // Iterate over the words
                for word in words {
                    // Check if the word is already in the cache
                    if !cache.contains_key(word) {
                        cache.insert(word.to_string(), Vec::new());
                    }
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

fn main() {
    // Track the start time
    let start = std::time::Instant::now();
    
    // Search the cache for the word "computer"
    let results: Vec<i16> = search("computer");
    indices_to_data(results);

    // Print the elapsed time
    println!("Elapsed: {:?}", start.elapsed());

    // run cargo run --release
}
