package customers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"project-se67/config"
	"project-se67/entity"
	"project-se67/services"
	"time"
    "fmt"
)


type (

   Authen struct {

       Email    string `json:"email"`
       Password string `json:"password"`
   }


   signUp struct {
       
       FirstName string    `json:"first_name"`
       LastName  string    `json:"last_name"`
       Email     string    `json:"email"`
       Phone     string    `json:"phone"`
       Age       uint8     `json:"age"`
       Address   string    `json:"address"`
       BirthDay  time.Time `json:"birthday"`
       Password  string    `json:"password"`
       Picture   string    `json:"picture"`
       GenderID  uint      `json:"gender_id"`
       RoleID    uint      `json:"role_id"`
   }

)


func SignUp(c *gin.Context) {
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


   // Hash the user's password

   hashedPassword, _ := config.HashPassword(payload.Password)


   // Create a new user

   user := entity.Customers{

       FirstName: payload.FirstName,
       LastName:  payload.LastName,
       Email:     payload.Email,
       Phone:     payload.Phone,
       Age:       payload.Age,
       Address:   payload.Address,
       BirthDay:  payload.BirthDay,
       Password:  hashedPassword,
       Picture:   payload.Picture,
       GenderID:  payload.GenderID,
       RoleID:    payload.RoleID,

   }


   // Save the user to the database

   if err := db.Create(&user).Error; err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

       return

   }


   c.JSON(http.StatusCreated, gin.H{"message": "Sign-up successful"})

}


func SignIn(c *gin.Context) {
    var payload Authen
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    db := config.DB()
    var customer entity.Customers
    var employee entity.Employees

    // ตรวจสอบในตาราง Customers
    if err := db.Where("email = ?", payload.Email).First(&customer).Error; err == nil {
        // ตรวจสอบรหัสผ่านของ Customers
        if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(payload.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
            return
        }

        // สร้าง JWT Token
        jwtWrapper := services.JwtWrapper{
            SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
            Issuer:          "AuthService",
            ExpirationHours: 24,
        }

        signedToken, err := jwtWrapper.GenerateToken(customer.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": customer.ID, "role": "customer"})
        return
    }

    // หากไม่พบใน Customers ให้ตรวจสอบใน Employees
    if err := db.Where("email = ?", payload.Email).First(&employee).Error; err == nil {
        // ตรวจสอบรหัสผ่านของ Employees
        if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(payload.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
            return
        }

        // สร้าง JWT Token
        jwtWrapper := services.JwtWrapper{
            SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
            Issuer:          "AuthService",
            ExpirationHours: 24,
        }

        signedToken, err := jwtWrapper.GenerateToken(employee.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        fmt.Println("Generated Token:", signedToken)
        c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": employee.ID, "role": "employee"})
        return
    }
    
    // หากไม่พบข้อมูลในทั้งสองตาราง
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}


