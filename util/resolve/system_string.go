// Code generated by "stringer -type System,VersionType"; DO NOT EDIT.

package resolve

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownSystem-0]
	_ = x[NPM-3]
	_ = x[Maven-6]
}

const (
	_System_name_0 = "UnknownSystem"
	_System_name_1 = "NPM"
	_System_name_2 = "Maven"
)

func (i System) String() string {
	switch {
	case i == 0:
		return _System_name_0
	case i == 3:
		return _System_name_1
	case i == 6:
		return _System_name_2
	default:
		return "System(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownVersionType-0]
	_ = x[Concrete-1]
	_ = x[Requirement-2]
}

const _VersionType_name = "UnknownVersionTypeConcreteRequirement"

var _VersionType_index = [...]uint8{0, 18, 26, 37}

func (i VersionType) String() string {
	if i >= VersionType(len(_VersionType_index)-1) {
		return "VersionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VersionType_name[_VersionType_index[i]:_VersionType_index[i+1]]
}
