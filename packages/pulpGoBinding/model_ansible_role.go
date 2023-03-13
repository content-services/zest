/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the AnsibleRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleRole{}

// AnsibleRole A serializer for Role versions.
type AnsibleRole struct {
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	Version string `json:"version"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewAnsibleRole instantiates a new AnsibleRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleRole(artifact string, version string, name string, namespace string) *AnsibleRole {
	this := AnsibleRole{}
	this.Artifact = artifact
	this.Version = version
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewAnsibleRoleWithDefaults instantiates a new AnsibleRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleRoleWithDefaults() *AnsibleRole {
	this := AnsibleRole{}
	return &this
}

// GetArtifact returns the Artifact field value
func (o *AnsibleRole) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *AnsibleRole) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *AnsibleRole) SetArtifact(v string) {
	o.Artifact = v
}

// GetVersion returns the Version field value
func (o *AnsibleRole) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AnsibleRole) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AnsibleRole) SetVersion(v string) {
	o.Version = v
}

// GetName returns the Name field value
func (o *AnsibleRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleRole) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *AnsibleRole) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *AnsibleRole) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *AnsibleRole) SetNamespace(v string) {
	o.Namespace = v
}

func (o AnsibleRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifact"] = o.Artifact
	toSerialize["version"] = o.Version
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableAnsibleRole struct {
	value *AnsibleRole
	isSet bool
}

func (v NullableAnsibleRole) Get() *AnsibleRole {
	return v.value
}

func (v *NullableAnsibleRole) Set(val *AnsibleRole) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleRole) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleRole(val *AnsibleRole) *NullableAnsibleRole {
	return &NullableAnsibleRole{value: val, isSet: true}
}

func (v NullableAnsibleRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

