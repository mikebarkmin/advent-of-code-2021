package day1

import (
	"fmt"

	"github.com/mikebarkmin/aoc-2021/utils"
)

func part1(measurements []int) int {
	last_measurement := measurements[0]

	count_increased := 0

	for i := 1; i < len(measurements); i++ {
		curr_measurement := measurements[i]

		if last_measurement < curr_measurement {
			count_increased++
		}
		last_measurement = curr_measurement
	}

	return count_increased
}

func part2(measurements []int) int {
	last_sliding_window := measurements[0] + measurements[1] + measurements[2]

	count_increased := 0

	for i := 1; i < len(measurements); i++ {
		if i+2 >= len(measurements) {
			break
		}
		curr_sliding_window := measurements[i] + measurements[i+1] + measurements[i+2]

		if last_sliding_window < curr_sliding_window {
			count_increased++
		}

		last_sliding_window = curr_sliding_window
	}

	return count_increased
}

func Run() {
  fmt.Println("Day 01\n======")

	measurements := utils.ReadLinesAsIntegers("input_day1")
	c := part1(measurements)
	fmt.Printf("Part 1: %d\n", c)

	c = part2(measurements)
	fmt.Printf("Part 2: %d\n", c)

  fmt.Println()
}
