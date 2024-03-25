package main

/*
 * Author:  Lamine Djibo, ldjibo2016@my.fit.edu
 * Course:  CSE 4251, Fall 2023
 * Project: Project #2, Manatee Transportation
 */

import (
	"fmt"
	"os"
	"strconv"
)

func convert(list []string) []int64 {

	sizes := make([]int64, len(list))

	for i, str := range list {
		num, err := strconv.ParseInt(str, 0, 64)
		if err != nil {
			//fmt.Printf("Error parsing element at index %d: %v\n", i, err)
			return sizes
		}
		sizes[i] = num
	}

	return sizes
}

func filterList(list []int64, N int64) []int64 {
	var result []int64

	for _, value := range list {
		if value > N {
			break
		}
		result = append(result, value)
	}

	if len(result) == 0 {
		//fmt.Println("Impossible to fit")
		os.Exit(0)
	}

	return result
}

func first_k_elements(list []int64, two_N int64) []int64 {

	if len(list) == 0 {
		//fmt.Println("List is empty")
		os.Exit(0)
	} else if len(list) == 1 {
		if list[0] <= int64(two_N/2) {
			fmt.Println("1 \nport")
			os.Exit(0)
		} else {
			//fmt.Println("Impossible")
			os.Exit(0)
		}
	}
	var first []int64
	var count int64
	count = 0

	for i := 0; i < len(list); i++ {
		if (count + list[i]) <= two_N {
			count += list[i]
			first = append(first, list[i])
		} else {
			break
		}
	}

	return first

}

func combinations_to_n(nums []int64, N int64) [][]int64 {
	var result [][]int64
	var dp func(index int64, currentCombination []int64, currentSum int64)
	dp = func(index int64, currentCombination []int64, currentSum int64) {
		if currentSum <= N {
			if currentSum <= N {
				result = append(result, append([]int64{}, currentCombination...))
			}

			for i := index; i < int64(len(nums)); i++ {
				if i > index && nums[i] == nums[i-1] {
					continue
				}

				currentCombination = append(currentCombination, nums[i])
				currentSum += nums[i]
				dp(i+1, currentCombination, currentSum)
				currentCombination = currentCombination[:len(currentCombination)-1]
				currentSum -= nums[i]
			}
		}
	}

	dp(0, []int64{}, 0)
	return result
}

func equal_to_n(combinations [][]int64, N int64) [][]int64 {
	var result [][]int64
	count := 0

	for j := 0; j < int(len(combinations)); j++ {
		for i := 0; i < int(len(combinations[j])); i++ {
			if (int(count) + int(combinations[j][i])) != int(N) {
				count += int(combinations[j][i])
			} else {
				if i == len(combinations[j])-1 {
					result = append(result, combinations[j])
				}
				break
			}
		}
		count = 0
	}

	return result
}

func findMaxLists(combinations [][]int64) [][]int64 {
	var maxLists [][]int64
	var maxLength int64
	maxLength = 0

	for _, row := range combinations {
		if int64(len(row)) > maxLength {
			maxLength = int64(len(row))
		}
	}

	for _, row := range combinations {
		if int64(len(row)) == maxLength {
			maxLists = append(maxLists, row)
		}
	}

	return maxLists
}

func findSublistWithLargestElement(list [][]int64) int64 {
	if len(list) == 0 {
		return -1
	}

	maxValue := list[0][0]
	maxIndex := 0

	for i, row := range list {
		for j, element := range row {
			if element > maxValue {
				maxValue = element
				maxIndex = i
				j = j
			}
		}
	}

	return int64(maxIndex)
}

func update_answer(answers []int64, cross []int64, side int64) []int64 {
	count := len(cross)
	for i := 0; i < len(answers); i++ {
		for _, value := range cross {
			if count == 0 {
				break
			}
			if answers[i] == value {
				answers[i] = side
				count--
				break
			}
		}

	}

	return answers
}

func update_original_list(list1 []int64, list2 []int64) []int64 {
	updatedList := make([]int64, 0)
	count := len(list2)

	for _, value1 := range list1 {
		matched := false
		for _, value2 := range list2 {
			if count == 0 {
				break
			}
			if value1 == value2 {
				matched = true
				count--
				break
			}
		}

		if !matched {
			updatedList = append(updatedList, value1)
		}
	}

	return updatedList
}

func update(current_list [][]int64, sizes []int64, answer []int64, side int64) ([]int64, []int64) {

	biggest_index := findSublistWithLargestElement(current_list)

	sizes = update_original_list(sizes, current_list[biggest_index])

	answer = update_answer(answer, current_list[biggest_index], side)

	return sizes, answer
}

func second_k_elements(list []int64, two_N int64) []int64 {

	if len(list) == 0 {
		return []int64{}
	}

	var first []int64
	var count int64
	count = 0

	for i := 0; i < len(list); i++ {
		if (count + list[i]) <= two_N {
			count += list[i]
			first = append(first, list[i])
		} else {
			break
		}
	}

	return first
}

func second_update(second []int64, sizes []int64, answer []int64, side int64) ([]int64, []int64) {

	sizes = update_original_list(sizes, second)
	answer = update_answer(answer, second, side)

	return sizes, answer
}

func findFirstZeroFromEnd(list []int64) int64 {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == 0 {
			return int64(i) + 1
		}
	}
	return -1
}

// Calling main
func main() {
	var list []string
	var input string
	var port int64
	var starboard int64
	var err error
	var answer []int64

	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		if input == "0" {
			break
		}
		list = append(list, input)
	}

	if len(list) != 0 {
		port, err = strconv.ParseInt(list[0], 0, 64)
		err = err
		port = port * 100
		starboard = port
		list = list[1:]
	} else {
		//fmt.Println("List is Empty")
		os.Exit(0)
	}

	sizes := convert(list)
	sizes = filterList(sizes, port)
	answer = sizes

	first := first_k_elements(sizes, (port + starboard))
	N := port
	combinations := combinations_to_n(first, N)
	list_to_N := equal_to_n(combinations, N)

	if len(list_to_N) == 0 {
		list_almost_N := findMaxLists(combinations)
		sizes, answer = update(list_almost_N, sizes, answer, 0)

	} else {
		sizes, answer = update(list_to_N, sizes, answer, 0)

	}

	second := second_k_elements(sizes, port)
	answer = update_answer(answer, second, -1)

	count := 0

	for _, items := range answer {
		if items == 0 || items == -1 {
			count++
		}
	}

	fmt.Println(count)

	for _, items := range answer {
		if items == 0 {
			fmt.Println("port")
		} else if items == -1 {
			fmt.Println("starboard")
		}
	}

}
