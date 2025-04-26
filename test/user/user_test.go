package user_test

import (
	"net/http"
	"testing"

	"github.com/KiranPawar0/coditas-test/pkg/helper/structvalidator"
	"github.com/KiranPawar0/coditas-test/pkg/user"
	"github.com/KiranPawar0/coditas-test/test/payload"
	"github.com/KiranPawar0/coditas-test/test/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateUser API", func() {
	var r *gin.Engine

	BeforeEach(func() {
		r = gin.Default()
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			structvalidator.RegisterCustomValidations(v)
		}
		r.POST("/user", user.CreateUser)
	})

	It("should create a user successfully", func() {
		w := utils.PerformRequest(r, "POST", "/user", payload.ValidUser)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.Body.String()).To(ContainSubstring("User created successfully"))
	})

	It("should return a bad request error when PAN is invalid", func() {
		w := utils.PerformRequest(r, "POST", "/user", payload.InvalidPANUser)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		Expect(w.Body.String()).To(ContainSubstring("Field validation for 'PAN' failed"))
	})

	It("should return a bad request error when mobile number is invalid", func() {
		w := utils.PerformRequest(r, "POST", "/user", payload.InvalidMobileUser)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		Expect(w.Body.String()).To(ContainSubstring("Field validation for 'Mobile' failed"))
	})
})

func TestUserAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User API Suite")
}
