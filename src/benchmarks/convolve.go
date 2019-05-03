package benchmarks

import (
	"math"
	"syscall/js"
)

func Convolve(canvas js.Value, matrix []float64, factor float64) {
	side := int(math.Sqrt(float64(len(matrix))))
	halfSide := int(side / 2)
	context := canvas.Call("getContext", "2d")
	source := context.Call("getImageData", 0.0, 0.0, canvas.Get("width").Int(), canvas.Get("height").Int())
	sourceData := source.Get("data")
	imageWidth := source.Get("width").Int()
	imageHeight := source.Get("height").Int()
	output := context.Call("createImageData", imageWidth, imageHeight)
	outputData := output.Get("data")

	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			outputIndex := (y*imageWidth + x) * 4
			r := 0.0
			g := 0.0
			b := 0.0
			for cy := 0; cy < side; cy++ {
				for cx := 0; cx < side; cx++ {
					scy := y + cy - halfSide
					scx := x + cx - halfSide
					if scy >= 0 && scy < imageHeight && scx >= 0 && scx < imageWidth {
						sourceIndex := (scy*imageWidth + scx) * 4
						modify := matrix[cy*side+cx]
						r += sourceData.Index(sourceIndex).Float() * modify
						g += sourceData.Index(sourceIndex+1).Float() * modify
						b += sourceData.Index(sourceIndex+2).Float() * modify
					}
				}
			}
			outputData.SetIndex(outputIndex, r*factor)
			outputData.SetIndex(outputIndex+1, g*factor)
			outputData.SetIndex(outputIndex+2, b*factor)
			outputData.SetIndex(outputIndex+3, sourceData.Index(outputIndex+3))
		}
	}

	context.Call("putImageData", output, 0, 0)
}
