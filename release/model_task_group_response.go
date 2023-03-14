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

// checks if the TaskGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskGroupResponse{}

// TaskGroupResponse Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type TaskGroupResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// A description of the task group.
	Description string `json:"description"`
	// Whether all tasks have been spawned for this task group.
	AllTasksDispatched bool `json:"all_tasks_dispatched"`
	// Number of tasks in the 'waiting' state
	Waiting *int32 `json:"waiting,omitempty"`
	// Number of tasks in the 'skipped' state
	Skipped *int32 `json:"skipped,omitempty"`
	// Number of tasks in the 'running' state
	Running *int32 `json:"running,omitempty"`
	// Number of tasks in the 'completed' state
	Completed *int32 `json:"completed,omitempty"`
	// Number of tasks in the 'canceled' state
	Canceled *int32 `json:"canceled,omitempty"`
	// Number of tasks in the 'failed' state
	Failed *int32 `json:"failed,omitempty"`
	// Number of tasks in the 'canceling' state
	Canceling *int32 `json:"canceling,omitempty"`
	GroupProgressReports []GroupProgressReportResponse `json:"group_progress_reports,omitempty"`
	Tasks []MinimalTaskResponse `json:"tasks,omitempty"`
}

// NewTaskGroupResponse instantiates a new TaskGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskGroupResponse(description string, allTasksDispatched bool) *TaskGroupResponse {
	this := TaskGroupResponse{}
	this.Description = description
	this.AllTasksDispatched = allTasksDispatched
	return &this
}

// NewTaskGroupResponseWithDefaults instantiates a new TaskGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskGroupResponseWithDefaults() *TaskGroupResponse {
	this := TaskGroupResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *TaskGroupResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetDescription returns the Description field value
func (o *TaskGroupResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TaskGroupResponse) SetDescription(v string) {
	o.Description = v
}

// GetAllTasksDispatched returns the AllTasksDispatched field value
func (o *TaskGroupResponse) GetAllTasksDispatched() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllTasksDispatched
}

// GetAllTasksDispatchedOk returns a tuple with the AllTasksDispatched field value
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetAllTasksDispatchedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllTasksDispatched, true
}

// SetAllTasksDispatched sets field value
func (o *TaskGroupResponse) SetAllTasksDispatched(v bool) {
	o.AllTasksDispatched = v
}

// GetWaiting returns the Waiting field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetWaiting() int32 {
	if o == nil || IsNil(o.Waiting) {
		var ret int32
		return ret
	}
	return *o.Waiting
}

// GetWaitingOk returns a tuple with the Waiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetWaitingOk() (*int32, bool) {
	if o == nil || IsNil(o.Waiting) {
		return nil, false
	}
	return o.Waiting, true
}

// HasWaiting returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasWaiting() bool {
	if o != nil && !IsNil(o.Waiting) {
		return true
	}

	return false
}

