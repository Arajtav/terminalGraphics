package main

import (
	"fmt"
	tg "terminalGraphics/src"
	"time"

	tc "github.com/Arajtav/terminalCanvas"
)

var frames [3]*tc.Canvas
var models [5]tg.Model

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

    for i := 0; i<len(models); i++ {
        models[i] = tg.Cube(tc.MaterialUV{})
        models[i].SetScale(tg.Vec3{X: 1.0, Y: 1.0, Z: 1.0})
    }
    models[0].SetScale(tg.Vec3{X: 2.0, Y: 2.0, Z: 2.0})
    models[1].Position.X =  5
    models[2].Position.X = -5
    models[3].Position.Y =  5
    models[4].Position.Y = -5

    w := tg.GetEmptyWorld(canvasSize)

    w.Cam.Position.Z = -8.0
    w.Cam.Position.Y = 8.0
    w.Cam.Rotation.X = tg.DegreesToRadians(45.0)

    for i := 0; i<len(models); i++ { w.AddModel(&models[i]) }

    frames[1].Print() // this is important, because later there is only PrintD
    c := float32(0.0)
    for {
        time_start := time.Now().Nanosecond()
        tris := w.Render(frames[0])
        _ = tris
        time_render := time.Now().Nanosecond()
        frames[0].PrintD(frames[1])
        time_print := time.Now().Nanosecond()
        for i := 0; i<len(models); i++ { models[i].SetRotation(tg.Vec3{Z: tg.DegreesToRadians(c)}) }
        time_update := time.Now().Nanosecond()
        fmt.Printf("render time: %10dns    print time: %10dns\n", time_render - time_start, time_print - time_render)
        fmt.Printf("update time: %10dns    total time: %10dns\n", time_update - time_print, time.Now().Nanosecond() - time_start)
        fmt.Printf("triangles:   %10d\n", tris)
        c += 0.5
        swapCanvas()
    }
}
