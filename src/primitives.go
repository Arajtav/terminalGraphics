package terminalGraphics

// TODO: UV MAP
func Cube() Model {
    m := GetEmptyModel();

    // back
    m.AddTriangle(Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1, -1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1, -1}, Vec2{0, 0}});
    // front
    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}});
    // // top
    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{-1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}});
    // bottom
    m.AddTriangle(Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}});
    // left
    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{ 1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 0}});
    // right
    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}});
    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}});

    return m;
}

func Plane() Model {
    m := GetEmptyModel();

    m.AddTriangle(Vertex{Vec3{-1,  0, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}});
    m.AddTriangle(Vertex{Vec3{ 1,  0,  1}, Vec2{1, 1}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}});

    return m;
}
