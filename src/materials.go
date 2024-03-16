package terminalGraphics

type Material interface {
    GetColor(Vec2) Color;
}

type MaterialFlat struct {
    C Color;
}

func (m *MaterialFlat) GetColor(coord Vec2) Color {
    return m.C;
}
