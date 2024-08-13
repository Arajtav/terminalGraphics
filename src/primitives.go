package terminalGraphics

import (
    tc "github.com/Arajtav/terminalCanvas"
)

func Cube(mt tc.Material) Model {
    m := GetEmptyModel()

    // bottom
    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  1, -1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{-1,  1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  1, -1}, Vec2{1, 0}}, mt)

    // right
    m.AddTriangle(Vertex{Vec3{ 1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{ 1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{1, 0}}, mt)

    // top
    m.AddTriangle(Vertex{Vec3{ 1, -1, -1}, Vec2{1, 1}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1, -1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1,  1}, Vec2{1, 0}}, mt)

    // left
    m.AddTriangle(Vertex{Vec3{-1,  1, -1}, Vec2{1, 1}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1, -1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{-1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1, -1}, Vec2{1, 0}}, mt)

    // front
    m.AddTriangle(Vertex{Vec3{ 1,  1, -1}, Vec2{1, 1}}, Vertex{Vec3{-1,  1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1, -1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{-1, -1, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  1, -1}, Vec2{0, 1}}, Vertex{Vec3{ 1, -1, -1}, Vec2{1, 0}}, mt)

    // back
    m.AddTriangle(Vertex{Vec3{-1,  1,  1}, Vec2{1, 1}}, Vertex{Vec3{ 1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1,  1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{ 1, -1,  1}, Vec2{0, 0}}, Vertex{Vec3{ 1,  1,  1}, Vec2{0, 1}}, Vertex{Vec3{-1, -1,  1}, Vec2{1, 0}}, mt)

    return m
}

func Plane(mt tc.Material) Model {
    m := GetEmptyModel()

    m.AddTriangle(Vertex{Vec3{-1,  0, -1}, Vec2{0, 0}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}}, mt)
    m.AddTriangle(Vertex{Vec3{ 1,  0,  1}, Vec2{1, 1}}, Vertex{Vec3{-1,  0,  1}, Vec2{0, 1}}, Vertex{Vec3{ 1,  0, -1}, Vec2{1, 0}}, mt)

    return m
}
