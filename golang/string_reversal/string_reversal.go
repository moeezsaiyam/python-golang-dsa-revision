package main

import "fmt"

func string_reversal(str string) string {
	runes := []rune(str)

	for i := range len(runes) / 2 {
		fmt.Println(i)
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	rev := string_reversal("moeez")
	fmt.Println(rev)
}
