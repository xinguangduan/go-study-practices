package main

import (
	"fmt"
	"image"
	"sync"
)

var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
	fmt.Println("load completed")
}

func loadIcon(s string) image.Image {
	return nil
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {
	Icon("left")
}
