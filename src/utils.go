package terminalGraphics

import (
    "math"
)

func interpolate16(a int16, b int16, t float32) float32 {
    return float32(a) + ((float32(b)-float32(a))*t);
}

// distance between two points
func dist2D(p0 Coord2D, p1 Coord2D) float32 {
    return float32(math.Sqrt(math.Pow(float64(p1.X) - float64(p0.X), 2.0) + math.Pow(float64(p1.Y) - float64(p0.Y), 2.0)));
}

func Point3DToPoint2D(v Vertex, fv float32) Coord2D {
    return Coord2D{int16(v.X/v.Z*fv), int16(v.Y/v.Z*fv)};
}
