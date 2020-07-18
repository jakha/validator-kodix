package validator_kodix

import "regexp"

// DefaultLen - constant to setting validity
// length of slices, arrays, strings and maps
const DefaultLen = 7

// Validator interface to check validity of data
type Validator interface {
	IsValid() bool
}

// EmptyChecker interface to check that is variable empty
type EmptyChecker interface {
	IsEmpty() bool
}

// RegularChecker interface to check that string is compliant to pattern
type RegularChecker interface {
	Compliant(string) bool
}

// KodixString wrapper for string data type, with interfaces implementation possibility
type KodixString string

// KodixIntMap wrapper for [int: int] map data type, with interfaces implementation possibility
type KodixIntMap map[int]int

// KodixSlice wrapper for slice data type, with interfaces implementation possibility
type KodixSlice []int

// KodixArray wrapper for  array of length 7 data type, with interfaces implementation possibility
type KodixArray [7]int

func (ks KodixString) Compliant(pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(string(ks))
}

func (ks KodixString) IsValid() bool {
	return len(ks) == validLen()
}

func (kim KodixIntMap) IsValid() bool {
	return len(kim) == validLen()
}

func (ks KodixSlice) IsValid() bool {
	return len(ks) == validLen()
}

func (ka KodixArray) IsValid() bool {
	return len(ka) == validLen()
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
