package graphqlstructs

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserAccesses : User struct to read/write user information
type UserAccesses struct {
	UUID       string `gorm:"type:char(36);primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EntityType string
	EntityUUID *string `gorm: "index"`
	UserUUID   *string `gorm: "index"`
}

// BeforeCreate : Populates the UUID value
func (access *UserAccesses) BeforeCreate(db *gorm.DB) {
	access.UUID = uuid.NewV4().String()
	return
}
