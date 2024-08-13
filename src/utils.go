package terminalGraphics

import (
    "fmt"
	"math"
)

// Linear interpolation
func Interpolate(a float32, b float32, t float32) float32 { return a + ((b-a)*t) }

// Distance between two points.
func dist2D(a Vec2, b Vec2) float32 {
    return float32(math.Sqrt(
        math.Pow(float64(b.X)-float64(a.X), 2.0) + math.Pow(float64(b.Y)-float64(a.Y), 2.0),
    ))
}

// Works like Math.Round
func roundF32ToI32(v float32) int32 {
    if v < 0 { return int32(v) }
    return int32(v + 0.5)
}

// Converts degrees to radians
func DegreesToRadians(degrees float32) float32 { return degrees * (math.Pi / 180.0) }

// Prints escape sequence which should move cursor to the top left corner of the terminal
func Chome() { fmt.Print("\033[H") }

// Prints escape sequence which should clear the terminal
func ClearTerm() { fmt.Print("\x1b\x5b\x48\x1b\x5b\x32\x4a\x1b\x5b\x33\x4a") }

// Adds 2 Vec3
func AddVec3(v0 Vec3, v1 Vec3) Vec3 {
    return Vec3{v0.X+v1.X, v0.Y+v1.Y, v0.Z+v1.Z}
}

// Subtracts 2 Vec3
func SubVec3(v0 Vec3, v1 Vec3) Vec3 {
    return Vec3{v0.X-v1.X, v0.Y-v1.Y, v0.Z-v1.Z}
}
