/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RpmPackageEnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackageEnvironmentResponse{}

// RpmPackageEnvironmentResponse PackageEnvironment serializer.
type RpmPackageEnvironmentResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Environment id.
	Id string `json:"id"`
	// Environment name.
	Name string `json:"name"`
	// Environment description.
	Description string `json:"description"`
	// Environment display order.
	DisplayOrder NullableInt64 `json:"display_order"`
	// Environment group list.
	GroupIds map[string]interface{} `json:"group_ids"`
	// Environment option ids
	OptionIds map[string]interface{} `json:"option_ids"`
	// Environment description by language.
	DescByLang map[string]interface{} `json:"desc_by_lang"`
	// Environment name by language.
	NameByLang map[string]interface{} `json:"name_by_lang"`
	// Environment digest.
	Digest string `json:"digest"`
}

// NewRpmPackageEnvironmentResponse instantiates a new RpmPackageEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackageEnvironmentResponse(id string, name string, description string, displayOrder NullableInt64, groupIds map[string]interface{}, optionIds map[string]interface{}, descByLang map[string]interface{}, nameByLang map[string]interface{}, digest string) *RpmPackageEnvironmentResponse {
	this := RpmPackageEnvironmentResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.DisplayOrder = displayOrder
	this.GroupIds = groupIds
	this.OptionIds = optionIds
	this.DescByLang = descByLang
	this.NameByLang = nameByLang
	this.Digest = digest
	return &this
}

// NewRpmPackageEnvironmentResponseWithDefaults instantiates a new RpmPackageEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmPackageEnvironmentResponseWithDefaults() *RpmPackageEnvironmentResponse {
	this := RpmPackageEnvironmentResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmPackageEnvironmentResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmPackageEnvironmentResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmPackageEnvironmentResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmPackageEnvironmentResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmPackageEnvironmentResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmPackageEnvironmentResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetId returns the Id field value
func (o *RpmPackageEnvironmentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RpmPackageEnvironmentResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RpmPackageEnvironmentResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmPackageEnvironmentResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RpmPackageEnvironmentResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RpmPackageEnvironmentResponse) SetDescription(v string) {
	o.Description = v
}

// GetDisplayOrder returns the DisplayOrder field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *RpmPackageEnvironmentResponse) GetDisplayOrder() int64 {
	if o == nil || o.DisplayOrder.Get() == nil {
		var ret int64
		return ret
	}

	return *o.DisplayOrder.Get()
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetDisplayOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayOrder.Get(), o.DisplayOrder.IsSet()
}

// SetDisplayOrder sets field value
func (o *RpmPackageEnvironmentResponse) SetDisplayOrder(v int64) {
	o.DisplayOrder.Set(&v)
}

// GetGroupIds returns the GroupIds field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetGroupIds() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetGroupIdsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return map[string]interface{}{}, false
	}
	return o.GroupIds, true
}

// SetGroupIds sets field value
func (o *RpmPackageEnvironmentResponse) SetGroupIds(v map[string]interface{}) {
	o.GroupIds = v
}

// GetOptionIds returns the OptionIds field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetOptionIds() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetOptionIdsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return map[string]interface{}{}, false
	}
	return o.OptionIds, true
}

// SetOptionIds sets field value
func (o *RpmPackageEnvironmentResponse) SetOptionIds(v map[string]interface{}) {
	o.OptionIds = v
}

// GetDescByLang returns the DescByLang field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetDescByLang() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DescByLang
}

// GetDescByLangOk returns a tuple with the DescByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetDescByLangOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DescByLang) {
		return map[string]interface{}{}, false
	}
	return o.DescByLang, true
}

// SetDescByLang sets field value
func (o *RpmPackageEnvironmentResponse) SetDescByLang(v map[string]interface{}) {
	o.DescByLang = v
}

// GetNameByLang returns the NameByLang field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetNameByLang() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.NameByLang
}

// GetNameByLangOk returns a tuple with the NameByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetNameByLangOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NameByLang) {
		return map[string]interface{}{}, false
	}
	return o.NameByLang, true
}

// SetNameByLang sets field value
func (o *RpmPackageEnvironmentResponse) SetNameByLang(v map[string]interface{}) {
	o.NameByLang = v
}

// GetDigest returns the Digest field value
func (o *RpmPackageEnvironmentResponse) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *RpmPackageEnvironmentResponse) SetDigest(v string) {
	o.Digest = v
}

func (o RpmPackageEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmPackageEnvironmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["display_order"] = o.DisplayOrder.Get()
	if o.GroupIds != nil {
		toSerialize["group_ids"] = o.GroupIds
	}
	if o.OptionIds != nil {
		toSerialize["option_ids"] = o.OptionIds
	}
	if o.DescByLang != nil {
		toSerialize["desc_by_lang"] = o.DescByLang
	}
	if o.NameByLang != nil {
		toSerialize["name_by_lang"] = o.NameByLang
	}
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableRpmPackageEnvironmentResponse struct {
	value *RpmPackageEnvironmentResponse
	isSet bool
}

func (v NullableRpmPackageEnvironmentResponse) Get() *RpmPackageEnvironmentResponse {
	return v.value
}

func (v *NullableRpmPackageEnvironmentResponse) Set(val *RpmPackageEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmPackageEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmPackageEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmPackageEnvironmentResponse(val *RpmPackageEnvironmentResponse) *NullableRpmPackageEnvironmentResponse {
	return &NullableRpmPackageEnvironmentResponse{value: val, isSet: true}
}

func (v NullableRpmPackageEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmPackageEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


