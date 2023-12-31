// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatCustomerEncryptionKey is an auto-generated flat version of CustomerEncryptionKey.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCustomerEncryptionKey struct {
	KmsKeyName *string `mapstructure:"kmsKeyName" json:"kmsKeyName,omitempty" cty:"kmsKeyName" hcl:"kmsKeyName"`
	RawKey     *string `mapstructure:"rawKey" json:"rawKey,omitempty" cty:"rawKey" hcl:"rawKey"`
}

// FlatMapstructure returns a new FlatCustomerEncryptionKey.
// FlatCustomerEncryptionKey is an auto-generated flat version of CustomerEncryptionKey.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CustomerEncryptionKey) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCustomerEncryptionKey)
}

// HCL2Spec returns the hcl spec of a CustomerEncryptionKey.
// This spec is used by HCL to read the fields of CustomerEncryptionKey.
// The decoded values from this spec will then be applied to a FlatCustomerEncryptionKey.
func (*FlatCustomerEncryptionKey) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"kmsKeyName": &hcldec.AttrSpec{Name: "kmsKeyName", Type: cty.String, Required: false},
		"rawKey":     &hcldec.AttrSpec{Name: "rawKey", Type: cty.String, Required: false},
	}
	return s
}
