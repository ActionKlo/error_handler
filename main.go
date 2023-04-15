package main

import "fmt"

func main() {
	categories := []string{"Category2", "Category4", "NoCategory"}

	/*
		for _, cat := range categories {
			total, err := Products.TotalPrice(cat)
			if err == nil {
				fmt.Println(cat, "Total:", ToCurrenct(total))
			} else {
				fmt.Println(cat, "(no such category)")
			}
		}
	*/

	channel := make(chan ChannelMessage, 10)

	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrenct(message.Total))
		} else {
			//fmt.Println(message.CategoryError, "(no such category)")
			panic(message.CategoryError)
		}
	}
}
