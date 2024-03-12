package terminalGraphics

type Camera struct {
    Fv          float32;
    Rotation    Vec3;
    Position    Vec3;
}

// Renderer to which you can add models
type World struct {
    models      []*Model;
    Cam         Camera;
}

func GetEmptyWorld(fv float32) World {
    var w World;
    w.Cam.Fv = fv;
    return w;
}

func (w *World) AddModel(m *Model) {
    w.models = append(w.models, m);
}

func (w *World) Render(s *Canvas, c Color) {
    rotNew := SubVec3(Vec3{0, 0, 0}, w.Cam.Rotation);
    for i := 0 ; i<len(w.models); i++ {
        for j := 0; j<len(w.models[i].triangles); j++ {
            // calculate vertex position in world space
            // for each vertex move it to position relative to camera (from Model.Position and Cam.Position), and rotate it relative to camera
            v0 := Rotate3D(AddVec3(w.models[i].vertices[w.models[i].triangles[j].i0], SubVec3(w.models[i].Position, w.Cam.Position)), rotNew);
            v1 := Rotate3D(AddVec3(w.models[i].vertices[w.models[i].triangles[j].i1], SubVec3(w.models[i].Position, w.Cam.Position)), rotNew);
            v2 := Rotate3D(AddVec3(w.models[i].vertices[w.models[i].triangles[j].i2], SubVec3(w.models[i].Position, w.Cam.Position)), rotNew);

            drawTriangle3D(s, v0, v1, v2, w.Cam.Fv, c);
        }
    }
}

// converts 3D position to position on screen
func point3DToPoint2D(v Vec3, fv float32) Vec2 {
    return Vec2{v.X/v.Z*fv, v.Y/v.Z*fv};
}

func drawPoint3D(s *Canvas, v Vec3, fv float32, c Color) {
    s.SetPixel(point3DToPoint2D(v, fv), c);
}

// Line from Vec3 to Vec3
func drawLine3D(s *Canvas, a Vec3, b Vec3, fv float32, c Color) {
    s.DrawLine(point3DToPoint2D(a, fv), point3DToPoint2D(b, fv), c)
}

func drawTriangle3D(s *Canvas, v0 Vec3, v1 Vec3, v2 Vec3, fv float32, c Color) {
    s.DrawTriangle(point3DToPoint2D(v0, fv), point3DToPoint2D(v1, fv), point3DToPoint2D(v2, fv), c);
}

// For debug, renders only points
func drawModelV(s *Canvas, m Model, fv float32, c Color) {
    for i := 0; i<len(m.vertices); i++ {
        drawPoint3D(s, m.vertices[i], fv, c);
    }
}

// Adds x of v0 to v1, y to y and so on
func AddVec3(v0 Vec3, v1 Vec3) Vec3 {
    v0.X += v1.X;
    v0.Y += v1.Y;
    v0.Z += v1.Z;
    return v0;
}

// Subtracts x of v1 from v0 and so on
func SubVec3(v0 Vec3, v1 Vec3) Vec3 {
    v0.X -= v1.X;
    v0.Y -= v1.Y;
    v0.Z -= v1.Z;
    return v0;
}
