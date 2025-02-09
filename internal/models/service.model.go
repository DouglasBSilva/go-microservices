package models

type Service struct {
	ServiceID string  `gorm:"primaryKey" json:"serviceID"`
	Name      string  `json:"name"`
	Price     float32 `gorm:"type:numeric" json:"price"`
}
