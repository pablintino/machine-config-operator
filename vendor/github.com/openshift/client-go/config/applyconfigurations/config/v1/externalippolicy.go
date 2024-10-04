// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ExternalIPPolicyApplyConfiguration represents a declarative configuration of the ExternalIPPolicy type for use
// with apply.
type ExternalIPPolicyApplyConfiguration struct {
	AllowedCIDRs  []string `json:"allowedCIDRs,omitempty"`
	RejectedCIDRs []string `json:"rejectedCIDRs,omitempty"`
}

// ExternalIPPolicyApplyConfiguration constructs a declarative configuration of the ExternalIPPolicy type for use with
// apply.
func ExternalIPPolicy() *ExternalIPPolicyApplyConfiguration {
	return &ExternalIPPolicyApplyConfiguration{}
}

// WithAllowedCIDRs adds the given value to the AllowedCIDRs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedCIDRs field.
func (b *ExternalIPPolicyApplyConfiguration) WithAllowedCIDRs(values ...string) *ExternalIPPolicyApplyConfiguration {
	for i := range values {
		b.AllowedCIDRs = append(b.AllowedCIDRs, values[i])
	}
	return b
}

// WithRejectedCIDRs adds the given value to the RejectedCIDRs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RejectedCIDRs field.
func (b *ExternalIPPolicyApplyConfiguration) WithRejectedCIDRs(values ...string) *ExternalIPPolicyApplyConfiguration {
	for i := range values {
		b.RejectedCIDRs = append(b.RejectedCIDRs, values[i])
	}
	return b
}
