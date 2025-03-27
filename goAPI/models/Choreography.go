package models

import "github.com/jackc/pgx/v5/pgtype"

type Choreography struct {
	ChoreoID    uint        `gorm:"primaryKey" json:"choreoID"`
	Dance       string      `json:"dance"`
	Description string      `json:"description"`
	Notes       string      `json:"notes"`
	Song        string      `json:"song"`
	Image       byte        `json:"image"`
	Date        pgtype.Date `json:"date"`
}
