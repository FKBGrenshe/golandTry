package main

import (
	"sync"
)

var mu sync.RWMutex
var icons map[string]string

func main() {

}

func traditionalInitialization(name string) string {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	mu.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

func loadIcons() {
	icons = make(map[string]string)
	icons["spades.png"] = "spades.png"
	icons["hearts.png"] = "hearts.png"
	icons["diamonds.png"] = "diamonds.png"
	icons["clubs.png"] = "clubs.png"
}
