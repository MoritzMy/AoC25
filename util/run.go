package util

import "fmt"

func Run[T any, R any](convert func(line string) (T, error), stars ...func(items []T) R) {
	input := ReadInputList(convert)
	for i, star := range stars {
		fmt.Printf("Solving Star %v:\n", i+1)
		result := star(input)
		fmt.Printf("Result: %v\n", result)
	}
}
