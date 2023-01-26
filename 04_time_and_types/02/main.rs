use std::io;

// compile and run
// 04_time_and_types\02> rustc main.rs
// 04_time_and_types\02> .\main.exe

fn main() {
    let mut input = String::new();
    println!("Enter a number: ");
    io::stdin().read_line(&mut input).unwrap();

    let input: i32 = input.trim().parse().unwrap();
    println!("Decimal: {}", input);
    println!("Binary: {:b}", input);
    println!("Hexadecimal: {:#X}", input);
}

