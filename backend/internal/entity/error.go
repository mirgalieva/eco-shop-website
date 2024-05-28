package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrProductNotFound = errors.New("product not found")
var ErrInvalidEmail = errors.New("email is incorrect")
var ErrInvalidPassword = errors.New("password is incorrect")
