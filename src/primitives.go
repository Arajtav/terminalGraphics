package terminalGraphics

// TODO: UV MAP
func Cube() Model {
    m := GetEmptyModel();

    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  1, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{-1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  1, -1}, Vec2{1, 0}});

    m.AddTriangle(Vertex{Vec3{ 1, -1,  1}, Vec2{1, 1}}, Vertex{Vec3{ 1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1, -1}, Vec2{1, 0}});

    m.AddTriangle(Vertex{Vec3{-1, -1,  1}, Vec2{1, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{ 1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1, -1}, Vec2{1, 0}});

    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{-1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1, -1}, Vec2{1, 0}});

    m.AddTriangle(Vertex{Vec3{ 1,  1, -1}, Vec2{1, 1}}, Vertex{Vec3{ 1, -1, -1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1, -1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1, -1}, Vec2{1, 0}});

    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1,  1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1,  1,  1}, Vec2{1, 0}});

    return m;
}

func Plane() Model {
    m := GetEmptyModel();

    m.AddTriangle(Vertex{Vec3{-1,  0, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{ 1,  0,  1}, Vec2{1, 1}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}});

    return m;
}
