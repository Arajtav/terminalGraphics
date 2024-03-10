package terminalGraphics

// Canvas related
type Color struct {
    R uint8;
    G uint8;
    B uint8;
}

type Coord2D struct {
    X int16;
    Y int16;
}

// 3D
type Vertex struct {
    X float32;
    Y float32;
    Z float32;
}
