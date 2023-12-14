package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type key struct {
	i, ans, si int
	row bool
}

func main() {
	file, err := os.Open("day13input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	squares := make([][]string, 0)
	square := make([]string, 0)
	for _, line := range lines {
		if line == "" && len(square) != 0 {
			squares = append(squares, square)
			square = make([]string, 0)
		} else {
			square = append(square, line)
		}
	}
	squares = append(squares, square)
	part2(squares)
}

var noSmudge bool = true

func part1(squares [][]string) {
	ans := 0
	for i, square := range squares {
		ans += getRowReflectionValue(square, i)
		ans += getColumnReflectionValue(square, i)
	}
	fmt.Println(ans)
}

func part2(squares [][]string) {
	ans := 0
	for i, square := range squares {
		getRowReflectionValue(square, i)
		getColumnReflectionValue(square, i)
	}
	noSmudge = false
	for i, square := range squares {
		ans += getRowReflectionValue(square, i)
		ans += getColumnReflectionValue(square, i)
	}
	fmt.Println(ans)
}

var cache = make(map[key]bool)
func getRowReflectionValue(square []string, squareIndex int) int {
	for i := 0; i < len(square)-1; i++ {
		cacheKey := key{i:i, ans:100*(i+1), row:true, si:squareIndex}
		if _, ok := cache[cacheKey]; ok {
			continue
		} else if (square[i] == square[i+1] || (!noSmudge && checkForSmudge(square[i], square[i+1]))) && checkRowReflection(square, i) {
			cache[cacheKey] = true
			return 100*(i+1)
		}
	}
	return 0
}

func getColumnReflectionValue(square []string, squareIndex int) int {
	s := condenseMatrix(transpose(square))
	for i := 0; i < len(s)-1; i++ {
		cacheKey := key{i:i, ans:i+1, row:false, si:squareIndex}
		if _, ok := cache[cacheKey]; ok {
			continue
		} else if (s[i] == s[i+1] || (!noSmudge && checkForSmudge(s[i], s[i+1]))) && checkRowReflection(s, i) {
			cache[cacheKey] = true
			return i+1
		}
	}
	return 0
}

func transpose(slice []string) [][]string {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]string, xl)
    for i := range result {
        result[i] = make([]string, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = string(slice[j][i])
        }
    }
    return result
}

func checkRowReflection(square []string, index int) bool {
	topHalf := removeEmptyLines(square[:index+1])
	bottomSize := len(square) - len(topHalf)
	if bottomSize < len(topHalf) {
		topHalf = topHalf[len(topHalf)-bottomSize:]
	}
	var bottomHalf []string = nil
	bottomHalf = removeEmptyLines(square[index+1:index + len(topHalf)+1])
	if len(topHalf) > len(bottomHalf) {
		diff := len(topHalf)-len(bottomHalf)
		topHalf = topHalf[diff:]
	}
	reverseBottom := reverse(bottomHalf)
	equal := testEq(topHalf, reverseBottom)
	return equal
}

func condenseMatrix(matrix [][]string) []string {
	ans := make([]string, 0)
	for _, row := range matrix {
		line := ""
		for _, col := range row {
			line += col
		}
		ans = append(ans, line)
	}
	return ans
}

func testEq(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] && (!checkForSmudge(a[i], b[i]) || noSmudge) {
            return false
        }
    }
    return true
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    	s[i], s[j] = s[j], s[i]
	}
	return s
}

func removeEmptyLines(s []string) []string {
	n := make([]string, 0)
	for _, v := range s {
		if v != "" {
			n = append(n, v)
		}
	}
	return n
}

func checkForSmudge(a string, b string) bool {
	diffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs == 1 || diffs == 0
}