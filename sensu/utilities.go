package sensu

func statusCodeToString(status int) string {
	var str string

	switch status {
	case 200:
		str = "OK"
	case 201:
		str = "Created"
	case 202:
		str = "Accepted"
	case 204:
		str = "No Content"
	case 400:
		str = "Bad Request"
	case 404:
		str = "Not Found"
	case 500:
		str = "Internal Server Error"
	case 503:
		str = "Service Unavailable"
	}

	return str
}
