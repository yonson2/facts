package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Fact struct {
	ID        uuid.UUID `db:"id"`
	Number    int       `db:"number"     rw:"r"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Content   string    `db:"content"`
}
