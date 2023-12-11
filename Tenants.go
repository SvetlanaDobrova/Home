package main

import "fmt"

func showTenants(h home) {
	for _, t := range h.Tenants {
		fmt.Printf("  %s, (%d)\n", t.Name, t.Age) //
	}
}

type tenant struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
