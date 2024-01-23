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
	"bytes"
	"fmt"
)

// checks if the CompositeContentGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompositeContentGuard{}

// CompositeContentGuard Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type CompositeContentGuard struct {
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// List of ContentGuards to ask for access-permission.
	Guards []*string `json:"guards,omitempty"`
}

type _CompositeContentGuard CompositeContentGuard

// NewCompositeContentGuard instantiates a new CompositeContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeContentGuard(name string) *CompositeContentGuard {
	this := CompositeContentGuard{}
	this.Name = name
	return &this
}

// NewCompositeContentGuardWithDefaults instantiates a new CompositeContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeContentGuardWithDefaults() *CompositeContentGuard {
	this := CompositeContentGuard{}
	return &this
}

// GetName returns the Name field value
func (o *CompositeContentGuard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeContentGuard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompositeContentGuard) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompositeContentGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompositeContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CompositeContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CompositeContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CompositeContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CompositeContentGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetGuards returns the Guards field value if set, zero value otherwise.
func (o *CompositeContentGuard) GetGuards() []*string {
	if o == nil || IsNil(o.Guards) {
		var ret []*string
		return ret
	}
	return o.Guards
}

// GetGuardsOk returns a tuple with the Guards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeContentGuard) GetGuardsOk() ([]*string, bool) {
	if o == nil || IsNil(o.Guards) {
		return nil, false
	}
	return o.Guards, true
}

// HasGuards returns a boolean if a field has been set.
func (o *CompositeContentGuard) HasGuards() bool {
	if o != nil && !IsNil(o.Guards) {
		return true
	}

	return false
}

// SetGuards gets a reference to the given []*string and assigns it to the Guards field.
func (o *CompositeContentGuard) SetGuards(v []*string) {
	o.Guards = v
}

func (o CompositeContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeContentGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Guards) {
		toSerialize["guards"] = o.Guards
	}
	return toSerialize, nil
}

func (o *CompositeContentGuard) UnmarshalJSON(data []byte) (err error) {
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

	varCompositeContentGuard := _CompositeContentGuard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompositeContentGuard)

	if err != nil {
		return err
	}

	*o = CompositeContentGuard(varCompositeContentGuard)

	return err
}

type NullableCompositeContentGuard struct {
	value *CompositeContentGuard
	isSet bool
}

func (v NullableCompositeContentGuard) Get() *CompositeContentGuard {
	return v.value
}

func (v *NullableCompositeContentGuard) Set(val *CompositeContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeContentGuard(val *CompositeContentGuard) *NullableCompositeContentGuard {
	return &NullableCompositeContentGuard{value: val, isSet: true}
}

func (v NullableCompositeContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

