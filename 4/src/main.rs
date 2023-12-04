use std::collections::HashSet;
use std::error::Error;
use std::fs;

fn calculate_winnings(winning: &HashSet<i32>, found: &HashSet<i32>) -> i32 {
    const WIN_MULTIPLE: i32 = 2;

    let mut winners = 0;

    for f in found {
        if winning.contains(f) {
            winners += 1;
        }
    }

    if winners > 0 {
        return WIN_MULTIPLE.pow(winners - 1);
    }

    winners as i32
}

fn read_lines(input: &str) -> Result<i32, Box<dyn Error>> {
    let mut points = 0;

    for line in input.lines() {
        let line = line.split(": ").skip(1).collect::<Vec<_>>();
        let line = line[0]
            .split(" | ")
            .map(|numbers| {
                numbers
                    .split(' ')
                    .filter_map(|n| n.parse::<i32>().ok())
                    .collect::<HashSet<_>>()
            })
            .collect::<Vec<_>>();

        let winning_numbers = &line[0];
        let found_numbers = &line[1];
        points += calculate_winnings(winning_numbers, found_numbers);
    }

    Ok(points)
}

fn main() -> Result<(), Box<dyn Error>> {
    let input = fs::read_to_string("input")?;
    let points = read_lines(&input)?;
    println!("{points}");

    Ok(())
}
