package validator_kodix

import (
	"regexp"
)

// DefaultLen - constant to setting validity
// length of slices, arrays, strings and maps
const DefaultLen = 7

// Validator interface to check validity of data
type Validator interface {
	IsValid() error
}

// EmptyChecker interface to check that is variable empty
type EmptyChecker interface {
	IsEmpty() bool
}

// RegularChecker interface to check that string is compliant to pattern
type RegularChecker interface {
	Compliant(string) error
}

// KodixString wrapper for string data type, with interfaces implementation possibility
type KodixString string

// KodixIntMap wrapper for [int: int] map data type, with interfaces implementation possibility
type KodixIntMap map[int]int

// KodixSlice wrapper for slice data type, with interfaces implementation possibility
type KodixSlice []int

// KodixArray wrapper for  array of length 7 data type, with interfaces implementation possibility
type KodixArray [7]int

func (ks KodixString) Compliant(pattern string) error {
	re := regexp.MustCompile(pattern)
	if re.MatchString(string(ks)) {
		return nil
	}
	return StringPatternCompliantError(ks)
}

func (ks KodixString) IsValid() error {
	if len(ks) == validLen() {
		return nil
	}
	return StringValidError("")
}

func (kim KodixIntMap) IsValid() error {
	if len(kim) == validLen() {
		return nil
	}
	return MapIntValidError("")
}

func (ks KodixSlice) IsValid() error {
	if len(ks) == validLen() {
		return nil
	}
	return SliceValidError("")
}

func (ka KodixArray) IsValid() error {
	if len(ka) == validLen() {
		return nil
	}
	return ArrayValidError("")
}

func (ks KodixString) IsEmpty() bool {
	return ks == ""
}

func (ka KodixArray) IsEmpty() bool {
	for _, val := range ka {
		if val != 0 {
			return false
		}
	}
	return true
}

func (ks KodixSlice) IsEmpty() bool {
	return ks == nil
}

func (kim KodixIntMap) IsEmpty() bool {
	return kim == nil
}

func validLen() int {
	return DefaultLen
}

// InIntArray - function to define that exist number in number slice or not
func InIntArray(v int, arrV []int) bool {
	for _, i := range arrV {
		if v == i {
			return true
		}
	}
	return false
}

// InStringArray - function to define that exist string in string slice or not
func InStringArray(v string, arrV []string) bool {
	for _, i := range arrV {
		if v == i {
			return true
		}
	}
	return false
}
