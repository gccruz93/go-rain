package main

type Size struct {
	h float32
	w float32
}

type Pos struct {
	x float32
	y float32
}

type SpawnRange struct {
	left  float32
	top   float32
	right float32
}

type Entity struct {
	Size
	Pos
	SpawnRange
}

func (e *Entity) Update() {}

func (e *Entity) Draw() {}

type Drop struct {
}
