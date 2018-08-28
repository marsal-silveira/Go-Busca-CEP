package statuscode

// StatusCode : TODO
type StatusCode int

//Status Code Enum...
const (
	OK                  StatusCode = 200
	NoContent           StatusCode = 204
	NotFound            StatusCode = 404
	UnprocessableEntity StatusCode = 422
	InternalServerError StatusCode = 500
)

// ToInt : TODO
func ToInt(statusCode StatusCode) int {
	return int(statusCode)
}
