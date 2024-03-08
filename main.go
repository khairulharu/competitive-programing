package main

import (
	"fmt"
	"math"
	"strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int64) int64 {
	n := len(arr)
	var dX int64
	var dY int64
	for i := 0; i < n; i++ {
		dX += arr[i][i]
		dY += arr[i][n-i-1]
	}
	return int64(math.Abs((float64(dX) - float64(dY))))
}

// coding interview fizzbuzz
func fizzBuzz(total int64) {
	for i := 1; i <= int(total); i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzBuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)

		}
	}
}

// coding interiview palindrome
func isPalindrome(request string) bool {
	tempString := strings.Split(request, "")

	for i := 0; i < len(tempString)/2; i++ {
		if tempString[i] != tempString[len(tempString)-i-1] {
			return false
		}
	}

	return true
}

func staircase(n int32) {
	for x := 1; x <= int(n); x++ {
		for i := 1; i <= int(n); i++ {
			if i <= int(n-int32(x)) {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func miniMaxSum(arr []int32) []int32 {
	var sortArr = []int32{}

	for i := 0; i < len(arr)-1; i++ {
	}

	return sortArr
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here

	n := len(candles)

	var t int32 = 0

	var l int32 = candles[0]

	for i := 0; i < n; i++ {
		if l == candles[i] {
			t++
		} else if l < candles[i] {
			l = candles[i]
			t = 1
		}
	}

	return t
}

func timeConversion(s string) string {
	// Write your code here

}

func main() {
	t := birthdayCakeCandles([]int32{3, 4, 2, 1})
	fmt.Println(t)

}
