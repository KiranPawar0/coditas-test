package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KiranPawar0/coditas-test/pkg/helper/structvalidator"
	"github.com/KiranPawar0/coditas-test/pkg/user"
	"github.com/KiranPawar0/coditas-test/pkg/user/config"
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

	Context("when the user data is valid", func() {
		It("should create a user successfully", func() {
			userData := config.User{
				Name:   "John Doe",
				PAN:    "ABCDE1234F",
				Mobile: "1234567890",
				Email:  "john.doe@example.com",
			}

			w := performRequest(r, "POST", "/user", userData)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring("User created successfully"))
		})
	})

	Context("when the PAN is invalid", func() {
		It("should return a bad request error", func() {
			userData := config.User{
				Name:   "John Doe",
				PAN:    "INVALIDPAN",
				Mobile: "1234567890",
				Email:  "john.doe@example.com",
			}

			w := performRequest(r, "POST", "/user", userData)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring("Field validation for 'PAN' failed"))
		})
	})

	Context("when the mobile number is invalid", func() {
		It("should return a bad request error", func() {
			userData := config.User{
				Name:   "John Doe",
				PAN:    "ABCDE1234F",
				Mobile: "12345",
				Email:  "john.doe@example.com",
			}

			w := performRequest(r, "POST", "/user", userData)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring("Field validation for 'Mobile' failed"))
		})
	})

})

func performRequest(r *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUserAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User API Suite")
}
