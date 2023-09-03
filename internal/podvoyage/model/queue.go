package model

type Queue struct {
	Episode   *Episode `json:"episode" gorm:"foreignKey:EpisodeId;references:Id;constraint:OnDelete:CASCADE;"`
	EpisodeId int      `json:"-"`
	UserId    int      `json:"-"`
	Pos       int      `json:"pos"`
}
