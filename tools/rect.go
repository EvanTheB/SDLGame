package tools

import ()

type ViewBox struct {
	Vector
	W, H float64
}

func (this *ViewBox) Zoom(zoom float64) {
	this.ZoomH(zoom)
	this.ZoomW(zoom)
}

func (this *ViewBox) ZoomW(zoom float64) {
	this.X = this.X + 0.5*this.W*(1-1/zoom)

	this.W = this.W / zoom

}

func (this *ViewBox) ZoomH(zoom float64) {

	this.Y = this.Y + 0.5*this.H*(1-1/zoom)

	this.H = this.H / zoom
}

type vecComp func(Vector, Vector) bool

func get(points []Vector, op vecComp) Vector {
	best := points[0]
	for _, p := range points {
		if !op(best, p) {
			best = p
		}
	}
	return best
}

func maxX(lhs Vector, rhs Vector) bool {
	return lhs.X > rhs.X
}

func minX(lhs Vector, rhs Vector) bool {
	return lhs.X < rhs.X
}

func maxY(lhs Vector, rhs Vector) bool {
	return lhs.Y > rhs.Y
}

func minY(lhs Vector, rhs Vector) bool {
	return lhs.Y < rhs.Y
}

func GetAutoView(points []Vector, screenRatio float64) ViewBox {
	right := get(points, maxX).X
	left := get(points, minX).X
	top := get(points, maxY).Y
	bot := get(points, minY).Y
	ret := ViewBox{Vector{left, bot}, right - left, top - bot}
	// scale to correct ratio
	if ret.H*screenRatio > ret.W {
		ret.ZoomW(ret.W / (ret.H * screenRatio))
	} else {
		ret.ZoomH(ret.H * screenRatio / ret.W)
	}
	//zoom out a bit for context.
	ret.Zoom(0.5)
	return ret
}
