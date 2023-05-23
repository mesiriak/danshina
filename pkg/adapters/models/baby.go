package models

type Baby struct {
	FirstName string
	LastName  string
	Age       string
	Photo     string // we'll save photo on server, and save link on it in base (idk)
}
