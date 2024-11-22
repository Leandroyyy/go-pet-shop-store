package input_fiber

// import (
// 	"bytes"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	output_database "github.com/leandroyyy/poc-golang/src/adapters/output/database/in_memory"
// 	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
// 	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
// 	test_factories "github.com/leandroyyy/poc-golang/tests/factories"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRegisterOwnerUseCase struct {
// 	ExecuteFn func(req use_cases.RegisterOwnerUseCaseRequest) (*entities.Owner, error)
// }

// func (m *MockRegisterOwnerUseCase) Execute(req use_cases.RegisterOwnerUseCaseRequest) (*entities.Owner, error) {
// 	return m.ExecuteFn(req)
// }

// func TestRegisterOwnerController_Execute(t *testing.T) {
// 	app := fiber.New()

// 	// Instancie o mock com a função substituível
// 	mockUseCase := &MockRegisterOwnerUseCase{}

// 	// Passe o mock para o controlador
// 	controller := RegisterOwnerController{registerOwnerUseCase: use_cases.NewRegisterOwnerUseCase(output_database.InMemoryOwnerRepository{})}

// 	app.Post("/register-owner", controller.Execute)

// 	t.Run("should return 400 when payload is invalid", func(t *testing.T) {
// 		requestBody := `{"name":"John Doe","email":"invalid-email-format"}`

// 		req := httptest.NewRequest(http.MethodPost, "/register-owner", bytes.NewBuffer([]byte(requestBody)))
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, _ := app.Test(req, -1)

// 		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// 	})

// 	t.Run("should return 500 when use case returns error", func(t *testing.T) {
// 		requestBody := `{"name":"John Doe","document":"123456789","birthday":"2000-01-01","email":"john@example.com"}`

// 		mockUseCase.On("Execute", mock.Anything).Return(nil, errors.New("use case error"))

// 		req := httptest.NewRequest(http.MethodPost, "/register-owner", bytes.NewBuffer([]byte(requestBody)))
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, _ := app.Test(req, -1)

// 		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

// 		mockUseCase.AssertCalled(t, "Execute", use_cases.RegisterOwnerUseCaseRequest{
// 			Name:     "John Doe",
// 			Document: "123456789",
// 			Birthday: "2000-01-01",
// 			Email:    "john@example.com",
// 		})
// 	})

// 	t.Run("should return 200 when owner is registered successfully", func(t *testing.T) {
// 		requestBody := `{"name":"John Doe","document":"123456789","birthday":"2000-01-01","email":"john@example.com"}`

// 		owner := test_factories.MakeOwner()

// 		mockUseCase.On("Execute", mock.Anything).Return(owner, nil)

// 		req := httptest.NewRequest(http.MethodPost, "/register-owner", bytes.NewBuffer([]byte(requestBody)))
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, _ := app.Test(req, -1)

// 		assert.Equal(t, http.StatusOK, resp.StatusCode)

// 		var responseBody map[string]interface{}
// 		err := fiber.New().Config().JSONDecoder(resp.Body, &responseBody)
// 		assert.NoError(t, err)

// 		assert.Equal(t, "123", responseBody["id"])
// 		assert.Equal(t, "John Doe", responseBody["name"])
// 		assert.Equal(t, "123456789", responseBody["document"])
// 		assert.Equal(t, "john@example.com", responseBody["email"])
// 		assert.Equal(t, "2000-01-01", responseBody["birthday"])

// 		mockUseCase.AssertCalled(t, "Execute", use_cases.RegisterOwnerUseCaseRequest{
// 			Name:     "John Doe",
// 			Document: "123456789",
// 			Birthday: "2000-01-01",
// 			Email:    "john@example.com",
// 		})
// 	})
// }
