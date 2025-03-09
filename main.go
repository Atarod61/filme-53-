package main

import (
	functioen "film-53/Functioen"
	structs "film-53/Structs" // or "project/structs" if you rename the package

	// Corrected spelling
	"fmt"
)

func main() {
	u := structs.Allusers{ // or structs.Allusers if you rename the package
		Name: "Alireza",
		Age:  27,
	}
	result := functioen.DisplayInfo(u.Name) // Assuming function name is DisplayInfo and package name is function
	fmt.Println(result)

}
