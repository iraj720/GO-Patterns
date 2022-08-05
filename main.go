package main

type x interface {
	plus() 
	// minus()
	// xx()
}

type h struct {
	x int
}

func (n h) plus() {
	n.x += 1
}

func (n *h) minus() {
	n.x -= 1
}

func (n *h) xx() {
	print(n.x)
}

func main() {
	int1 := h{}
	inter(int1)
	println(int1.x)
}

func inter(in x) {
	//in2, _ := in.(int)
	// in.plus().xx()
	// in.xx()
}
