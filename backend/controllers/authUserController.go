package controllers

import (
	"regexp"
	"strings"

	"github.com/HiteshKumarMeghwar/L-M-S-2/bcryptPassword"
	"github.com/HiteshKumarMeghwar/L-M-S-2/database"
	"github.com/HiteshKumarMeghwar/L-M-S-2/models"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]string
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Check if password is less than 6 characters .......
	if len(data["password"]) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}

	// Validating Email Address .......
	if !validateEmail(strings.TrimSpace(data["email"])) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}

	// Check if email already exist in database ........
	database.DB.Where("email = ?", strings.TrimSpace(data["email"])).First(&userData)
	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	pass, _ := bcryptPassword.HashPassword(data["password"])

	// Fetch the role with the name "student" from the database.
	var studentRole models.Role
	database.DB.Where("name = ?", "Student").First(&studentRole)

	// If the role does not exist, create it.
	if studentRole.ID == 0 {
		studentRole = models.Role{Name: "Student"}
		database.DB.Create(&studentRole)
	}

	user := models.User{
		FirstName:      data["first_name"],
		LastName:       data["last_name"],
		Email:          strings.TrimSpace(data["email"]),
		Password:       pass,
		Phone:          data["phone"],
		ProfilePicture: data["profile_picture"],
		Roles:          []*models.Role{&studentRole}, // Assign the role as a slice
	}

	database.DB.Create(&user)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Account has been created successfully ...!",
		"user":    user,
	})
}
