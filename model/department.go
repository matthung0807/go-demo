package model

type Department struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Employees []Employee `json:"employees"`
}
