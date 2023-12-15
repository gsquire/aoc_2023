use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let input = fs::read_to_string("input")?;
    let input = input.trim();

    let mut sum = 0;

    for value in input.split(',') {
        let cur_hash = value.chars().map(|c| c as u8).fold(0, |acc, code| {
            let mut value = acc;
            value += code as u32;
            value *= 17;
            value %= 256;
            value
        });

        sum += cur_hash;
    }

    println!("{sum}");

    Ok(())
}
