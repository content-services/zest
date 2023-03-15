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

// checks if the AnsibleRepositoryRebuild type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleRepositoryRebuild{}

// AnsibleRepositoryRebuild Serializer for Ansible Repository Rebuild.
type AnsibleRepositoryRebuild struct {
	Namespace NullableString `json:"namespace,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Version NullableString `json:"version,omitempty"`
}

// NewAnsibleRepositoryRebuild instantiates a new AnsibleRepositoryRebuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleRepositoryRebuild() *AnsibleRepositoryRebuild {
	this := AnsibleRepositoryRebuild{}
	return &this
}

// NewAnsibleRepositoryRebuildWithDefaults instantiates a new AnsibleRepositoryRebuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleRepositoryRebuildWithDefaults() *AnsibleRepositoryRebuild {
	this := AnsibleRepositoryRebuild{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRepositoryRebuild) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRepositoryRebuild) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *AnsibleRepositoryRebuild) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *AnsibleRepositoryRebuild) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *AnsibleRepositoryRebuild) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *AnsibleRepositoryRebuild) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRepositoryRebuild) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRepositoryRebuild) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AnsibleRepositoryRebuild) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AnsibleRepositoryRebuild) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AnsibleRepositoryRebuild) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AnsibleRepositoryRebuild) UnsetName() {
	o.Name.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRepositoryRebuild) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRepositoryRebuild) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *AnsibleRepositoryRebuild) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *AnsibleRepositoryRebuild) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *AnsibleRepositoryRebuild) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *AnsibleRepositoryRebuild) UnsetVersion() {
	o.Version.Unset()
}

func (o AnsibleRepositoryRebuild) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleRepositoryRebuild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableAnsibleRepositoryRebuild struct {
	value *AnsibleRepositoryRebuild
	isSet bool
}

func (v NullableAnsibleRepositoryRebuild) Get() *AnsibleRepositoryRebuild {
	return v.value
}

func (v *NullableAnsibleRepositoryRebuild) Set(val *AnsibleRepositoryRebuild) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleRepositoryRebuild) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleRepositoryRebuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleRepositoryRebuild(val *AnsibleRepositoryRebuild) *NullableAnsibleRepositoryRebuild {
	return &NullableAnsibleRepositoryRebuild{value: val, isSet: true}
}

func (v NullableAnsibleRepositoryRebuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleRepositoryRebuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


