package benchmarks

import "sort"

type User struct {
	Id   int
	Name string
}

type ByName []User

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Sort(data []User) []User {
	sort.Sort(ByName(data))
	return data
}
