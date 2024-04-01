package main

import (
    tg "terminalGraphics/src"
    tc "github.com/Arajtav/terminalCanvas"
)

func main() {
    tg.ClearTerm();
    s := *tc.NewCanvas(96, 96);
    s.Ccolor.B = 255;

    m := tg.Cube(tc.MaterialUV{});
    m.SetScale(tg.Vec3{X:2.0, Y:2.0, Z:2.0});

    w := tg.GetEmptyWorld(96.0);

    w.Cam.Position.Z = -8.0;
    w.Cam.Position.Y =  8.0;
    w.Cam.Rotation.X = tg.DegreesToRadians(45.0);

    w.AddModel(&m);

    i := float32(0.0)
    for {
        s.Clear();
        w.Render(&s);
        m.SetRotation(tg.Vec3{X:0, Y:0, Z:tg.DegreesToRadians(i)});
        tg.Chome();
        s.Print();
        i += 1.0;
    }
}
