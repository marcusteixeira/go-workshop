// Looping Over Arrays and Slices
package main

import (
	"fmt"
)

func main() {
	times := []string{"Flamengo", "Botafogo", "Vasco", "Fluminense"}
	sgbd := []string{"MSSQL", "MongoDB", "PostgreSQL", "MariaDB"}

	fmt.Println("############## - ######################")
	for i := 0; i < len(times); i++ {
		fmt.Println("Time: ", times[i])
	}
	fmt.Println("############## - ######################")
	for x := 0; x < len(sgbd); x++ {
		fmt.Println("SGBD: ", sgbd[x])
	}
}
