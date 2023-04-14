package main

import "fmt"

func main() {
	categories := []string{"Category2", "Category4", "NoCategory"}

	for _, cat := range categories {
		total, err := Products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, "Total:", ToCurrenct(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}
}
