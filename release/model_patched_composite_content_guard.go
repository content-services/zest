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

// checks if the PatchedCompositeContentGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCompositeContentGuard{}

// PatchedCompositeContentGuard Base serializer for use with [pulpcore.app.models.Model][]This ensures that all Serializers provide values for the 'pulp_href` field.The class provides a default for the ``ref_name`` attribute in theModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitionsof plugins are namespaced properly.
type PatchedCompositeContentGuard struct {
	// The unique name.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// List of ContentGuards to ask for access-permission.
	Guards []*string `json:"guards,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCompositeContentGuard PatchedCompositeContentGuard

// NewPatchedCompositeContentGuard instantiates a new PatchedCompositeContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCompositeContentGuard() *PatchedCompositeContentGuard {
	this := PatchedCompositeContentGuard{}
	return &this
}

// NewPatchedCompositeContentGuardWithDefaults instantiates a new PatchedCompositeContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCompositeContentGuardWithDefaults() *PatchedCompositeContentGuard {
	this := PatchedCompositeContentGuard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCompositeContentGuard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCompositeContentGuard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCompositeContentGuard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCompositeContentGuard) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCompositeContentGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCompositeContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedCompositeContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedCompositeContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedCompositeContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedCompositeContentGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetGuards returns the Guards field value if set, zero value otherwise.
func (o *PatchedCompositeContentGuard) GetGuards() []*string {
	if o == nil || IsNil(o.Guards) {
		var ret []*string
		return ret
	}
	return o.Guards
}

// GetGuardsOk returns a tuple with the Guards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCompositeContentGuard) GetGuardsOk() ([]*string, bool) {
	if o == nil || IsNil(o.Guards) {
		return nil, false
	}
	return o.Guards, true
}

// HasGuards returns a boolean if a field has been set.
func (o *PatchedCompositeContentGuard) HasGuards() bool {
	if o != nil && !IsNil(o.Guards) {
		return true
	}

	return false
}

// SetGuards gets a reference to the given []*string and assigns it to the Guards field.
func (o *PatchedCompositeContentGuard) SetGuards(v []*string) {
	o.Guards = v
}

func (o PatchedCompositeContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCompositeContentGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Guards) {
		toSerialize["guards"] = o.Guards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedCompositeContentGuard) UnmarshalJSON(data []byte) (err error) {
	varPatchedCompositeContentGuard := _PatchedCompositeContentGuard{}

	err = json.Unmarshal(data, &varPatchedCompositeContentGuard)

	if err != nil {
		return err
	}

	*o = PatchedCompositeContentGuard(varPatchedCompositeContentGuard)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "guards")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCompositeContentGuard struct {
	value *PatchedCompositeContentGuard
	isSet bool
}

func (v NullablePatchedCompositeContentGuard) Get() *PatchedCompositeContentGuard {
	return v.value
}

func (v *NullablePatchedCompositeContentGuard) Set(val *PatchedCompositeContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCompositeContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCompositeContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCompositeContentGuard(val *PatchedCompositeContentGuard) *NullablePatchedCompositeContentGuard {
	return &NullablePatchedCompositeContentGuard{value: val, isSet: true}
}

func (v NullablePatchedCompositeContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCompositeContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


