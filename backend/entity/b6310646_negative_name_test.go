package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeName(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "",
		Email:      "Supanan@gmail.com",
		EmployeeID: "M12345678",
	}
	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("ต้องไม่เป็นค่าว่าง"))
}
