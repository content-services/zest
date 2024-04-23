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

// checks if the PatchedostreeOstreeRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedostreeOstreeRepository{}

// PatchedostreeOstreeRepository A Serializer class for an OSTree repository.
type PatchedostreeOstreeRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	ComputeDelta *bool `json:"compute_delta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedostreeOstreeRepository PatchedostreeOstreeRepository

// NewPatchedostreeOstreeRepository instantiates a new PatchedostreeOstreeRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedostreeOstreeRepository() *PatchedostreeOstreeRepository {
	this := PatchedostreeOstreeRepository{}
	var computeDelta bool = true
	this.ComputeDelta = &computeDelta
	return &this
}

// NewPatchedostreeOstreeRepositoryWithDefaults instantiates a new PatchedostreeOstreeRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedostreeOstreeRepositoryWithDefaults() *PatchedostreeOstreeRepository {
	this := PatchedostreeOstreeRepository{}
	var computeDelta bool = true
	this.ComputeDelta = &computeDelta
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedostreeOstreeRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedostreeOstreeRepository) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedostreeOstreeRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedostreeOstreeRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedostreeOstreeRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRepository) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRepository) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *PatchedostreeOstreeRepository) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *PatchedostreeOstreeRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *PatchedostreeOstreeRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedostreeOstreeRepository) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedostreeOstreeRepository) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *PatchedostreeOstreeRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *PatchedostreeOstreeRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *PatchedostreeOstreeRepository) UnsetRemote() {
	o.Remote.Unset()
}

// GetComputeDelta returns the ComputeDelta field value if set, zero value otherwise.
func (o *PatchedostreeOstreeRepository) GetComputeDelta() bool {
	if o == nil || IsNil(o.ComputeDelta) {
		var ret bool
		return ret
	}
	return *o.ComputeDelta
}

// GetComputeDeltaOk returns a tuple with the ComputeDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedostreeOstreeRepository) GetComputeDeltaOk() (*bool, bool) {
	if o == nil || IsNil(o.ComputeDelta) {
		return nil, false
	}
	return o.ComputeDelta, true
}

// HasComputeDelta returns a boolean if a field has been set.
func (o *PatchedostreeOstreeRepository) HasComputeDelta() bool {
	if o != nil && !IsNil(o.ComputeDelta) {
		return true
	}

	return false
}

// SetComputeDelta gets a reference to the given bool and assigns it to the ComputeDelta field.
func (o *PatchedostreeOstreeRepository) SetComputeDelta(v bool) {
	o.ComputeDelta = &v
}

func (o PatchedostreeOstreeRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedostreeOstreeRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedostreeOstreeRepository) UnmarshalJSON(data []byte) (err error) {
	varPatchedostreeOstreeRepository := _PatchedostreeOstreeRepository{}

	err = json.Unmarshal(data, &varPatchedostreeOstreeRepository)

	if err != nil {
		return err
	}

	*o = PatchedostreeOstreeRepository(varPatchedostreeOstreeRepository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "retain_repo_versions")
		delete(additionalProperties, "remote")
		delete(additionalProperties, "compute_delta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedostreeOstreeRepository struct {
	value *PatchedostreeOstreeRepository
	isSet bool
}

func (v NullablePatchedostreeOstreeRepository) Get() *PatchedostreeOstreeRepository {
	return v.value
}

func (v *NullablePatchedostreeOstreeRepository) Set(val *PatchedostreeOstreeRepository) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedostreeOstreeRepository) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedostreeOstreeRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedostreeOstreeRepository(val *PatchedostreeOstreeRepository) *NullablePatchedostreeOstreeRepository {
	return &NullablePatchedostreeOstreeRepository{value: val, isSet: true}
}

func (v NullablePatchedostreeOstreeRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedostreeOstreeRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


