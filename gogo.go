package main

import "fmt"

type home struct {
	area    float64
	address string
	rooms   []room
	tenants []tenant
}

func showHome(h home) {
	fmt.Println(h.address)
	fmt.Println(" Проживают:")
	showTenants(h)
	fmt.Println(" Комнаты:")
	for _, r := range h.rooms {
		showRoom(r)
	}
}
