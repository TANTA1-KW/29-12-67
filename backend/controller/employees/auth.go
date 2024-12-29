package employees


import (
	
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"project-se67/config"
	"project-se67/entity"

)

type (
 
	signUp struct {
		FirstName  string    `json:"first_name"`
		LastName   string    `json:"last_name"`
		Email      string    `json:"email"`
		Phone      string    `json:"phone"`
		Age        uint8     `json:"age"`
		Address    string    `json:"address"`
		Password   string    `json:"password"`
		BirthDay   time.Time `json:"birthday"`
		GenderID   uint      `json:"gender_id"`
		Salary     float32   `json:"salary"`
		Picture    string	 `json:"picture"`
		RoleID     uint      `json:"role_id"`
		StatID     uint      `json:"stat_id"`
		ShipID     uint      `json:"ship_id"`
	}
 
 )

 func SignUpEmployee(c *gin.Context) {
	var payload signUp
	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
 
	}
 
	db := config.DB()
	var userCheck1 entity.Customers
	var userCheck2 entity.Employees
	// Check if the user with the provided email already exists
	result1 := db.Where("email = ?", payload.Email).First(&userCheck1)
	if result1.Error != nil && !errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		// If there's a database error other than "record not found"
		c.JSON(http.StatusInternalServerError, gin.H{"error": result1.Error.Error()})
		return
 
	}

	result2 := db.Where("email = ?", payload.Email).First(&userCheck2)
	if result2.Error != nil && !errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		// If there's a database error other than "record not found"
		c.JSON(http.StatusInternalServerError, gin.H{"error": result2.Error.Error()})
		return
 
	}
 
	if userCheck1.ID != 0 || userCheck2.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
 
	}

	hashedPassword, _ := config.HashPassword(payload.Password)

	user := entity.Employees{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Age:       payload.Age,
		Address:   payload.Address,
		Password:  hashedPassword,
		BirthDay:  payload.BirthDay,
		GenderID:  payload.GenderID,
		Salary:    payload.Salary,
		Picture:   payload.Picture,
		RoleID:    payload.RoleID,
		StatID:    payload.StatID,
		ShipID:    payload.ShipID,
	}
 
 
	// Save the user to the database
 
	if err := db.Create(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Sign-up successful"})
 }