package entity

type Saga struct {
	Id    string // Correlation id
	Name  string // Saga name
	State string // Saga state
}
