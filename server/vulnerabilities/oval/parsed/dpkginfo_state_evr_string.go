package oval_parsed

import (
	"fmt"
	"strings"
)

// This is just one possible children of <dpkginfo_state>, that said
// all deb package tests are written againts evr strings
// see: https://oval.mitre.org/language/version5.10.1/ovaldefinition/documentation/oval-definitions-schema.html#EntityStateEVRStringType

// op|value
type ObjectStateEvrString string

func NewObjectState(op string, evr string) ObjectStateEvrString {
	return ObjectStateEvrString(fmt.Sprintf("%s|%s", op, evr))
}

func (sta ObjectStateEvrString) unpack() (OperationType, string) {
	parts := strings.Split(string(sta), "|")
	return NewOperationType(parts[0]), parts[1]
}

func (sta ObjectStateEvrString) Eval(ver string, cmp func(string, string) int) bool {
	op, evr := sta.unpack()

	r := cmp(ver, evr)
	switch op {
	case LessThan:
		return r == -1
	case Equals:
		return r == 0
	case NotEqual:
		return r != 0
	case GreaterThan:
		return r == 1
	case GreaterThanOrEqual:
		return r == 1 || r == 0
	case LessThanOrEqual:
		return r == -1 || r == 0
	}

	return false
}
