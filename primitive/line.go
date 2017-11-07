package primitive

import (
	"fmt"

	"github.com/fogleman/gg"
)

type VerticalLine struct {
	Worker *Worker
	X1, Y1 int
	X2, Y2 int
}

func NewRandomVerticalLine(worker *Worker) *VerticalLine {
	rnd := worker.Rnd
	x1 := rnd.Intn(worker.W)
	y1 := rnd.Intn(worker.H)
	x2 := clampInt(x1+rnd.Intn(4)+1, 0, worker.W-1)
	y2 := clampInt(y1+rnd.Intn(32)+1, 0, worker.H-1)
	return &VerticalLine{worker, x1, y1, x2, y2}
}

func (r *VerticalLine) bounds() (x1, y1, x2, y2 int) {
	x1, y1 = r.X1, r.Y1
	x2, y2 = r.X2, r.Y2
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return
}

func (r *VerticalLine) Draw(dc *gg.Context, scale float64) {
	x1, y1, x2, y2 := r.bounds()
	dc.DrawRectangle(float64(x1), float64(y1), float64(x2-x1+1), float64(y2-y1+1))
	dc.Fill()
}

func (r *VerticalLine) SVG(attrs string) string {
	x1, y1, x2, y2 := r.bounds()
	w := x2 - x1 + 1
	h := y2 - y1 + 1
	return fmt.Sprintf(
		"<rect %s x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />",
		attrs, x1, y1, w, h)
}

func (r *VerticalLine) Copy() Shape {
	a := *r
	return &a
}

func (r *VerticalLine) Mutate() {
	w := r.Worker.W
	h := r.Worker.H
	rnd := r.Worker.Rnd
	o := rnd.Intn(4)+1
	switch rnd.Intn(2) {
	case 0:
		r.X1 = clampInt(r.X1+int(rnd.NormFloat64()*16), 0, w-1)
		r.Y1 = clampInt(r.Y1+int(rnd.NormFloat64()*16), 0, h-1)
		r.X2 = clampInt(r.X1+o, 0, w-1)
	case 1:
		r.X2 = clampInt(r.X2+int(rnd.NormFloat64()*16), 0, w-1)
		r.Y2 = clampInt(r.Y2+int(rnd.NormFloat64()*16), 0, h-1)
		r.X1 = clampInt(r.X2-o, 0, w-1)
	}
}

func (r *VerticalLine) Rasterize() []Scanline {
	x1, y1, x2, y2 := r.bounds()
	lines := r.Worker.Lines[:0]
	for y := y1; y <= y2; y++ {
		lines = append(lines, Scanline{y, x1, x2, 0xffff})
	}
	return lines
}


type HorizontalLine struct {
	Worker *Worker
	X1, Y1 int
	X2, Y2 int
}

func NewRandomHorizontalLine(worker *Worker) *HorizontalLine {
	rnd := worker.Rnd
	x1 := rnd.Intn(worker.W)
	y1 := rnd.Intn(worker.H)
	x2 := clampInt(x1+rnd.Intn(32)+1, 0, worker.W-1)
	y2 := clampInt(y1+rnd.Intn(4)+1, 0, worker.H-1)
	return &HorizontalLine{worker, x1, y1, x2, y2}
}

func (r *HorizontalLine) bounds() (x1, y1, x2, y2 int) {
	x1, y1 = r.X1, r.Y1
	x2, y2 = r.X2, r.Y2
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return
}

func (r *HorizontalLine) Draw(dc *gg.Context, scale float64) {
	x1, y1, x2, y2 := r.bounds()
	dc.DrawRectangle(float64(x1), float64(y1), float64(x2-x1+1), float64(y2-y1+1))
	dc.Fill()
}

func (r *HorizontalLine) SVG(attrs string) string {
	x1, y1, x2, y2 := r.bounds()
	w := x2 - x1 + 1
	h := y2 - y1 + 1
	return fmt.Sprintf(
		"<rect %s x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />",
		attrs, x1, y1, w, h)
}

func (r *HorizontalLine) Copy() Shape {
	a := *r
	return &a
}

func (r *HorizontalLine) Mutate() {
	w := r.Worker.W
	h := r.Worker.H
	rnd := r.Worker.Rnd
	o := rnd.Intn(4)+1
	switch rnd.Intn(2) {
	case 0:
		r.X1 = clampInt(r.X1+int(rnd.NormFloat64()*16), 0, w-1)
		r.Y1 = clampInt(r.Y1+int(rnd.NormFloat64()*16), 0, h-1)
		r.Y2 = clampInt(r.Y1+o, 0, h-1)
	case 1:
		r.X2 = clampInt(r.X2+int(rnd.NormFloat64()*16), 0, w-1)
		r.Y2 = clampInt(r.Y2+int(rnd.NormFloat64()*16), 0, h-1)
		r.Y1 = clampInt(r.Y2-o, 0, h-1)
	}
}

func (r *HorizontalLine) Rasterize() []Scanline {
	x1, y1, x2, y2 := r.bounds()
	lines := r.Worker.Lines[:0]
	for y := y1; y <= y2; y++ {
		lines = append(lines, Scanline{y, x1, x2, 0xffff})
	}
	return lines
}

