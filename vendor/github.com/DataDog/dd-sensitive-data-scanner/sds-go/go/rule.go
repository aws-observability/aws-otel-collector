package dd_sds

/*
#include <stdlib.h>
#include <dd_sds.h>
*/
import "C"

type RuleConfig interface {
	CreateRule() (*Rule, error)
}

type Rule struct {
	// pointer to the native rule
	nativeRulePtr int64
}

func CreateRuleFromRawPtr(ptr int64) Rule {
	return Rule{
		nativeRulePtr: ptr,
	}
}

// Delete deletes the native resources associated with this Rule.
// It is safe to delete it while it is still being used by a scanner.
func (r Rule) Delete() {
	C.free_any_rule(C.long(r.nativeRulePtr))
}

type RuleList struct {
	// pointer to the native list
	nativePtr int64
}

func CreateRuleList() RuleList {
	var ptr = C.create_rule_list()
	return RuleList{
		nativePtr: int64(ptr),
	}
}

func (l RuleList) AppendRule(r *Rule) {
	C.append_rule_to_list(C.long(r.nativeRulePtr), C.long(l.nativePtr))
}

// Delete deletes the native resources associated with this RuleList.
// It is safe to delete it while it is still being used by a scanner.
func (r RuleList) Delete() {
	C.free_rule_list(C.long(r.nativePtr))
}
