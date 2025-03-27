package models

import "github.com/jackc/pgx/v5/pgtype"

type Users struct {
	Id           uint        `gorm:"primaryKey" json:"userID"`
	Username     string      `json:"username"`
	PasswordHash string      `json:"password"`
	Email        string      `json:"email"`
	CreatedAt    pgtype.Date `json:"created_at"`
}
