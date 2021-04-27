package platforms

// StringError is a basic error with static information.
type StringError string

func (err StringError) Error() string {
	return string(err)
}

const (
	// ErrUnsupportedClientAPI is used in case the API is not available by the platform.
	ErrUnsupportedClientAPI = StringError("unsupported ClientAPI")
)
