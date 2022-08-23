package main
import "fmt"

func main() {
	n := 10
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf("// %d genap \n", i)
		} else {
			fmt.Printf("// %d ganjil \n", i)
		}
	}
}