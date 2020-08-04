package models

// VinResponse from NHTSA API
type VinResponse struct {
	Count          int         `json:"Count"`
	Meaasge        string      `json:"Message"`
	SearchCriteria string      `json:"SearchCriteria"`
	Attributes     []Attribute `json:"Results"`
}

// Attribute for a vehicle
type Attribute struct {
	Value      string `json:"Value"`
	ValueID    string `json:"ValueId"`
	Variable   string `json:"Variable"`
	VariableID int    `json:"VariableId"`
}

// Vehicle internal vehicle representation
type Vehicle struct {
	Year        string
	Make        string
	Model       string
	Trim        string
	VehicleType string
	DriveType   string
}