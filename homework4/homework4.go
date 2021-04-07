package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type sliceInt64 []int64

func (m sliceInt64) Len() int {
	return len(m)
}

func (m sliceInt64) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m sliceInt64) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func getListOfNumbers() ([]int64, error) {
	arr := []int64{}
	scaner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Type number or S to stop.")
		if scaner.Scan() {

			if scaner.Text() == "S" || scaner.Text() == "s" {
				break
			}

			n, err := strconv.ParseInt(scaner.Text(), 10, 64)

			if err != nil {
				fmt.Println("Error: You can add only (1) number to the list.")
			} else {
				arr = append(arr, n)
			}

		}
	}
	return arr, nil
}

func main() {

	arr, err := getListOfNumbers()
	fmt.Println("")

	if err == nil {
		fmt.Println("unsorted list:", arr)

		var sortedArr sliceInt64 = arr
		sort.Sort(sortedArr)
		fmt.Println(sortedArr)
	}

}
