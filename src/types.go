package terminalGraphics

// Vec2 with float32
type Vec2 struct {
    X float32
    Y float32
}

// Vec3 with float32
type Vec3 struct {
    X float32
    Y float32
    Z float32
}

// Vertex
type Vertex struct {
    Position Vec3
    UV       Vec2
}
