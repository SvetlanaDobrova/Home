package main

import "fmt"

type home struct {
	Area    float64  `json:"area"`
	Address string   `json:"address"`
	Rooms   []room   `json:"rooms"`
	Tenants []tenant `json:"tenants"`
}

func showHome(h home) {
	fmt.Println(h.Address)
	fmt.Println(" Проживают:")
	showTenants(h)
	fmt.Println(" Комнаты:")
	for _, r := range h.Rooms {
		showRoom(r)
	}
}
