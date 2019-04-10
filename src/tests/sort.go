package tests

type Person struct {
	Id   int
	Name string
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func sorter() {
	//return sort.Sort(ByName(data))
}
