package main

import (
	tg "terminalGraphics/src"
)

func main() {
    // I have no idea, random lines.
    var s tg.Screen;
    s.Init(64, 64);
    s.DrawLine(4, 4, 59, 59, tg.Color{0, 127, 255});
    s.DrawLine(4, 59, 59, 4, tg.Color{0, 255, 127});
    s.Print();
}
