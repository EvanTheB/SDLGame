import re
import json
import sys

re_name = re.compile("Target body name: (.*) {source")
re_radius = [
    (1E3, re.compile(r"radius \(km\) *= *([0-9]+\.?[0-9]+)")),
    (1E3, re.compile(r"Radius, km *= *([0-9]+\.?[0-9]+)")),
    (1E3, re.compile(r"Mean Radius \(km\) *= *([0-9]+\.?[0-9]+)")),
    (1E5 * 1E3,
     re.compile(r"Radius \(photosphere\) *= *([0-9]+\.?[0-9]+)\(10\^5\) km")),
]
re_mass = [
    (1E20, re.compile(r"Mass,? \(?10\^20 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    (1E21, re.compile(r"Mass,? \(?10\^21 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    (1E22, re.compile(r"Mass,? \(?10\^22 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    (1E23, re.compile(r"Mass,? \(?10\^23 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    (1E24, re.compile(r"Mass,? \(?10\^24 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    (1E26, re.compile(r"Mass,? \(?10\^26 kg *\)? *= +([0-9]+\.?[0-9]+)")),
    # silly special cases
    (1E22, re.compile(r"Mass Pluto \(10\^22 kg\) *= +([0-9]+\.?[0-9]+)")),
    (1E19, re.compile(r"Mass \(10\^22 g\) *= +([0-9]+\.?[0-9]+)")),
    (1E30, re.compile(r"Mass \(10\^30 kg *\) *~ +([0-9]+\.?[0-9]+)")),
]


def convert(lines):
    # these are optional because bad data format
    radius = None
    mass = None
    for i in range(len(lines)):
        if re_name.match(lines[i]):
            # print re_name.match(lines[i]).group(1)
            the_name = re_name.match(lines[i]).group(1).strip()

        for conv_factor, reg in re_radius:
            if reg.search(lines[i]):
                radius = float(
                    reg.search(lines[i]).group(1).strip()) * conv_factor
                # print radius, lines[i]

        for conv_factor, reg in re_mass:
            if reg.search(lines[i]):
                mass = float(
                    reg.search(lines[i]).group(1).strip()) * conv_factor
                # print mass, lines[i]

        if lines[i].startswith("$$SOE"):
            # 2 dates at start
            floats = [float(x) * 1000 for x in lines[i + 1].split(',')[2:8]]
            x, y, z, vx, vy, vz = floats

    if mass is None:# or radius is None:
        print "Failed", the_name, mass, radius

    planet = {}
    planet["name"] = the_name
    planet["position"] = {"X": x, "Y": y, "Z": z}
    planet["velocity"] = {"X": vx, "Y": vy, "Z": vz}
    planet["radius"] = radius
    planet["mass"] = mass

    return planet

planets = []
for fname in sys.argv[1:]:
    print fname
    with open(fname) as raw:
        raw_lines = raw.readlines()
    planet = convert(raw_lines)
    planets.append(planet)

with open("planets.json", 'w') as out:
    json.dump(planets, out, indent=4,
              separators=(',', ': '))
