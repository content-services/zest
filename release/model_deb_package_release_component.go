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
)

// checks if the DebPackageReleaseComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebPackageReleaseComponent{}

// DebPackageReleaseComponent A Serializer for PackageReleaseComponent.
type DebPackageReleaseComponent struct {
	// Package that is contained in release_comonent.
	Package string `json:"package"`
	// ReleaseComponent this package is contained in.
	ReleaseComponent string `json:"release_component"`
}

// NewDebPackageReleaseComponent instantiates a new DebPackageReleaseComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebPackageReleaseComponent(package_ string, releaseComponent string) *DebPackageReleaseComponent {
	this := DebPackageReleaseComponent{}
	this.Package = package_
	this.ReleaseComponent = releaseComponent
	return &this
}

// NewDebPackageReleaseComponentWithDefaults instantiates a new DebPackageReleaseComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebPackageReleaseComponentWithDefaults() *DebPackageReleaseComponent {
	this := DebPackageReleaseComponent{}
	return &this
}

// GetPackage returns the Package field value
func (o *DebPackageReleaseComponent) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *DebPackageReleaseComponent) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *DebPackageReleaseComponent) SetPackage(v string) {
	o.Package = v
}

// GetReleaseComponent returns the ReleaseComponent field value
func (o *DebPackageReleaseComponent) GetReleaseComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseComponent
}

// GetReleaseComponentOk returns a tuple with the ReleaseComponent field value
// and a boolean to check if the value has been set.
func (o *DebPackageReleaseComponent) GetReleaseComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseComponent, true
}

// SetReleaseComponent sets field value
func (o *DebPackageReleaseComponent) SetReleaseComponent(v string) {
	o.ReleaseComponent = v
}

func (o DebPackageReleaseComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebPackageReleaseComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package"] = o.Package
	toSerialize["release_component"] = o.ReleaseComponent
	return toSerialize, nil
}

type NullableDebPackageReleaseComponent struct {
	value *DebPackageReleaseComponent
	isSet bool
}

func (v NullableDebPackageReleaseComponent) Get() *DebPackageReleaseComponent {
	return v.value
}

func (v *NullableDebPackageReleaseComponent) Set(val *DebPackageReleaseComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableDebPackageReleaseComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableDebPackageReleaseComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebPackageReleaseComponent(val *DebPackageReleaseComponent) *NullableDebPackageReleaseComponent {
	return &NullableDebPackageReleaseComponent{value: val, isSet: true}
}

func (v NullableDebPackageReleaseComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebPackageReleaseComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


