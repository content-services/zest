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
	"time"
)

// checks if the RpmPackageGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackageGroupResponse{}

// RpmPackageGroupResponse PackageGroup serializer.
type RpmPackageGroupResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// PackageGroup id.
	Id string `json:"id"`
	// PackageGroup default.
	Default *bool `json:"default,omitempty"`
	// PackageGroup user visibility.
	UserVisible *bool `json:"user_visible,omitempty"`
	// PackageGroup display order.
	DisplayOrder NullableInt32 `json:"display_order"`
	// PackageGroup name.
	Name string `json:"name"`
	// PackageGroup description.
	Description string `json:"description"`
	// PackageGroup package list.
	Packages map[string]interface{} `json:"packages"`
	// PackageGroup biarch only.
	BiarchOnly *bool `json:"biarch_only,omitempty"`
	// PackageGroup description by language.
	DescByLang map[string]interface{} `json:"desc_by_lang"`
	// PackageGroup name by language.
	NameByLang map[string]interface{} `json:"name_by_lang"`
	// PackageGroup digest.
	Digest string `json:"digest"`
}

// NewRpmPackageGroupResponse instantiates a new RpmPackageGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackageGroupResponse(id string, displayOrder NullableInt32, name string, description string, packages map[string]interface{}, descByLang map[string]interface{}, nameByLang map[string]interface{}, digest string) *RpmPackageGroupResponse {
	this := RpmPackageGroupResponse{}
	this.Id = id
	this.DisplayOrder = displayOrder
	this.Name = name
	this.Description = description
	this.Packages = packages
	this.DescByLang = descByLang
	this.NameByLang = nameByLang
	this.Digest = digest
	return &this
}

// NewRpmPackageGroupResponseWithDefaults instantiates a new RpmPackageGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmPackageGroupResponseWithDefaults() *RpmPackageGroupResponse {
	this := RpmPackageGroupResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmPackageGroupResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmPackageGroupResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmPackageGroupResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmPackageGroupResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmPackageGroupResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmPackageGroupResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetId returns the Id field value
func (o *RpmPackageGroupResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RpmPackageGroupResponse) SetId(v string) {
	o.Id = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RpmPackageGroupResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RpmPackageGroupResponse) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RpmPackageGroupResponse) SetDefault(v bool) {
	o.Default = &v
}

// GetUserVisible returns the UserVisible field value if set, zero value otherwise.
func (o *RpmPackageGroupResponse) GetUserVisible() bool {
	if o == nil || IsNil(o.UserVisible) {
		var ret bool
		return ret
	}
	return *o.UserVisible
}

// GetUserVisibleOk returns a tuple with the UserVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetUserVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.UserVisible) {
		return nil, false
	}
	return o.UserVisible, true
}

// HasUserVisible returns a boolean if a field has been set.
func (o *RpmPackageGroupResponse) HasUserVisible() bool {
	if o != nil && !IsNil(o.UserVisible) {
		return true
	}

	return false
}

// SetUserVisible gets a reference to the given bool and assigns it to the UserVisible field.
func (o *RpmPackageGroupResponse) SetUserVisible(v bool) {
	o.UserVisible = &v
}

// GetDisplayOrder returns the DisplayOrder field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RpmPackageGroupResponse) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DisplayOrder.Get()
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageGroupResponse) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayOrder.Get(), o.DisplayOrder.IsSet()
}

// SetDisplayOrder sets field value
func (o *RpmPackageGroupResponse) SetDisplayOrder(v int32) {
	o.DisplayOrder.Set(&v)
}

// GetName returns the Name field value
func (o *RpmPackageGroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmPackageGroupResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RpmPackageGroupResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RpmPackageGroupResponse) SetDescription(v string) {
	o.Description = v
}

// GetPackages returns the Packages field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageGroupResponse) GetPackages() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageGroupResponse) GetPackagesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Packages) {
		return map[string]interface{}{}, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *RpmPackageGroupResponse) SetPackages(v map[string]interface{}) {
	o.Packages = v
}

// GetBiarchOnly returns the BiarchOnly field value if set, zero value otherwise.
func (o *RpmPackageGroupResponse) GetBiarchOnly() bool {
	if o == nil || IsNil(o.BiarchOnly) {
		var ret bool
		return ret
	}
	return *o.BiarchOnly
}

// GetBiarchOnlyOk returns a tuple with the BiarchOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetBiarchOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.BiarchOnly) {
		return nil, false
	}
	return o.BiarchOnly, true
}

// HasBiarchOnly returns a boolean if a field has been set.
func (o *RpmPackageGroupResponse) HasBiarchOnly() bool {
	if o != nil && !IsNil(o.BiarchOnly) {
		return true
	}

	return false
}

// SetBiarchOnly gets a reference to the given bool and assigns it to the BiarchOnly field.
func (o *RpmPackageGroupResponse) SetBiarchOnly(v bool) {
	o.BiarchOnly = &v
}

// GetDescByLang returns the DescByLang field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageGroupResponse) GetDescByLang() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DescByLang
}

// GetDescByLangOk returns a tuple with the DescByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageGroupResponse) GetDescByLangOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DescByLang) {
		return map[string]interface{}{}, false
	}
	return o.DescByLang, true
}

// SetDescByLang sets field value
func (o *RpmPackageGroupResponse) SetDescByLang(v map[string]interface{}) {
	o.DescByLang = v
}

// GetNameByLang returns the NameByLang field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageGroupResponse) GetNameByLang() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.NameByLang
}

// GetNameByLangOk returns a tuple with the NameByLang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageGroupResponse) GetNameByLangOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NameByLang) {
		return map[string]interface{}{}, false
	}
	return o.NameByLang, true
}

// SetNameByLang sets field value
func (o *RpmPackageGroupResponse) SetNameByLang(v map[string]interface{}) {
	o.NameByLang = v
}

// GetDigest returns the Digest field value
func (o *RpmPackageGroupResponse) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *RpmPackageGroupResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *RpmPackageGroupResponse) SetDigest(v string) {
	o.Digest = v
}

func (o RpmPackageGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmPackageGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["id"] = o.Id
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.UserVisible) {
		toSerialize["user_visible"] = o.UserVisible
	}
	toSerialize["display_order"] = o.DisplayOrder.Get()
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.BiarchOnly) {
		toSerialize["biarch_only"] = o.BiarchOnly
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

type NullableRpmPackageGroupResponse struct {
	value *RpmPackageGroupResponse
	isSet bool
}

func (v NullableRpmPackageGroupResponse) Get() *RpmPackageGroupResponse {
	return v.value
}

func (v *NullableRpmPackageGroupResponse) Set(val *RpmPackageGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmPackageGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmPackageGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmPackageGroupResponse(val *RpmPackageGroupResponse) *NullableRpmPackageGroupResponse {
	return &NullableRpmPackageGroupResponse{value: val, isSet: true}
}

func (v NullableRpmPackageGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmPackageGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


