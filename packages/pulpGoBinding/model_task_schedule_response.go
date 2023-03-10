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
	"time"
)

// checks if the TaskScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskScheduleResponse{}

// TaskScheduleResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type TaskScheduleResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The name of the task schedule.
	Name string `json:"name"`
	// The name of the task to be scheduled.
	TaskName string `json:"task_name"`
	// Periodicity of the schedule.
	DispatchInterval string `json:"dispatch_interval"`
	// Timestamp of the next time the task will be dispatched.
	NextDispatch *time.Time `json:"next_dispatch,omitempty"`
	// The last task dispatched by this schedule.
	LastTask *string `json:"last_task,omitempty"`
}

// NewTaskScheduleResponse instantiates a new TaskScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskScheduleResponse(name string, taskName string, dispatchInterval string) *TaskScheduleResponse {
	this := TaskScheduleResponse{}
	this.Name = name
	this.TaskName = taskName
	this.DispatchInterval = dispatchInterval
	return &this
}

// NewTaskScheduleResponseWithDefaults instantiates a new TaskScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskScheduleResponseWithDefaults() *TaskScheduleResponse {
	this := TaskScheduleResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *TaskScheduleResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *TaskScheduleResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *TaskScheduleResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *TaskScheduleResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *TaskScheduleResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *TaskScheduleResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *TaskScheduleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskScheduleResponse) SetName(v string) {
	o.Name = v
}

// GetTaskName returns the TaskName field value
func (o *TaskScheduleResponse) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value
func (o *TaskScheduleResponse) SetTaskName(v string) {
	o.TaskName = v
}

// GetDispatchInterval returns the DispatchInterval field value
func (o *TaskScheduleResponse) GetDispatchInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DispatchInterval
}

// GetDispatchIntervalOk returns a tuple with the DispatchInterval field value
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetDispatchIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DispatchInterval, true
}

// SetDispatchInterval sets field value
func (o *TaskScheduleResponse) SetDispatchInterval(v string) {
	o.DispatchInterval = v
}

// GetNextDispatch returns the NextDispatch field value if set, zero value otherwise.
func (o *TaskScheduleResponse) GetNextDispatch() time.Time {
	if o == nil || IsNil(o.NextDispatch) {
		var ret time.Time
		return ret
	}
	return *o.NextDispatch
}

// GetNextDispatchOk returns a tuple with the NextDispatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetNextDispatchOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextDispatch) {
		return nil, false
	}
	return o.NextDispatch, true
}

// HasNextDispatch returns a boolean if a field has been set.
func (o *TaskScheduleResponse) HasNextDispatch() bool {
	if o != nil && !IsNil(o.NextDispatch) {
		return true
	}

	return false
}

// SetNextDispatch gets a reference to the given time.Time and assigns it to the NextDispatch field.
func (o *TaskScheduleResponse) SetNextDispatch(v time.Time) {
	o.NextDispatch = &v
}

// GetLastTask returns the LastTask field value if set, zero value otherwise.
func (o *TaskScheduleResponse) GetLastTask() string {
	if o == nil || IsNil(o.LastTask) {
		var ret string
		return ret
	}
	return *o.LastTask
}

// GetLastTaskOk returns a tuple with the LastTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskScheduleResponse) GetLastTaskOk() (*string, bool) {
	if o == nil || IsNil(o.LastTask) {
		return nil, false
	}
	return o.LastTask, true
}

// HasLastTask returns a boolean if a field has been set.
func (o *TaskScheduleResponse) HasLastTask() bool {
	if o != nil && !IsNil(o.LastTask) {
		return true
	}

	return false
}

// SetLastTask gets a reference to the given string and assigns it to the LastTask field.
func (o *TaskScheduleResponse) SetLastTask(v string) {
	o.LastTask = &v
}

func (o TaskScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["name"] = o.Name
	toSerialize["task_name"] = o.TaskName
	toSerialize["dispatch_interval"] = o.DispatchInterval
	// skip: next_dispatch is readOnly
	// skip: last_task is readOnly
	return toSerialize, nil
}

type NullableTaskScheduleResponse struct {
	value *TaskScheduleResponse
	isSet bool
}

func (v NullableTaskScheduleResponse) Get() *TaskScheduleResponse {
	return v.value
}

func (v *NullableTaskScheduleResponse) Set(val *TaskScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskScheduleResponse(val *TaskScheduleResponse) *NullableTaskScheduleResponse {
	return &NullableTaskScheduleResponse{value: val, isSet: true}
}

func (v NullableTaskScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


