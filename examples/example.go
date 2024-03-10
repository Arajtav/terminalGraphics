package main

import (
	tg "terminalGraphics/src"
)

func main() {
    var s tg.Canvas;
    s.Init(96, 96);

    fv := float32(32.0);

    s.Fill(tg.Color{255, 255, 255})
    c := tg.Color{0, 0, 0};

    tg.DrawPoint3D(&s, tg.Vertex{0, 0, 0}, fv, c);

    i := float32(1.0);
    // back
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1, i+1}, tg.Vertex{-1,  1, i+1}, tg.Vertex{ 1, -1, i+1}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{ 1,  1, i+1}, tg.Vertex{-1,  1, i+1}, tg.Vertex{ 1, -1, i+1}, fv, c);
    // front
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{ 1,  1,   i}, tg.Vertex{-1,  1,   i}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{ 1,  1,   i}, tg.Vertex{ 1, -1,   i}, fv, c);
    // left
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{-1, -1, i+1}, tg.Vertex{-1,  1, i+1}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{-1,  1,   i}, tg.Vertex{-1,  1, i+1}, fv, c);
    // right
    tg.DrawTriangle3D(&s, tg.Vertex{ 1, -1,   i}, tg.Vertex{ 1,  1,   i}, tg.Vertex{ 1, -1, i+1}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{ 1,  1, i+1}, tg.Vertex{ 1,  1,   i}, tg.Vertex{ 1, -1, i+1}, fv, c);
    // bottom
    tg.DrawTriangle3D(&s, tg.Vertex{ 1,  1,   i}, tg.Vertex{ 1,  1, i+1}, tg.Vertex{-1,  1, i+1}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{-1,  1,   i}, tg.Vertex{ 1,  1,   i}, tg.Vertex{-1,  1, i+1}, fv, c);
    // top
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{ 1, -1,   i}, tg.Vertex{ 1, -1, i+1}, fv, c);
    tg.DrawTriangle3D(&s, tg.Vertex{-1, -1,   i}, tg.Vertex{-1, -1, i+1}, tg.Vertex{ 1, -1, i+1}, fv, c);
    
    s.Print();
}
