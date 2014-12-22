import math

plans = [
	
	(	"Mercury"	,	29.16965847	,	22496.95378	,	0.2056387727	),
								
	(	"Venus"	,	54.96444137	,	8832.946404	,	0.006770594639	),
								
	(	"EM"	,	102.9858964	,	5469.570155	,	0.01670467958	),
								
	(	"Mars"	,	-73.3932514	,	2800.600038	,	0.09340585556	),
								
	(	"Jupiter"	,	-85.74426062	,	386.5074589	,	0.04836647389	),
								
	(	"Saturn"	,	-21.08297883	,	118.6629787	,	0.05378573977	),
								
	(	"Uranus"	,	96.99188547	,	303.1205591	,	0.04725088212	),
								
	(	"Neptune"	,	-86.86679117	,	-154.321488	,	0.008598093823	),
								
	(	"Pluto"	,	113.7606849	,	150.2838007	,	0.2488350108	),
]

for tup in plans:
	name,w,M,e = tup
	estar = e * math.pi / 180

	E0 = M + estar * math.sin(M)
	deltaE = 100

	while(deltaE > 0.000001):
		deltaM = M - (E0 - estar*math.sin(E0))
		deltaE = deltaM / (1 - e*math.cos(E0))
		E0 = deltaE + E0
		print deltaE, deltaM

	print name, E0

