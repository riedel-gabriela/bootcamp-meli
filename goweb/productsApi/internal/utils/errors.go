package utils

import "errors"

var ErrNoFieldsToUpdate = errors.New("no fields to update")
var ErrProductNotFound = errors.New("product not found")
var ErrInvalidProductID = errors.New("invalid product ID")
var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrProductAlreadyExists = errors.New("product already exists")
var ErrInvalidDateFormat = errors.New("invalid date format, expected DD/MM/YYYY")
var ErrProductNotCreated = errors.New("product not created, check the request body")
var ErrProductNotUpdated = errors.New("product not updated, check the request body")
var ErrProductNotPatched = errors.New("product not patched, check the request body")
var ErrProductNotDeleted = errors.New("product not deleted, check the request body")
var ErrNoProductsFound = errors.New("no products found for the given criteria")
var ErrInternalServerError = errors.New("internal server error, please try again later")
var ErrInvalidPrice = errors.New("invalid price, must be greater than zero")
var ErrInvalidQuantity = errors.New("invalid quantity, must be greater than or equal to zero")
var ErrInvalidCodeValue = errors.New("invalid code value, must be a non-empty string")
var ErrInvalidExpirationDate = errors.New("invalid expiration date, must be in DD/MM/YYYY format")
var ErrCouldntReadFile = errors.New("couldn't read the file, please check the file path and permissions")
var ErrCouldntUnmarshalJSON = errors.New("couldn't unmarshal JSON, please check the JSON format")
