package app





type Ticket struct {
	Id int32 `json:"id"`
	Catalog Catalog `json:"catalog"`
	Slot Slot `json:"slot"`
}