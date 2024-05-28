package models

type Employee struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	DepartmentID uint    `json:"department_id"`
	VrID         uint    `json:"vr_id"`
	Salary       float64 `json:"salary"`
	RankID       uint    `json:"rank_id"`
}

type CreateEmployeeRequest struct {
	DepartmentName string  `json:"department_name"`
	Salary         float64 `json:"salary"`
	Name           string  `json:"name"`
	RankName       string  `json:"rank_name"`
}

type UpdateEmployeeRequest struct {
	Name  string
	Value float64
}

type DeleteEmployeeRequest struct {
	EmployeeName string `json:"employee_name"`
}

type PromoteEmployeeRequest struct {
	EmployeeName string `json:"employee_name"`
	RankName     string `json:"rank_name"`
}
