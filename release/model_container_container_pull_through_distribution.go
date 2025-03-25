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

// checks if the ContainerContainerPullThroughDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerPullThroughDistribution{}

// ContainerContainerPullThroughDistribution A serializer for a specialized pull-through distribution referencing sub-distributions.
type ContainerContainerPullThroughDistribution struct {
	// The latest RepositoryVersion for this Repository will be served.
	Repository NullableString `json:"repository,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Whether this distribution should be shown in the content app.
	Hidden *bool `json:"hidden,omitempty"`
	// A unique name. Ex, `rawhide` and `stable`.
	Name string `json:"name"`
	// The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \"foo\" and \"foo/bar\")
	BasePath string `json:"base_path"`
	// An optional content-guard. If none is specified, a default one will be used.
	ContentGuard *string `json:"content_guard,omitempty"`
	// Remote that can be used to fetch content when using pull-through caching.
	Remote string `json:"remote"`
	// Distributions created after pulling content through cache
	Distributions []string `json:"distributions,omitempty"`
	// Restrict pull access to explicitly authorized users. Related distributions inherit this value. Defaults to unrestricted pull access.
	Private *bool `json:"private,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerContainerPullThroughDistribution ContainerContainerPullThroughDistribution

// NewContainerContainerPullThroughDistribution instantiates a new ContainerContainerPullThroughDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerPullThroughDistribution(name string, basePath string, remote string) *ContainerContainerPullThroughDistribution {
	this := ContainerContainerPullThroughDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	this.Name = name
	this.BasePath = basePath
	this.Remote = remote
	return &this
}

// NewContainerContainerPullThroughDistributionWithDefaults instantiates a new ContainerContainerPullThroughDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerPullThroughDistributionWithDefaults() *ContainerContainerPullThroughDistribution {
	this := ContainerContainerPullThroughDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPullThroughDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPullThroughDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *ContainerContainerPullThroughDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *ContainerContainerPullThroughDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *ContainerContainerPullThroughDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerPullThroughDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistribution) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *ContainerContainerPullThroughDistribution) SetHidden(v bool) {
	o.Hidden = &v
}

// GetName returns the Name field value
func (o *ContainerContainerPullThroughDistribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerPullThroughDistribution) SetName(v string) {
	o.Name = v
}

// GetBasePath returns the BasePath field value
func (o *ContainerContainerPullThroughDistribution) GetBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePath, true
}

// SetBasePath sets field value
func (o *ContainerContainerPullThroughDistribution) SetBasePath(v string) {
	o.BasePath = v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard) {
		var ret string
		return ret
	}
	return *o.ContentGuard
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil || IsNil(o.ContentGuard) {
		return nil, false
	}
	return o.ContentGuard, true
}

// HasContentGuard returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasContentGuard() bool {
	if o != nil && !IsNil(o.ContentGuard) {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given string and assigns it to the ContentGuard field.
func (o *ContainerContainerPullThroughDistribution) SetContentGuard(v string) {
	o.ContentGuard = &v
}

// GetRemote returns the Remote field value
func (o *ContainerContainerPullThroughDistribution) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *ContainerContainerPullThroughDistribution) SetRemote(v string) {
	o.Remote = v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistribution) GetDistributions() []string {
	if o == nil || IsNil(o.Distributions) {
		var ret []string
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetDistributionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Distributions) {
		return nil, false
	}
	return o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasDistributions() bool {
	if o != nil && !IsNil(o.Distributions) {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []string and assigns it to the Distributions field.
func (o *ContainerContainerPullThroughDistribution) SetDistributions(v []string) {
	o.Distributions = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ContainerContainerPullThroughDistribution) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPullThroughDistribution) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ContainerContainerPullThroughDistribution) SetPrivate(v bool) {
	o.Private = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPullThroughDistribution) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPullThroughDistribution) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerPullThroughDistribution) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerPullThroughDistribution) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerPullThroughDistribution) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerPullThroughDistribution) UnsetDescription() {
	o.Description.Unset()
}

func (o ContainerContainerPullThroughDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerPullThroughDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	toSerialize["name"] = o.Name
	toSerialize["base_path"] = o.BasePath
	if !IsNil(o.ContentGuard) {
		toSerialize["content_guard"] = o.ContentGuard
	}
	toSerialize["remote"] = o.Remote
	if !IsNil(o.Distributions) {
		toSerialize["distributions"] = o.Distributions
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerContainerPullThroughDistribution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"base_path",
		"remote",
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

	varContainerContainerPullThroughDistribution := _ContainerContainerPullThroughDistribution{}

	err = json.Unmarshal(data, &varContainerContainerPullThroughDistribution)

	if err != nil {
		return err
	}

	*o = ContainerContainerPullThroughDistribution(varContainerContainerPullThroughDistribution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "name")
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "remote")
		delete(additionalProperties, "distributions")
		delete(additionalProperties, "private")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerContainerPullThroughDistribution struct {
	value *ContainerContainerPullThroughDistribution
	isSet bool
}

func (v NullableContainerContainerPullThroughDistribution) Get() *ContainerContainerPullThroughDistribution {
	return v.value
}

func (v *NullableContainerContainerPullThroughDistribution) Set(val *ContainerContainerPullThroughDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerPullThroughDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerPullThroughDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerPullThroughDistribution(val *ContainerContainerPullThroughDistribution) *NullableContainerContainerPullThroughDistribution {
	return &NullableContainerContainerPullThroughDistribution{value: val, isSet: true}
}

func (v NullableContainerContainerPullThroughDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerPullThroughDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


