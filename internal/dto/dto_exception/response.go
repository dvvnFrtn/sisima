package dtoexception

// Exception
type ExceptionTitle string

const (
	NotFound          ExceptionTitle = "NOT_FOUND"
	InvalidQueryParam ExceptionTitle = "INVALID_QUERY_PARAM"
	InvalidRequest    ExceptionTitle = "INVALID_REQUEST"
	ValidationErr     ExceptionTitle = "VALIDATION_ERROR"
	InternalErr       ExceptionTitle = "INTERNAL_ERROR"
)

type exceptionResponse struct {
	Title ExceptionTitle `json:"title"`
	Error interface{}    `json:"errors,omitempty"`
}

func NewExceptionResponse(title ExceptionTitle, error interface{}) exceptionResponse {
	return exceptionResponse{
		Title: title,
		Error: error,
	}
}
