package models

type Product struct {
	ID               int    `json:"id"`
	ProductID        string `json:"product_id"`
	ProductTypeID    string `json:"product_type_id"`
	ProductTextureID string `json:"product_texture_id"`
	ModelID          string `json:"model_id"`
	MasterID         string `json:"master_id"`
	// Price           float64 `json:"price"`
}
