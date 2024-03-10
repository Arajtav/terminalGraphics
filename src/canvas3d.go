package terminalGraphics

func DrawPoint3D(s *Canvas, v Vertex, fv float32, c Color) {
    s.SetPixel(Point3DToPoint2D(v, fv), c);
}

func DrawLine3D(s *Canvas, v0 Vertex, v1 Vertex, fv float32, c Color) {
    s.DrawLine(Point3DToPoint2D(v0, fv), Point3DToPoint2D(v1, fv), c)
}

func DrawTriangle3D(s *Canvas, v0 Vertex, v1 Vertex, v2 Vertex, fv float32, c Color) {
    DrawLine3D(s, v0, v1, fv, c);
    DrawLine3D(s, v1, v2, fv, c);
    DrawLine3D(s, v2, v0, fv, c);
}
