package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["GO"] = "golang"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["GO"])

	delete(languages, "PY")
	fmt.Println(languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
