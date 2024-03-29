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

// checks if the FileFileRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileFileRepository{}

// FileFileRepository Serializer for File Repositories.
type FileFileRepository struct {
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository.
	Autopublish *bool `json:"autopublish,omitempty"`
	// Filename to use for manifest file containing metadata for all the files.
	Manifest NullableString `json:"manifest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileFileRepository FileFileRepository

// NewFileFileRepository instantiates a new FileFileRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileFileRepository(name string) *FileFileRepository {
	this := FileFileRepository{}
	this.Name = name
	var autopublish bool = false
	this.Autopublish = &autopublish
	var manifest string = "PULP_MANIFEST"
	this.Manifest = *NewNullableString(&manifest)
	return &this
}

// NewFileFileRepositoryWithDefaults instantiates a new FileFileRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileFileRepositoryWithDefaults() *FileFileRepository {
	this := FileFileRepository{}
	var autopublish bool = false
	this.Autopublish = &autopublish
	var manifest string = "PULP_MANIFEST"
	this.Manifest = *NewNullableString(&manifest)
	return &this
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *FileFileRepository) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileRepository) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *FileFileRepository) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *FileFileRepository) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetName returns the Name field value
func (o *FileFileRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileFileRepository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileFileRepository) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileFileRepository) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileFileRepository) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *FileFileRepository) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *FileFileRepository) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *FileFileRepository) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *FileFileRepository) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileFileRepository) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileFileRepository) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *FileFileRepository) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *FileFileRepository) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *FileFileRepository) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *FileFileRepository) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileFileRepository) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileFileRepository) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *FileFileRepository) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *FileFileRepository) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *FileFileRepository) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *FileFileRepository) UnsetRemote() {
	o.Remote.Unset()
}

// GetAutopublish returns the Autopublish field value if set, zero value otherwise.
func (o *FileFileRepository) GetAutopublish() bool {
	if o == nil || IsNil(o.Autopublish) {
		var ret bool
		return ret
	}
	return *o.Autopublish
}

// GetAutopublishOk returns a tuple with the Autopublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileRepository) GetAutopublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Autopublish) {
		return nil, false
	}
	return o.Autopublish, true
}

// HasAutopublish returns a boolean if a field has been set.
func (o *FileFileRepository) HasAutopublish() bool {
	if o != nil && !IsNil(o.Autopublish) {
		return true
	}

	return false
}

// SetAutopublish gets a reference to the given bool and assigns it to the Autopublish field.
func (o *FileFileRepository) SetAutopublish(v bool) {
	o.Autopublish = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileFileRepository) GetManifest() string {
	if o == nil || IsNil(o.Manifest.Get()) {
		var ret string
		return ret
	}
	return *o.Manifest.Get()
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileFileRepository) GetManifestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manifest.Get(), o.Manifest.IsSet()
}

// HasManifest returns a boolean if a field has been set.
func (o *FileFileRepository) HasManifest() bool {
	if o != nil && o.Manifest.IsSet() {
		return true
	}

	return false
}

// SetManifest gets a reference to the given NullableString and assigns it to the Manifest field.
func (o *FileFileRepository) SetManifest(v string) {
	o.Manifest.Set(&v)
}
// SetManifestNil sets the value for Manifest to be an explicit nil
func (o *FileFileRepository) SetManifestNil() {
	o.Manifest.Set(nil)
}

// UnsetManifest ensures that no value is present for Manifest, not even an explicit nil
func (o *FileFileRepository) UnsetManifest() {
	o.Manifest.Unset()
}

func (o FileFileRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileFileRepository) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Autopublish) {
		toSerialize["autopublish"] = o.Autopublish
	}
	if o.Manifest.IsSet() {
		toSerialize["manifest"] = o.Manifest.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileFileRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varFileFileRepository := _FileFileRepository{}

	err = json.Unmarshal(data, &varFileFileRepository)

	if err != nil {
		return err
	}

	*o = FileFileRepository(varFileFileRepository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "retain_repo_versions")
		delete(additionalProperties, "remote")
		delete(additionalProperties, "autopublish")
		delete(additionalProperties, "manifest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileFileRepository struct {
	value *FileFileRepository
	isSet bool
}

func (v NullableFileFileRepository) Get() *FileFileRepository {
	return v.value
}

func (v *NullableFileFileRepository) Set(val *FileFileRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFileRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFileRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFileRepository(val *FileFileRepository) *NullableFileFileRepository {
	return &NullableFileFileRepository{value: val, isSet: true}
}

func (v NullableFileFileRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFileRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


