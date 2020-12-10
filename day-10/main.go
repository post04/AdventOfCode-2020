package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var exampleinput = `0
16
10
15
5
1
11
7
19
6
12
4`
var input = `84
60
10
23
126
2
128
63
59
69
127
73
140
55
154
133
36
139
4
70
110
97
153
105
41
106
79
145
35
134
146
148
13
77
49
107
46
138
88
152
83
120
52
114
159
158
53
76
16
28
89
25
42
66
119
3
17
67
94
99
7
56
85
122
18
20
43
160
54
113
29
130
19
135
30
80
116
91
161
115
141
102
37
157
129
34
147
142
151
68
78
24
90
121
123
33
98
1
40`

func convertAll(strings []string) []int {
	var final []int
	for _, num := range strings {
		kek, err := strconv.Atoi(num)
		if err != nil {
		} else {
			final = append(final, kek)
		}
	}
	return final
}

//sort list lowest to highest and grab the differences between jolts
func solve(in []int) int {
	var ans = 0
	var threes = []int{1233333}
	var ones = []int{}
	var jolt = 0
	for _, num := range in {
		if (num - 1) != jolt {
			threes = append(threes, num)
			jolt = num
		} else {
			ones = append(ones, num)
			jolt = num
		}

	}
	ans = len(ones) * len(threes)
	return ans
}

// I COULDN'T DO PART TWO
// CREDIT: https://github.com/L33m4n123/adventOfCodeGo/blob/master/2020/day10/solution.go#L44
func findIterationCount(currIndex int, input []int, cache map[int]int) int {
	if currIndex == len(input)-1 {
		return 1
	}

	if val, ok := cache[currIndex]; ok {
		return val
	}
	count := 0
	for i := currIndex + 1; i < len(input); i++ {
		if input[i]-input[currIndex] <= 3 {
			count += findIterationCount(i, input, cache)
		}
	}
	cache[currIndex] = count
	return count
}
func solve2(base []int) int {
	cache := map[int]int{}
	possibleCombinations := findIterationCount(0, base, cache)
	return possibleCombinations
}

func main() {
	input1 := strings.Split(input, "\n")
	input := convertAll(input1)
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	a := solve(input)
	fmt.Println(a)
	a2 := solve2(input)
	fmt.Println(a2)
}
