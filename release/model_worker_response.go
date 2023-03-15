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

// checks if the WorkerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkerResponse{}

// WorkerResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type WorkerResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The name of the worker.
	Name *string `json:"name,omitempty"`
	// Timestamp of the last time the worker talked to the service.
	LastHeartbeat *time.Time `json:"last_heartbeat,omitempty"`
	// The task this worker is currently executing, or empty if the worker is not currently assigned to a task.
	CurrentTask *string `json:"current_task,omitempty"`
}

// NewWorkerResponse instantiates a new WorkerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerResponse() *WorkerResponse {
	this := WorkerResponse{}
	return &this
}

// NewWorkerResponseWithDefaults instantiates a new WorkerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerResponseWithDefaults() *WorkerResponse {
	this := WorkerResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *WorkerResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *WorkerResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *WorkerResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *WorkerResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *WorkerResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *WorkerResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkerResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkerResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkerResponse) SetName(v string) {
	o.Name = &v
}

// GetLastHeartbeat returns the LastHeartbeat field value if set, zero value otherwise.
func (o *WorkerResponse) GetLastHeartbeat() time.Time {
	if o == nil || IsNil(o.LastHeartbeat) {
		var ret time.Time
		return ret
	}
	return *o.LastHeartbeat
}

// GetLastHeartbeatOk returns a tuple with the LastHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetLastHeartbeatOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastHeartbeat) {
		return nil, false
	}
	return o.LastHeartbeat, true
}

// HasLastHeartbeat returns a boolean if a field has been set.
func (o *WorkerResponse) HasLastHeartbeat() bool {
	if o != nil && !IsNil(o.LastHeartbeat) {
		return true
	}

	return false
}

// SetLastHeartbeat gets a reference to the given time.Time and assigns it to the LastHeartbeat field.
func (o *WorkerResponse) SetLastHeartbeat(v time.Time) {
	o.LastHeartbeat = &v
}

// GetCurrentTask returns the CurrentTask field value if set, zero value otherwise.
func (o *WorkerResponse) GetCurrentTask() string {
	if o == nil || IsNil(o.CurrentTask) {
		var ret string
		return ret
	}
	return *o.CurrentTask
}

// GetCurrentTaskOk returns a tuple with the CurrentTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetCurrentTaskOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentTask) {
		return nil, false
	}
	return o.CurrentTask, true
}

// HasCurrentTask returns a boolean if a field has been set.
func (o *WorkerResponse) HasCurrentTask() bool {
	if o != nil && !IsNil(o.CurrentTask) {
		return true
	}

	return false
}

// SetCurrentTask gets a reference to the given string and assigns it to the CurrentTask field.
func (o *WorkerResponse) SetCurrentTask(v string) {
	o.CurrentTask = &v
}

func (o WorkerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	// skip: name is readOnly
	// skip: last_heartbeat is readOnly
	// skip: current_task is readOnly
	return toSerialize, nil
}

type NullableWorkerResponse struct {
	value *WorkerResponse
	isSet bool
}

func (v NullableWorkerResponse) Get() *WorkerResponse {
	return v.value
}

func (v *NullableWorkerResponse) Set(val *WorkerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerResponse(val *WorkerResponse) *NullableWorkerResponse {
	return &NullableWorkerResponse{value: val, isSet: true}
}

func (v NullableWorkerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


