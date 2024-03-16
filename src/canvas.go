package terminalGraphics

import (
	"fmt"
	"sort"
)

// Canvas, stores pixels and own size
type Canvas struct {
    sizeX   int16;
    sizeY   int16;
    data    [][]Color;
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
func (s *Canvas) Clear() { s.Fill(Color{0, 0, 0}); }

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

// Vec2 but with int32 for DrawLine
type i32vec2 struct {
    X int32;
    Y int32;
};

func pxPosToU(p Vec2, sizex int16, sizey int16) i32vec2 {
    return i32vec2{roundF32ToI32(p.X)+int32(sizex/2), roundF32ToI32(p.Y)+int32(sizey/2)};
}

func (s *Canvas) DrawLine(ra Vec2, rb Vec2, c Color) {
    s.drawLinei32(pxPosToU(ra, s.sizeX, s.sizeY), pxPosToU(rb, s.sizeX, s.sizeY), c);
}

// Draws line from point ra to point rb
// based on Bresenham's line algorithm
func (s *Canvas) drawLinei32(a i32vec2, b i32vec2, c Color) {
    // Not using getLine form utils because it would slow things down a lot
    d := i32vec2{b.X - a.X, b.Y - a.Y};
    g := i32vec2{1, 1};

    if d.X < 0 {
        d.X = -d.X;
        g.X = -1;
    }
    if d.Y < 0 {
        d.Y = -d.Y;
        g.Y = -1;
    }

    err := d.X - d.Y;
    cp := a;

    for {
        // TODO, CLIP LINE TO CANVAS EDGES FIRST, INSTEAD OF THAT CHECKS IN LOOP
        if cp.X >= int32(s.sizeX) || cp.Y >= int32(s.sizeY) || cp.X < 0 || cp.Y < 0 { break; }
        s.data[cp.Y][cp.X] = c;
        if cp.X == b.X && cp.Y == b.Y { break; }

        e2 := 2 * err;
        if e2 > -d.Y {
            err -= d.Y;
            cp.X += g.X;
        }
        if e2 < d.X {
            err += d.X;
            cp.Y += g.Y;
        }
    }
}

// Filled triangle (scan line algorithm)
func (s *Canvas) DrawTriangleFull(p0 Vec2, p1 Vec2, p2 Vec2, c Color) {
    var points []i32vec2;
    points = append(points, getLine(pxPosToU(p0, s.sizeX, s.sizeY), pxPosToU(p1, s.sizeX, s.sizeY))...);
    points = append(points, getLine(pxPosToU(p1, s.sizeX, s.sizeY), pxPosToU(p2, s.sizeX, s.sizeY))...);
    points = append(points, getLine(pxPosToU(p2, s.sizeX, s.sizeY), pxPosToU(p0, s.sizeX, s.sizeY))...);

    sort.Slice(points, func(i int, j int) bool {
        return points[i].Y < points[j].Y || (points[i].Y == points[j].Y && points[i].X <= points[j].X);
    });

    i := 0;
    for i<len(points) {
        startp := points[i];
        j := i;
        for points[j].Y == points[i].Y {    // finding last position with same y
            j++;
            if j >= len(points) { break; }
        }
        endp := points[j-1];
        s.drawLinei32(startp, endp, c);
        i = j;
    }
}
