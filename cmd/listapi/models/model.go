package models

type ListEntry struct {
	Selected bool   `json:"selected,omitempty" example:"false"`
	Value    string `json:"value" example:"Renew insurance"`
}

type List struct {
	Id      string      `json:"id,omitempty" example:"8d314920-b0f9-423c-acfa-17cb610370f6"`
	Name    string      `json:"name,omitempty" example:"Todo"`
	Entries []ListEntry `json:"entries,omitempty"`
}
