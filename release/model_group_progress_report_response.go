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

// checks if the GroupProgressReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupProgressReportResponse{}

// GroupProgressReportResponse Base serializer for use with :class:`pulpcore.app.models.Model`This ensures that all Serializers provide values for the 'pulp_href` field.The class provides a default for the ``ref_name`` attribute in theModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitionsof plugins are namespaced properly.
type GroupProgressReportResponse struct {
	// The message shown to the user for the group progress report.
	Message *string `json:"message,omitempty"`
	// Identifies the type of group progress report'.
	Code *string `json:"code,omitempty"`
	// The total count of items.
	Total *int64 `json:"total,omitempty"`
	// The count of items already processed. Defaults to 0.
	Done *int64 `json:"done,omitempty"`
	// The suffix to be shown with the group progress report.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewGroupProgressReportResponse instantiates a new GroupProgressReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupProgressReportResponse() *GroupProgressReportResponse {
	this := GroupProgressReportResponse{}
	return &this
}

// NewGroupProgressReportResponseWithDefaults instantiates a new GroupProgressReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupProgressReportResponseWithDefaults() *GroupProgressReportResponse {
	this := GroupProgressReportResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GroupProgressReportResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProgressReportResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GroupProgressReportResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GroupProgressReportResponse) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GroupProgressReportResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProgressReportResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GroupProgressReportResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GroupProgressReportResponse) SetCode(v string) {
	o.Code = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GroupProgressReportResponse) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProgressReportResponse) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GroupProgressReportResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GroupProgressReportResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *GroupProgressReportResponse) GetDone() int64 {
	if o == nil || IsNil(o.Done) {
		var ret int64
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProgressReportResponse) GetDoneOk() (*int64, bool) {
	if o == nil || IsNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *GroupProgressReportResponse) HasDone() bool {
	if o != nil && !IsNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given int64 and assigns it to the Done field.
func (o *GroupProgressReportResponse) SetDone(v int64) {
	o.Done = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupProgressReportResponse) GetSuffix() string {
	if o == nil || IsNil(o.Suffix.Get()) {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupProgressReportResponse) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *GroupProgressReportResponse) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *GroupProgressReportResponse) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *GroupProgressReportResponse) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *GroupProgressReportResponse) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o GroupProgressReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupProgressReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return toSerialize, nil
}

type NullableGroupProgressReportResponse struct {
	value *GroupProgressReportResponse
	isSet bool
}

func (v NullableGroupProgressReportResponse) Get() *GroupProgressReportResponse {
	return v.value
}

func (v *NullableGroupProgressReportResponse) Set(val *GroupProgressReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupProgressReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupProgressReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupProgressReportResponse(val *GroupProgressReportResponse) *NullableGroupProgressReportResponse {
	return &NullableGroupProgressReportResponse{value: val, isSet: true}
}

func (v NullableGroupProgressReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupProgressReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


