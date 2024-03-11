package terminalGraphics

func Cube() Model {
    m := GetEmptyModel();

    // back
    m.AddTriangle(Vec3{ 1,  1, -1}, Vec3{-1, -1, -1}, Vec3{ 1, -1, -1});
    m.AddTriangle(Vec3{ 1,  1, -1}, Vec3{-1, -1, -1}, Vec3{-1,  1, -1});
    // front
    m.AddTriangle(Vec3{-1,  1,  1}, Vec3{-1, -1,  1}, Vec3{ 1, -1,  1});
    m.AddTriangle(Vec3{-1,  1,  1}, Vec3{ 1,  1,  1}, Vec3{ 1, -1,  1});
    // // top
    m.AddTriangle(Vec3{ 1,  1,  1}, Vec3{-1,  1,  1}, Vec3{ 1,  1, -1});
    m.AddTriangle(Vec3{-1,  1, -1}, Vec3{-1,  1,  1}, Vec3{ 1,  1, -1});
    // bottom
    m.AddTriangle(Vec3{ 1, -1,  1}, Vec3{-1, -1,  1}, Vec3{-1, -1, -1});
    m.AddTriangle(Vec3{ 1, -1,  1}, Vec3{ 1, -1, -1}, Vec3{-1, -1, -1});
    // left
    m.AddTriangle(Vec3{ 1,  1,  1}, Vec3{ 1, -1,  1}, Vec3{ 1,  1, -1});
    m.AddTriangle(Vec3{ 1, -1, -1}, Vec3{ 1, -1,  1}, Vec3{ 1,  1, -1});
    // right
    m.AddTriangle(Vec3{-1,  1,  1}, Vec3{-1, -1,  1}, Vec3{-1, -1, -1});
    m.AddTriangle(Vec3{-1,  1,  1}, Vec3{-1,  1, -1}, Vec3{-1, -1, -1});

    return m;
}
