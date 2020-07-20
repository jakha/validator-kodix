package validator_kodix

import "reflect"

type SliceValidError string

type ArrayValidError string

type MapIntValidError string

type StringValidError string

type StringPatternCompliantError string

// SliceValidError - implementation Error interface
func (sve SliceValidError) Error() string {
	return getLengthErrorMsg(reflect.TypeOf(sve).String())
}

// ArrayValidError - implementation Error interface
func (ave ArrayValidError) Error() string {
	return getLengthErrorMsg(reflect.TypeOf(ave).String())
}

// MapIntValidError - implementation Error interface
func (mive MapIntValidError) Error() string {
	return getLengthErrorMsg(reflect.TypeOf(mive).String())
}

// StringValidError - implementation Error interface
func (sve StringValidError) Error() string {
	return getLengthErrorMsg(reflect.TypeOf(sve).String())
}

// getLengthErrorMsg - message pattern for iterable errors
func getLengthErrorMsg(name string) string {
	return "Not valid " + name + " length."
}

// StringPatternCompliantError - implementation Error interface
func (spce StringPatternCompliantError) Error() string {
	return "String not compliant to pattern " + string(spce)
}
