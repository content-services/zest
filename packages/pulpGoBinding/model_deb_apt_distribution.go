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

// checks if the DebAptDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebAptDistribution{}

// DebAptDistribution Serializer for AptDistributions.
type DebAptDistribution struct {
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// Publication to be served
	Publication NullableString `json:"publication,omitempty"`
}

// NewDebAptDistribution instantiates a new DebAptDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebAptDistribution(basePath string, name string) *DebAptDistribution {
	this := DebAptDistribution{}
	this.BasePath = basePath
	this.Name = name
	return &this
}

// NewDebAptDistributionWithDefaults instantiates a new DebAptDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebAptDistributionWithDefaults() *DebAptDistribution {
	this := DebAptDistribution{}
	return &this
}

// GetBasePath returns the BasePath field value
func (o *DebAptDistribution) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *DebAptDistribution) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *DebAptDistribution) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *DebAptDistribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *DebAptDistribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *DebAptDistribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *DebAptDistribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *DebAptDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebAptDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *DebAptDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *DebAptDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *DebAptDistribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DebAptDistribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DebAptDistribution) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *DebAptDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *DebAptDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *DebAptDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *DebAptDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebAptDistribution) GetPublication() string {
	if o == nil || IsNil(o.Publication.Get()) {
		var ret string
		return ret
	}
	return *o.Publication.Get()
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebAptDistribution) GetPublicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Publication.Get(), o.Publication.IsSet()
}

// HasPublication returns a boolean if a field has been set.
func (o *DebAptDistribution) HasPublication() bool {
	if o != nil && o.Publication.IsSet() {
		return true
	}

	return false
}

// SetPublication gets a reference to the given NullableString and assigns it to the Publication field.
func (o *DebAptDistribution) SetPublication(v string) {
	o.Publication.Set(&v)
}
// SetPublicationNil sets the value for Publication to be an explicit nil
func (o *DebAptDistribution) SetPublicationNil() {
	o.Publication.Set(nil)
}

// UnsetPublication ensures that no value is present for Publication, not even an explicit nil
func (o *DebAptDistribution) UnsetPublication() {
	o.Publication.Unset()
}

func (o DebAptDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebAptDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_path"] = o.BasePath
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["name"] = o.Name
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Publication.IsSet() {
		toSerialize["publication"] = o.Publication.Get()
	}
	return toSerialize, nil
}

type NullableDebAptDistribution struct {
	value *DebAptDistribution
	isSet bool
}

func (v NullableDebAptDistribution) Get() *DebAptDistribution {
	return v.value
}

func (v *NullableDebAptDistribution) Set(val *DebAptDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableDebAptDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableDebAptDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebAptDistribution(val *DebAptDistribution) *NullableDebAptDistribution {
	return &NullableDebAptDistribution{value: val, isSet: true}
}

func (v NullableDebAptDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebAptDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


