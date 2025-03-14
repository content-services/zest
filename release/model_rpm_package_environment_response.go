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
	"fmt"
)

// checks if the RpmPackageEnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackageEnvironmentResponse{}

// RpmPackageEnvironmentResponse PackageEnvironment serializer.
type RpmPackageEnvironmentResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Environment id.
	Id string `json:"id"`
	// Environment name.
	Name string `json:"name"`
	// Environment description.
	Description string `json:"description"`
	// Environment display order.
	DisplayOrder NullableInt64 `json:"display_order"`
	// Environment group list.
	GroupIds interface{} `json:"group_ids"`
	// Environment option ids
	OptionIds interface{} `json:"option_ids"`
	// Environment description by language.
	DescByLang interface{} `json:"desc_by_lang"`
	// Environment name by language.
	NameByLang interface{} `json:"name_by_lang"`
	// Environment digest.
	Digest string `json:"digest"`
	AdditionalProperties map[string]interface{}
}

type _RpmPackageEnvironmentResponse RpmPackageEnvironmentResponse

// NewRpmPackageEnvironmentResponse instantiates a new RpmPackageEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackageEnvironmentResponse(id string, name string, description string, displayOrder NullableInt64, groupIds interface{}, optionIds interface{}, descByLang interface{}, nameByLang interface{}, digest string) *RpmPackageEnvironmentResponse {
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

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *RpmPackageEnvironmentResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *RpmPackageEnvironmentResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *RpmPackageEnvironmentResponse) SetPrn(v string) {
	o.Prn = &v
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

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *RpmPackageEnvironmentResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *RpmPackageEnvironmentResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *RpmPackageEnvironmentResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *RpmPackageEnvironmentResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageEnvironmentResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *RpmPackageEnvironmentResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *RpmPackageEnvironmentResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
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
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetGroupIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetGroupIdsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return &o.GroupIds, true
}

// SetGroupIds sets field value
func (o *RpmPackageEnvironmentResponse) SetGroupIds(v interface{}) {
	o.GroupIds = v
}

// GetOptionIds returns the OptionIds field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetOptionIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetOptionIdsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return &o.OptionIds, true
}

// SetOptionIds sets field value
func (o *RpmPackageEnvironmentResponse) SetOptionIds(v interface{}) {
	o.OptionIds = v
}

// GetDescByLang returns the DescByLang field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetDescByLang() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DescByLang
}

// GetDescByLangOk returns a tuple with the DescByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetDescByLangOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DescByLang) {
		return nil, false
	}
	return &o.DescByLang, true
}

// SetDescByLang sets field value
func (o *RpmPackageEnvironmentResponse) SetDescByLang(v interface{}) {
	o.DescByLang = v
}

// GetNameByLang returns the NameByLang field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RpmPackageEnvironmentResponse) GetNameByLang() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.NameByLang
}

// GetNameByLangOk returns a tuple with the NameByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageEnvironmentResponse) GetNameByLangOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NameByLang) {
		return nil, false
	}
	return &o.NameByLang, true
}

// SetNameByLang sets field value
func (o *RpmPackageEnvironmentResponse) SetNameByLang(v interface{}) {
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
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.Prn) {
		toSerialize["prn"] = o.Prn
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RpmPackageEnvironmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"display_order",
		"group_ids",
		"option_ids",
		"desc_by_lang",
		"name_by_lang",
		"digest",
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

	varRpmPackageEnvironmentResponse := _RpmPackageEnvironmentResponse{}

	err = json.Unmarshal(data, &varRpmPackageEnvironmentResponse)

	if err != nil {
		return err
	}

	*o = RpmPackageEnvironmentResponse(varRpmPackageEnvironmentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "display_order")
		delete(additionalProperties, "group_ids")
		delete(additionalProperties, "option_ids")
		delete(additionalProperties, "desc_by_lang")
		delete(additionalProperties, "name_by_lang")
		delete(additionalProperties, "digest")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


