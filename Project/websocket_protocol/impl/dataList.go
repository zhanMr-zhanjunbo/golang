package impl

type User struct {
	Name string
}

type Tidings struct {
	Name string         `json:"username"`
	Information string  `json:"message"`
}

type DataSet struct {
	Users []User
	News  []Tidings
}
