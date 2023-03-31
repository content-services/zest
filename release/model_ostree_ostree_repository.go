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

// checks if the OstreeOstreeRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeRepository{}

// OstreeOstreeRepository A Serializer class for an OSTree repository.
type OstreeOstreeRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	ComputeDelta *bool `json:"compute_delta,omitempty"`
}

// NewOstreeOstreeRepository instantiates a new OstreeOstreeRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeRepository(name string) *OstreeOstreeRepository {
	this := OstreeOstreeRepository{}
	this.Name = name
	var computeDelta bool = true
	this.ComputeDelta = &computeDelta
	return &this
}

// NewOstreeOstreeRepositoryWithDefaults instantiates a new OstreeOstreeRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeRepositoryWithDefaults() *OstreeOstreeRepository {
	this := OstreeOstreeRepository{}
	var computeDelta bool = true
	this.ComputeDelta = &computeDelta
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *OstreeOstreeRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *OstreeOstreeRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *OstreeOstreeRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *OstreeOstreeRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OstreeOstreeRepository) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OstreeOstreeRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OstreeOstreeRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OstreeOstreeRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OstreeOstreeRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeRepository) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeRepository) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *OstreeOstreeRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *OstreeOstreeRepository) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *OstreeOstreeRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *OstreeOstreeRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OstreeOstreeRepository) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OstreeOstreeRepository) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *OstreeOstreeRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *OstreeOstreeRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *OstreeOstreeRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *OstreeOstreeRepository) UnsetRemote() {
	o.Remote.Unset()
}

// GetComputeDelta returns the ComputeDelta field value if set, zero value otherwise.
func (o *OstreeOstreeRepository) GetComputeDelta() bool {
	if o == nil || IsNil(o.ComputeDelta) {
		var ret bool
		return ret
	}
	return *o.ComputeDelta
}

// GetComputeDeltaOk returns a tuple with the ComputeDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeRepository) GetComputeDeltaOk() (*bool, bool) {
	if o == nil || IsNil(o.ComputeDelta) {
		return nil, false
	}
	return o.ComputeDelta, true
}

// HasComputeDelta returns a boolean if a field has been set.
func (o *OstreeOstreeRepository) HasComputeDelta() bool {
	if o != nil && !IsNil(o.ComputeDelta) {
		return true
	}

	return false
}

// SetComputeDelta gets a reference to the given bool and assigns it to the ComputeDelta field.
func (o *OstreeOstreeRepository) SetComputeDelta(v bool) {
	o.ComputeDelta = &v
}

func (o OstreeOstreeRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}
	if !IsNil(o.ComputeDelta) {
		toSerialize["compute_delta"] = o.ComputeDelta
	}
	return toSerialize, nil
}

type NullableOstreeOstreeRepository struct {
	value *OstreeOstreeRepository
	isSet bool
}

func (v NullableOstreeOstreeRepository) Get() *OstreeOstreeRepository {
	return v.value
}

func (v *NullableOstreeOstreeRepository) Set(val *OstreeOstreeRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeRepository(val *OstreeOstreeRepository) *NullableOstreeOstreeRepository {
	return &NullableOstreeOstreeRepository{value: val, isSet: true}
}

func (v NullableOstreeOstreeRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


