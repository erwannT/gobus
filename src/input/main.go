package main

import (
	"fmt"
)

func main() {

	var t1 string
	_, err := fmt.Scan(&t1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Hello %s\n", t1)

}
