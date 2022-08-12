package main

import (
	"bernoulli"
	"fmt"
)

func main() {
	bernoulliApp := bernoulli.New()
	bernoulliApp.AddFunc("helloWorld", func(...any) {
		fmt.Println("Hello World")
	})

	bernoulliApp.Exec("helloWorld")
}
