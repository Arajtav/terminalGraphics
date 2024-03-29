package main

import (
    tg "terminalGraphics/src"
)

func main() {
    tg.ClearTerm();
    var s tg.Canvas;
    s.Init(96, 96);

    cube := tg.Cube();
    cube.SetScale(tg.Vec3{X:2.0, Y:2.0, Z:2.0});

    mflat := tg.MaterialFlat{C: tg.Color{R: 0, G: 0, B: 0}};

    w := tg.GetEmptyWorld(96.0);

    w.AddMaterial(&mflat);

    w.AddModel(&cube);

    w.Cam.Position.Z = -8.0;
    w.Cam.Position.Y =  8.0;
    w.Cam.Rotation.X = tg.DegreesToRadians(45.0);

    i := float32(0.0)
    for {
        s.Clear();
        w.Render(&s, tg.Color{R:0, G:0, B:0});
        cube.SetRotation(tg.Vec3{X:0, Y:0, Z:tg.DegreesToRadians(i)})
        tg.Chome();
        s.Print();
        i += 1.0;
    }
}
