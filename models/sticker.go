package models

import "bromo/utils"

// Sticker is struct for modeling data sticker
type Sticker struct {
	ID       *string `gorm:"primary_key;column:id" json:"id,omitempty"`
	Name     string  `gorm:"not null;column:name" json:"name" binding:"required"`
	ImageURL string  `gorm:"not null;column:image_url" json:"image_url" binding:"required"`
	PackID   string  `gorm:"column:pack_id" json:"-"`
}

// CheckDataStickerList used
func CheckDataStickerList(s []Sticker) []Sticker {
	parse := make([]Sticker, 0, len(s))
	for i := range s {
		parse = append(parse, CheckDataSticker(s[i]))
	}

	return parse
}

// CheckDataSticker used
func CheckDataSticker(s Sticker) Sticker {
	sticker := Sticker{}

	if s.ID == nil {
		id := utils.GeneratorUID64()
		sticker.ID = &id
	} else {
		sticker.ID = s.ID
	}

	sticker.Name = s.Name
	sticker.ImageURL = s.ImageURL
	sticker.PackID = s.PackID

	return sticker
}
