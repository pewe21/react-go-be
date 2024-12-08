package model

import (
	"database/sql"
	"time"
)

type ProductModel struct {
	Id          int          `json:"id" goqu:"skipinsert,skipupdate"`
	Name        string       `json:"name"`
	Stock       int          `json:"stock"`
	Price       int64        `json:"price"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
}
