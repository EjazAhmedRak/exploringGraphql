package graphqlstructs

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// User : User struct to read/write user information
type User struct {
	UUID      string `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     *string `gorm: "index"`
	Age       uint8
	Birthday  *time.Time
}

// BeforeCreate : Populates the UUID value
func (user *User) BeforeCreate(db *gorm.DB) {
	user.UUID = uuid.NewV4().String()
	return
}

