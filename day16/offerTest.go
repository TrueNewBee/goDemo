package main

type S struct {
	m string
}

func f()*S{
	return &S{"foo"}	// A
}

func main() {
	p := f()	// B
	print(p.m)
}