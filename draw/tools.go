package draw

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"
)

// getRandomColor returns a random color between c1 and c2
func randomColor(c1, c2 color.RGBA) color.RGBA {
	r := rand.Intn(2)
	if r == 1 {
		return c1
	}
	return c2
}

func fillFromRGBA(c color.RGBA) string {
	return fmt.Sprintf("fill:rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func colorFromKey(key string, color1, color2 color.RGBA, index int) color.RGBA {
	s := hex.EncodeToString([]byte{key[index]})
	if r, err := strconv.ParseInt(s, 16, 0); err == nil {
		if r%2 == 0 {
			return color1
		}
		return color2
	} else {
		log.Println("Error calling ParseInt(%v, 16, 0)", s, err)
	}
	return color1
}