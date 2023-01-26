use std::time::{SystemTime, UNIX_EPOCH};

/**
* # Compile with 
* cargo build
* # Run 
* cargo run 
*/


fn main() {
    // Get current Unix timestamp
    let current_timestamp = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs();
    println!("Current Unix timestamp: {}", current_timestamp);

    // Get current UTC time
    let current_utc_time = time::now_utc();
    println!("UTC time: {}", current_utc_time.rfc822());
}