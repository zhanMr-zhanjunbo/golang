package impl

type User struct {
	Name string
}

type Tidings struct {
	Name string
	Information string
}

type DataSet struct {
	Users []User
	News  []Tidings
}
