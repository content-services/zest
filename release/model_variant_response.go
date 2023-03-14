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

// checks if the VariantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantResponse{}

// VariantResponse Variant serializer.
type VariantResponse struct {
	// Variant id.
	VariantId string `json:"variant_id"`
	// Variant uid.
	Uid string `json:"uid"`
	// Variant name.
	Name string `json:"name"`
	// Variant type.
	Type string `json:"type"`
	// Relative path to directory with binary RPMs.
	Packages string `json:"packages"`
	// Relative path to directory with source RPMs.
	SourcePackages NullableString `json:"source_packages"`
	// Relative path to YUM repository with source RPMs.
	SourceRepository NullableString `json:"source_repository"`
	// Relative path to directory with debug RPMs.
	DebugPackages NullableString `json:"debug_packages"`
	// Relative path to YUM repository with debug RPMs.
	DebugRepository NullableString `json:"debug_repository"`
	// Relative path to a pem file that identifies a product.
	Identity NullableString `json:"identity"`
}

// NewVariantResponse instantiates a new VariantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantResponse(variantId string, uid string, name string, type_ string, packages string, sourcePackages NullableString, sourceRepository NullableString, debugPackages NullableString, debugRepository NullableString, identity NullableString) *VariantResponse {
	this := VariantResponse{}
	this.VariantId = variantId
	this.Uid = uid
	this.Name = name
	this.Type = type_
	this.Packages = packages
	this.SourcePackages = sourcePackages
	this.SourceRepository = sourceRepository
	this.DebugPackages = debugPackages
	this.DebugRepository = debugRepository
	this.Identity = identity
	return &this
}

// NewVariantResponseWithDefaults instantiates a new VariantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantResponseWithDefaults() *VariantResponse {
	this := VariantResponse{}
	return &this
}

// GetVariantId returns the VariantId field value
func (o *VariantResponse) GetVariantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value
// and a boolean to check if the value has been set.
func (o *VariantResponse) GetVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantId, true
}

// SetVariantId sets field value
func (o *VariantResponse) SetVariantId(v string) {
	o.VariantId = v
}

// GetUid returns the Uid field value
func (o *VariantResponse) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *VariantResponse) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *VariantResponse) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *VariantResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariantResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariantResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *VariantResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VariantResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VariantResponse) SetType(v string) {
	o.Type = v
}

// GetPackages returns the Packages field value
func (o *VariantResponse) GetPackages() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *VariantResponse) GetPackagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packages, true
}

// SetPackages sets field value
func (o *VariantResponse) SetPackages(v string) {
	o.Packages = v
}

// GetSourcePackages returns the SourcePackages field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariantResponse) GetSourcePackages() string {
	if o == nil || o.SourcePackages.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourcePackages.Get()
}

// GetSourcePackagesOk returns a tuple with the SourcePackages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariantResponse) GetSourcePackagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourcePackages.Get(), o.SourcePackages.IsSet()
}

// SetSourcePackages sets field value
func (o *VariantResponse) SetSourcePackages(v string) {
	o.SourcePackages.Set(&v)
}

// GetSourceRepository returns the SourceRepository field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariantResponse) GetSourceRepository() string {
	if o == nil || o.SourceRepository.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourceRepository.Get()
}

// GetSourceRepositoryOk returns a tuple with the SourceRepository field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariantResponse) GetSourceRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceRepository.Get(), o.SourceRepository.IsSet()
}

// SetSourceRepository sets field value
func (o *VariantResponse) SetSourceRepository(v string) {
	o.SourceRepository.Set(&v)
}

// GetDebugPackages returns the DebugPackages field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariantResponse) GetDebugPackages() string {
	if o == nil || o.DebugPackages.Get() == nil {
		var ret string
		return ret
	}

	return *o.DebugPackages.Get()
}

// GetDebugPackagesOk returns a tuple with the DebugPackages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariantResponse) GetDebugPackagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugPackages.Get(), o.DebugPackages.IsSet()
}

// SetDebugPackages sets field value
func (o *VariantResponse) SetDebugPackages(v string) {
	o.DebugPackages.Set(&v)
}

// GetDebugRepository returns the DebugRepository field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariantResponse) GetDebugRepository() string {
	if o == nil || o.DebugRepository.Get() == nil {
		var ret string
		return ret
	}

	return *o.DebugRepository.Get()
}

// GetDebugRepositoryOk returns a tuple with the DebugRepository field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariantResponse) GetDebugRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugRepository.Get(), o.DebugRepository.IsSet()
}

// SetDebugRepository sets field value
func (o *VariantResponse) SetDebugRepository(v string) {
	o.DebugRepository.Set(&v)
}

// GetIdentity returns the Identity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VariantResponse) GetIdentity() string {
	if o == nil || o.Identity.Get() == nil {
		var ret string
		return ret
	}

	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariantResponse) GetIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// SetIdentity sets field value
func (o *VariantResponse) SetIdentity(v string) {
	o.Identity.Set(&v)
}

func (o VariantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant_id"] = o.VariantId
	toSerialize["uid"] = o.Uid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["packages"] = o.Packages
	toSerialize["source_packages"] = o.SourcePackages.Get()
	toSerialize["source_repository"] = o.SourceRepository.Get()
	toSerialize["debug_packages"] = o.DebugPackages.Get()
	toSerialize["debug_repository"] = o.DebugRepository.Get()
	toSerialize["identity"] = o.Identity.Get()
	return toSerialize, nil
}

type NullableVariantResponse struct {
	value *VariantResponse
	isSet bool
}

func (v NullableVariantResponse) Get() *VariantResponse {
	return v.value
}

func (v *NullableVariantResponse) Set(val *VariantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantResponse(val *VariantResponse) *NullableVariantResponse {
	return &NullableVariantResponse{value: val, isSet: true}
}

func (v NullableVariantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

