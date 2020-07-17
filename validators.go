package validator_kodix

import "regexp"

const DefaultLen = 7

type Validator interface {
	IsValid() bool
}

type EmptyChecker interface {
	IsEmpty() bool
}

type RegularChecker interface {
	Compliant(string) bool
}

type KodixString string

type KodixIntMap map[int]int

type KodixSlice []int

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
	l := len(ka)
	return l == validLen()
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

func InIntArray(v int, arrV []int) bool {
	for _, i := range arrV {
		if v == i {
			return true
		}
	}
	return false
}

func InStringArray(v string, arrV []string) bool {
	for _, i := range arrV {
		if v == i {
			return true
		}
	}
	return false
}
