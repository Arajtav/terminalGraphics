package main

import (
	tg "terminalGraphics/src"
)

func main() {
    var s tg.Canvas;
    s.Init(96, 96);

    fv := float32(32.0);

    s.Fill(tg.Color{255, 255, 255});

    c := tg.Color{0, 0, 0};

    cube := tg.Cube(1.0);
    cube.Move(tg.Vec3{0.0, 0.0, -2.0});
    tg.DrawModel(&s, cube, fv, c);

    s.Print();
}
