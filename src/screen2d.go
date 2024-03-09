package terminalGraphics


// TODO: REWRITE THAT. ITS TEMPORARY
func (s *Screen) DrawLine(x0 uint16, y0 uint16, x1 uint16, y1 uint16, c Color) {
    if x0 >= s.sizeX || y0 >= s.sizeY || x1 >= s.sizeX || y1 >= s.sizeY { return; }
    ll := length(x0, y0, x1, y1);
    for t := float32(0); t<1.0; t+=(1/ll) {
        s.SetPixelUnsafe(uint16(interpolateU16(x0, x1, t)), uint16(interpolateU16(y0, y1, t)), c);
    }
    s.SetPixel(x1, y1, c);
}
