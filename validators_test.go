package validator_kodix

import (
	"reflect"
	"testing"
)

func TestKodixString_Compliant(t *testing.T) {
	var pattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	var stringText KodixString = "bb@mail.ru"
	err := stringText.Compliant(pattern)
	if err != nil {
		t.Error("Expected that text compliant to pattern")
	}
}

func TestInIntArray(t *testing.T) {
	if !InIntArray(1, []int{1, 2}) {
		t.Error("Expected that i2 in i1 array")
	}
}

func TestInStringArray(t *testing.T) {
	if !InStringArray("we", []string{"s", "we"}) {
		t.Error("Expected that i2 in i1 array")
	}
}

func TestKodixString_IsValid(t *testing.T) {
	checkValid(KodixString("1234567"), KodixString("12"), t)
}

func TestKodixArray_IsValid(t *testing.T) {
	var ka = KodixArray{1, 2, 3}
	err := ka.IsValid()
	if err != nil {
		t.Error(errMsgCorrectValid(reflect.TypeOf(ka).String()))
	}
}

func TestKodixIntMap_IsValid(t *testing.T) {
	checkValid(KodixIntMap{1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8}, KodixIntMap{1: 1, 2: 2}, t)
}

func TestKodixSlice_IsValid(t *testing.T) {
	checkValid(KodixSlice{1, 2, 3, 4, 5, 6, 7}, KodixSlice{1, 2}, t)
}

func TestKodixArray_IsEmpty(t *testing.T) {
	var empty KodixArray
	checkEmpty(empty, KodixArray{1}, t)
}

func TestKodixIntMap_IsEmpty(t *testing.T) {
	var empty KodixIntMap
	checkEmpty(empty, KodixIntMap{1: 2}, t)
}

func TestKodixSlice_IsEmpty(t *testing.T) {
	var empty KodixSlice
	checkEmpty(empty, KodixSlice{1, 2}, t)
}

func TestKodixString_IsEmpty(t *testing.T) {
	var empty KodixString
	checkEmpty(empty, KodixString("1"), t)
}

func errMsgForEmpty(typeName string) string {
	return "IsEmpty func says that var by type: " + typeName + " is not empty, but expected empty var."
}

func errMsgForNotEmpty(typeName string) string {
	return "IsEmpty func says that var by type" + typeName + " is empty, but expected  not empty var."
}

func errMsgIncorrectValid(typeName string) string {
	return "Not valid data by type: " + typeName + ", but method define as valid data."
}

func errMsgCorrectValid(typeName string) string {
	return "Valid data by type: " + typeName + ", but method define as not valid data."
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
	err := valid.IsValid()
	if err != nil {
		t.Error(errMsgCorrectValid(reflect.TypeOf(valid).String()))
	}

	err = notValid.IsValid()
	if err == nil {
		t.Error(errMsgIncorrectValid(reflect.TypeOf(notValid).String()))
	}
}
