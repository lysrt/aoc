package main

func asmInitial() int {
	var a, h int
	b := 93 // set b 93
	c := b  // set c b

	if a != 0 {
		goto l1 // jnz a 2
	}

	goto l2 // jnz 1 5

l1:
	b = b * 100       // mul b 100
	b = b - (-100000) // sub b -100000
	c = b             // set c b
	c = c - (-17000)  // sub c -17000

l2:
	f := 1 // set f 1
	d := 2 // set d 2

l5:
	e := 2 // set e 2
l4:
	g := d // set g d

	g = g * e // mul g e
	g = g - b // sub g b

	if g != 0 {
		goto l3 //jnz g 2
	}

	f = 0 //set f 0

l3:
	e = e - (-1) // sub e -1
	g = e        // set g e
	g = g - b    // sub g b

	if g != 0 {
		goto l4 // jnz g -8
	}

	d = d - (-1) // sub d -1
	g = d        // set g d
	g = g - b    // sub g b

	if g != 0 {
		goto l5 // jnz g -13
	}

	if f != 0 {
		goto l6 // jnz f 2
	}

	h = h - (-1) // sub h -1

l6:
	g = b     // set g b
	g = g - c // sub g c

	if g != 0 {
		goto l7 // jnz g 2
	}

	return h // jnz 1 3

l7:
	b = b - (-17) // sub b -17
	goto l2       // jnz 1 -23
}
