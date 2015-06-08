import telnetlib

def jpl_it(planet):
    print tel.read_eager()

    print tel.read_until("Horizons>", 10)
    tel.write(planet + "\n")

    print tel.read_until("Revised", 10)
    save_general = tel.read_until("***", 10)
    tel.read_until("Select ...", 10)
    tel.write("e\n")

    print tel.read_until("Observe", 10)
    tel.write("v\n")

    index, match, text = tel.expect(["Coordinate", "Use Previous"], 10)
    print text
    if index == 0:
        tel.write("@0\n")
    else:
        tel.write("y\n")

    print tel.read_until("Reference", 10)
    tel.write("eclip\n")

    print tel.read_until("Starting", 10)
    tel.write("2014-12-05 00:00\n")

    print tel.read_until("Ending", 10)
    tel.write("2014-12-06 00:00\n")

    print tel.read_until("Output", 10)
    tel.write("1d\n")

    print tel.read_until("Accept", 10)
    tel.write("n\n")

    print tel.read_until("Output", 10)
    tel.write("\n")
    print tel.read_until("Corrections", 10)
    tel.write("\n")
    print tel.read_until("Output units", 10)
    tel.write("1\n")
    print tel.read_until("Spreadsheet", 10)
    tel.write("YES\n")
    print tel.read_until("Label", 10)
    tel.write("\n")
    print tel.read_until("Select", 10)
    tel.write("\n")

    tel.read_until("Ephemeris", 10)
    save_eph = tel.read_until("Jon.Giorgini@jpl.nasa.gov", 10)
    with open(planet, 'w') as out:
        out.write(save_general)
        out.write(save_eph)
    tel.write("N\n")

tel = telnetlib.Telnet("ssd.jpl.nasa.gov", 6775)

# biggest things in the system
# http://en.wikipedia.org/wiki/List_of_Solar_System_objects_by_size
# jpl_it("10") # sun
# jpl_it("199") # mercury
# jpl_it("299") # venus
# jpl_it("399") # earth
# jpl_it("499") # mars
# jpl_it("599") # jupiter
# jpl_it("699") # saturn
# jpl_it("799") # uranus
# jpl_it("899") # neptune
# jpl_it("999") # pluto

# jpl_it("301") # moon
# jpl_it("401") # Phobos
# jpl_it("402") # Deimos
# jpl_it("901") # charon
# jpl_it("902") # Nix
# jpl_it("903") # Hydra
# jpl_it("904") # Kerberos
# jpl_it("905") # Styx
# jpl_it("503") # ganymede
# jpl_it("606") # titan
# jpl_it("504") # callisto
# jpl_it("501") # io
# jpl_it("502") # europa
# jpl_it("801") # triton
# jpl_it("703") # titania
# jpl_it("605") # rhea
# jpl_it("704") # oberon
# jpl_it("608") # iapetus
# jpl_it("136472") # makemake
# jpl_it("136108") # haumea
# jpl_it("901") # charon
# jpl_it("702") # umbriel
# jpl_it("701") # ariel
# jpl_it("604") # dione
# jpl_it("50000") # quaoar
# jpl_it("603") # tethys
# jpl_it("90377") # sedna
jpl_it("1;") # ceres
jpl_it("90482") # orcus
jpl_it("120347") # salacia
jpl_it("136199") # eris
jpl_it("557581") # 2007 OR
jpl_it("73435") # 2002 MS

tel.write("exit\n")

tel.read_all()
tel.close()
