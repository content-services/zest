/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PulpExportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpExportResponse{}

// PulpExportResponse Serializer for PulpExports.
type PulpExportResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A URI of the task that ran the Export.
	Task NullableString `json:"task,omitempty"`
	// Resources that were exported.
	ExportedResources []string `json:"exported_resources,omitempty"`
	// Any additional parameters that were used to create the export.
	Params map[string]interface{} `json:"params,omitempty"`
	// Dictionary of filename: sha256hash entries for export-output-file(s)
	OutputFileInfo map[string]interface{} `json:"output_file_info,omitempty"`
	// Filename and sha256-checksum of table-of-contents for this export
	TocInfo map[string]interface{} `json:"toc_info,omitempty"`
}

// NewPulpExportResponse instantiates a new PulpExportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpExportResponse() *PulpExportResponse {
	this := PulpExportResponse{}
	return &this
}

// NewPulpExportResponseWithDefaults instantiates a new PulpExportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpExportResponseWithDefaults() *PulpExportResponse {
	this := PulpExportResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *PulpExportResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *PulpExportResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *PulpExportResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *PulpExportResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *PulpExportResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *PulpExportResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetTask returns the Task field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PulpExportResponse) GetTask() string {
	if o == nil || IsNil(o.Task.Get()) {
		var ret string
		return ret
	}
	return *o.Task.Get()
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PulpExportResponse) GetTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Task.Get(), o.Task.IsSet()
}

// HasTask returns a boolean if a field has been set.
func (o *PulpExportResponse) HasTask() bool {
	if o != nil && o.Task.IsSet() {
		return true
	}

	return false
}

// SetTask gets a reference to the given NullableString and assigns it to the Task field.
func (o *PulpExportResponse) SetTask(v string) {
	o.Task.Set(&v)
}
// SetTaskNil sets the value for Task to be an explicit nil
func (o *PulpExportResponse) SetTaskNil() {
	o.Task.Set(nil)
}

// UnsetTask ensures that no value is present for Task, not even an explicit nil
func (o *PulpExportResponse) UnsetTask() {
	o.Task.Unset()
}

// GetExportedResources returns the ExportedResources field value if set, zero value otherwise.
func (o *PulpExportResponse) GetExportedResources() []string {
	if o == nil || IsNil(o.ExportedResources) {
		var ret []string
		return ret
	}
	return o.ExportedResources
}

// GetExportedResourcesOk returns a tuple with the ExportedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetExportedResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExportedResources) {
		return nil, false
	}
	return o.ExportedResources, true
}

// HasExportedResources returns a boolean if a field has been set.
func (o *PulpExportResponse) HasExportedResources() bool {
	if o != nil && !IsNil(o.ExportedResources) {
		return true
	}

	return false
}

// SetExportedResources gets a reference to the given []string and assigns it to the ExportedResources field.
func (o *PulpExportResponse) SetExportedResources(v []string) {
	o.ExportedResources = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *PulpExportResponse) GetParams() map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *PulpExportResponse) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *PulpExportResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetOutputFileInfo returns the OutputFileInfo field value if set, zero value otherwise.
func (o *PulpExportResponse) GetOutputFileInfo() map[string]interface{} {
	if o == nil || IsNil(o.OutputFileInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputFileInfo
}

// GetOutputFileInfoOk returns a tuple with the OutputFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetOutputFileInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OutputFileInfo) {
		return map[string]interface{}{}, false
	}
	return o.OutputFileInfo, true
}

// HasOutputFileInfo returns a boolean if a field has been set.
func (o *PulpExportResponse) HasOutputFileInfo() bool {
	if o != nil && !IsNil(o.OutputFileInfo) {
		return true
	}

	return false
}

// SetOutputFileInfo gets a reference to the given map[string]interface{} and assigns it to the OutputFileInfo field.
func (o *PulpExportResponse) SetOutputFileInfo(v map[string]interface{}) {
	o.OutputFileInfo = v
}

// GetTocInfo returns the TocInfo field value if set, zero value otherwise.
func (o *PulpExportResponse) GetTocInfo() map[string]interface{} {
	if o == nil || IsNil(o.TocInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.TocInfo
}

// GetTocInfoOk returns a tuple with the TocInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PulpExportResponse) GetTocInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TocInfo) {
		return map[string]interface{}{}, false
	}
	return o.TocInfo, true
}

// HasTocInfo returns a boolean if a field has been set.
func (o *PulpExportResponse) HasTocInfo() bool {
	if o != nil && !IsNil(o.TocInfo) {
		return true
	}

	return false
}

// SetTocInfo gets a reference to the given map[string]interface{} and assigns it to the TocInfo field.
func (o *PulpExportResponse) SetTocInfo(v map[string]interface{}) {
	o.TocInfo = v
}

func (o PulpExportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpExportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	if o.Task.IsSet() {
		toSerialize["task"] = o.Task.Get()
	}
	// skip: exported_resources is readOnly
	// skip: params is readOnly
	// skip: output_file_info is readOnly
	// skip: toc_info is readOnly
	return toSerialize, nil
}

type NullablePulpExportResponse struct {
	value *PulpExportResponse
	isSet bool
}

func (v NullablePulpExportResponse) Get() *PulpExportResponse {
	return v.value
}

func (v *NullablePulpExportResponse) Set(val *PulpExportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpExportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpExportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpExportResponse(val *PulpExportResponse) *NullablePulpExportResponse {
	return &NullablePulpExportResponse{value: val, isSet: true}
}

func (v NullablePulpExportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpExportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


