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

// checks if the RpmUpdateCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmUpdateCollectionResponse{}

// RpmUpdateCollectionResponse A Serializer for UpdateCollection.
type RpmUpdateCollectionResponse struct {
	// Collection name.
	Name NullableString `json:"name"`
	// Collection short name.
	Shortname NullableString `json:"shortname"`
	// Collection modular NSVCA.
	Module map[string]interface{} `json:"module"`
	// List of packages
	Packages []map[string]interface{} `json:"packages,omitempty"`
}

type _RpmUpdateCollectionResponse RpmUpdateCollectionResponse

// NewRpmUpdateCollectionResponse instantiates a new RpmUpdateCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmUpdateCollectionResponse(name NullableString, shortname NullableString, module map[string]interface{}) *RpmUpdateCollectionResponse {
	this := RpmUpdateCollectionResponse{}
	this.Name = name
	this.Shortname = shortname
	this.Module = module
	return &this
}

// NewRpmUpdateCollectionResponseWithDefaults instantiates a new RpmUpdateCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmUpdateCollectionResponseWithDefaults() *RpmUpdateCollectionResponse {
	this := RpmUpdateCollectionResponse{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmUpdateCollectionResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpdateCollectionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *RpmUpdateCollectionResponse) SetName(v string) {
	o.Name.Set(&v)
}

// GetShortname returns the Shortname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmUpdateCollectionResponse) GetShortname() string {
	if o == nil || o.Shortname.Get() == nil {
		var ret string
		return ret
	}

	return *o.Shortname.Get()
}

// GetShortnameOk returns a tuple with the Shortname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpdateCollectionResponse) GetShortnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortname.Get(), o.Shortname.IsSet()
}

// SetShortname sets field value
func (o *RpmUpdateCollectionResponse) SetShortname(v string) {
	o.Shortname.Set(&v)
}

// GetModule returns the Module field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmUpdateCollectionResponse) GetModule() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmUpdateCollectionResponse) GetModuleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Module) {
		return map[string]interface{}{}, false
	}
	return o.Module, true
}

// SetModule sets field value
func (o *RpmUpdateCollectionResponse) SetModule(v map[string]interface{}) {
	o.Module = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *RpmUpdateCollectionResponse) GetPackages() []map[string]interface{} {
	if o == nil || IsNil(o.Packages) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmUpdateCollectionResponse) GetPackagesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *RpmUpdateCollectionResponse) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []map[string]interface{} and assigns it to the Packages field.
func (o *RpmUpdateCollectionResponse) SetPackages(v []map[string]interface{}) {
	o.Packages = v
}

func (o RpmUpdateCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmUpdateCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["shortname"] = o.Shortname.Get()
	if o.Module != nil {
		toSerialize["module"] = o.Module
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	return toSerialize, nil
}

func (o *RpmUpdateCollectionResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"shortname",
		"module",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRpmUpdateCollectionResponse := _RpmUpdateCollectionResponse{}

	err = json.Unmarshal(bytes, &varRpmUpdateCollectionResponse)

	if err != nil {
		return err
	}

	*o = RpmUpdateCollectionResponse(varRpmUpdateCollectionResponse)

	return err
}

type NullableRpmUpdateCollectionResponse struct {
	value *RpmUpdateCollectionResponse
	isSet bool
}

func (v NullableRpmUpdateCollectionResponse) Get() *RpmUpdateCollectionResponse {
	return v.value
}

func (v *NullableRpmUpdateCollectionResponse) Set(val *RpmUpdateCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmUpdateCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmUpdateCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmUpdateCollectionResponse(val *RpmUpdateCollectionResponse) *NullableRpmUpdateCollectionResponse {
	return &NullableRpmUpdateCollectionResponse{value: val, isSet: true}
}

func (v NullableRpmUpdateCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmUpdateCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


