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

// checks if the PatchedcontainerContainerRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedcontainerContainerRepository{}

// PatchedcontainerContainerRepository Serializer for Container Repositories.
type PatchedcontainerContainerRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// A reference to an associated signing service.
	ManifestSigningService NullableString `json:"manifest_signing_service,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedcontainerContainerRepository PatchedcontainerContainerRepository

// NewPatchedcontainerContainerRepository instantiates a new PatchedcontainerContainerRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedcontainerContainerRepository() *PatchedcontainerContainerRepository {
	this := PatchedcontainerContainerRepository{}
	return &this
}

// NewPatchedcontainerContainerRepositoryWithDefaults instantiates a new PatchedcontainerContainerRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedcontainerContainerRepositoryWithDefaults() *PatchedcontainerContainerRepository {
	this := PatchedcontainerContainerRepository{}
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedcontainerContainerRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcontainerContainerRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedcontainerContainerRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedcontainerContainerRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcontainerContainerRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedcontainerContainerRepository) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedcontainerContainerRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedcontainerContainerRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedcontainerContainerRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerRepository) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerRepository) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *PatchedcontainerContainerRepository) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *PatchedcontainerContainerRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *PatchedcontainerContainerRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerRepository) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerRepository) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *PatchedcontainerContainerRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *PatchedcontainerContainerRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *PatchedcontainerContainerRepository) UnsetRemote() {
	o.Remote.Unset()
}

// GetManifestSigningService returns the ManifestSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerRepository) GetManifestSigningService() string {
	if o == nil || IsNil(o.ManifestSigningService.Get()) {
		var ret string
		return ret
	}
	return *o.ManifestSigningService.Get()
}

// GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerRepository) GetManifestSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManifestSigningService.Get(), o.ManifestSigningService.IsSet()
}

// HasManifestSigningService returns a boolean if a field has been set.
func (o *PatchedcontainerContainerRepository) HasManifestSigningService() bool {
	if o != nil && o.ManifestSigningService.IsSet() {
		return true
	}

	return false
}

// SetManifestSigningService gets a reference to the given NullableString and assigns it to the ManifestSigningService field.
func (o *PatchedcontainerContainerRepository) SetManifestSigningService(v string) {
	o.ManifestSigningService.Set(&v)
}
// SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil
func (o *PatchedcontainerContainerRepository) SetManifestSigningServiceNil() {
	o.ManifestSigningService.Set(nil)
}

// UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
func (o *PatchedcontainerContainerRepository) UnsetManifestSigningService() {
	o.ManifestSigningService.Unset()
}

func (o PatchedcontainerContainerRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedcontainerContainerRepository) ToMap() (map[string]interface{}, error) {
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
	if o.ManifestSigningService.IsSet() {
		toSerialize["manifest_signing_service"] = o.ManifestSigningService.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedcontainerContainerRepository) UnmarshalJSON(data []byte) (err error) {
	varPatchedcontainerContainerRepository := _PatchedcontainerContainerRepository{}

	err = json.Unmarshal(data, &varPatchedcontainerContainerRepository)

	if err != nil {
		return err
	}

	*o = PatchedcontainerContainerRepository(varPatchedcontainerContainerRepository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "retain_repo_versions")
		delete(additionalProperties, "remote")
		delete(additionalProperties, "manifest_signing_service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedcontainerContainerRepository struct {
	value *PatchedcontainerContainerRepository
	isSet bool
}

func (v NullablePatchedcontainerContainerRepository) Get() *PatchedcontainerContainerRepository {
	return v.value
}

func (v *NullablePatchedcontainerContainerRepository) Set(val *PatchedcontainerContainerRepository) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedcontainerContainerRepository) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedcontainerContainerRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedcontainerContainerRepository(val *PatchedcontainerContainerRepository) *NullablePatchedcontainerContainerRepository {
	return &NullablePatchedcontainerContainerRepository{value: val, isSet: true}
}

func (v NullablePatchedcontainerContainerRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedcontainerContainerRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


