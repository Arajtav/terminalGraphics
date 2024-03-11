package terminalGraphics

import (
    "math"
)

// t=0 is a, t=1 is b
func Interpolate(a float32, b float32, t float32) float32 {
    return a + ((b-a)*t);
}

// t=0 is a, t=1 is b
func InterpolateVec2(a Vec2, b Vec2, t float32) Vec2 {
    return Vec2{Interpolate(a.X, b.X, t), Interpolate(a.Y, b.Y, t)};
}

// Distance between two points. from a² + b² = c²
func dist2D(a Vec2, b Vec2) float32 {
    return float32(math.Sqrt(
        math.Pow(float64(b.X) - float64(a.X), 2.0) + math.Pow(float64(b.Y) - float64(a.Y), 2.0),    // I HATE U GO, WHY I NEED `,` ON END OF THIS LINE
    ));
}

// Works like Math.Round
func roundF32ToI32(v float32) int32 {
    if v < 0 { return int32(v); }
    return int32(v + 0.5);
}

// Converts degrees to radians
func DegreesToRadians(degrees float32) float32 {
	return degrees * (math.Pi / 180.0)
}
