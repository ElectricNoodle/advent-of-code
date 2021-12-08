package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFileOfInts(fname string) (nums []int, err error) {

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func challengeOne() {

	nums, err := readFileOfInts("inputs/1")

	if err != nil {
		panic(err)
	}

	numOfIncreases := 0

	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] > nums[i-1] {
				numOfIncreases++
			}
		}
	}

	fmt.Printf("(1) Number of increases: %d \n", numOfIncreases)

	numOfIncreases = 0

	rollingSum := make([]int, 20)

	for i := 0; i < len(nums); i++ {
		if i > 1 {

			sum := nums[i] + nums[i-1] + nums[i-2]
			rollingSum = append(rollingSum, sum)

		}
	}

	for i := 0; i < len(rollingSum); i++ {
		if i > 0 {
			if rollingSum[i] > rollingSum[i-1] {
				numOfIncreases++
			}
		}
	}

	fmt.Printf("(1.5) Number of increases: %d \n", numOfIncreases)

}

func challengeTwo() {

}

func main() {
	challengeOne()
	challengeTwo()
}
