package main

import (
	"fmt"

	ink "github.com/Yandelf00/GhostInk/ghostink"
)

func main() {
	ink.Haunt("something")
	shade_instance := ink.Get_Shades()

	fmt.Println("this is todo : ", shade_instance.TODO)
}
