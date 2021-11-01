package main

import "fmt"

func square(x int) int {
	return x * x
}

func inc(x int) int {
	return x + 1
}

func main() {
	2 |> square |> inc |> fmt.Println
}
