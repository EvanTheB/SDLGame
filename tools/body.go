package tools

import (
	"math"
)

type Body struct {
	Name     string
	Mass     float64
	Radius   float64
	Position Vector
	Velocity Vector
}

func UpdateBodiesSeconds(bodies []*Body, seconds float64) {
	timeStep := seconds / 1000
	for seconds > timeStep {
		UpdateBodies(bodies, timeStep)
		seconds -= timeStep
	}
	if seconds > 0 {
		UpdateBodies(bodies, seconds)
	}
}

func UpdateBodies(bodies []*Body, stepSize float64) {
	for _, a := range bodies {
		totalDeltaA := Vector{}
		for _, b := range bodies {
			if a == b {
				continue
			}
			dir := a.Position.To(b.Position)
			accel := 6.67E-11 * b.Mass / math.Pow(a.Position.Dist(b.Position), 3)
			dir.Mul(accel)
			totalDeltaA.Add(dir)
		}
		totalDeltaA.Mul(stepSize)
		a.Velocity.Add(totalDeltaA)
		deltaV := Vector{a.Velocity.X, a.Velocity.Y, a.Velocity.Z}
		deltaV.Mul(stepSize)
		a.Position.Add(deltaV)
	}
}
