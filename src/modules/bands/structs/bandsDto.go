package structs

type SaveBandDto struct {
	Title   string `json:"title" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Country string `json:"country" binding:"required"`
	UF      string `json:"uf"`
	Style   string `json:"style" binding:"required"`
}

type UpdateBandDto struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Country string `json:"country"`
	UF      string `json:"uf"`
	Style   string `json:"style"`
}
