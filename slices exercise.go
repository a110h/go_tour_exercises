package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var ss [][]uint8
	for i := 0; i < dy; i++ {

		var s []uint8
		for j := 0; j < dx; j++ {
			s = append(s, uint8(j^i))
		}
		ss = append(ss, s)
	}
	return ss

}

func main() {
	pic.Show(Pic)
}