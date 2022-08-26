package main
import "fmt"

func main() {
	friends := []string{
		"muthus",
		"hilmi",
		"bayu",
		"barru",
		"satrio",
		"Pak Tata",
		"hasan",
		"clara",
		"giva",
		"cecep",
	}

	fmt.Println("Para Sahabat Training GO: ")
	for count, friend := range friends {
		fmt.Printf("%d %s \n", count+1, friend)
	}
}