package terminalGraphics

// Stores Canvas and data needed to render 3d objects
type Canvas3D struct {
    cv Canvas;
    fv float32;
}

// Initialize Canvas3D
func (s *Canvas3D) Init(x int16, y int16, fv float32) {
    s.cv.Init(x, y);
    s.fv = fv;
}

// Sets every pixel in Canvas to given color
func (s *Canvas3D) Fill(c Color) {
    s.cv.Fill(c);
}

// Sets every pixel in Canvas to #000000
func (s *Canvas3D) Clear() {
    s.cv.Clear();
}

// Prints Canvas to the terminal
func (s *Canvas3D) Print() {
    s.cv.Print();
}


func (s *Canvas3D) DrawPoint(v Vertex, c Color) {
    s.cv.SetPixel(Point3DToPoint2D(v, s.fv), c);
}

func (s *Canvas3D) DrawLine(v0 Vertex, v1 Vertex, c Color) {
    s.cv.DrawLine(Point3DToPoint2D(v0, s.fv), Point3DToPoint2D(v1, s.fv), c)
}
