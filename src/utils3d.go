package terminalGraphics

// TODO: ACTUALLY USE CAMERA
type Camera struct {
    Fv float32;
    Rot Vec3;
    Pos Vec3;
}

type World struct {
    models []*Model;
    Cam Camera;
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
    for i := 0 ; i<len(w.models); i++ {
        for j := 0; j<len(w.models[i].triangles); j++ {
            // calculate Vertex position in world space
            v0 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i0], AddVec3(w.models[i].Position, w.Cam.Pos));
            v1 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i1], AddVec3(w.models[i].Position, w.Cam.Pos));
            v2 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i2], AddVec3(w.models[i].Position, w.Cam.Pos));
            v0 = Rotate3D(v0, w.Cam.Rot);
            v1 = Rotate3D(v1, w.Cam.Rot);
            v2 = Rotate3D(v2, w.Cam.Rot);

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
