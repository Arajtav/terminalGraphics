package terminalGraphics

func Point3DToPoint2D(v Vertex, fv float32) Coord2D {
    return Coord2D{int16(float32(v.X)/float32(v.Z)*fv), int16(float32(v.Y)/float32(v.Z)*fv)};
}
