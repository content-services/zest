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

// checks if the TaskGroupOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskGroupOperationResponse{}

// TaskGroupOperationResponse Serializer for asynchronous operations that return a task group.
type TaskGroupOperationResponse struct {
	// The href of the task group.
	TaskGroup string `json:"task_group"`
}

type _TaskGroupOperationResponse TaskGroupOperationResponse

// NewTaskGroupOperationResponse instantiates a new TaskGroupOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskGroupOperationResponse(taskGroup string) *TaskGroupOperationResponse {
	this := TaskGroupOperationResponse{}
	this.TaskGroup = taskGroup
	return &this
}

// NewTaskGroupOperationResponseWithDefaults instantiates a new TaskGroupOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskGroupOperationResponseWithDefaults() *TaskGroupOperationResponse {
	this := TaskGroupOperationResponse{}
	return &this
}

// GetTaskGroup returns the TaskGroup field value
func (o *TaskGroupOperationResponse) GetTaskGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskGroup
}

// GetTaskGroupOk returns a tuple with the TaskGroup field value
// and a boolean to check if the value has been set.
func (o *TaskGroupOperationResponse) GetTaskGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskGroup, true
}

// SetTaskGroup sets field value
func (o *TaskGroupOperationResponse) SetTaskGroup(v string) {
	o.TaskGroup = v
}

func (o TaskGroupOperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskGroupOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["task_group"] = o.TaskGroup
	return toSerialize, nil
}

func (o *TaskGroupOperationResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"task_group",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTaskGroupOperationResponse := _TaskGroupOperationResponse{}

	err = json.Unmarshal(bytes, &varTaskGroupOperationResponse)

	if err != nil {
		return err
	}

	*o = TaskGroupOperationResponse(varTaskGroupOperationResponse)

	return err
}

type NullableTaskGroupOperationResponse struct {
	value *TaskGroupOperationResponse
	isSet bool
}

func (v NullableTaskGroupOperationResponse) Get() *TaskGroupOperationResponse {
	return v.value
}

func (v *NullableTaskGroupOperationResponse) Set(val *TaskGroupOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskGroupOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskGroupOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskGroupOperationResponse(val *TaskGroupOperationResponse) *NullableTaskGroupOperationResponse {
	return &NullableTaskGroupOperationResponse{value: val, isSet: true}
}

func (v NullableTaskGroupOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskGroupOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


