package main

import (
	tg "terminalGraphics/src"
)

func main() {
    // 3D cube
    var s tg.Canvas3D;
    s.Init(96, 96, 32.0);

    s.Fill(tg.Color{255, 255, 255})
    c := tg.Color{0, 0, 0};


    s.DrawPoint(tg.Vertex{0, 0, 0}, c);

    s.DrawLine(tg.Vertex{-1, -1, 2}, tg.Vertex{-1,  1, 2}, c);
    s.DrawLine(tg.Vertex{-1,  1, 2}, tg.Vertex{ 1,  1, 2}, c);
    s.DrawLine(tg.Vertex{ 1,  1, 2}, tg.Vertex{ 1, -1, 2}, c);
    s.DrawLine(tg.Vertex{ 1, -1, 2}, tg.Vertex{-1, -1, 2}, c);
    s.DrawLine(tg.Vertex{-1, -1, 1}, tg.Vertex{-1,  1, 1}, c);
    s.DrawLine(tg.Vertex{-1,  1, 1}, tg.Vertex{ 1,  1, 1}, c);
    s.DrawLine(tg.Vertex{ 1,  1, 1}, tg.Vertex{ 1, -1, 1}, c);
    s.DrawLine(tg.Vertex{ 1, -1, 1}, tg.Vertex{-1, -1, 1}, c);
    s.DrawLine(tg.Vertex{ 1, -1, 1}, tg.Vertex{ 1, -1, 2}, c);
    s.DrawLine(tg.Vertex{-1, -1, 1}, tg.Vertex{-1, -1, 2}, c);
    s.DrawLine(tg.Vertex{ 1,  1, 1}, tg.Vertex{ 1,  1, 2}, c);
    s.DrawLine(tg.Vertex{-1,  1, 1}, tg.Vertex{-1,  1, 2}, c);

    s.Print();
}
