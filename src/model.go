package terminalGraphics

type Triangle struct {
    i0 uint16;
    i1 uint16;
    i2 uint16;
}

type Model struct {
    vertices    []Vec3;
    triangles   []Triangle;     // stores indexes of vertices, making triangles
}

// TODO FUNCTIONS TO ADD FACES

func Cube(size float32) Model {
    var m Model;
    m.vertices = append(m.vertices, Vec3{ size,  size,  size});
    m.vertices = append(m.vertices, Vec3{-size,  size,  size});
    m.vertices = append(m.vertices, Vec3{ size, -size,  size});
    m.vertices = append(m.vertices, Vec3{-size, -size,  size});
    m.vertices = append(m.vertices, Vec3{ size,  size, -size});
    m.vertices = append(m.vertices, Vec3{-size,  size, -size});
    m.vertices = append(m.vertices, Vec3{ size, -size, -size});
    m.vertices = append(m.vertices, Vec3{-size, -size, -size});
    // back
    m.triangles = append(m.triangles, Triangle{4, 7, 6});
    m.triangles = append(m.triangles, Triangle{4, 7, 5});
    // front
    m.triangles = append(m.triangles, Triangle{1, 3, 2});
    m.triangles = append(m.triangles, Triangle{1, 0, 2});
    // top
    m.triangles = append(m.triangles, Triangle{0, 1, 4});
    m.triangles = append(m.triangles, Triangle{5, 1, 4});
    // bottom
    m.triangles = append(m.triangles, Triangle{2, 3, 7});
    m.triangles = append(m.triangles, Triangle{2, 6, 7});
    // left
    m.triangles = append(m.triangles, Triangle{0, 2, 4});
    m.triangles = append(m.triangles, Triangle{6, 2, 4});
    // right
    m.triangles = append(m.triangles, Triangle{1, 3, 7});
    m.triangles = append(m.triangles, Triangle{1, 5, 7});
    return m;
}

// THIS IS FOR TEST, PROBABLY WILL BE REMOVED BECAUSE IT'S NOT NEEDED
func (m *Model) Move(v Vec3) {
    for i := 0; i<len(m.vertices); i++ {
        m.vertices[i].X += v.X;
        m.vertices[i].Y += v.Y;
        m.vertices[i].Z += v.Z;
    }
}
