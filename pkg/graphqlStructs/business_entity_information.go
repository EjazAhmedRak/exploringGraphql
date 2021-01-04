package graphqlstructs

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// BusinessEntity : User struct to read/write user information
type BusinessEntity struct {
	UUID       string `gorm:"type:char(36);primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EntityType string
	EntityName *string `gorm: "index"`
}

// BeforeCreate : Populates the UUID value
func (be *BusinessEntity) BeforeCreate(db *gorm.DB) {
	be.UUID = uuid.NewV4().String()
	return
}
