package main

import "fmt"

func showRoom(r room) {
	fmt.Printf("  Комната %s, (%4.1f кв.м)\n", r.name, r.area)
	fmt.Println("  Состав:")
	showFurniture(r.furniture)

}

type room struct {
	height, width, length, area float64
	name                        string
	furniture                   []furniture
}
