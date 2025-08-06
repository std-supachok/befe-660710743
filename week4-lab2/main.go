package main

import (
	"fmt"
)

func main() {
	// var name string = "supachok"
	var age int = 20

	email := "hangnak_s@silpakorn.edu"
	gpa := 2.69

	firstname, lastname := "Supachok", "Hangnak"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)

}