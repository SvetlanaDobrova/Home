package main

import "fmt"

func showFurniture(f []furniture) {
	for _, fi := range f {
		fmt.Printf("   %s ( Высота %4.2fм, ширина %4.2fм, длина %4.2fм) \n", fi.name, fi.height, fi.width, fi.length)
	}
}

type furniture struct {
	height, width, length float64
	name                  string
}
