/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
	"fmt"
)

// PackageTypesEnum * `bdist_dmg` - bdist_dmg* `bdist_dumb` - bdist_dumb* `bdist_egg` - bdist_egg* `bdist_msi` - bdist_msi* `bdist_rpm` - bdist_rpm* `bdist_wheel` - bdist_wheel* `bdist_wininst` - bdist_wininst* `sdist` - sdist
type PackageTypesEnum string

// List of PackageTypesEnum
const (
	PACKAGETYPESENUM_BDIST_DMG PackageTypesEnum = "bdist_dmg"
	PACKAGETYPESENUM_BDIST_DUMB PackageTypesEnum = "bdist_dumb"
	PACKAGETYPESENUM_BDIST_EGG PackageTypesEnum = "bdist_egg"
	PACKAGETYPESENUM_BDIST_MSI PackageTypesEnum = "bdist_msi"
	PACKAGETYPESENUM_BDIST_RPM PackageTypesEnum = "bdist_rpm"
	PACKAGETYPESENUM_BDIST_WHEEL PackageTypesEnum = "bdist_wheel"
	PACKAGETYPESENUM_BDIST_WININST PackageTypesEnum = "bdist_wininst"
	PACKAGETYPESENUM_SDIST PackageTypesEnum = "sdist"
)

// All allowed values of PackageTypesEnum enum
var AllowedPackageTypesEnumEnumValues = []PackageTypesEnum{
	"bdist_dmg",
	"bdist_dumb",
	"bdist_egg",
	"bdist_msi",
	"bdist_rpm",
	"bdist_wheel",
	"bdist_wininst",
	"sdist",
}

func (v *PackageTypesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageTypesEnum(value)
	for _, existing := range AllowedPackageTypesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageTypesEnum", value)
}

// NewPackageTypesEnumFromValue returns a pointer to a valid PackageTypesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageTypesEnumFromValue(v string) (*PackageTypesEnum, error) {
	ev := PackageTypesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageTypesEnum: valid values are %v", v, AllowedPackageTypesEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageTypesEnum) IsValid() bool {
	for _, existing := range AllowedPackageTypesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackageTypesEnum value
func (v PackageTypesEnum) Ptr() *PackageTypesEnum {
	return &v
}

type NullablePackageTypesEnum struct {
	value *PackageTypesEnum
	isSet bool
}

func (v NullablePackageTypesEnum) Get() *PackageTypesEnum {
	return v.value
}

func (v *NullablePackageTypesEnum) Set(val *PackageTypesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageTypesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageTypesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageTypesEnum(val *PackageTypesEnum) *NullablePackageTypesEnum {
	return &NullablePackageTypesEnum{value: val, isSet: true}
}

func (v NullablePackageTypesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageTypesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

