package utils

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type BaseErrorResponse struct {
	BaseResponse
	Errors interface{} `json:"errors,omitempty"`
}

// Fiber context response functions (for direct use in handlers)
func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, code string, errors interface{}) error {
	if code == "" {
		code = "0001"
	}

	if message == "" {
		message = errorMessage(code)
	}

	return c.Status(statusCode).JSON(BaseErrorResponse{
		BaseResponse: BaseResponse{
			Success: false,
			Message: message,
			Code:    code,
		},
		Errors: errors,
	})
}

func Error(message string, code string) BaseResponse {
	if code == "" {
		code = "0001"
	}

	if message == "" {
		message = errorMessage(code)
	}

	return BaseResponse{
		Success: false,
		Message: message,
		Code:    code,
	}
}

func errorMessage(code string) string {
	if message, ok := codeList[code]; ok {
		return message
	}
	return "Something went wrong"
}

var codeList = map[string]string{
	"0001": "General error",
}
