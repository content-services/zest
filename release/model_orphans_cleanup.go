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

// checks if the OrphansCleanup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrphansCleanup{}

// OrphansCleanup struct for OrphansCleanup
type OrphansCleanup struct {
	// Will delete specified content and associated Artifacts if they are orphans.
	ContentHrefs []interface{} `json:"content_hrefs,omitempty"`
	// The time in minutes for how long Pulp will hold orphan Content and Artifacts before they become candidates for deletion by this orphan cleanup task. This should ideally be longer than your longest running task otherwise any content created during that task could be cleaned up before the task finishes. If not specified, a default value is taken from the setting ORPHAN_PROTECTION_TIME.
	OrphanProtectionTime NullableInt64 `json:"orphan_protection_time,omitempty"`
}

// NewOrphansCleanup instantiates a new OrphansCleanup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrphansCleanup() *OrphansCleanup {
	this := OrphansCleanup{}
	return &this
}

// NewOrphansCleanupWithDefaults instantiates a new OrphansCleanup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrphansCleanupWithDefaults() *OrphansCleanup {
	this := OrphansCleanup{}
	return &this
}

// GetContentHrefs returns the ContentHrefs field value if set, zero value otherwise.
func (o *OrphansCleanup) GetContentHrefs() []interface{} {
	if o == nil || IsNil(o.ContentHrefs) {
		var ret []interface{}
		return ret
	}
	return o.ContentHrefs
}

// GetContentHrefsOk returns a tuple with the ContentHrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphansCleanup) GetContentHrefsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ContentHrefs) {
		return nil, false
	}
	return o.ContentHrefs, true
}

// HasContentHrefs returns a boolean if a field has been set.
func (o *OrphansCleanup) HasContentHrefs() bool {
	if o != nil && !IsNil(o.ContentHrefs) {
		return true
	}

	return false
}

// SetContentHrefs gets a reference to the given []interface{} and assigns it to the ContentHrefs field.
func (o *OrphansCleanup) SetContentHrefs(v []interface{}) {
	o.ContentHrefs = v
}

// GetOrphanProtectionTime returns the OrphanProtectionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrphansCleanup) GetOrphanProtectionTime() int64 {
	if o == nil || IsNil(o.OrphanProtectionTime.Get()) {
		var ret int64
		return ret
	}
	return *o.OrphanProtectionTime.Get()
}

// GetOrphanProtectionTimeOk returns a tuple with the OrphanProtectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrphansCleanup) GetOrphanProtectionTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrphanProtectionTime.Get(), o.OrphanProtectionTime.IsSet()
}

// HasOrphanProtectionTime returns a boolean if a field has been set.
func (o *OrphansCleanup) HasOrphanProtectionTime() bool {
	if o != nil && o.OrphanProtectionTime.IsSet() {
		return true
	}

	return false
}

// SetOrphanProtectionTime gets a reference to the given NullableInt64 and assigns it to the OrphanProtectionTime field.
func (o *OrphansCleanup) SetOrphanProtectionTime(v int64) {
	o.OrphanProtectionTime.Set(&v)
}
// SetOrphanProtectionTimeNil sets the value for OrphanProtectionTime to be an explicit nil
func (o *OrphansCleanup) SetOrphanProtectionTimeNil() {
	o.OrphanProtectionTime.Set(nil)
}

// UnsetOrphanProtectionTime ensures that no value is present for OrphanProtectionTime, not even an explicit nil
func (o *OrphansCleanup) UnsetOrphanProtectionTime() {
	o.OrphanProtectionTime.Unset()
}

func (o OrphansCleanup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrphansCleanup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentHrefs) {
		toSerialize["content_hrefs"] = o.ContentHrefs
	}
	if o.OrphanProtectionTime.IsSet() {
		toSerialize["orphan_protection_time"] = o.OrphanProtectionTime.Get()
	}
	return toSerialize, nil
}

type NullableOrphansCleanup struct {
	value *OrphansCleanup
	isSet bool
}

func (v NullableOrphansCleanup) Get() *OrphansCleanup {
	return v.value
}

func (v *NullableOrphansCleanup) Set(val *OrphansCleanup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrphansCleanup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrphansCleanup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrphansCleanup(val *OrphansCleanup) *NullableOrphansCleanup {
	return &NullableOrphansCleanup{value: val, isSet: true}
}

func (v NullableOrphansCleanup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrphansCleanup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


