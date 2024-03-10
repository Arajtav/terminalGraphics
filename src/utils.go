package terminalGraphics

import (
    "math"
)

// t=0 is a, t=1 is b
func Interpolate(a float32, b float32, t float32) float32 {
    return a + ((b-a)*t);
}

// t=0 is a, t=1 is b
func InterpolateCoord(a Coord2D, b Coord2D, t float32) Coord2D {
    return Coord2D{Interpolate(a.X, b.X, t), Interpolate(a.Y, b.Y, t)};
}

// distance between two points, from a² + b² = c²
func dist2D(a Coord2D, b Coord2D) float32 {
    return float32(math.Sqrt(
        math.Pow(float64(b.X) - float64(a.X), 2.0) + math.Pow(float64(b.Y) - float64(a.Y), 2.0),    // I HATE U GO, WHY I NEED `,` ON END OF THIS LINE
    ));
}

// works like Math.Round
func roundF32ToI32(v float32) int32 {
    if v < 0 { return int32(v); }
    return int32(v + 0.5);
}
