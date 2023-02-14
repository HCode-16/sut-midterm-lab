package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required"`
	Url  string `gorm:"UniqueIndex" valid:"url"`
}

func TestVideoValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Name", func(t *testing.T) {

		v := Video{
			Name: "",
			Url:  "https://www.youtube.com/watch?v=q9wUNgqVK0I",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name: non zero value required"))
	})

}
