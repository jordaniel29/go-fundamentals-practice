package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	colors := map[string]string {
		"red": "#ff0000",
		"green": "#aa000",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap (m map[string]string){
	for key,value := range m {
		fmt.Println(key, ":", value)
	}
}