package structs

type BandDto struct {
	Title   string `json:"title" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Country string `json:"country" binding:"required"`
	UF      string `json:"uf"`
	Style   string `json:"style" binding:"required"`
}
