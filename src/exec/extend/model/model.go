package model

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	score float64
}

type C struct {
	A
	B
}

type D struct {
	A A
}
