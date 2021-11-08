package main

type person struct {
	first string
}

type mongo map[int]person
type postg map[int]person

func (m mongo) save(n, int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postg) save(n, int, p person) {
	pg[n] = p
}

func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person)  {
	a.save(n int) person
}

func get(a accessor, n int) person  {
	return a.retrieve(n)
}






