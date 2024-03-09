package terminalGraphics

import (
    "math"
)

func interpolateU16(a uint16, b uint16, t float32) float32 {
    return float32(a) + ((float32(b)-float32(a))*t);
}

// distance between two points
func dist2D(p0 Coord2D, p1 Coord2D) float32 {
    return float32(math.Sqrt(math.Pow(float64(p1.X) - float64(p0.X), 2.0) + math.Pow(float64(p1.Y) - float64(p0.Y), 2.0)));
}
