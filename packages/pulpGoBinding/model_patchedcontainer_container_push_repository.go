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
)

// checks if the PatchedcontainerContainerPushRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedcontainerContainerPushRepository{}

// PatchedcontainerContainerPushRepository Serializer for Container Push Repositories.
type PatchedcontainerContainerPushRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future.
	RetainRepoVersions NullableInt32 `json:"retain_repo_versions,omitempty"`
	// A unique name for this repository.
	Name *string `json:"name,omitempty"`
	// A reference to an associated signing service.
	ManifestSigningService NullableString `json:"manifest_signing_service,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewPatchedcontainerContainerPushRepository instantiates a new PatchedcontainerContainerPushRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedcontainerContainerPushRepository() *PatchedcontainerContainerPushRepository {
	this := PatchedcontainerContainerPushRepository{}
	return &this
}

// NewPatchedcontainerContainerPushRepositoryWithDefaults instantiates a new PatchedcontainerContainerPushRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedcontainerContainerPushRepositoryWithDefaults() *PatchedcontainerContainerPushRepository {
	this := PatchedcontainerContainerPushRepository{}
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedcontainerContainerPushRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcontainerContainerPushRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedcontainerContainerPushRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedcontainerContainerPushRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerPushRepository) GetRetainRepoVersions() int32 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int32
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerPushRepository) GetRetainRepoVersionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *PatchedcontainerContainerPushRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt32 and assigns it to the RetainRepoVersions field.
func (o *PatchedcontainerContainerPushRepository) SetRetainRepoVersions(v int32) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *PatchedcontainerContainerPushRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *PatchedcontainerContainerPushRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedcontainerContainerPushRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcontainerContainerPushRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedcontainerContainerPushRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedcontainerContainerPushRepository) SetName(v string) {
	o.Name = &v
}

// GetManifestSigningService returns the ManifestSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerPushRepository) GetManifestSigningService() string {
	if o == nil || IsNil(o.ManifestSigningService.Get()) {
		var ret string
		return ret
	}
	return *o.ManifestSigningService.Get()
}

// GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerPushRepository) GetManifestSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManifestSigningService.Get(), o.ManifestSigningService.IsSet()
}

// HasManifestSigningService returns a boolean if a field has been set.
func (o *PatchedcontainerContainerPushRepository) HasManifestSigningService() bool {
	if o != nil && o.ManifestSigningService.IsSet() {
		return true
	}

	return false
}

// SetManifestSigningService gets a reference to the given NullableString and assigns it to the ManifestSigningService field.
func (o *PatchedcontainerContainerPushRepository) SetManifestSigningService(v string) {
	o.ManifestSigningService.Set(&v)
}
// SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil
func (o *PatchedcontainerContainerPushRepository) SetManifestSigningServiceNil() {
	o.ManifestSigningService.Set(nil)
}

// UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
func (o *PatchedcontainerContainerPushRepository) UnsetManifestSigningService() {
	o.ManifestSigningService.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcontainerContainerPushRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcontainerContainerPushRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedcontainerContainerPushRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedcontainerContainerPushRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedcontainerContainerPushRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedcontainerContainerPushRepository) UnsetDescription() {
	o.Description.Unset()
}

func (o PatchedcontainerContainerPushRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedcontainerContainerPushRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ManifestSigningService.IsSet() {
		toSerialize["manifest_signing_service"] = o.ManifestSigningService.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullablePatchedcontainerContainerPushRepository struct {
	value *PatchedcontainerContainerPushRepository
	isSet bool
}

func (v NullablePatchedcontainerContainerPushRepository) Get() *PatchedcontainerContainerPushRepository {
	return v.value
}

func (v *NullablePatchedcontainerContainerPushRepository) Set(val *PatchedcontainerContainerPushRepository) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedcontainerContainerPushRepository) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedcontainerContainerPushRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedcontainerContainerPushRepository(val *PatchedcontainerContainerPushRepository) *NullablePatchedcontainerContainerPushRepository {
	return &NullablePatchedcontainerContainerPushRepository{value: val, isSet: true}
}

func (v NullablePatchedcontainerContainerPushRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedcontainerContainerPushRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

