package domain

import "errors"

func errorDataSourse(cause error, text string) error {
	return errors.New(text + "\nCause: " + cause.Error())
}

func ErrorUnexpectedAPIResponse(cause error) error {
	return errorDataSourse(cause, "Unexpected external API response")
}

func ErrorStorageFailure(cause error) error {
	return errorDataSourse(cause, "Unexpected storage error")
}

func ErrorUnableToSendMessage(reason string) error {
	return errors.New("Unable to send a message: " + reason)
}
