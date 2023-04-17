package view

import "time"

type CreateEmployeeRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// Employee title:
	// * engineer - developer, programer.
	// * manager - product manger, project manager, supervisor.
	Title string `json:"title" enums:"engineer,manager"`
}

type CreateEmployeeResponse struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
