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

// checks if the GalaxyRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalaxyRoleResponse{}

// GalaxyRoleResponse A serializer for Galaxy's representation of Roles.
type GalaxyRoleResponse struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewGalaxyRoleResponse instantiates a new GalaxyRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalaxyRoleResponse(name string, namespace string) *GalaxyRoleResponse {
	this := GalaxyRoleResponse{}
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewGalaxyRoleResponseWithDefaults instantiates a new GalaxyRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalaxyRoleResponseWithDefaults() *GalaxyRoleResponse {
	this := GalaxyRoleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GalaxyRoleResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalaxyRoleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GalaxyRoleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GalaxyRoleResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *GalaxyRoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GalaxyRoleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GalaxyRoleResponse) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *GalaxyRoleResponse) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *GalaxyRoleResponse) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *GalaxyRoleResponse) SetNamespace(v string) {
	o.Namespace = v
}

func (o GalaxyRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalaxyRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableGalaxyRoleResponse struct {
	value *GalaxyRoleResponse
	isSet bool
}

func (v NullableGalaxyRoleResponse) Get() *GalaxyRoleResponse {
	return v.value
}

func (v *NullableGalaxyRoleResponse) Set(val *GalaxyRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGalaxyRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGalaxyRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalaxyRoleResponse(val *GalaxyRoleResponse) *NullableGalaxyRoleResponse {
	return &NullableGalaxyRoleResponse{value: val, isSet: true}
}

func (v NullableGalaxyRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalaxyRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


