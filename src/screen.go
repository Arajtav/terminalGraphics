package terminalGraphics

import "fmt"

type Screen struct {
    sizeX int16;
    sizeY int16;
    data [][]Color;
}

// Initialize Screen with given size
func (s *Screen) Init(x int16, y int16) {
    if x<0 || y<0 {
        panic("Screen size cannot be negative");
    }

    s.sizeX = x;
    s.sizeY = y;
    s.data = make([][]Color, y);
    for i := int16(0); i<y; i++ {
        s.data[i] = make([]Color, x);
    }
}

// Sets every pixel in Screen to give color
func (s *Screen) Fill(c Color) {
    for i := int16(0); i<s.sizeY; i++ {
        for j := int16(0); j<s.sizeX; j++ {
            s.data[i][j] = c;
        }
    }
}

// Fills Screen with #000000
func (s *Screen) Clear() {
    s.Fill(Color{0, 0, 0});
}

// Prints Screen to the terminal
func (s *Screen) Print() {
    for i := int16(0); i<s.sizeY; i+=2 {
        for j := int16(0); j<s.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", s.data[i][j].R, s.data[i][j].G, s.data[i][j].B, s.data[i+1][j].R, s.data[i+1][j].G, s.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}

// Sets value of a pixel at given position to c
func (s *Screen) SetPixel(p Coord2D, c Color) {
    if p.X >= s.sizeX/2 || p.Y >= s.sizeY/2 { return; }
    s.data[int32(p.Y)+int32(s.sizeY/2)][int32(p.X)+int32(s.sizeX/2)] = c;
}

// Same as SetPixel, but it doesn't check if pixel position is correct
func (s *Screen) SetPixelUnsafe(p Coord2D, c Color) {
    s.data[int32(p.Y)+int32(s.sizeY/2)][int32(p.X)+int32(s.sizeX/2)] = c;
}
