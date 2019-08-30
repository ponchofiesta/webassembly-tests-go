package benchmarks

func ConvolveVideo(data []byte, width int, height int, matrix [][]float64, factor float64, count int) []byte {
	side := len(matrix)
	half := side / 2
	for i := 0; i < count; i++ {
		for y := 1; y < height-1; y++ {
			for x := 1; x < width-1; x++ {
				index := (y*width + x) * 4
				r := 0.0
				g := 0.0
				b := 0.0
				for cy := 0; cy < side; cy++ {
					for cx := 0; cx < side; cx++ {
						cpx := ((y+(cy-half))*width + (x + (cx - half))) * 4
						r += float64(data[cpx]) * matrix[cy][cx]
						g += float64(data[cpx+1]) * matrix[cy][cx]
						b += float64(data[cpx+2]) * matrix[cy][cx]
					}
				}
				data[index] = byte(factor * r)
				data[index+1] = byte(factor * g)
				data[index+2] = byte(factor * b)
			}
		}
	}
	return data
}
