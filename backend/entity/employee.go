package entity


import (
	
   "time"
   "gorm.io/gorm"
)

type Employees struct {

   gorm.Model

	FirstName   string  `json:"first_name" valid:"required~FirstName is required"`
	LastName    string  `json:"last_name" valid:"required~LastName is required"`
	Email       string    `json:"email" valid:"required~Email is required, email~Email is invalid"`
	Phone       string    `json:"phone" valid:"required~Phone is required, stringlength(10|10)"`
	Age         uint8     `json:"age" valid:"required~Age is required, range(1|150)"`
	Address     string    `json:"Address" valid:"required~Address is required"`
	BirthDay    time.Time `json:"birthday" valid:"required~BirthDay is required"`
	Password    string    `json:"-" valid:"required~Password is required"`
   Salary    float32   `json:"salary" valid:"required~Salary is required" `
   Picture     string `json:"picture" gorm:"type:longtext" `
   
   GenderID  uint      `json:"gender_id" valid:"required~GenderID is required"`
   Gender    *Genders  `gorm:"foreignKey: gender_id" json:"gender"`

   RoleID    uint      `json:"role_id" valid:"required~RoleID is required"`
   Role    *Roles  `gorm:"foreignKey: role_id" json:"role"`

   StatID uint	 `json:"stat_id" valid:"required~StatID is required"`
   Stat  *Stats `gorm:"foreignKey: StatID" json:"stat"`

   ShipID uint	 `json:"ship_id" valid:"required~ShipID is required"`
//   Ship  *Ship `gorm:"foreignKey: ShipID" json:"ship"`
} 