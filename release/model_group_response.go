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

// checks if the GroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupResponse{}

// GroupResponse Serializer for Group.
type GroupResponse struct {
	// Name
	Name string `json:"name"`
	PulpHref *string `json:"pulp_href,omitempty"`
	Prn *string `json:"prn,omitempty"`
	Id *int64 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupResponse GroupResponse

// NewGroupResponse instantiates a new GroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResponse(name string) *GroupResponse {
	this := GroupResponse{}
	this.Name = name
	return &this
}

// NewGroupResponseWithDefaults instantiates a new GroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResponseWithDefaults() *GroupResponse {
	this := GroupResponse{}
	return &this
}

// GetName returns the Name field value
func (o *GroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupResponse) SetName(v string) {
	o.Name = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *GroupResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *GroupResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *GroupResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *GroupResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *GroupResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *GroupResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GroupResponse) SetId(v int64) {
	o.Id = &v
}

func (o GroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.Prn) {
		toSerialize["prn"] = o.Prn
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupResponse := _GroupResponse{}

	err = json.Unmarshal(data, &varGroupResponse)

	if err != nil {
		return err
	}

	*o = GroupResponse(varGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupResponse struct {
	value *GroupResponse
	isSet bool
}

func (v NullableGroupResponse) Get() *GroupResponse {
	return v.value
}

func (v *NullableGroupResponse) Set(val *GroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResponse(val *GroupResponse) *NullableGroupResponse {
	return &NullableGroupResponse{value: val, isSet: true}
}

func (v NullableGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


