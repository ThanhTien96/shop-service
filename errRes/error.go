package errRes

import "errors"

var (
	// Error validating, binding
	ErrBindingData      = errors.New("err binding data")
	ErrValidatingData   = errors.New("err validating data")
	ErrValidatingHeader = errors.New("err validating header")
	ErrBindingHeader    = errors.New("err binding header")
	ErrUuidInvalid      = errors.New("uuid is invalid")
	ErrBindingParam = errors.New("err param binding data")
	// Error users, tokens
	ErrUnauthorized   = errors.New("unauthorized token")
	ErrInvalidToken   = errors.New("invalid token")
	ErrNoToken        = errors.New("missing token")
	ErrBearerRequired = errors.New("authorization Bearer is required")
	ErrNoPermission   = errors.New("you have no permissions to do this")
	// item
	ErrItemNotFound = errors.New("Item not found")
	ErrFailDecode = errors.New("Failed to decode item")
	ErrCreatePayloadFail = errors.New("Bindding item payload error")
	ErrCreateItemFail = errors.New("Create item faild")
	ErrFailQuery = errors.New("Fail to quer item")
	ErrUpdateFail = errors.New("No item was updated")
)
