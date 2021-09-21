package main

import "fmt"

func main() {
	//var mapMethod2 map[string]string
	mapMethod2 := make(map[string]string)
	mapMethod2["ray"] = "Name"
	mapMethod2["box"] = "Thing"
	fmt.Print(mapMethod2)
	mapMethod1 := map[string]string{
		"color": "red",
		"choco": "brown",
	}
	fmt.Print(mapMethod1)
}
