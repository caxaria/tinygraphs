// Package extract provides a set of functions to extract parameters from an http:Request.
package extract

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/format"
)

// Size returns the value of size param from HTTP request.
// Default value: 240
func Size(r *http.Request) int {
	strSize := r.FormValue("size")
	if len(strSize) > 0 {
		if size, errSize := strconv.ParseInt(strSize, 0, 64); errSize == nil {
			isize := int(size)
			if isize > 0 && isize < 1000 {
				return isize
			}
		}
	}
	return 240
}

// Format returns the specified format parameter defined in the http.Request.
// possible values: svg or jpeg
// default value JPEG
func Format(r *http.Request) format.Format {
	strFmt := strings.ToLower(r.FormValue("fmt"))
	if len(strFmt) > 0 {
		if strFmt == "svg" {
			return format.SVG
		} else if strFmt == "jpeg" || strFmt == "jpg" {
			return format.JPEG
		}
	}
	return format.JPEG
}

// Theme returns the theme param if defined in the http request.
// Default value base.
func Theme(r *http.Request) string {
	strTheme := strings.ToLower(r.FormValue("theme"))
	if len(strTheme) > 0 {
		return strTheme
	}
	return "base"
}

// Hexalines return the value of the hexalines parameter in the http.Request.
// possible values: 6 or 8. Default value : 6
func Hexalines(r *http.Request) int {
	s := strings.ToLower(r.FormValue("hexalines"))
	if len(s) > 0 {
		if n, err := strconv.ParseInt(s, 0, 64); err == nil {
			if n%6 == 0 || n%4 == 0 {
				return int(n)
			}
		}
	}
	return 6
}

// Lines return the value of the lines parameter in the http.Request.
// Default value is 6
func Lines(r *http.Request) int {
	s := strings.ToLower(r.FormValue("lines"))
	if len(s) > 0 {
		if n, err := strconv.ParseInt(s, 0, 64); err == nil {
			if n >= 4 {
				return int(n)
			}
		}
	}
	return 6
}

// Width returns the value of the 'w' parameter in the http.Request.
func Width(r *http.Request) int {
	strW := r.FormValue("w")
	if len(strW) > 0 {
		if w, errW := strconv.ParseInt(strW, 0, 64); errW == nil {
			iw := int(w)
			if iw > 0 {
				return iw
			}
		}
	}
	return 720
}

// WidthOrDefault returns the value of the 'w' parameter in the http.Request.
func WidthOrDefault(r *http.Request, v int) int {
	strW := r.FormValue("w")
	if len(strW) > 0 {
		if w, errW := strconv.ParseInt(strW, 0, 64); errW == nil {
			iw := int(w)
			if iw > 0 {
				return iw
			}
		}
	}
	return v
}

// Height returns the value of the 'h' parameter in the http.Request.
func Height(r *http.Request) int {
	strH := r.FormValue("h")
	if len(strH) > 0 {
		if h, errH := strconv.ParseInt(strH, 0, 64); errH == nil {
			ih := int(h)
			if ih > 0 {
				return ih
			}
		}
	}
	return 300
}

// HeightOrDefault returns the value of the 'h' parameter in the http.Request.
func HeightOrDefault(r *http.Request, v int) int {
	strH := r.FormValue("h")
	if len(strH) > 0 {
		if h, errH := strconv.ParseInt(strH, 0, 64); errH == nil {
			ih := int(h)
			if ih > 0 {
				return ih
			}
		}
	}
	return v
}

// XSquares returns the value of the 'xs' parameter in the http.Request.
// Used to defined the number of squares that are wanted in the x axis of an image
func XSquares(r *http.Request) int {
	strXS := r.FormValue("xs")
	if len(strXS) > 0 {
		if xs, errXS := strconv.ParseInt(strXS, 0, 64); errXS == nil {
			ixs := int(xs)
			if ixs > 0 {
				return ixs
			}
		}
	}
	return 50
}

// XTriangles returns the value of the 'xt' parameter in the http.Request.
// Used to defined the number of triangles that are wanted in the x axis of an image
func XTriangles(r *http.Request) int {
	strXT := r.FormValue("xt")
	if len(strXT) > 0 {
		if xt, errXT := strconv.ParseInt(strXT, 0, 64); errXT == nil {
			ixt := int(xt)
			if ixt > 0 {
				return ixt
			}
		}
	}
	return 50
}

// GX1 returns the value of the 'gx1' parameter in the http.Request.
// Used to defined the x coordinate of first point of the gradient vector.
func GX1OrDefault(r *http.Request, d uint8) uint8 {
	strX := r.FormValue("gx1")
	if len(strX) > 0 {
		if x, errX := strconv.ParseInt(strX, 0, 64); errX == nil {
			ix := uint8(x)
			if ix > 0 {
				return ix
			}
		}
	}
	return d
}

// GX2 returns the value of the 'gx2' parameter in the http.Request.
// Used to defined the x coordinate of second point of the gradient vector.
func GX2OrDefault(r *http.Request, d uint8) uint8 {
	strX := r.FormValue("gx2")
	if len(strX) > 0 {
		if x, errX := strconv.ParseInt(strX, 0, 64); errX == nil {
			ix := uint8(x)
			if ix > 0 {
				return ix
			}
		}
	}
	return d
}

// GY1 returns the value of the 'gy1' parameter in the http.Request.
// Used to defined the y coordinate of first point of the gradient vector.
func GY1OrDefault(r *http.Request, d uint8) uint8 {
	strY := r.FormValue("gy1")
	if len(strY) > 0 {
		if y, errY := strconv.ParseInt(strY, 0, 64); errY == nil {
			iy := uint8(y)
			if iy > 0 {
				return iy
			}
		}
	}
	return d
}

// GY2 returns the value of the 'gy2' parameter in the http.Request.
// Used to defined the y coordinate of the second point of the gradient vector.
func GY2OrDefault(r *http.Request, d uint8) uint8 {
	strY := r.FormValue("gy2")
	if len(strY) > 0 {
		if y, errY := strconv.ParseInt(strY, 0, 64); errY == nil {
			iy := uint8(y)
			if iy > 0 {
				return iy
			}
		}
	}
	return d
}

func GradientVector(r *http.Request, gx1, gy1, gx2, gy2 uint8) colors.GradientVector {
	x1 := GX1OrDefault(r, gx1)
	y1 := GY1OrDefault(r, gy1)
	x2 := GX2OrDefault(r, gx2)
	y2 := GY2OrDefault(r, gy2)
	return colors.GradientVector{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

func Probability(r *http.Request, dp float64) float64 {
	strP := r.FormValue("p")
	if len(strP) > 0 {
		if p, errP := strconv.ParseFloat(strP, 64); errP == nil {
			if p >= 0 && p <= 1 {
				return p
			}
		}
	}
	return dp
}
