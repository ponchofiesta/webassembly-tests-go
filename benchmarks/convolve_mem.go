package benchmarks

import (
	"math"
	"syscall/js"
)

func ConvolveMem(data []byte, width int, height int, matrix []float64, factor float64) {
	//js.Global().Get("console").Call("debug", len(data))
	side := int(math.Sqrt(float64(len(matrix))))
	halfSide := int(side / 2)
	newData := make([]byte, width*height*4)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			outputIndex := (y*width + x) * 4
			r := 0.0
			g := 0.0
			b := 0.0
			for cy := 0; cy < side; cy++ {
				for cx := 0; cx < side; cx++ {
					scy := y + cy - halfSide
					scx := x + cx - halfSide
					if scy >= 0 && scy < height && scx >= 0 && scx < width {
						sourceIndex := (scy*width + scx) * 4
						modify := matrix[cy*side+cx]
						r += float64(data[sourceIndex]) * modify
						g += float64(data[sourceIndex+1]) * modify
						b += float64(data[sourceIndex+2]) * modify
					}
				}
			}
			newData[outputIndex] = byte(r * factor)
			newData[outputIndex+1] = byte(g * factor)
			newData[outputIndex+2] = byte(b * factor)
			newData[outputIndex+3] = data[outputIndex+3]
		}
	}
	copy(data, newData)
	js.Global().Get("console").Call("debug", "Go: convolve_mem: done")
}
