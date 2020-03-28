package main

import "fmt"

func slices() {
	valar := [7]string{"Manwe", "Ulmo", "Aule", "Yavanna", "Varda", "Este", "Nienna"}
	var ladyValar []string
	ladyValar = append(ladyValar, valar[3:6]...) // []string{"Yavanna", "Varda", "Este"}
	fmt.Println(ladyValar[1])
}
