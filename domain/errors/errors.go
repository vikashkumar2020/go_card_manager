package errors

import "errors"

var (
	UserNotFound    = errors.New("user not found")
	UserExists      = errors.New("this user already exists")
	InvalidEmail    = errors.New("email is invalid")
	InvalidPassword = errors.New("password must be alphanumeric and must consists of at least 6 characters and not more than 15 characters")
	InvalidParams   = errors.New("please make sure all required fields are filled out and try again")
	InvalidFields   = errors.New("one or more fields in the data you sent are invalids")
)