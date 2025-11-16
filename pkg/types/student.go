package types

type Student struct {
	ID    string
	Email string `validator:"required"`
	Name  string `validator:"required"`
	Age   int    `validator:"required"`
}
