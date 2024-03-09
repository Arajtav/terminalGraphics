package terminalGraphics

import "fmt"

type Screen struct {
    sizeX uint16;
    sizeY uint16;
    data [][]Color;
}

// Initialize Screen with given size
func (s *Screen) Init(x uint16, y uint16) {
    s.sizeX = x;
    s.sizeY = y;
    s.data = make([][]Color, y);
    for i := uint16(0); i<y; i++ {
        s.data[i] = make([]Color, x);
    }
}

// Sets every pixel in Screen to give color
func (s *Screen) Fill(c Color) {
    for i := uint16(0); i<s.sizeY; i++ {
        for j := uint16(0); j<s.sizeX; j++ {
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
    for i := uint16(0); i<s.sizeY; i+=2 {
        for j := uint16(0); j<s.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", s.data[i][j].R, s.data[i][j].G, s.data[i][j].B, s.data[i+1][j].R, s.data[i+1][j].G, s.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}

// Sets value of a pixel at given position to c
func (s *Screen) SetPixel(x uint16, y uint16, c Color) {
    if x >= s.sizeX || y >= s.sizeY { return; }
    s.data[y][x] = c;
}

// Same as SetPixel, but it doesn't check if pixel position is correct
func (s *Screen) SetPixelUnsafe(x uint16, y uint16, c Color) {
    s.data[y][x] = c;
}
