package main
import "fmt"


type People struct {
	name string
}

func main() {
	peoples := []*People{
		{name: "Hilmi"}, 
		{name: "Bayu"},
		{name: "Muthus"}, 
		{name: "Eka"},
		{name: "Barru"}, 
		{name: "Asep"},
		{name: "Satrio"}, 
		{name: "Tata"},
		{name: "Iqbal"}, 
		{name: "Other Friends"},
	}
	printFriends := func(friends []*People) {
		for count, friend := range friends {
			fmt.Printf("%d %s \n", count+1, friend.name)
		}
	}

	printFriends(peoples)
}