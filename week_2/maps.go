package main

import "fmt"

var kingdoms map[string]string

func maps() {
	kingdoms = make(map[string]string)
	kingdoms["Noldor"] = "feanor"
	kingdoms["Teleri"] = "Fingolfin"
	fmt.Println(len(kingdoms))
	delete(kingdoms, "Teleri")
	fmt.Println(len(kingdoms))

	leaders := map[string][]string{
		"Noldor": []string{"Feanor"},
		"Teleri": []string{"Finwe"},
	}
	fmt.Println(leaders)
}
