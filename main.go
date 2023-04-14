package main

import "fmt"

func main() {
	categories := []string{"Category2", "Category4"}

	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrenct(total))
	}
}
