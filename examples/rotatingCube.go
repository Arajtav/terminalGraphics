package main

import (
    tg "terminalGraphics/src"
    tc "github.com/Arajtav/terminalCanvas"
)

var frames [3]*tc.Canvas

func swapCanvas() {
    frames[2] = frames[0]
    frames[0] = frames[1]
    frames[1] = frames[2]
}

func main() {
    const canvasSize = 96
    tg.ClearTerm()

    // 3, because frames[2] will be used for swapping
    frames[0] = tc.NewCanvas(canvasSize, canvasSize)
    frames[1] = tc.NewCanvas(canvasSize, canvasSize)
    frames[0].Ccolor.B = 255
    frames[1].Ccolor.B = 255

    m := tg.Cube(tc.MaterialUV{})
    m.SetScale(tg.Vec3{X: 2.0, Y: 2.0, Z: 2.0})

    w := tg.GetEmptyWorld(96.0)

    w.Cam.Position.Z = -8.0
    w.Cam.Position.Y = 8.0
    w.Cam.Rotation.X = tg.DegreesToRadians(45.0)

    w.AddModel(&m)

    frames[1].Print() // this is important, because later there is only PrintD
    i := float32(0.0)
    for {
        frames[0].Clear()
        w.Render(frames[0])
        m.SetRotation(tg.Vec3{X: 0, Y: 0, Z: tg.DegreesToRadians(i)})
        frames[0].PrintD(frames[1])
        i += 0.5
        swapCanvas()
    }
}
