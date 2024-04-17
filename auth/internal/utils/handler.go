package utils

const (
	GET  = "GET"
	POST = "POST"
)

func IsSupportedHandlerMethod(httpMethodType string) bool {
	switch httpMethodType {
	case GET, POST:
		return true
	}
	return false
}
