package tools

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := []Vector{Vector{1, 2}, Vector{3, 4}, Vector{0, 5}}
	res := get(a, maxX)
	if res.X != 3 {
		t.Error("3, got ", res)
	}
	res = get(a, minX)
	if res.X != 0 {
		t.Error("0, got ", res)
	}
	res = get(a, maxY)
	if res.Y != 5 {
		t.Error("5, got ", res)
	}
	res = get(a, minY)
	if res.Y != 2 {
		t.Error("2, got ", res)
	}
}

func TestAutoZoom(t *testing.T) {
	a := []Vector{Vector{0, 0}, Vector{1, 1}, Vector{0, 3}}
	zoom := GetAutoView(a, 1)
	if zoom.H != zoom.W {
		t.Error("Expected same W and H, got ", zoom)
	}
}
