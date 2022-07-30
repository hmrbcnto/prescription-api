package entities

type Drug struct {
	ID string `json:"_id"`
	GenericName string `json:"generic_name"`
	Dosage string `json:"dosage"`
	CurrentStock int `json:"current_stock"`
	BrandName string `json:"brand_name"`
}