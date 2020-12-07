// https://adventofcode.com/2020/day/6
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countDistinct(str string) int {
	var counts [26]int

	for _, chr := range str {
		counts[chr-'a'] = 1
	}

	count := 0
	for i := 0; i < len(counts); i++ {
		count += counts[i]
	}

	return count
}

func countExactN(str string, expected int) int {
	var counts [26]int

	for _, chr := range str {
		counts[chr-'a'] ++
	}

	count := 0
	for i := 0; i < len(counts); i++ {
		if counts[i] == expected {
			count++
		}
	}

	return count
}

func main() {
	fc, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		return
	}

	lines := strings.Split(string(fc), "\n")

	sum := 0
	ans := ""
	for _, line := range lines {
		ans += strings.TrimSpace(line)

		if len(line) == 0 {
			sum += countDistinct(ans)

			ans = ""
			continue
		}
	}

	sum += countDistinct(ans)

	fmt.Println(sum)

	// Part 2
	sum = 0
	ans = ""
	personCount := 0
	for _, line := range lines {
		ans += strings.TrimSpace(line)

		if len(line) > 0 {
			personCount++
		}

		if len(line) == 0 {
			sum += countExactN(ans, personCount)

			ans = ""
			personCount = 0
			continue
		}
	}

	sum += countExactN(ans, personCount)

	fmt.Println(sum)
}
