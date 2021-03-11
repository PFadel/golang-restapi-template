package errors

import "fmt"

// ErrorConnectingToAPI TODO
type ErrorConnectingToAPI struct {
	BaseError error
}

func (e ErrorConnectingToAPI) Error() string {
	return fmt.Sprintf("error connecting to API: %s", e.BaseError.Error())
}

// UnexpectedStatusCode TODO
type UnexpectedStatusCode struct {
	Expected int
	Actual   int
}

func (e UnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d was expecting: %d", e.Actual, e.Expected)
}

// ErrorParsingResponse TODO
type ErrorParsingResponse struct {
	BaseError error
}

func (e ErrorParsingResponse) Error() string {
	return fmt.Sprintf("error parsing response: %s", e.BaseError.Error())
}
