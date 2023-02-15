package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeEmployee(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "supanan",
		Email:      "Supanan@gmail.com",
		EmployeeID: "L12345678",
	}
	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("รหัสพนักงานขึ้นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจำนวน 8 ตัว"))
}
