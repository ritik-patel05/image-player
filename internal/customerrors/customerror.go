package customerrors

type CustomError struct {
	Message        string
	HttpStatusCode int
	ErrorCode      string
}

func (c *CustomError) Error() string {
	return c.Message
}
