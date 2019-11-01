// Code generated by "string2enum -linecomment -type ParamAttr ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.ParamAttrImmArg-0]
	_ = x[enum.ParamAttrInAlloca-1]
	_ = x[enum.ParamAttrInReg-2]
	_ = x[enum.ParamAttrNest-3]
	_ = x[enum.ParamAttrNoAlias-4]
	_ = x[enum.ParamAttrNoCapture-5]
	_ = x[enum.ParamAttrNonNull-6]
	_ = x[enum.ParamAttrReadNone-7]
	_ = x[enum.ParamAttrReadOnly-8]
	_ = x[enum.ParamAttrReturned-9]
	_ = x[enum.ParamAttrSignExt-10]
	_ = x[enum.ParamAttrSRet-11]
	_ = x[enum.ParamAttrSwiftError-12]
	_ = x[enum.ParamAttrSwiftSelf-13]
	_ = x[enum.ParamAttrWriteOnly-14]
	_ = x[enum.ParamAttrZeroExt-15]
}

const _ParamAttr_name = "immarginallocainregnestnoaliasnocapturenonnullreadnonereadonlyreturnedsignextsretswifterrorswiftselfwriteonlyzeroext"

var _ParamAttr_index = [...]uint8{0, 6, 14, 19, 23, 30, 39, 46, 54, 62, 70, 77, 81, 91, 100, 109, 116}

// ParamAttrFromString returns the ParamAttr enum corresponding to s.
func ParamAttrFromString(s string) enum.ParamAttr {
	if len(s) == 0 {
		return 0
	}
	for i := range _ParamAttr_index[:len(_ParamAttr_index)-1] {
		if s == _ParamAttr_name[_ParamAttr_index[i]:_ParamAttr_index[i+1]] {
			return enum.ParamAttr(i)
		}
	}
	panic(fmt.Errorf("unable to locate ParamAttr enum corresponding to %q", s))
}
