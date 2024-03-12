package terminalGraphics

import (
	"math"
	"fmt"
)

// Interpolation of a and b, when t is 0, function will return a, when t is 1, function will return b
func Interpolate(a float32, b float32, t float32) float32 {
    return a + ((b-a)*t);
}

// Same as Interpolate() but for Vec2
func InterpolateVec2(a Vec2, b Vec2, t float32) Vec2 {
    return Vec2{Interpolate(a.X, b.X, t), Interpolate(a.Y, b.Y, t)};
}

// Distance between two points. From a² + b² = c²
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

// Prints escape sequence which should move cursor to top left corner of terminal
func Chome() {
    fmt.Print("\033[H");
}

// Prints escape sequence which should clear terminal
func ClearTerm() {
    fmt.Print("\x1b\x5b\x48\x1b\x5b\x32\x4a\x1b\x5b\x33\x4a");
}
