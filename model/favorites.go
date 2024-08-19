/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package model

import (
	"encoding/json"
)

type Favorite struct {
	Pk              int             `gorm:"column:pk_favorite; PrimaryKey;" json:"pk_favorite,omitempty"`
	VideoPlaylistId string          `gorm:"column:video_playlist_id;" json:"video_playlist_id,omitempty"`
	Item            json.RawMessage `gorm:"column:item;" json:"item,omitempty"`
}

func (Favorite) TableName() string {
	return "tr_favorites"
}

type Favorites []Favorite
