package handlers_test

import (
	"bytes"
	. "handlers"
	"net/http"
	"net/http/httptest"
	"users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserHandler", func() {
	var handler *UserHandler
	var mockUserService *mockUserService

	BeforeEach(func() {
		mockUserService = newMockUserService()
		handler = NewUserHandler(mockUserService)
	})

	It("registers a new user", func() {
		mockUserService.RegisterUserOutput.Ret0 <- nil
		recorder := httptest.NewRecorder()

		json := `{
			"EmailAddress": "bobby@gmail.com",
			"FirstName": "bobby",
			"LastName": "fisher",
			"Password": "password"
		}`

		expectedUser := &users.User{
			EmailAddress: "bobby@gmail.com",
			FirstName:    "bobby",
			LastName:     "fisher",
			Password:     "password",
		}

		request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(json)))
		Expect(err).NotTo(HaveOccurred())

		handler.ServeHttp(recorder, request)

		Expect(recorder.Code).To(Equal(http.StatusOK))
		Expect(len(mockUserService.RegisterUserCalled)).To(Equal(1))
		receivedUser := <-mockUserService.RegisterUserInput.NewUser
		Expect(receivedUser.EmailAddress).To(Equal(expectedUser.EmailAddress))
	})
})
