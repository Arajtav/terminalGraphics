package main

import (
	tg "terminalGraphics/src"
)

func main() {
    tg.ClearTerm();
    var s tg.Canvas;
    s.Init(96, 96);

    fv := float32(32.0);

    c := tg.Color{0, 0, 0};

    cube := tg.Cube();
    cube.Position.Z = 3.0;
    for i := 0; true; i++ {
        s.Fill(tg.Color{255, 255, 255});
        cube.SetRotation(tg.Vec3{tg.DegreesToRadians(float32(2*i%360)), tg.DegreesToRadians(float32(5*i%360)), 0.0});
        tg.DrawModel(&s, cube, fv, c);
        tg.Chome();
        s.Print();
    }
}
