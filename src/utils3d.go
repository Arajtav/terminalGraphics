package terminalGraphics

// converts 3D position to position on screen
func Point3DToPoint2D(v Vertex, fv float32) Coord2D {
    return Coord2D{v.X/v.Z*fv, v.Y/v.Z*fv};
}

func DrawPoint3D(s *Canvas, v Vertex, fv float32, c Color) {
    s.SetPixel(Point3DToPoint2D(v, fv), c);
}

// Line from Vertex to Vertex
func DrawLine3D(s *Canvas, v0 Vertex, v1 Vertex, fv float32, c Color) {
    s.DrawLine(Point3DToPoint2D(v0, fv), Point3DToPoint2D(v1, fv), c)
}

func DrawTriangle3D(s *Canvas, v0 Vertex, v1 Vertex, v2 Vertex, fv float32, c Color) {
    s.DrawTriangle(Point3DToPoint2D(v0, fv), Point3DToPoint2D(v1, fv), Point3DToPoint2D(v2, fv), c);
}

func DrawModel(s *Canvas, m Model, fv float32, c Color) {
    for i := 0; i<len(m.triangles); i++ {
        DrawTriangle3D(s, m.vertices[m.triangles[i].i0], m.vertices[m.triangles[i].i1], m.vertices[m.triangles[i].i2], fv, c);
    }
}

// For debug, renders only points
func DrawModelV(s *Canvas, m Model, fv float32, c Color) {
    for i := 0; i<len(m.vertices); i++ {
        DrawPoint3D(s, m.vertices[i], fv, c);
    }
}
