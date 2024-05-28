package models

import "time"

type Frequency struct {
	ID             uint
	DataReferencia time.Time
	TotalDias      int
}

type GetFrequencyRequest struct {
	Date         time.Time
	EmployeeName string
}

type UpdateFrequencyRequest struct {
	Date         time.Time
	EmployeeName string
	Frequency    int
}

type EmployeeFrequency struct {
	ID             uint
	DataReferencia time.Time
	FuncionarioID  uint
	FrequencyID    uint
}
