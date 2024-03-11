package terminalGraphics

// converts 3D position to position on screen
func Point3DToPoint2D(v Vec3, fv float32) Vec2 {
    return Vec2{v.X/v.Z*fv, v.Y/v.Z*fv};
}

func DrawPoint3D(s *Canvas, v Vec3, fv float32, c Color) {
    s.SetPixel(Point3DToPoint2D(v, fv), c);
}

// Line from Vec3 to Vec3
func DrawLine3D(s *Canvas, a Vec3, b Vec3, fv float32, c Color) {
    s.DrawLine(Point3DToPoint2D(a, fv), Point3DToPoint2D(b, fv), c)
}

func DrawTriangle3D(s *Canvas, v0 Vec3, v1 Vec3, v2 Vec3, fv float32, c Color) {
    s.DrawTriangle(Point3DToPoint2D(v0, fv), Point3DToPoint2D(v1, fv), Point3DToPoint2D(v2, fv), c);
}

func DrawModel(s *Canvas, m Model, fv float32, c Color) {
    for i := 0; i<len(m.triangles); i++ {
        // calculate Vertex position in world space
        v0 := AddVec3(m.vertices[m.triangles[i].i0], m.Position);
        v1 := AddVec3(m.vertices[m.triangles[i].i1], m.Position);
        v2 := AddVec3(m.vertices[m.triangles[i].i2], m.Position);
        DrawTriangle3D(s, v0, v1, v2, fv, c);
    }
}

// For debug, renders only points
func DrawModelV(s *Canvas, m Model, fv float32, c Color) {
    for i := 0; i<len(m.vertices); i++ {
        DrawPoint3D(s, m.vertices[i], fv, c);
    }
}

func AddVec3(v0 Vec3, v1 Vec3) Vec3 {
    v0.X += v1.X;
    v0.Y += v1.Y;
    v0.Z += v1.Z;
    return v0;
}
