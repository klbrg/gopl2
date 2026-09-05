// Package lazyicon loads a table of icons on first use.
package lazyicon

import (
	"image"
	"sync"
)

// icons is a function, not a map: the table can be reached only by
// calling it, and the call performs the one-time initialization.
var icons = sync.OnceValue(func() map[string]image.Image {
	return map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
})

// Icon is concurrency-safe.
func Icon(name string) image.Image { return icons()[name] }

// loadIcon is a stand-in for a slow decoder.
func loadIcon(name string) image.Image {
	return image.NewRGBA(image.Rect(0, 0, 1, 1))
}
