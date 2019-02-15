package main

import "fmt"

func main() {


	test := 444444;

	for {
		go func() {
			fmt.Println(test)
		}()
	}

}
