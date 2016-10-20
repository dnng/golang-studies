// gopl.io/ch3/surface2.go
// surface2 serves the SVG image via a web server
package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (= 30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		// No validation at all! yay!
		// Call the server with
		//    http://localhost:8000/?peaks=%23ff0000&valleys=%230000ff
		// for it to work!
		// Explanation: http://i.imgur.com/dzDsW.jpg?fb
		peaks := r.FormValue("peaks")
		valleys := r.FormValue("valleys")

		fmt.Println("start")
		fmt.Println(peaks)
		fmt.Println(valleys)
		var surf string = surface(peaks, valleys)
		fmt.Fprintf(w, "%s", surf)
		fmt.Println("end")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(peaks string, valleys string) string {
	var s string
	var u string
	s = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j, peaks, valleys)
			bx, by, _ := corner(i, j, peaks, valleys)
			cx, cy, _ := corner(i, j+1, peaks, valleys)
			dx, dy, color := corner(i+1, j+1, peaks, valleys)

			// Solution for exercise 3.1
			if math.IsInf(ax, 0) || math.IsInf(bx, 0) || math.IsInf(cx, 0) || math.IsInf(dx, 0) {
				continue
			}
			u = fmt.Sprintf("<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
			s += u
		}
	}
	s += "</svg>"
	return s
}

func corner(i int, j int, pe string, va string) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	var co string

	// Compute surface height z.
	z := f(x, y)
	if z >= 0 {
		co = pe
	} else {
		co = va
	}

	// Solution for exercise 3.1
	if math.IsInf(z, 0) {
		return math.Inf(0), math.Inf(0), co
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, co
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
