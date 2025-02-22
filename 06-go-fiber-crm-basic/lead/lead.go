package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/lokesh1jha/go-fiber-crm-basic/database"
)

// Lead struct
type Lead struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// GetLeads func
func GetLeads(c *fiber.Ctx) error {
	db := database.DB
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

// GetLead func
func GetLead(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

// NewLead func
func NewLead(c *fiber.Ctx) error {
	db := database.DB
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(400).SendString(err.Error())
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}

// DeleteLead func
func DeleteLead(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	db.Delete(&Lead{}, id)
	c.SendString("Record deleted")
	return nil
}

// UpdateLead func
func UpdateLead(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(400).SendString(err.Error())
	}
	db.Model(&Lead{}).Where("id = ?", id).Updates(lead)
	c.JSON(lead)
	return nil
}
