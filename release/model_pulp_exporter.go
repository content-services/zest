/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"encoding/json"
)

// checks if the PulpExporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpExporter{}

// PulpExporter Serializer for pulp exporters.
type PulpExporter struct {
	// Unique name of the file system exporter.
	Name string `json:"name"`
	// File system directory to store exported tar.gzs.
	Path string `json:"path"`
	Repositories []string `json:"repositories"`
	// Last attempted export for this PulpExporter
	LastExport NullableString `json:"last_export,omitempty"`
}

// NewPulpExporter instantiates a new PulpExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpExporter(name string, path string, repositories []string) *PulpExporter {
	this := PulpExporter{}
	this.Name = name
	this.Path = path
	this.Repositories = repositories
	return &this
}

// NewPulpExporterWithDefaults instantiates a new PulpExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpExporterWithDefaults() *PulpExporter {
	this := PulpExporter{}
	return &this
}

// GetName returns the Name field value
func (o *PulpExporter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PulpExporter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PulpExporter) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *PulpExporter) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PulpExporter) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PulpExporter) SetPath(v string) {
	o.Path = v
}

// GetRepositories returns the Repositories field value
func (o *PulpExporter) GetRepositories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *PulpExporter) GetRepositoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repositories, true
}

// SetRepositories sets field value
func (o *PulpExporter) SetRepositories(v []string) {
	o.Repositories = v
}

// GetLastExport returns the LastExport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PulpExporter) GetLastExport() string {
	if o == nil || IsNil(o.LastExport.Get()) {
		var ret string
		return ret
	}
	return *o.LastExport.Get()
}

// GetLastExportOk returns a tuple with the LastExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PulpExporter) GetLastExportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExport.Get(), o.LastExport.IsSet()
}

// HasLastExport returns a boolean if a field has been set.
func (o *PulpExporter) HasLastExport() bool {
	if o != nil && o.LastExport.IsSet() {
		return true
	}

	return false
}

// SetLastExport gets a reference to the given NullableString and assigns it to the LastExport field.
func (o *PulpExporter) SetLastExport(v string) {
	o.LastExport.Set(&v)
}
// SetLastExportNil sets the value for LastExport to be an explicit nil
func (o *PulpExporter) SetLastExportNil() {
	o.LastExport.Set(nil)
}

// UnsetLastExport ensures that no value is present for LastExport, not even an explicit nil
func (o *PulpExporter) UnsetLastExport() {
	o.LastExport.Unset()
}

func (o PulpExporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpExporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["repositories"] = o.Repositories
	if o.LastExport.IsSet() {
		toSerialize["last_export"] = o.LastExport.Get()
	}
	return toSerialize, nil
}

type NullablePulpExporter struct {
	value *PulpExporter
	isSet bool
}

func (v NullablePulpExporter) Get() *PulpExporter {
	return v.value
}

func (v *NullablePulpExporter) Set(val *PulpExporter) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpExporter) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpExporter(val *PulpExporter) *NullablePulpExporter {
	return &NullablePulpExporter{value: val, isSet: true}
}

func (v NullablePulpExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


