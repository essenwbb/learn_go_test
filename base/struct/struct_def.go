package _struct

type Product struct {
	Name       string  `json:"name"`
	Weight     float64 `json:"weight,string"`
	WeightUnit string  `json:"weight_unit"`
	OnSale     bool    `json:"on_sale"`
}

type Payload struct {
	Objects     []interface{} `json:"objects,omitempty"`
	Description string        `json:"description"`
}

type ResData struct {
	Status  string  `json:"status"`
	Payload Payload `json:"payload"`
}
