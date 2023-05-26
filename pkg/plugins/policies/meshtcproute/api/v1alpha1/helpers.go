package v1alpha1

import (
	"github.com/kumahq/kuma/pkg/util/pointer"
)

func (x *To) GetDefault() interface{} {
	if len(x.Rules) == 0 {
		return Rule{
			Default: RuleConf{
				BackendRefs: []BackendRef{{
					TargetRef: x.TargetRef,
					Weight:    pointer.To(uint(1)),
				}},
			},
		}
	}

	return x.Rules[0]
}
