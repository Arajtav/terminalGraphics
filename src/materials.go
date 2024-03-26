package terminalGraphics

type Material interface {
    GetColor(Vec2) Color;
}

// Material that will always return one color
type MaterialFlat struct { C Color; }
func (m *MaterialFlat) GetColor(coord Vec2) Color { return m.C; }

// Material that maps UV to red and green channels
type MaterialUV struct {}
func (m *MaterialUV) GetColor(coord Vec2) Color { return Color{uint8(coord.X*255), uint8(coord.Y*255), 0}; }
