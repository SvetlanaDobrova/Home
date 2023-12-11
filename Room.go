package main

import "fmt"

func showRoom(r room) {
	fmt.Printf("  Комната %s, (%4.1f кв.м)\n", r.Name, r.Area)
	fmt.Println("  Состав:") //
	showFurniture(r.Furniture)

}

type room struct {
	Height    float64     `json:"height,omitempty"`
	Width     float64     `json:"width,omitempty"`
	Length    float64     `json:"length,omitempty"`
	Area      float64     `json:"area,omitempty"`
	Name      string      `json:"name,omitempty"`
	Furniture []furniture `json:"furniture,omitempty"`
}
