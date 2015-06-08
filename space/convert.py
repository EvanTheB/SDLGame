import re
import json
import sys

re_name = re.compile("Target body name: (.*) {source")
re_radius = re.compile(r"Mean radius.+= +([0-9]+\.[0-9]+)")
re_mass = re.compile(r"Mass.*= +([0-9]+\.[0-9]+)")


def convert(lines):
    # these are optional because bad data format
    radius = None
    mass = None
    for i in range(len(lines)):
        if re_name.match(lines[i]):
            # print re_name.match(lines[i]).group(1)
            the_name = re_name.match(lines[i]).group(1).strip()

        if re_radius.search(lines[i]):
            radius = re_radius.search(lines[i]).group(1).strip()
            # print radius

        if re_mass.search(lines[i]):
            mass = re_mass.search(lines[i]).group(1).strip()
            # print mass

        if lines[i].startswith("$$SOE"):
            # 2 dates and an empty at the end
            floats = [float(x) for x in lines[i + 1].split(',')[2:-1]]
            x, y, z, vx, vy, vz, lt, rg, rr = floats
            # print x
    planet = {}
    planet["name"] = the_name
    planet["position"] = (x, y, z)
    planet["velocity"] = (vx, vy, vz)
    planet["radius"] = radius
    planet["mass"] = mass

    print json.dumps(planet)
    return planet

planets = {}
for fname in sys.argv[1:]:
    print fname
    with open(fname) as raw:
        raw_lines = raw.readlines()
    planet = convert(raw_lines)
    planets[planet["name"]] = planet

with open("planets.json", 'w') as out:
    json.dump(planets, out, indent=4,
              separators=(',', ': '))
