package main

import "fmt"

/**
コメント

 */
//hoge
func main() {
	fmt.Println("test")
	aAa := 1
	defer aAa = 2

	fmt.Println(aAa)
}
