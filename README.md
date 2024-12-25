# 🎄 Advent of Code | int y=2024;

This repository contains my solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges, implemented in **Go (Golang)**.

## 🚀 Usage

```bash
    # Add your input.txt file in the folder of the day you want to run

    # run specific day (x = number of day)
    $ make run-day day=x
    
    # run all days
    $ make run-all
    
    # Also you run directly using go:
    $ go run <day>/main.go
```

## ⭐ Solutions

| **Day**                                                                | **Solution**                                                              | **Comments**                                                                                                                                                                                                                                                                                       |
|------------------------------------------------------------------------|---------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Day 1: Historian Hysteria](https://adventofcode.com/2024/day/1)       | [Link](https://github.com/dcorto/adventofcode2024/blob/master/1/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 2: Red-Nosed Reports](https://adventofcode.com/2024/day/2)        | [Link](https://github.com/dcorto/adventofcode2024/blob/master/2/main.go)  |                                                                                                                                                                                                                                                                                                    |                                                                      
| [Day 3: Mull It Over](https://adventofcode.com/2024/day/3)             | [Link](https://github.com/dcorto/adventofcode2024/blob/master/3/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 4: Ceres Search](https://adventofcode.com/2024/day/4)             | [Link](https://github.com/dcorto/adventofcode2024/blob/master/4/main.go)  |                                                                                                                                                                                                                                                                                                    |     
| [Day 5: Print Queue](https://adventofcode.com/2024/day/5)              | [Link](https://github.com/dcorto/adventofcode2024/blob/master/5/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 6: Guard Gallivant](https://adventofcode.com/2024/day/6)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/6/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 7: Bridge Repair](https://adventofcode.com/2024/day/7)            | [Link](https://github.com/dcorto/adventofcode2024/blob/master/7/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 8: Resonant Collinearity](https://adventofcode.com/2024/day/8)    | [Link](https://github.com/dcorto/adventofcode2024/blob/master/8/main.go)  |                                                                                                                                                                                                                                                                                                    |
| [Day 9: Disk Fragmenter](https://adventofcode.com/2024/day/9)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/9/main.go)  | I used this [solution](https://github.com/shraddhaag/aoc/blob/main/2024/day9/main.go) for the part B, i got stuck a lot of time, that solution save the day                                                                                                                                        |
| [Day 10: Hoof It](https://adventofcode.com/2024/day/10)                | [Link](https://github.com/dcorto/adventofcode2024/blob/master/10/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 11: Plutonian Pebbles](https://adventofcode.com/2024/day/11)      | [Link](https://github.com/dcorto/adventofcode2024/blob/master/11/main.go) | The first version for part A, I made using slices and for 25 iterations it was OK. For part B (75 iterations), it no longer worked, so in the end, I had to redo it using the recursive function 'blink' and storing previously calculated results with a map (cache) to avoid recalculating them. |
| [Day 12: Garden Groups](https://adventofcode.com/2024/day/12)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/12/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 13: Claw Contraption](https://adventofcode.com/2024/day/13)       | [Link](https://github.com/dcorto/adventofcode2024/blob/master/13/main.go) | [Here is](https://www.reddit.com/r/adventofcode/comments/1hd7irq/2024_day_13_an_explanation_of_the_mathematics/) the explanation of the solution. Spoiler: Cramer's Rule                                                                                                                           |
| [Day 14: Restroom Redoubt](https://adventofcode.com/2024/day/14)       | [Link](https://github.com/dcorto/adventofcode2024/blob/master/14/main.go) | The tip for the part 2 is move robots until their do not overlaps and count the iterations as seconds                                                                                                                                                                                              |
| [Day 15: Warehouse Woes](https://adventofcode.com/2024/day/15)         | [Link](https://github.com/dcorto/adventofcode2024/blob/master/15/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 16: Reindeer Maze](https://adventofcode.com/2024/day/16)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/16/main.go) | Part A solves using [Dijkstra](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)                                                                                                                                                                                                               |
| [Day 17: Chronospatial Computer](https://adventofcode.com/2024/day/17) | [Link](https://github.com/dcorto/adventofcode2024/blob/master/17/main.go) | Today is my b-day (39y) :tada: Today puzzle was fun, but a small bug in one operation took me a while...                                                                                                                                                                                           |
| [Day 18: RAM Run](https://adventofcode.com/2024/day/18)                | [Link](https://github.com/dcorto/adventofcode2024/blob/master/18/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 19: Linen Layout](https://adventofcode.com/2024/day/19)           | [Link](https://github.com/dcorto/adventofcode2024/blob/master/19/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 20: Race Condition](https://adventofcode.com/2024/day/20)         | [Link](https://github.com/dcorto/adventofcode2024/blob/master/20/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 21: Keypad Conundrum](https://adventofcode.com/2024/day/21)       | [Link](https://github.com/dcorto/adventofcode2024/blob/master/21/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 22: Monkey Market](https://adventofcode.com/2024/day/22)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/22/main.go) |                                                                                                                                                                                                                                                                                                    |
| [Day 23: LAN Party](https://adventofcode.com/2024/day/23)              | [Link](https://github.com/dcorto/adventofcode2024/blob/master/23/main.go) | For the part B i used the [Bron-Kerbosch algorithm](https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm) to find all maximal cliques in a graph                                                                                                                                          |
| [Day 24: Crossed Wires](https://adventofcode.com/2024/day/24)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/24/main.go) | [Guide on the idea behind the solution B](https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/)                                                                                                                                               |
| [Day 25: Code Chronicle](https://adventofcode.com/2024/day/25)          | [Link](https://github.com/dcorto/adventofcode2024/blob/master/25/main.go) |                                                                                                                                              |

## 📝 Notes

- Each solution is designed to work with the input provided by the Advent of Code website.
- Make sure to place the input file `input.txt` (if required) in the day folder.
- I have not included my own input sets in the repository as [recommended](https://www.reddit.com/r/adventofcode/comments/e7khy8/comment/fa13hb9/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button) by one of the mods in the official subreddit.

## 🔄 My Previous Years

[Advent of Code 2023](https://github.com/dcorto/adventofcode2023)

[Advent of Code 2022](https://github.com/dcorto/adventofcode2022)

[Advent of Code 2021](https://github.com/dcorto/adventofcode2021)

---

Happy coding and good luck with Advent of Code 2024! 🎉


