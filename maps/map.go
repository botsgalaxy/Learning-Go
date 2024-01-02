package main

import "fmt"

func main() {
	websites := map[string]string { 
		"google": "https://google.com",
		"amazon": "https://amazon.com",
		"yahoo": "https://yahoo.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["google"])
	fmt.Println(websites["amazon"])
	fmt.Println(websites["yahoo"])


	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites,"google")
	fmt.Println(websites)
}