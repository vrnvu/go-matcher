package main

import "fmt"

type holder struct {
	operation string
	value     int
	l, r      *holder
}

type matcher func(h *holder) bool

func bin(op string, l, r matcher) matcher {
	return func(h *holder) bool {
		return h.operation == op && l(h.l) && r(h.r)
	}
}

func num(n *int) matcher {
	return func(h *holder) bool {
		if h.operation == "num" {
			*n = h.value
			return true
		}
		return false
	}
}

func main() {
	a := &holder{operation: "num", value: 10, l: nil, r: nil}
	b := &holder{operation: "num", value: 20, l: nil, r: nil}

	opAdd := &holder{operation: "+", value: 0, l: a, r: b}
	opMul := &holder{operation: "*", value: 0, l: a, r: b}

	for _, h := range []*holder{opAdd, opMul} {
		var a, b int
		switch {
		case bin("+", num(&a), num(&b))(h):
			fmt.Printf("%d + %d = %d\n", a, b, a+b)
		case bin("*", num(&a), num(&b))(h):
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
		default:
			panic("op not found")
		}
	}

}
