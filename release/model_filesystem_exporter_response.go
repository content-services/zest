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
	"time"
	"bytes"
	"fmt"
)

// checks if the FilesystemExporterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesystemExporterResponse{}

// FilesystemExporterResponse Serializer for FilesystemExporters.
type FilesystemExporterResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// Unique name of the file system exporter.
	Name string `json:"name"`
	// File system location to export to.
	Path string `json:"path"`
	// Method of exporting* `write` - Export by writing* `hardlink` - Export by hardlinking* `symlink` - Export by symlinking
	Method *MethodEnum `json:"method,omitempty"`
}

type _FilesystemExporterResponse FilesystemExporterResponse

// NewFilesystemExporterResponse instantiates a new FilesystemExporterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemExporterResponse(name string, path string) *FilesystemExporterResponse {
	this := FilesystemExporterResponse{}
	this.Name = name
	this.Path = path
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// NewFilesystemExporterResponseWithDefaults instantiates a new FilesystemExporterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemExporterResponseWithDefaults() *FilesystemExporterResponse {
	this := FilesystemExporterResponse{}
	var method MethodEnum = METHODENUM_WRITE
	this.Method = &method
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *FilesystemExporterResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *FilesystemExporterResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *FilesystemExporterResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *FilesystemExporterResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *FilesystemExporterResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *FilesystemExporterResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *FilesystemExporterResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *FilesystemExporterResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *FilesystemExporterResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetName returns the Name field value
func (o *FilesystemExporterResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesystemExporterResponse) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *FilesystemExporterResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FilesystemExporterResponse) SetPath(v string) {
	o.Path = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *FilesystemExporterResponse) GetMethod() MethodEnum {
	if o == nil || IsNil(o.Method) {
		var ret MethodEnum
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemExporterResponse) GetMethodOk() (*MethodEnum, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *FilesystemExporterResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given MethodEnum and assigns it to the Method field.
func (o *FilesystemExporterResponse) SetMethod(v MethodEnum) {
	o.Method = &v
}

func (o FilesystemExporterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesystemExporterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

func (o *FilesystemExporterResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"path",
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

	varFilesystemExporterResponse := _FilesystemExporterResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilesystemExporterResponse)

	if err != nil {
		return err
	}

	*o = FilesystemExporterResponse(varFilesystemExporterResponse)

	return err
}

type NullableFilesystemExporterResponse struct {
	value *FilesystemExporterResponse
	isSet bool
}

func (v NullableFilesystemExporterResponse) Get() *FilesystemExporterResponse {
	return v.value
}

func (v *NullableFilesystemExporterResponse) Set(val *FilesystemExporterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemExporterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemExporterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemExporterResponse(val *FilesystemExporterResponse) *NullableFilesystemExporterResponse {
	return &NullableFilesystemExporterResponse{value: val, isSet: true}
}

func (v NullableFilesystemExporterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemExporterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


