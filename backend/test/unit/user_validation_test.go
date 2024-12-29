package unit

import (
	"fmt"
	"testing"
	"time"
	"project-se67/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomersEmailValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("email is required", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "", // Email ว่าง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Email is required"))
	})

	t.Run("email is invalid format", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "invalid-email", // รูปแบบ Email ไม่ถูกต้อง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Email is invalid"))
	})

	t.Run("email is valid", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com", // Email ถูกต้อง
			Phone:     "0800000000",
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestCustomersPhoneValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("phone is required", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "", // Phone ว่าง
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Phone is required"))
	})

	t.Run("phone is invalid format", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "12345", // Phone สั้นเกินไป
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("phone: %s does not validate as stringlength(10|10)", customer.Phone)))
	})

	t.Run("phone is valid", func(t *testing.T) {
		Date, _ := time.Parse("2006-01-02", "2024-01-01")

		customer := entity.Customers{
			FirstName: "Unit",
			LastName:  "Test",
			Email:     "test@example.com",
			Phone:     "0801234567", // Phone ถูกต้อง
			Age:       20,
			Address:   "123 Test St",
			BirthDay:  Date,
			Password:  "password123",
			Picture:   "",
			GenderID:  1,
			RoleID:    1,
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
