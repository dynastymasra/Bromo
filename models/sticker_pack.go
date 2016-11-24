package models

// StickerPack is struct for define table pack stricker and merchant uses
type StickerPack struct {
	MerchantID string `gorm:"not null;column:merchant_id" json:"-"`
	PackID     string `gorm:"not null;column:pack_id" json:"pack_id" binding:"required"`
	Status     bool   `gorm:"not null;column:status" json:"is_active"`
}
