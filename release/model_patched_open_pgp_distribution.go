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

// checks if the PatchedOpenPGPDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedOpenPGPDistribution{}

// PatchedOpenPGPDistribution The Serializer for the Distribution model.The serializer deliberately omits the `publication` and `repository_version` field due toplugins typically requiring one or the other but not both.To include the ``publication`` field, it is recommended plugins define the field::  publication = DetailRelatedField(      required=False,      help_text=_(\"Publication to be served\"),      view_name_pattern=r\"publications(-.*_/.*)?-detail\",      queryset=models.Publication.objects.exclude(complete=False),      allow_null=True,  )To include the ``repository_version`` field, it is recommended plugins define the field::  repository_version = RepositoryVersionRelatedField(      required=False, help_text=_(\"RepositoryVersion to be served\"), allow_null=True  )Additionally, the serializer omits the ``remote`` field, which is used for pull-through cachingfeature and only by plugins which use publications. Plugins implementing a pull-through cachingshould define the field in their derived serializer class like this::  remote = DetailRelatedField(      required=False,      help_text=_('Remote that can be used to fetch content when using pull-through caching.'),      queryset=models.Remote.objects.all(),      allow_null=True  )
type PatchedOpenPGPDistribution struct {
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
	// RepositoryVersion to be served
	RepositoryVersion NullableString `json:"repository_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedOpenPGPDistribution PatchedOpenPGPDistribution

// NewPatchedOpenPGPDistribution instantiates a new PatchedOpenPGPDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOpenPGPDistribution() *PatchedOpenPGPDistribution {
	this := PatchedOpenPGPDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// NewPatchedOpenPGPDistributionWithDefaults instantiates a new PatchedOpenPGPDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOpenPGPDistributionWithDefaults() *PatchedOpenPGPDistribution {
	this := PatchedOpenPGPDistribution{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *PatchedOpenPGPDistribution) GetBasePath() string {
	if o == nil || IsNil(o.BasePath) {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOpenPGPDistribution) GetBasePathOk() (*string, bool) {
	if o == nil || IsNil(o.BasePath) {
		return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasBasePath() bool {
	if o != nil && !IsNil(o.BasePath) {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *PatchedOpenPGPDistribution) SetBasePath(v string) {
	o.BasePath = &v
}

// GetContentGuard returns the ContentGuard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOpenPGPDistribution) GetContentGuard() string {
	if o == nil || IsNil(o.ContentGuard.Get()) {
		var ret string
		return ret
	}
	return *o.ContentGuard.Get()
}

// GetContentGuardOk returns a tuple with the ContentGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOpenPGPDistribution) GetContentGuardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentGuard.Get(), o.ContentGuard.IsSet()
}

// HasContentGuard returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasContentGuard() bool {
	if o != nil && o.ContentGuard.IsSet() {
		return true
	}

	return false
}

// SetContentGuard gets a reference to the given NullableString and assigns it to the ContentGuard field.
func (o *PatchedOpenPGPDistribution) SetContentGuard(v string) {
	o.ContentGuard.Set(&v)
}
// SetContentGuardNil sets the value for ContentGuard to be an explicit nil
func (o *PatchedOpenPGPDistribution) SetContentGuardNil() {
	o.ContentGuard.Set(nil)
}

// UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
func (o *PatchedOpenPGPDistribution) UnsetContentGuard() {
	o.ContentGuard.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *PatchedOpenPGPDistribution) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOpenPGPDistribution) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *PatchedOpenPGPDistribution) SetHidden(v bool) {
	o.Hidden = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedOpenPGPDistribution) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOpenPGPDistribution) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedOpenPGPDistribution) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOpenPGPDistribution) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOpenPGPDistribution) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOpenPGPDistribution) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOpenPGPDistribution) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOpenPGPDistribution) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *PatchedOpenPGPDistribution) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *PatchedOpenPGPDistribution) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *PatchedOpenPGPDistribution) UnsetRepository() {
	o.Repository.Unset()
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOpenPGPDistribution) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion.Get()
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOpenPGPDistribution) GetRepositoryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryVersion.Get(), o.RepositoryVersion.IsSet()
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *PatchedOpenPGPDistribution) HasRepositoryVersion() bool {
	if o != nil && o.RepositoryVersion.IsSet() {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given NullableString and assigns it to the RepositoryVersion field.
func (o *PatchedOpenPGPDistribution) SetRepositoryVersion(v string) {
	o.RepositoryVersion.Set(&v)
}
// SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil
func (o *PatchedOpenPGPDistribution) SetRepositoryVersionNil() {
	o.RepositoryVersion.Set(nil)
}

// UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
func (o *PatchedOpenPGPDistribution) UnsetRepositoryVersion() {
	o.RepositoryVersion.Unset()
}

func (o PatchedOpenPGPDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedOpenPGPDistribution) ToMap() (map[string]interface{}, error) {
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
	if o.RepositoryVersion.IsSet() {
		toSerialize["repository_version"] = o.RepositoryVersion.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedOpenPGPDistribution) UnmarshalJSON(data []byte) (err error) {
	varPatchedOpenPGPDistribution := _PatchedOpenPGPDistribution{}

	err = json.Unmarshal(data, &varPatchedOpenPGPDistribution)

	if err != nil {
		return err
	}

	*o = PatchedOpenPGPDistribution(varPatchedOpenPGPDistribution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base_path")
		delete(additionalProperties, "content_guard")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "repository_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedOpenPGPDistribution struct {
	value *PatchedOpenPGPDistribution
	isSet bool
}

func (v NullablePatchedOpenPGPDistribution) Get() *PatchedOpenPGPDistribution {
	return v.value
}

func (v *NullablePatchedOpenPGPDistribution) Set(val *PatchedOpenPGPDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOpenPGPDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOpenPGPDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOpenPGPDistribution(val *PatchedOpenPGPDistribution) *NullablePatchedOpenPGPDistribution {
	return &NullablePatchedOpenPGPDistribution{value: val, isSet: true}
}

func (v NullablePatchedOpenPGPDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOpenPGPDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


