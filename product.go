package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

type ProductSlice []*Product

var Products = ProductSlice{
	{"Name1", "Category1", 10},
	{"Name2", "Category1", 20},
	{"Name3", "Category2", 30},
	{"Name4", "Category2", 40},
	{"Name5", "Category3", 50},
	{"Name6", "Category3", 60},
	{"Name7", "Category4", 70},
	{"Name8", "Category4", 80},
}

func ToCurrenct(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}
