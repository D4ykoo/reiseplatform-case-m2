package entities

type AddressEntity struct {
	Street string `json:"street" gorm:"uniqueIndex"`
	State  string `json:"state" gorm:"uniqueIndex"`
	Land   string `json:"land" gorm:"uniqueIndex"`
}
