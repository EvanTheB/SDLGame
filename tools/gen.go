package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

const JPL_ADDR = "http://ssd.jpl.nasa.gov/txt/p_elem_t1.txt"
const JPL_FILE = "p_elem_t1.txt"
const DAY1 = (2456992.5 - 2451545.0) / 36525.0
const DAY2 = (2456993.5 - 2451545.0) / 36525.0
const DAY1_eps = (2456992.500011574 - 2451545.0) / 36525.0

type JPLPlanet struct {
	Name     string
	Position Vector
}

func degMod(in float64) float64 {
	return anyMod(in, -180, 180)
}

func radMod(in float64) float64 {
	return anyMod(in, -1*math.Pi, math.Pi)
}

func anyMod(in float64, low float64, high float64) float64 {
	width := high - low
	offsetValue := in - low

	return (offsetValue - (math.Floor(offsetValue/width) * width)) + low
}

func GetJPLString() []string {
	resp, err := http.Get(JPL_ADDR)
	if err != nil {
		fmt.Println("failed to get from JPL")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//convert to string with fix for earth
	pageString := strings.Replace(string(body), "EM Bary", "Earth", -1)
	n := strings.Index(pageString, "Mercury")
	lines := strings.Split(pageString[n:], "\n")
	return lines[:len(lines)-1]
}

func GetJPLStringStored() []string {
	file, err := os.Open(JPL_FILE)
	if err != nil {
		fmt.Println("failed to get file JPL")
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//convert to string with fix for earth
	pageString := strings.Replace(string(body), "EM Bary", "Earth", -1)
	n := strings.Index(pageString, "Mercury")
	lines := strings.Split(pageString[n:], "\n")
	return lines[:len(lines)-1]
}

func GetJPLXYZ(jplStrings []string, T_Now float64) []JPLPlanet {
	if len(jplStrings)%2 != 0 {
		log.Fatal("odd number of JPL Planet lines")
	}
	retPlanets := make([]JPLPlanet, 0, len(jplStrings)/2)
	for i := 0; i < len(jplStrings); i += 2 {
		var name string
		var (
			a      float64
			e      float64
			I      float64
			L      float64
			peri   float64
			node   float64
			a_d    float64
			e_d    float64
			I_d    float64
			L_d    float64
			peri_d float64
			node_d float64
		)
		readNum, _ := fmt.Sscanf(jplStrings[i], "%s %f %f %f %f %f %f", &name, &a, &e, &I, &L, &peri, &node)
		readNum2, _ := fmt.Sscanf(jplStrings[i+1], "%f %f %f %f %f %f", &a_d, &e_d, &I_d, &L_d, &peri_d, &node_d)
		if readNum != 7 || readNum2 != 6 {
			panic("Help!")
		}
		a = a + a_d*T_Now
		e = e + e_d*T_Now
		I = (I + I_d*T_Now) * math.Pi / 180
		L = (L + L_d*T_Now) * math.Pi / 180
		peri = (peri + peri_d*T_Now) * math.Pi / 180
		node = (node + node_d*T_Now) * math.Pi / 180
		peri2 := peri - node
		M := radMod(L - peri)
		E0 := M + e*math.Sin(M)
		E_delta := 1.0
		for E_delta > 0.00000001*math.Pi {
			M_delta := M - (E0 - e*math.Sin(E0))
			E_delta = M_delta / (1 - e*math.Cos(E0))
			E0 = E0 + E_delta
		}
		x_dash := a * (math.Cos(E0) - e)
		y_dash := a * math.Sqrt(1-e*e) * math.Sin(E0)
		// z_dash := 0

		eps := 23.43928 * math.Pi / 180

		x_ecl := (math.Cos(peri2)*math.Cos(node)-math.Sin(peri2)*math.Sin(node)*math.Cos(I))*x_dash + (-math.Sin(peri2)*math.Cos(node)-math.Cos(peri2)*math.Sin(node)*math.Cos(I))*y_dash

		y_ecl := (math.Cos(peri2)*math.Sin(node)+math.Sin(peri2)*math.Cos(node)*math.Cos(I))*x_dash + (-math.Sin(peri2)*math.Sin(node)+math.Cos(peri2)*math.Cos(node)*math.Cos(I))*y_dash

		z_ecl := (math.Sin(peri2)*math.Sin(I))*x_dash + (math.Cos(peri2)*math.Sin(I))*y_dash

		x_eq := x_ecl
		y_eq := math.Cos(eps)*y_ecl - math.Sin(eps)*z_ecl
		z_eq := math.Sin(eps)*y_ecl + math.Cos(eps)*z_ecl

		// fmt.Println(x_dash, y_dash, z_dash)
		// fmt.Println(x_ecl, y_ecl, z_ecl)
		// fmt.Println(x_eq, y_eq, z_eq)

		retPlanets = append(retPlanets, JPLPlanet{name, Vector{x_eq, y_eq, z_eq}})
	}
	return retPlanets
}

func JPLToBody(jplPlanets []JPLPlanet) []Body {
	// P and V are not correct...
	planets := []Body{
		Body{"Sun", 1.988544E30, 6.963E5 * 1E3, Vector{0, 0, 0}, Vector{0, 0, 0}},
		Body{"Mercury", 3.302E23, 2440 * 1E3, Vector{1.5E11, 0, 0}, Vector{0, 2.9E4, 0}},
		Body{"Venus", 48.685E23, 6051.8 * 1E3, Vector{1.5E11, 0, 0}, Vector{0, 2.9E4, 0}},
		Body{"Earth", 5.97219E24, 6371 * 1E3, Vector{1.5E11, 0, 0}, Vector{0, 2.9E4, 0}},
		Body{"Jupiter", 1898.13E24, 71492 * 1E3, Vector{2.25e11, 0, 0}, Vector{0, 2.41e4, 0}},
		Body{"Saturn", 5.68319E26, 60268 * 1E3, Vector{2.25e11, 0, 0}, Vector{0, 2.41e4, 0}},
		Body{"Uranus", 86.8103E24, 25559 * 1E3, Vector{2.25e11, 0, 0}, Vector{0, 2.41e4, 0}},
		Body{"Neptune", 102.41E24, 24766 * 1E3, Vector{2.25e11, 0, 0}, Vector{0, 2.41e4, 0}},
		Body{"Pluto", 1.307E22, 1195 * 1E3, Vector{2.25e11, 0, 0}, Vector{0, 2.41e4, 0}},
	}
	for i := range jplPlanets {
		planets[i+1].Position = jplPlanets.Position
	}
	return planets
}
