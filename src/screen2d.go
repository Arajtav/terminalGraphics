package terminalGraphics


// TODO: REWRITE THAT. IT'S TEMPORARY
func (s *Screen) DrawLine(p0 Coord2D, p1 Coord2D, c Color) {
    if p0.X >= s.sizeX || p0.Y >= s.sizeY || p1.X >= s.sizeX || p1.Y >= s.sizeY { return; }
    ll := dist2D(p0, p1);
    for t := float32(0); t<1.0; t+=(1/ll) {
        s.SetPixelUnsafe(Coord2D{uint16(interpolateU16(p0.X, p1.X, t)+0.5), uint16(interpolateU16(p0.Y, p1.Y, t)+0.5)}, c);
    }
    s.SetPixel(p1, c);
}