// SetWaiting gets a reference to the given int32 and assigns it to the Waiting field.
func (o *TaskGroupResponse) SetWaiting(v int32) {
	o.Waiting = &v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetSkipped() int32 {
	if o == nil || IsNil(o.Skipped) {
		var ret int32
		return ret
	}
	return *o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetSkippedOk() (*int32, bool) {
	if o == nil || IsNil(o.Skipped) {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasSkipped() bool {
	if o != nil && !IsNil(o.Skipped) {
		return true
	}

	return false
}

// SetSkipped gets a reference to the given int32 and assigns it to the Skipped field.
func (o *TaskGroupResponse) SetSkipped(v int32) {
	o.Skipped = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetRunning() int32 {
	if o == nil || IsNil(o.Running) {
		var ret int32
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetRunningOk() (*int32, bool) {
	if o == nil || IsNil(o.Running) {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasRunning() bool {
	if o != nil && !IsNil(o.Running) {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int32 and assigns it to the Running field.
func (o *TaskGroupResponse) SetRunning(v int32) {
	o.Running = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetCompleted() int32 {
	if o == nil || IsNil(o.Completed) {
		var ret int32
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int32 and assigns it to the Completed field.
func (o *TaskGroupResponse) SetCompleted(v int32) {
	o.Completed = &v
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetCanceled() int32 {
	if o == nil || IsNil(o.Canceled) {
		var ret int32
		return ret
	}
	return *o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetCanceledOk() (*int32, bool) {
	if o == nil || IsNil(o.Canceled) {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasCanceled() bool {
	if o != nil && !IsNil(o.Canceled) {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given int32 and assigns it to the Canceled field.
func (o *TaskGroupResponse) SetCanceled(v int32) {
	o.Canceled = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *TaskGroupResponse) SetFailed(v int32) {
	o.Failed = &v
}

// GetCanceling returns the Canceling field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetCanceling() int32 {
	if o == nil || IsNil(o.Canceling) {
		var ret int32
		return ret
	}
	return *o.Canceling
}

// GetCancelingOk returns a tuple with the Canceling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetCancelingOk() (*int32, bool) {
	if o == nil || IsNil(o.Canceling) {
		return nil, false
	}
	return o.Canceling, true
}

// HasCanceling returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasCanceling() bool {
	if o != nil && !IsNil(o.Canceling) {
		return true
	}

	return false
}

// SetCanceling gets a reference to the given int32 and assigns it to the Canceling field.
func (o *TaskGroupResponse) SetCanceling(v int32) {
	o.Canceling = &v
}

// GetGroupProgressReports returns the GroupProgressReports field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetGroupProgressReports() []GroupProgressReportResponse {
	if o == nil || IsNil(o.GroupProgressReports) {
		var ret []GroupProgressReportResponse
		return ret
	}
	return o.GroupProgressReports
}

// GetGroupProgressReportsOk returns a tuple with the GroupProgressReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetGroupProgressReportsOk() ([]GroupProgressReportResponse, bool) {
	if o == nil || IsNil(o.GroupProgressReports) {
		return nil, false
	}
	return o.GroupProgressReports, true
}

// HasGroupProgressReports returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasGroupProgressReports() bool {
	if o != nil && !IsNil(o.GroupProgressReports) {
		return true
	}

	return false
}

// SetGroupProgressReports gets a reference to the given []GroupProgressReportResponse and assigns it to the GroupProgressReports field.
func (o *TaskGroupResponse) SetGroupProgressReports(v []GroupProgressReportResponse) {
	o.GroupProgressReports = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *TaskGroupResponse) GetTasks() []MinimalTaskResponse {
	if o == nil || IsNil(o.Tasks) {
		var ret []MinimalTaskResponse
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroupResponse) GetTasksOk() ([]MinimalTaskResponse, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *TaskGroupResponse) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MinimalTaskResponse and assigns it to the Tasks field.
func (o *TaskGroupResponse) SetTasks(v []MinimalTaskResponse) {
	o.Tasks = v
}

func (o TaskGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	toSerialize["description"] = o.Description
	toSerialize["all_tasks_dispatched"] = o.AllTasksDispatched
	// skip: waiting is readOnly
	// skip: skipped is readOnly
	// skip: running is readOnly
	// skip: completed is readOnly
	// skip: canceled is readOnly
	// skip: failed is readOnly
	// skip: canceling is readOnly
	// skip: group_progress_reports is readOnly
	// skip: tasks is readOnly
	return toSerialize, nil
}

type NullableTaskGroupResponse struct {
	value *TaskGroupResponse
	isSet bool
}

func (v NullableTaskGroupResponse) Get() *TaskGroupResponse {
	return v.value
}

func (v *NullableTaskGroupResponse) Set(val *TaskGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskGroupResponse(val *TaskGroupResponse) *NullableTaskGroupResponse {
	return &NullableTaskGroupResponse{value: val, isSet: true}
}

func (v NullableTaskGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


