package main

import "fmt"

func main() {
	const (
		FlagA = 1 << iota
		FlagB
		FlagC
		FlagD
	)
	fmt.Println(FlagA, FlagB, FlagC, FlagD)
}
