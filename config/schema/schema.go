package schema

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// User schema
type User struct {
	gorm.Model
	FirstName    string `gorm:"size:255"`
	LastName     string `gorm:"size:255"`
	Email        string `gorm:"type:varchar(100);unique"`
	PasswordHash string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:255"`
	Disabled     bool
}

// BeforeCreate assigns a UUID for the user before creation
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// Organization schema
type Organization struct {
	gorm.Model
	TeamName     string    `gorm:"size:255"`
	ContactName  string    `gorm:"size:255"`
	ContactEmail string    `gorm:"size:255"`
	ContactPhone string    `gorm:"size:255"`
	Projects     []Project `gorm:"ForeignKey:OrganizationID;AssociationForeignKey:Refer"`
}

// Project schema
type Project struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Client         string `gorm:"size:255"`
	SlackChannel   string `gorm:"size:255"`
	StartDate      string `gorm:"size:255"`
	OrganizationID int
	Platforms      []Platform `gorm:"ForeignKey:ProjectID;AssociationForeignKey:Refer"`
	Pages          []Page     `gorm:"ForeignKey:ProjectID;AssociationForeignKey:Refer"`
	Tasks          []Task     `gorm:"ForeignKey:ProjectID;AssociationForeignKey:Refer"`
	Roles          []Role     `gorm:"ForeignKey:ProjectID;AssociationForeignKey:Refer"`
	Groups         []Group    `gorm:"ForeignKey:ProjectID;AssociationForeignKey:Refer"`
}

// Platform schema
type Platform struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

// Page schema
type Page struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

// Task schema
type Task struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

// Role schema
type Role struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

// Group schema
type Group struct {
	gorm.Model
	Name string `gorm:"size:255"`
}
