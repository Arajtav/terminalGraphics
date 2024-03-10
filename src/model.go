package terminalGraphics

type triangle struct {
    i0 uint16;
    i1 uint16;
    i2 uint16;
}

/*  model stores data in ready to process form
    that means rotation and scale is applied when adding data to model, not when reading from it
    that also means rotation and scale needs to be private, and only changeable by method which will also recalculate every Vertex
*/

// Model, stores transformed 3d data, covers transformed in object space
type Model struct {
    vertices    []Vec3;
    triangles   []triangle;     // stores indexes of vertices, making triangles
    rotation    Vec3;
    scale       Vec3;
}

// Creates empty model
func GetEmptyModel() Model {
    var m Model;
    m.scale = Vec3{1.0, 1.0, 1.0};
    return m;
}

// Returns current scale of the model
func (m *Model) GetScale() Vec3 { return m.scale; }

// Changes scale of Model, recalculate every vertex
func (m *Model) SetScale(s Vec3) {
    r := Vec3{s.X/m.scale.X, s.Y/m.scale.Y, s.Z/m.scale.Z};
    for i := 0; i<len(m.vertices); i++ {
        m.vertices[i].X *= r.X;
        m.vertices[i].Y *= r.Y;
        m.vertices[i].Z *= r.Z;
    }
    m.scale = s;
}

// Checks if Vertex exists in model, will return 65535 if Vertex doesn't exist. Vertex MUST already be transformed
func (m *Model) getVertexId(v Vec3) uint16 {
    for i := 0; i<len(m.vertices); i++ {
        if m.vertices[i] == v {
            return uint16(i);
        }
    }
    return 65535;
}

// If vertex doesn't exist, this method will add it, if it does, function will return id of that vertex
func (m *Model) SetVertex(v Vec3) uint16 {
    v.X *= m.scale.X;
    v.Y *= m.scale.Y;
    v.Z *= m.scale.Z;
    // TODO: APPLY ROTATION
    i := m.getVertexId(v); if i != 65535 { return i; }
    m.vertices = append(m.vertices, v);
    return uint16(len(m.vertices)) - 1;
}

func (m *Model) AddTriangle(v0 Vec3, v1 Vec3, v2 Vec3) {
    m.triangles = append(m.triangles, triangle{m.SetVertex(v0), m.SetVertex(v1), m.SetVertex(v2)});
}

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

// TODO: delete that after adding word space, moving in object space is useless
func (m *Model) Move(v Vec3) {
    for i := 0; i<len(m.vertices); i++ {
        m.vertices[i].X += v.X;
        m.vertices[i].Y += v.Y;
        m.vertices[i].Z += v.Z;
    }
}
