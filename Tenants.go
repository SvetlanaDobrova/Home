package main

import "fmt"

func showTenants(h home) {
	for _, t := range h.tenants {
		fmt.Printf("  %s, (%d)\n", t.name, t.age) //
	}
}

type tenant struct {
	name string
	age  int
}
