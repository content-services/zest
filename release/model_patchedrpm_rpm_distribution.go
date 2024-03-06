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

// checks if the PatchedrpmRpmDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedrpmRpmDistribution{}

// PatchedrpmRpmDistribution Serializer for RPM Distributions.
type PatchedrpmRpmDistribution struct {
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath *string `json:"base_path,omitempty"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name *string `json:"name,omitempty"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// Publication to be served
	Publication NullableString `json:"publication,omitempty"`
	// An option specifying whether Pulp should generate *.repo files.
	GenerateRepoConfig *bool `json:"generate_repo_config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedrpmRpmDistribution PatchedrpmRpmDistribution

// NewPatchedrpmRpmDistribution instantiates a new PatchedrpmRpmDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedrpmRpmDistribution() *PatchedrpmRpmDistribution {
	this := PatchedrpmRpmDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	var generateRepoConfig bool = false
	this.GenerateRepoConfig = &generateRepoConfig
	return &this
}

// NewPatchedrpmRpmDistributionWithDefaults instantiates a new PatchedrpmRpmDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedrpmRpmDistributionWithDefaults() *PatchedrpmRpmDistribution {
	this := PatchedrpmRpmDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	var generateRepoConfig bool = false
	this.GenerateRepoConfig = &generateRepoConfig
	return &this
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *PatchedrpmRpmDistribution) GetBasePath() string {
	if o == nil || IsNil(o.BasePath) {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmDistribution) GetBasePathOk() (*string, bool) {
	if o == nil || IsNil(o.BasePath) {
		return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasBasePath() bool {
	if o != nil && !IsNil(o.BasePath) {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *PatchedrpmRpmDistribution) SetBasePath(v string) {
	o.BasePath = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedrpmRpmDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedrpmRpmDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *PatchedrpmRpmDistribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *PatchedrpmRpmDistribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *PatchedrpmRpmDistribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *PatchedrpmRpmDistribution) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmDistribution) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *PatchedrpmRpmDistribution) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedrpmRpmDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedrpmRpmDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedrpmRpmDistribution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmDistribution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedrpmRpmDistribution) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedrpmRpmDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedrpmRpmDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *PatchedrpmRpmDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *PatchedrpmRpmDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *PatchedrpmRpmDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedrpmRpmDistribution) GetPublication() string {
	if o == nil || IsNil(o.Publication.Get()) {
		var ret string
		return ret
	}
	return *o.Publication.Get()
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedrpmRpmDistribution) GetPublicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Publication.Get(), o.Publication.IsSet()
}

// HasPublication returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasPublication() bool {
	if o != nil && o.Publication.IsSet() {
		return true
	}

	return false
}

// SetPublication gets a reference to the given NullableString and assigns it to the Publication field.
func (o *PatchedrpmRpmDistribution) SetPublication(v string) {
	o.Publication.Set(&v)
}
// SetPublicationNil sets the value for Publication to be an explicit nil
func (o *PatchedrpmRpmDistribution) SetPublicationNil() {
	o.Publication.Set(nil)
}

// UnsetPublication ensures that no value is present for Publication, not even an explicit nil
func (o *PatchedrpmRpmDistribution) UnsetPublication() {
	o.Publication.Unset()
}

// GetGenerateRepoConfig returns the GenerateRepoConfig field value if set, zero value otherwise.
func (o *PatchedrpmRpmDistribution) GetGenerateRepoConfig() bool {
	if o == nil || IsNil(o.GenerateRepoConfig) {
		var ret bool
		return ret
	}
	return *o.GenerateRepoConfig
}

// GetGenerateRepoConfigOk returns a tuple with the GenerateRepoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmDistribution) GetGenerateRepoConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateRepoConfig) {
		return nil, false
	}
	return o.GenerateRepoConfig, true
}

// HasGenerateRepoConfig returns a boolean if a field has been set.
func (o *PatchedrpmRpmDistribution) HasGenerateRepoConfig() bool {
	if o != nil && !IsNil(o.GenerateRepoConfig) {
		return true
	}

	return false
}

// SetGenerateRepoConfig gets a reference to the given bool and assigns it to the GenerateRepoConfig field.
func (o *PatchedrpmRpmDistribution) SetGenerateRepoConfig(v bool) {
	o.GenerateRepoConfig = &v
}

func (o PatchedrpmRpmDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedrpmRpmDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasePath) {
		toSerialize["base_path"] = o.BasePath
	}
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Publication.IsSet() {
		toSerialize["publication"] = o.Publication.Get()
	}
	if !IsNil(o.GenerateRepoConfig) {
		toSerialize["generate_repo_config"] = o.GenerateRepoConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedrpmRpmDistribution) UnmarshalJSON(data []byte) (err error) {
	varPatchedrpmRpmDistribution := _PatchedrpmRpmDistribution{}

	err = json.Unmarshal(data, &varPatchedrpmRpmDistribution)

	if err != nil {
		return err
	}

	*o = PatchedrpmRpmDistribution(varPatchedrpmRpmDistribution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "publication")
		delete(additionalProperties, "generate_repo_config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedrpmRpmDistribution struct {
	value *PatchedrpmRpmDistribution
	isSet bool
}

func (v NullablePatchedrpmRpmDistribution) Get() *PatchedrpmRpmDistribution {
	return v.value
}

func (v *NullablePatchedrpmRpmDistribution) Set(val *PatchedrpmRpmDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedrpmRpmDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedrpmRpmDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedrpmRpmDistribution(val *PatchedrpmRpmDistribution) *NullablePatchedrpmRpmDistribution {
	return &NullablePatchedrpmRpmDistribution{value: val, isSet: true}
}

func (v NullablePatchedrpmRpmDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedrpmRpmDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


