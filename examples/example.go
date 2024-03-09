package main

import (
	tg "terminalGraphics/src"
)

func main() {
    // 3D cube
    var s tg.Screen;
    s.Init(96, 96);

    s.Fill(tg.Color{255, 255, 255})

    fv := float32(32.0);
    c := tg.Color{0, 0, 0};

    s.SetPixel(tg.Point3DToPoint2D(tg.Vertex{0, 0, 0}, fv), c);

    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1, -1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{-1, -1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1, -1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  1}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{-1, -1,  1}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1,  1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{-1,  1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1,  1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  1}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{-1,  1,  1}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  1}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{ 1,  1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{ 1, -1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1,  1,  2}, fv), tg.Point3DToPoint2D(tg.Vertex{-1, -1,  2}, fv), c);
    s.DrawLine(tg.Point3DToPoint2D(tg.Vertex{-1,  1,  1}, fv), tg.Point3DToPoint2D(tg.Vertex{-1, -1,  1}, fv), c);

    s.Print();
}
