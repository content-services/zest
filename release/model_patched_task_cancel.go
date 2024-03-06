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

// checks if the PatchedTaskCancel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedTaskCancel{}

// PatchedTaskCancel Base serializer for use with :class:`pulpcore.app.models.Model`This ensures that all Serializers provide values for the 'pulp_href` field.The class provides a default for the ``ref_name`` attribute in theModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitionsof plugins are namespaced properly.
type PatchedTaskCancel struct {
	// The desired state of the task. Only 'canceled' is accepted.
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedTaskCancel PatchedTaskCancel

// NewPatchedTaskCancel instantiates a new PatchedTaskCancel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTaskCancel() *PatchedTaskCancel {
	this := PatchedTaskCancel{}
	return &this
}

// NewPatchedTaskCancelWithDefaults instantiates a new PatchedTaskCancel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTaskCancelWithDefaults() *PatchedTaskCancel {
	this := PatchedTaskCancel{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PatchedTaskCancel) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTaskCancel) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PatchedTaskCancel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PatchedTaskCancel) SetState(v string) {
	o.State = &v
}

func (o PatchedTaskCancel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedTaskCancel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedTaskCancel) UnmarshalJSON(data []byte) (err error) {
	varPatchedTaskCancel := _PatchedTaskCancel{}

	err = json.Unmarshal(data, &varPatchedTaskCancel)

	if err != nil {
		return err
	}

	*o = PatchedTaskCancel(varPatchedTaskCancel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedTaskCancel struct {
	value *PatchedTaskCancel
	isSet bool
}

func (v NullablePatchedTaskCancel) Get() *PatchedTaskCancel {
	return v.value
}

func (v *NullablePatchedTaskCancel) Set(val *PatchedTaskCancel) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTaskCancel) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTaskCancel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTaskCancel(val *PatchedTaskCancel) *NullablePatchedTaskCancel {
	return &NullablePatchedTaskCancel{value: val, isSet: true}
}

func (v NullablePatchedTaskCancel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTaskCancel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


