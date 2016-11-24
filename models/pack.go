package models

import "bromo/utils"

// Pack is struct for modeling stricker package
type Pack struct {
	ID           string        `grom:"primary_key;column:id" json:"id,omitempty"`
	Name         string        `grom:"not null;column:name" json:"name" binding:"required"`
	LogoURL      *string       `gorm:"not null;column:logo_url" json:"logo_url,omitempty"`
	CreatedBy    string        `gorm:"not null;column:created_by" json:"created_by" binding:"required"`
	IsPublic     bool          `gorm:"not null;column:is_public" json:"is_public"`
	Stickers     []Sticker     `gorm:"ForeignKey:PackID" json:"stickers" binding:"required"`
	StickerPacks []StickerPack `gorm:"ForeignKey:PackID" json:"-"`
}

// CheckDataPack used
func CheckDataPack(p Pack) Pack {
	pack := Pack{}

	if p.ID != "" {
		pack.ID = p.ID
	} else {
		id := utils.GeneratorUID64()
		pack.ID = id
	}
	pack.Name = p.Name

	if p.LogoURL != nil {
		pack.LogoURL = p.LogoURL
	} else {
		logoURL := p.Stickers[0].ImageURL
		pack.LogoURL = &logoURL
	}

	pack.CreatedBy = p.CreatedBy
	pack.IsPublic = p.IsPublic
	pack.Stickers = CheckDataStickerList(p.Stickers)

	return pack
}
