package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Function to check if all values are distinct
func allDistinct(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}

// Function to generate permutations of an array of digits
func permute(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// Make a copy of nums to avoid referencing the same slice
		permCopy := make([]int, len(nums))
		copy(permCopy, nums)
		*result = append(*result, permCopy)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permute(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func cipherNormal() {

	start := time.Now()

	for w := 1; w < 10; w++ {
		for o := 0; o < 10; o++ {
			for t := 0; t < 10; t++ {
				for l := 0; l < 10; l++ {
					for d := 1; d < 10; d++ {
						for g := 1; g < 10; g++ {
							for e := 0; e < 10; e++ {
								for c := 0; c < 10; c++ {
									for m := 0; m < 10; m++ {
										numbers := []int{w, o, t, l, d, g, e, c, m}

										if allDistinct(numbers) {
											minValue := strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t)
											subValue := strconv.Itoa(g) + strconv.Itoa(o) + strconv.Itoa(o) + strconv.Itoa(g) + strconv.Itoa(l) + strconv.Itoa(e)
											difValue := strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t) + strconv.Itoa(c) + strconv.Itoa(o) + strconv.Itoa(m)
											minValueInt, err1 := strconv.Atoi(minValue)
											subValueInt, err2 := strconv.Atoi(subValue)
											difValueInt, err3 := strconv.Atoi(difValue)
											if err1 != nil || err2 != nil || err3 != nil {
												fmt.Println("Error converting int to string")
											}

											res := minValueInt - subValueInt

											if res == difValueInt {
												fmt.Println(minValue)
												fmt.Println(subValue)
												fmt.Println(difValue)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s\n", elapsed)
}

func cipherPerm() {

	start := time.Now()

	// Create an array of digits from 0-9
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var permutations [][]int

	// Generate all permutations of the digits 0-9
	permute(digits, 0, &permutations)

	// Loop through all permutations and check the condition
	for _, perm := range permutations {
		// Extract specific digits for w, o, t, l, d, g, e, c, m
		w := perm[0]
		o := perm[1]
		t := perm[2]
		l := perm[3]
		d := perm[4]
		g := perm[5]
		e := perm[6]
		c := perm[7]
		m := perm[8]

		// Check if all values are distinct
		values := []int{w, o, t, l, d, g, e, c, m}
		if allDistinct(values) {
			// Construct the min_value, sub_value, and dif_value as strings
			minValue := strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t)
			subValue := strconv.Itoa(g) + strconv.Itoa(o) + strconv.Itoa(o) + strconv.Itoa(g) + strconv.Itoa(l) + strconv.Itoa(e)
			difValue := strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t) + strconv.Itoa(c) + strconv.Itoa(o) + strconv.Itoa(m)

			// Convert to integers
			minValueInt, _ := strconv.Atoi(minValue)
			subValueInt, _ := strconv.Atoi(subValue)
			difValueInt, _ := strconv.Atoi(difValue)

			// Check if the condition holds
			if minValueInt-subValueInt == difValueInt {
				fmt.Println("minValue:", minValue)
				fmt.Println("subValue:", subValue)
				fmt.Println("Result:", minValueInt-subValueInt)
				fmt.Println()
			}
		}
	}

	elapsed := time.Since(start)
	ms := elapsed.Milliseconds()
	fmt.Printf("Execution time: %d msec\n", ms)
}

// Worker function to process a chunk of permutations
func processChunk(permutations [][]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, perm := range permutations {
		// Extract specific digits for w, o, t, l, d, g, e, c, m
		w := perm[0]
		o := perm[1]
		t := perm[2]
		l := perm[3]
		d := perm[4]
		g := perm[5]
		e := perm[6]
		c := perm[7]
		m := perm[8]

		// Check if all values are distinct
		values := []int{w, o, t, l, d, g, e, c, m}
		if allDistinct(values) {
			// Construct the min_value, sub_value, and dif_value as strings
			minValue := strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(w) + strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t)
			subValue := strconv.Itoa(g) + strconv.Itoa(o) + strconv.Itoa(o) + strconv.Itoa(g) + strconv.Itoa(l) + strconv.Itoa(e)
			difValue := strconv.Itoa(d) + strconv.Itoa(o) + strconv.Itoa(t) + strconv.Itoa(c) + strconv.Itoa(o) + strconv.Itoa(m)

			// Convert to integers
			minValueInt, _ := strconv.Atoi(minValue)
			subValueInt, _ := strconv.Atoi(subValue)
			difValueInt, _ := strconv.Atoi(difValue)

			// Check if the condition holds
			if minValueInt-subValueInt == difValueInt {
				fmt.Println("minValue:", minValue)
				fmt.Println("subValue:", subValue)
				fmt.Println("Result:", minValueInt-subValueInt)
				fmt.Println()
			}
		}
	}
}

func cipherPermConcurrent() {
	start := time.Now()

	// Create an array of digits from 0-9
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var permutations [][]int
	permute(digits, 0, &permutations)

	// Split the permutations into chunks for concurrent processing
	numWorkers := 8 // Number of goroutines
	chunkSize := len(permutations) / numWorkers
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize
		if i == numWorkers-1 {
			endIdx = len(permutations) // Include any remaining permutations in the last chunk
		}

		wg.Add(1)
		go processChunk(permutations[startIdx:endIdx], &wg)
	}

	wg.Wait() // Wait for all goroutines to finish

	elapsed := time.Since(start)
	ms := elapsed.Milliseconds()
	fmt.Printf("Execution time with goroutines: %d msec\n", ms)
}

func main() {

	cipherNormal()
	cipherPerm()
	cipherPermConcurrent()

}
