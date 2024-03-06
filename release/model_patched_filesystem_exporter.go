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

// checks if the PatchedFilesystemExporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedFilesystemExporter{}

// PatchedFilesystemExporter Serializer for FilesystemExporters.
type PatchedFilesystemExporter struct {
	// Unique name of the file system exporter.
	Name *string `json:"name,omitempty"`
	// File system location to export to.
	Path *string `json:"path,omitempty"`
	// Method of exporting* `write` - Export by writing* `hardlink` - Export by hardlinking* `symlink` - Export by symlinking
	Method *MethodEnum `json:"method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedFilesystemExporter PatchedFilesystemExporter

// NewPatchedFilesystemExporter instantiates a new PatchedFilesystemExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFilesystemExporter() *PatchedFilesystemExporter {
	this := PatchedFilesystemExporter{}
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// NewPatchedFilesystemExporterWithDefaults instantiates a new PatchedFilesystemExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFilesystemExporterWithDefaults() *PatchedFilesystemExporter {
	this := PatchedFilesystemExporter{}
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedFilesystemExporter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFilesystemExporter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedFilesystemExporter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedFilesystemExporter) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PatchedFilesystemExporter) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFilesystemExporter) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PatchedFilesystemExporter) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PatchedFilesystemExporter) SetPath(v string) {
	o.Path = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PatchedFilesystemExporter) GetMethod() MethodEnum {
	if o == nil || IsNil(o.Method) {
		var ret MethodEnum
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFilesystemExporter) GetMethodOk() (*MethodEnum, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PatchedFilesystemExporter) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given MethodEnum and assigns it to the Method field.
func (o *PatchedFilesystemExporter) SetMethod(v MethodEnum) {
	o.Method = &v
}

func (o PatchedFilesystemExporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedFilesystemExporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedFilesystemExporter) UnmarshalJSON(data []byte) (err error) {
	varPatchedFilesystemExporter := _PatchedFilesystemExporter{}

	err = json.Unmarshal(data, &varPatchedFilesystemExporter)

	if err != nil {
		return err
	}

	*o = PatchedFilesystemExporter(varPatchedFilesystemExporter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "path")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedFilesystemExporter struct {
	value *PatchedFilesystemExporter
	isSet bool
}

func (v NullablePatchedFilesystemExporter) Get() *PatchedFilesystemExporter {
	return v.value
}

func (v *NullablePatchedFilesystemExporter) Set(val *PatchedFilesystemExporter) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFilesystemExporter) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFilesystemExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFilesystemExporter(val *PatchedFilesystemExporter) *NullablePatchedFilesystemExporter {
	return &NullablePatchedFilesystemExporter{value: val, isSet: true}
}

func (v NullablePatchedFilesystemExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFilesystemExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


