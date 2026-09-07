package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
    a float64
    b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
    a := n1.a
    b := n1.b
    c := n2.a
    d := n2.b
    return Number {a + c, b + d}
}

func (n1 Number) Subtract(n2 Number) Number {
    a := n1.a
    b := n1.b
    c := n2.a
    d := n2.b
    return Number {a - c, b - d}
}

func (n1 Number) Multiply(n2 Number) Number {
    a := n1.a
    b := n1.b
    c := n2.a
    d := n2.b
    return Number {a * c - b * d, b * c + a * d}
}

func (n Number) Times(factor float64) Number {
    return Number {n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
    a := n1.a
    b := n1.b
    c := n2.a
    d := n2.b
    return Number {(a * c + b * d) / (c * c + d * d), (b * c - a * d) / (c * c + d * d)}
}

func (n Number) Conjugate() Number {
	return Number {n.a, -n.b}
}

func (n Number) Abs() float64 {
    a := n.a
    b := n.b
	return math.Sqrt(a * a + b * b)
}

func (n Number) Exp() Number {
    a := n.a
    b := n.b
	return Number {math.Exp(a) * math.Cos(b), math.Exp(a) * math.Sin(b)}
}
