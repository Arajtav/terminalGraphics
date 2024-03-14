package main

import (
    "fmt"
    tg "terminalGraphics/src"
    "time"
)

func main() {
    tg.ClearTerm();
    var s tg.Canvas;
    s.Init(96, 96);

    cube1 := tg.Cube();

    cube2 := tg.Cube();
    cube2.SetScale(tg.Vec3{2.0, 2.0, 2.0});
    cube2.SetRotation(tg.Vec3{tg.DegreesToRadians(45.0), tg.DegreesToRadians(45.0), 0})

    w := tg.GetEmptyWorld(96.0);

    w.AddModel(&cube1);
    w.AddModel(&cube2);

    w.Cam.Position.Z = -8.0;

    i := 0
    ts := time.Now().UnixNano();
    for {
        s.Fill(tg.Color{255, 255, 255});
        cube1.SetRotation(tg.Vec3{tg.DegreesToRadians(float32(2*i%360)), tg.DegreesToRadians(float32(5*i%360)), 0.0});
        w.Render(&s, tg.Color{0, 0, 0});
        w.Cam.Rotation.Z += tg.DegreesToRadians(1.0);
        tg.Chome();
        fmt.Println();
        s.Print();
        i++;
        if time.Now().UnixNano()-ts > 1000000000 { break; }
    }
    fmt.Printf("%d frames in 1 second\n", i);
}
