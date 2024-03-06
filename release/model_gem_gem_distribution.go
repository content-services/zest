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
	"fmt"
)

// checks if the GemGemDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GemGemDistribution{}

// GemGemDistribution A Serializer for GemDistribution.
type GemGemDistribution struct {
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard.
	ContentGuard NullableString `json:"content_guard,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	// Publication to be served
	Publication NullableString `json:"publication,omitempty"`
	// Remote that can be used to fetch content when using pull-through caching.
	Remote NullableString `json:"remote,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GemGemDistribution GemGemDistribution

// NewGemGemDistribution instantiates a new GemGemDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGemGemDistribution(basePath string, name string) *GemGemDistribution {
	this := GemGemDistribution{}
	this.BasePath = basePath
	var hidden bool = false
	this.Hidden = &hidden
	this.Name = name
	return &this
}

// NewGemGemDistributionWithDefaults instantiates a new GemGemDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGemGemDistributionWithDefaults() *GemGemDistribution {
	this := GemGemDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetBasePath returns the BasePath field value
func (o *GemGemDistribution) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *GemGemDistribution) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *GemGemDistribution) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GemGemDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GemGemDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *GemGemDistribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *GemGemDistribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *GemGemDistribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *GemGemDistribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *GemGemDistribution) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemDistribution) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *GemGemDistribution) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *GemGemDistribution) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *GemGemDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GemGemDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *GemGemDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *GemGemDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *GemGemDistribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GemGemDistribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GemGemDistribution) SetName(v string) {
	o.Name = v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GemGemDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GemGemDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *GemGemDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *GemGemDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *GemGemDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *GemGemDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetPublication returns the Publication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GemGemDistribution) GetPublication() string {
	if o == nil || IsNil(o.Publication.Get()) {
		var ret string
		return ret
	}
	return *o.Publication.Get()
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GemGemDistribution) GetPublicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Publication.Get(), o.Publication.IsSet()
}

// HasPublication returns a boolean if a field has been set.
func (o *GemGemDistribution) HasPublication() bool {
	if o != nil && o.Publication.IsSet() {
		return true
	}

	return false
}

// SetPublication gets a reference to the given NullableString and assigns it to the Publication field.
func (o *GemGemDistribution) SetPublication(v string) {
	o.Publication.Set(&v)
}
// SetPublicationNil sets the value for Publication to be an explicit nil
func (o *GemGemDistribution) SetPublicationNil() {
	o.Publication.Set(nil)
}

// UnsetPublication ensures that no value is present for Publication, not even an explicit nil
func (o *GemGemDistribution) UnsetPublication() {
	o.Publication.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GemGemDistribution) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GemGemDistribution) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *GemGemDistribution) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *GemGemDistribution) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *GemGemDistribution) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *GemGemDistribution) UnsetRemote() {
	o.Remote.Unset()
}

func (o GemGemDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GemGemDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_path"] = o.BasePath
	if o.ContentGuard.IsSet() {
		toSerialize["content_guard"] = o.ContentGuard.Get()
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
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
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GemGemDistribution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_path",
		"name",
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

	varGemGemDistribution := _GemGemDistribution{}

	err = json.Unmarshal(data, &varGemGemDistribution)

	if err != nil {
		return err
	}

	*o = GemGemDistribution(varGemGemDistribution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "publication")
		delete(additionalProperties, "remote")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGemGemDistribution struct {
	value *GemGemDistribution
	isSet bool
}

func (v NullableGemGemDistribution) Get() *GemGemDistribution {
	return v.value
}

func (v *NullableGemGemDistribution) Set(val *GemGemDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableGemGemDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableGemGemDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGemGemDistribution(val *GemGemDistribution) *NullableGemGemDistribution {
	return &NullableGemGemDistribution{value: val, isSet: true}
}

func (v NullableGemGemDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGemGemDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


