package domain

import "errors"

func errDataSourse(cause error, text string) error {
	return errors.New(text + "\nCause: " + cause.Error())
}

func ErrUnexpectedAPIResponse(cause error) error {
	return errDataSourse(cause, "Unexpected external API response")
}

func ErrStorageFailure(cause error) error {
	return errDataSourse(cause, "Unexpected storage error")
}

func ErrUnableToSendMessage(reason string) error {
	return errors.New("Unable to send a message: " + reason)
}
