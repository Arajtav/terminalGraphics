package terminalGraphics

import (
	"math"

	tc "github.com/Arajtav/terminalCanvas"
)

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

func (w *World) AddModel(m *Model) { w.models = append(w.models, m); }

func (w *World) Render(s *tc.Canvas) {
    rotNew := SubVec3(Vec3{0, 0, 0}, w.Cam.Rotation);

    cosa := float32(math.Cos(float64(rotNew.Z)));
    sina := float32(math.Sin(float64(rotNew.Z)));
    cosb := float32(math.Cos(float64(rotNew.Y)));
    sinb := float32(math.Sin(float64(rotNew.Y)));
    cosc := float32(math.Cos(float64(rotNew.X)));
    sinc := float32(math.Sin(float64(rotNew.X)));

    Axx := cosa*cosb;
    Axy := cosa*sinb*sinc - sina*cosc;
    Axz := cosa*sinb*cosc + sina*sinc;
    Ayx := sina*cosb;
    Ayy := sina*sinb*sinc + cosa*cosc;
    Ayz := sina*sinb*cosc - cosa*sinc;
    Azx := -sinb;
    Azy := cosb*sinc;
    Azz := cosb*cosc;

    for i := 0 ; i<len(w.models); i++ {
        for j := 0; j<len(w.models[i].triangles); j++ {
            // calculate vertex position in world space
            // for each vertex move it to position relative to camera (from Model.Position and Cam.Position), and rotate it relative to camera
            // I know this code looks bad, but it's just Rotate3D without need to recalculate all stuff.
            v0 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i0].Position, SubVec3(w.models[i].Position, w.Cam.Position));
            v0u := w.models[i].vertices[w.models[i].triangles[j].i0].UV;
            vcopy := v0;
            v0.X = Axx*vcopy.X + Axy*vcopy.Y + Axz*vcopy.Z;
            v0.Y = Ayx*vcopy.X + Ayy*vcopy.Y + Ayz*vcopy.Z;
            v0.Z = Azx*vcopy.X + Azy*vcopy.Y + Azz*vcopy.Z;
            v1 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i1].Position, SubVec3(w.models[i].Position, w.Cam.Position));
            v1u := w.models[i].vertices[w.models[i].triangles[j].i1].UV;
            vcopy = v1;
            v1.X = Axx*vcopy.X + Axy*vcopy.Y + Axz*vcopy.Z;
            v1.Y = Ayx*vcopy.X + Ayy*vcopy.Y + Ayz*vcopy.Z;
            v1.Z = Azx*vcopy.X + Azy*vcopy.Y + Azz*vcopy.Z;
            v2 := AddVec3(w.models[i].vertices[w.models[i].triangles[j].i2].Position, SubVec3(w.models[i].Position, w.Cam.Position));
            v2u := w.models[i].vertices[w.models[i].triangles[j].i2].UV;
            vcopy = v2;
            v2.X = Axx*vcopy.X + Axy*vcopy.Y + Axz*vcopy.Z;
            v2.Y = Ayx*vcopy.X + Ayy*vcopy.Y + Ayz*vcopy.Z;
            v2.Z = Azx*vcopy.X + Azy*vcopy.Y + Azz*vcopy.Z;

            drawTriangle3D(s, Vertex{v0, v0u}, Vertex{v1, v1u}, Vertex{v2, v2u}, w.Cam.Fv, w.Cam.Position, tc.MaterialUV{});
        }
    }
}

// converts 3D position to position on screen
func point3DToPoint2D(v Vec3, fv float32) Vec2 {
    return Vec2{v.X/v.Z*fv, v.Y/v.Z*fv};
}

func Vec2ToCvI16Vec2(v Vec2) tc.I16Vec2 {
    return tc.I16Vec2{X: int16(roundF32ToI32(v.X)), Y: int16(roundF32ToI32(v.Y))};
}

func dist3D(a Vec3, b Vec3) float32 {
    return float32(math.Sqrt(math.Pow(float64(a.X)-float64(b.X), 2.0)) + math.Pow(float64(a.Y)-float64(b.Y), 2.0) + math.Pow(float64(a.Z)-float64(b.Z), 2.0));
}

func drawTriangle3D(s *tc.Canvas, v0 Vertex, v1 Vertex, v2 Vertex, fv float32, camPos Vec3, f tc.Material) {
    s.DrawTriangleFC(   tc.I16Frag{Pos: Vec2ToCvI16Vec2(point3DToPoint2D(v0.Position, fv)), UV: tc.F32Vec2{X: v0.UV.X, Y: v0.UV.Y}, Z: dist3D(camPos, v0.Position)},
                        tc.I16Frag{Pos: Vec2ToCvI16Vec2(point3DToPoint2D(v1.Position, fv)), UV: tc.F32Vec2{X: v1.UV.X, Y: v1.UV.Y}, Z: dist3D(camPos, v1.Position)},
                        tc.I16Frag{Pos: Vec2ToCvI16Vec2(point3DToPoint2D(v2.Position, fv)), UV: tc.F32Vec2{X: v2.UV.X, Y: v2.UV.Y}, Z: dist3D(camPos, v2.Position)},
                        f);
}
