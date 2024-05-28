package models

type DefaultVR struct {
	ID           uint
	DepartmentID uint
	Value        float64
}

type VR struct {
	ID    uint
	Value float64
}

type GetVrRequest struct {
	VrID uint `json:"vr_id"`
}

type UpdateVrRequest struct {
	EmployeeName string
	Value        float64
}
