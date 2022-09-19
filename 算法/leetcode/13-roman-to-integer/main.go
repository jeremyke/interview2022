package main

import "fmt"

func main() {
	//s := "MCMXCIV"
	s := "III"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var romanIntMap = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	if len(s) == 0 {
		return 0
	}
	var total int
	for i := 0; i < len(s); {
		if i < len(s)-1 {
			if v, ok := romanIntMap[s[i:i+2]]; ok {
				total += v
				i += 2
			} else {
				if v1, ok1 := romanIntMap[s[i:i+1]]; ok1 {
					total += v1
				}
				i++
			}
		} else {
			if v1, ok1 := romanIntMap[s[i:i+1]]; ok1 {
				total += v1
			}
			i++
		}

	}
	return total
}
