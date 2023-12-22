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
	"time"
	"bytes"
	"fmt"
)

// checks if the CompositeContentGuardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompositeContentGuardResponse{}

// CompositeContentGuardResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type CompositeContentGuardResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// List of ContentGuards to ask for access-permission.
	Guards []*string `json:"guards,omitempty"`
}

type _CompositeContentGuardResponse CompositeContentGuardResponse

// NewCompositeContentGuardResponse instantiates a new CompositeContentGuardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeContentGuardResponse(name string) *CompositeContentGuardResponse {
	this := CompositeContentGuardResponse{}
	this.Name = name
	return &this
}

// NewCompositeContentGuardResponseWithDefaults instantiates a new CompositeContentGuardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeContentGuardResponseWithDefaults() *CompositeContentGuardResponse {
	this := CompositeContentGuardResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *CompositeContentGuardResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeContentGuardResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *CompositeContentGuardResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *CompositeContentGuardResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *CompositeContentGuardResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeContentGuardResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *CompositeContentGuardResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *CompositeContentGuardResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *CompositeContentGuardResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeContentGuardResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompositeContentGuardResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompositeContentGuardResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompositeContentGuardResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CompositeContentGuardResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CompositeContentGuardResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CompositeContentGuardResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CompositeContentGuardResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetGuards returns the Guards field value if set, zero value otherwise.
func (o *CompositeContentGuardResponse) GetGuards() []*string {
	if o == nil || IsNil(o.Guards) {
		var ret []*string
		return ret
	}
	return o.Guards
}

// GetGuardsOk returns a tuple with the Guards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeContentGuardResponse) GetGuardsOk() ([]*string, bool) {
	if o == nil || IsNil(o.Guards) {
		return nil, false
	}
	return o.Guards, true
}

// HasGuards returns a boolean if a field has been set.
func (o *CompositeContentGuardResponse) HasGuards() bool {
	if o != nil && !IsNil(o.Guards) {
		return true
	}

	return false
}

// SetGuards gets a reference to the given []*string and assigns it to the Guards field.
func (o *CompositeContentGuardResponse) SetGuards(v []*string) {
	o.Guards = v
}

func (o CompositeContentGuardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeContentGuardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Guards) {
		toSerialize["guards"] = o.Guards
	}
	return toSerialize, nil
}

func (o *CompositeContentGuardResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCompositeContentGuardResponse := _CompositeContentGuardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompositeContentGuardResponse)

	if err != nil {
		return err
	}

	*o = CompositeContentGuardResponse(varCompositeContentGuardResponse)

	return err
}

type NullableCompositeContentGuardResponse struct {
	value *CompositeContentGuardResponse
	isSet bool
}

func (v NullableCompositeContentGuardResponse) Get() *CompositeContentGuardResponse {
	return v.value
}

func (v *NullableCompositeContentGuardResponse) Set(val *CompositeContentGuardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeContentGuardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeContentGuardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeContentGuardResponse(val *CompositeContentGuardResponse) *NullableCompositeContentGuardResponse {
	return &NullableCompositeContentGuardResponse{value: val, isSet: true}
}

func (v NullableCompositeContentGuardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeContentGuardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


