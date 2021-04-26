package main

import "fmt"

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is bool\n", i)
		case int:
			fmt.Printf("param #%d is int\n", i)
		case string:
			fmt.Printf("param #%d is string\n", i)
		}
	}
}
func main() {
	classifier(1, "test", true)
}
