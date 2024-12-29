package unit

import (
	"fmt"
	"testing"
	"time"
	"project-se67/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmailEmValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("emailEm is required", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "", // Email ว่าง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Email is required"))
	})

	t.Run("emailEm is invalid format", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "invalid-email", // รูปแบบ Email ไม่ถูกต้อง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			StatID: 1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Email is invalid"))
	})

	t.Run("emailEm is valid", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com", // Email ถูกต้อง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			StatID: 1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestPhoneValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("phoneEm is required", func(t *testing.T) {

		Date, _ := time.Parse("2006-01-02", "2024-01-01")
		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "", // Phone ว่าง
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			StatID: 1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Phone is required"))
	})

	t.Run("phoneEm is invalid format", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "12345", // Phone สั้นเกินไป
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			StatID: 1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("phone: %s does not validate as stringlength(10|10)", employee.Phone)))
	})

	t.Run("phoneEm is valid", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		employee := entity.Employees{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "0801234567", // Phone ถูกต้อง
			Age:       20,
			Address:   "123 Main St",
			BirthDay:  Date,
			Password:  "password123",
			Salary:    15000.00,
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
			StatID: 1,
			ShipID: 1,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
