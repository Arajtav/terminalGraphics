package main

import (
	"fmt"
	tg "terminalGraphics/src"
	"time"
)

func main() {
    var s tg.Canvas;
    s.Init(96, 96);

    fv := float32(32.0);

    c := tg.Color{0, 0, 0};

    cube := tg.Cube();
    cube.Position.Z = 3.0;
    for i := 0; true; i++ {
        s.Fill(tg.Color{255, 255, 255});
        cube.SetRotation(tg.Vec3{float32(i)/4.0, float32(i)/4.0, 0.0});
        tg.DrawModel(&s, cube, fv, c);
        fmt.Print("\x1b\x5b\x48\x1b\x5b\x32\x4a\x1b\x5b\x33\x4a");  // clearing screen
        s.Print();
        time.Sleep(125000000);
    }
}
