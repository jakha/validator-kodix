package validator_kodix

import (
	"reflect"
	"testing"
)

func TestKodixString_Compliant(t *testing.T) {
	var pattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	var stringText KodixString = "bb@mail.ru"
	if !stringText.Compliant(pattern) {
		t.Error("Expected that text compliant to pattern")
	}
}

func TestInIntArray(t *testing.T) {
	var i1 = []int{1, 2}
	var i2 = 1

	if !InIntArray(i2, i1) {
		t.Error("Expected that i2 in i1 array")
	}
}

func TestInStringArray(t *testing.T) {
	var i1 = []string{"s", "we"}
	var i2 = "we"

	if !InStringArray(i2, i1) {
		t.Error("Expected that i2 in i1 array")
	}
}

func TestKodixString_IsValid(t *testing.T) {
	var valid KodixString = "1234567"
	var notValid = KodixString("12")
	checkValid(valid, notValid, t)
}

func TestKodixArray_IsValid(t *testing.T) {
	var ka = KodixArray{1, 2, 3}
	if !ka.IsValid() {
		t.Error(errMsgCorrectValid(reflect.TypeOf(ka).String()))
	}
}

func TestKodixIntMap_IsValid(t *testing.T) {
	var notValid = KodixIntMap{1: 1, 2: 2}
	var valid = KodixIntMap{1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8}
	checkValid(valid, notValid, t)
}

func TestKodixSlice_IsValid(t *testing.T) {
	var notValid = KodixSlice{1, 2}
	var valid = KodixSlice{1, 2, 3, 4, 5, 6, 7}
	checkValid(valid, notValid, t)
}

func TestKodixArray_IsEmpty(t *testing.T) {
	var empty KodixArray
	notEmpty := KodixArray{1}
	checkEmpty(empty, notEmpty, t)
}

func TestKodixIntMap_IsEmpty(t *testing.T) {
	var empty KodixIntMap
	var notEmpty = KodixIntMap{1: 2}
	checkEmpty(empty, notEmpty, t)
}

func TestKodixSlice_IsEmpty(t *testing.T) {
	var empty KodixSlice
	var notEmpty = KodixSlice{1, 2}
	checkEmpty(empty, notEmpty, t)
}

func TestKodixString_IsEmpty(t *testing.T) {
	var empty KodixString
	var notEmpty = KodixString("1")
	checkEmpty(empty, notEmpty, t)
}

func errMsgForEmpty(typeName string) string {
	return "IsEmpty func says that var by type: " + typeName + " is not empty, but expected empty var."
}

func errMsgForNotEmpty(typeName string) string {
	return "IsEmpty func says that var by type" + typeName + " is empty, but expected  not empty var."
}

func errMsgIncorrectValid(typeName string) string {
	return "Not valid data by type: " + typeName + ", expected valid data."
}

func errMsgCorrectValid(typeName string) string {
	return "Valid data by type: " + typeName + ", expected not valid data."
}

func checkEmpty(empty EmptyChecker, notEmpty EmptyChecker, t *testing.T) {
	if !empty.IsEmpty() {
		t.Error(errMsgForNotEmpty(reflect.TypeOf(empty).String()))
	}

	if notEmpty.IsEmpty() {
		t.Error(errMsgForEmpty(reflect.TypeOf(notEmpty).String()))
	}
}

func checkValid(valid Validator, notValid Validator, t *testing.T) {
	if !valid.IsValid() {
		t.Error(errMsgIncorrectValid(reflect.TypeOf(valid).String()))
	}

	if notValid.IsValid() {
		t.Error(errMsgCorrectValid(reflect.TypeOf(notValid).String()))
	}
}
