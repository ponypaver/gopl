package main

import (
	"sync"
	"image"
)

var loadIconOnce sync.Once

func Icon(name string) image.Image {
	loadIconOnce.Do(loaIcons)
	return icons[name]
}

func loaIcons() {

}

func main() {
	var mu sync.Mutex

	mu.Lock()
}
