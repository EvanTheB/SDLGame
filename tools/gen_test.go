package tools

import (
	"testing"
)

func TestPlanets(t *testing.T) {
	data := GetJPLStringStored()
	planets := GetPlanets(data, DAY1, DAY2)
	t.Error(planets, data)
}

func TestDegMod(t *testing.T) {
	if 30 != degMod(30) {
		t.Error("30")
	}
	if -30 != degMod(330) {
		t.Error("330", degMod(330))
	}
	if -30 != degMod(-30) {
		t.Error("-30")
	}
	if 30 != degMod(-330) {
		t.Error("-330", degMod(-330))
	}
	if 30 != degMod(390) {
		t.Error("390", degMod(390))
	}
	if -30 != degMod(-390) {
		t.Error("-390", degMod(-390))
	}
}
