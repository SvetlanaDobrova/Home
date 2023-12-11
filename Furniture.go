package main

import "fmt"

func showFurniture(f []furniture) {
	for _, fi := range f {
		fmt.Printf("   %s ( Высота %4.2fм, ширина %4.2fм, длина %4.2fм) \n", fi.Name, fi.Height, fi.Width, fi.Length)
	}
}

type furniture struct {
	Height float64 `json:"height,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Length float64 `json:"length,omitempty"`
	Name   string  `json:"name,omitempty"`
}
