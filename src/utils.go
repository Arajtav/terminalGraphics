package terminalGraphics

import (
    "math"
)

func interpolateU16(a uint16, b uint16, t float32) float32 {
    return float32(a) + ((float32(b)-float32(a))*t);
}

// length of the line from (x0 y0) to (x1 y1)
func length(x0 uint16, y0 uint16, x1 uint16, y1 uint16) float32 {
    return float32(math.Sqrt(math.Pow(float64(x1) - float64(x0), 2.0) + math.Pow(float64(y1) - float64(y0), 2.0)));
}
