package terminalGraphics

import "fmt"

// Canvas, stores pixels and own size
type Canvas struct {
    sizeX int16;
    sizeY int16;
    data [][]Color;
}

// Initialize Canvas with given size
func (s *Canvas) Init(x int16, y int16) {
    if x<0 || y<0 {
        panic("Canvas size cannot be negative");
    }

    s.sizeX = x;
    s.sizeY = y;
    s.data = make([][]Color, y);
    for i := int16(0); i<y; i++ {
        s.data[i] = make([]Color, x);
    }
}

// Sets every pixel in Canvas to given color
func (s *Canvas) Fill(c Color) {
    for i := int16(0); i<s.sizeY; i++ {
        for j := int16(0); j<s.sizeX; j++ {
            s.data[i][j] = c;
        }
    }
}

// Sets every pixel in Canvas to #000000
func (s *Canvas) Clear() {
    s.Fill(Color{0, 0, 0});
}

// Prints Canvas to the terminal
func (s *Canvas) Print() {
    for i := int16(0); i<s.sizeY-1; i+=2 {
        for j := int16(0); j<s.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", s.data[i][j].R, s.data[i][j].G, s.data[i][j].B, s.data[i+1][j].R, s.data[i+1][j].G, s.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}

// Sets value of a pixel at given position to c
func (s *Canvas) SetPixel(p Vec2, c Color) {
    x := roundF32ToI32(p.X);
    y := roundF32ToI32(p.Y);
    if x >=  int32(s.sizeX/2)   || y >=  int32(s.sizeY/2)   { return; }
    if x <= -int32(s.sizeX/2)-1 || y <= -int32(s.sizeY/2)-1 { return; }
    s.data[y+int32(s.sizeY/2)][x+int32(s.sizeX/2)] = c;
}

// Same as SetPixel(p, c), but it doesn't check if pixel position is correct
func (s *Canvas) SetPixelUnsafe(p Vec2, c Color) {
    s.data[roundF32ToI32(p.Y)+int32(s.sizeY/2)][roundF32ToI32(p.X)+int32(s.sizeX/2)] = c;
}

// Draws line from point a to point b
func (s *Canvas) DrawLine(a Vec2, b Vec2, c Color) {
    // TODO: REAL LINE CLIPPING INSTEAD RELYING ON SetPixel SAFETY
    ll := dist2D(a, b);
    for t := float32(0); t<=1.0; t+=(1/ll) {
        s.SetPixel(Vec2{Interpolate(a.X, b.X, t), Interpolate(a.Y, b.Y, t)}, c);
    }
    s.SetPixel(b, c);
}

// Outline of triangle
func (s *Canvas) DrawTriangle(p0 Vec2, p1 Vec2, p2 Vec2, c Color) {
    s.DrawLine(p0, p1, c);
    s.DrawLine(p1, p2, c);
    s.DrawLine(p2, p0, c);
}

// Filled triangle
func (s *Canvas) DrawTriangleFull(p0 Vec2, p1 Vec2, p2 Vec2, c Color) {
    // TODO, THIS IS VERY VERY BAD IMPLEMENTATION REWRITE IT
    ll := (dist2D(p0, p2)+dist2D(p0, p1)/2.0)+dist2D(p1, p2);   // I`m not sure
    for i := float32(0); i<=1.0; i+=(1/ll) {
        s.DrawLine(p0, InterpolateVec2(p1, p2, i), c);
    }
    s.DrawTriangle(p0, p1, p2, c);  // this shouldn't be even needed
}
