package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dataStr := string(data)
	dataLines := strings.Split(dataStr, "\n")

	elves := make([]int64, len(dataLines))
	var elf int = 0

	for _, line := range dataLines {
		if line == "" {
			elf++
			continue
		}

		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		elves[elf] += val
	}

	maxCal := slices.Max(elves)
	fmt.Println(maxCal)
}
