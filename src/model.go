package terminalGraphics

import "math"

type triangle struct {
    i0 uint16;
    i1 uint16;
    i2 uint16;
}

/*  model stores data in ready to process form
    that means rotation and scale is applied when adding data to model, not when reading from it
    that also means rotation and scale needs to be private, and only changeable by method which will also recalculate every Vertex
*/

// Model, stores transformed 3d data, covers transformations in object space. Position can be changed, it will be calculated in render time.
type Model struct {
    vertices    []Vertex;
    triangles   []triangle;     // stores indexes of vertices, making triangles
    rotation    Vec3;           // object rotation in radians
    scale       Vec3;
    Position    Vec3;
}

// Creates empty model
func GetEmptyModel() Model {
    var m Model;
    m.scale = Vec3{1.0, 1.0, 1.0};
    return m;
}

// Returns current scale of the model
func (m *Model) GetScale() Vec3 { return m.scale; }

// Changes scale of the model, recalculates every vertex
func (m *Model) SetScale(s Vec3) {
    r := Vec3{s.X/m.scale.X, s.Y/m.scale.Y, s.Z/m.scale.Z}; // normalize then multiply to the new value
    for i := 0; i<len(m.vertices); i++ {
        m.vertices[i].Position.X *= r.X;
        m.vertices[i].Position.Y *= r.Y;
        m.vertices[i].Position.Z *= r.Z;
    }
    m.scale = s;
}

// Returns current rotation of the model
func (m *Model) GetRotation() Vec3 { return m.rotation; }

// Changes rotation of the model, recalculates every vertex
func (m *Model) SetRotation(rot Vec3) {
    for i := 0; i<len(m.vertices); i++ {
        // TODO: IDK, REVERSING ROTATION DOESN'T WORK WHEN DOING ALL AT ONE
        Rotate3D(&(m.vertices[i].Position), Vec3{0, 0, -m.rotation.Z});
        Rotate3D(&(m.vertices[i].Position), Vec3{0, -m.rotation.Y, 0});
        Rotate3D(&(m.vertices[i].Position), Vec3{-m.rotation.X, 0, 0});
        Rotate3D(&(m.vertices[i].Position), rot);
    }
    m.rotation = rot;
}

// If vertex doesn't exist, this method will add it, if it does, function will return id of that vertex
// Checks if Vertex exists in model, will return 65535 if Vertex doesn't exist. Vertex MUST already be transformed
func (m *Model) getVertexId(v Vertex) uint16 {
    for i := 0; i<len(m.vertices); i++ {
        if m.vertices[i] == v {
            return uint16(i);
        }
    }
    return 65535;
}

func (m *Model) SetVertex(v Vertex) uint16 {
    v.Position.X *= m.scale.X;
    v.Position.Y *= m.scale.Y;
    v.Position.Z *= m.scale.Z;
    Rotate3D(&v.Position, m.rotation);
    i := m.getVertexId(v); if i != 65535 { return i; }
    m.vertices = append(m.vertices, v);
    return uint16(len(m.vertices)) - 1;
}

func (m *Model) AddTriangle(v0 Vertex, v1 Vertex, v2 Vertex) {
    m.triangles = append(m.triangles, triangle{m.SetVertex(v0), m.SetVertex(v1), m.SetVertex(v2)});
}

// Rotates vertex v, by rot. Rotation is in radians
func Rotate3D(v *Vec3, rot Vec3) {
    cosa := float32(math.Cos(float64(rot.Z)));
    sina := float32(math.Sin(float64(rot.Z)));

    cosb := float32(math.Cos(float64(rot.Y)));
    sinb := float32(math.Sin(float64(rot.Y)));

    cosc := float32(math.Cos(float64(rot.X)));
    sinc := float32(math.Sin(float64(rot.X)));

    Axx := cosa*cosb;
    Axy := cosa*sinb*sinc - sina*cosc;
    Axz := cosa*sinb*cosc + sina*sinc;

    Ayx := sina*cosb;
    Ayy := sina*sinb*sinc + cosa*cosc;
    Ayz := sina*sinb*cosc - cosa*sinc;

    Azx := -sinb;
    Azy := cosb*sinc;
    Azz := cosb*cosc;

    // TODO: THIS CAN BE EXECUTED IN LOOP FOR EVERY VERTEX, NO NEED TO RECALCULATE ALL OF VARIABLES ABOVE
    px := v.X;
    py := v.Y;
    pz := v.Z;

    v.X = Axx*px + Axy*py + Axz*pz;
    v.Y = Ayx*px + Ayy*py + Ayz*pz;
    v.Z = Azx*px + Azy*py + Azz*pz;
}
