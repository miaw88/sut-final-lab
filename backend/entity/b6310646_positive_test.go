package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "supanan rueangsook",
		Email:      "Supanan@gmail.com",
		EmployeeID: "M12345678",
	}
	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
