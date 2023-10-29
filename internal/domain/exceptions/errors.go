package exceptions

import "errors"

var DbError error = errors.New("error while saving on db")
var UnkownErrror = errors.New("unknown error")
