package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func dalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) < float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	lingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	lingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if lingkaran1 && lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
