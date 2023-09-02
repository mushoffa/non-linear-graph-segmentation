package main

import (
	"fmt"
)

func main() {
	data := read_dataset()

	fmt.Println(data)
}

// This is a helper function, you can make your own function
func read_dataset() []uint16 {
	var T int

	// This is example dataset array, you can make your custom struct
	var data []uint16

	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		var x, y uint16
		fmt.Scanf("%d %d", &x, &y)
		
		data = append(data, y)
	}

	return data
}