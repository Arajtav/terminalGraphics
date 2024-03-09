package terminalGraphics

import "fmt"

// Stores Pixels and own size
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
    for i := int16(0); i<s.sizeY; i+=2 {
        for j := int16(0); j<s.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", s.data[i][j].R, s.data[i][j].G, s.data[i][j].B, s.data[i+1][j].R, s.data[i+1][j].G, s.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}

// Sets value of a pixel at given position to c
func (s *Canvas) SetPixel(p Coord2D, c Color) {
    if p.X >= s.sizeX/2 || p.Y >= s.sizeY/2 { return; }
    s.data[int32(p.Y)+int32(s.sizeY/2)][int32(p.X)+int32(s.sizeX/2)] = c;
}

// Same as SetPixel, but it doesn't check if pixel position is correct
func (s *Canvas) SetPixelUnsafe(p Coord2D, c Color) {
    s.data[int32(p.Y)+int32(s.sizeY/2)][int32(p.X)+int32(s.sizeX/2)] = c;
}

// TODO: REWRITE THAT. IT'S TEMPORARY
func (s *Canvas) DrawLine(p0 Coord2D, p1 Coord2D, c Color) {
    if p0.X >= s.sizeX/2 || p0.Y >= s.sizeY/2 || p1.X >= s.sizeX/2 || p1.Y >= s.sizeY/2 { return; }
    if p0.X < -s.sizeX/2 || p0.Y < -s.sizeY/2 || p1.X < -s.sizeX/2 || p1.Y < -s.sizeY/2 { return; }
    ll := dist2D(p0, p1);
    for t := float32(0); t<1.0; t+=(1/ll) {
        s.SetPixelUnsafe(Coord2D{int16(interpolate16(p0.X, p1.X, t)), int16(interpolate16(p0.Y, p1.Y, t))}, c);
    }
}
