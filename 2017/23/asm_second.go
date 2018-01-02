package main

func asmSecondVersion() int {
	var a, h int
	b := 93 // set b 93
	c := b  // set c b

	if a != 0 {
		b = b*100 + 100000 // mul b 100 // sub b -100000
		c = b + 17000      // set c b // sub c -17000
	}

	for {
		f := 1 // set f 1
		d := 2 // set d 2

	l5:
		e := 2 // set e 2
	l4:
		g := (d * e) - b // set g d // mul g e // sub g b
		if g == 0 {
			f = 0
		}

		e++       // sub e -1
		g = e - b // set g e // sub g b

		if g != 0 {
			goto l4 // jnz g -8
		}

		d++       // sub d -1
		g = d - b // set g d // sub g b
		if g != 0 {
			goto l5 // jnz g -13
		}

		if f == 0 {
			h++ // sub h -1
		}

		if b == c {
			return h // jnz 1 3
		}

		b += 17 // sub b -17
	}
}
