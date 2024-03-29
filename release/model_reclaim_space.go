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

// checks if the ReclaimSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReclaimSpace{}

// ReclaimSpace Serializer for reclaim disk space operation.
type ReclaimSpace struct {
	// Will reclaim space for the specified list of repos. Use ['*'] to specify all repos.
	RepoHrefs []interface{} `json:"repo_hrefs"`
	// Will exclude repo versions from space reclaim.
	RepoVersionsKeeplist []string `json:"repo_versions_keeplist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReclaimSpace ReclaimSpace

// NewReclaimSpace instantiates a new ReclaimSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReclaimSpace(repoHrefs []interface{}) *ReclaimSpace {
	this := ReclaimSpace{}
	this.RepoHrefs = repoHrefs
	return &this
}

// NewReclaimSpaceWithDefaults instantiates a new ReclaimSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReclaimSpaceWithDefaults() *ReclaimSpace {
	this := ReclaimSpace{}
	return &this
}

// GetRepoHrefs returns the RepoHrefs field value
func (o *ReclaimSpace) GetRepoHrefs() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.RepoHrefs
}

// GetRepoHrefsOk returns a tuple with the RepoHrefs field value
// and a boolean to check if the value has been set.
func (o *ReclaimSpace) GetRepoHrefsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepoHrefs, true
}

// SetRepoHrefs sets field value
func (o *ReclaimSpace) SetRepoHrefs(v []interface{}) {
	o.RepoHrefs = v
}

// GetRepoVersionsKeeplist returns the RepoVersionsKeeplist field value if set, zero value otherwise.
func (o *ReclaimSpace) GetRepoVersionsKeeplist() []string {
	if o == nil || IsNil(o.RepoVersionsKeeplist) {
		var ret []string
		return ret
	}
	return o.RepoVersionsKeeplist
}

// GetRepoVersionsKeeplistOk returns a tuple with the RepoVersionsKeeplist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReclaimSpace) GetRepoVersionsKeeplistOk() ([]string, bool) {
	if o == nil || IsNil(o.RepoVersionsKeeplist) {
		return nil, false
	}
	return o.RepoVersionsKeeplist, true
}

// HasRepoVersionsKeeplist returns a boolean if a field has been set.
func (o *ReclaimSpace) HasRepoVersionsKeeplist() bool {
	if o != nil && !IsNil(o.RepoVersionsKeeplist) {
		return true
	}

	return false
}

// SetRepoVersionsKeeplist gets a reference to the given []string and assigns it to the RepoVersionsKeeplist field.
func (o *ReclaimSpace) SetRepoVersionsKeeplist(v []string) {
	o.RepoVersionsKeeplist = v
}

func (o ReclaimSpace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReclaimSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repo_hrefs"] = o.RepoHrefs
	if !IsNil(o.RepoVersionsKeeplist) {
		toSerialize["repo_versions_keeplist"] = o.RepoVersionsKeeplist
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReclaimSpace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repo_hrefs",
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

	varReclaimSpace := _ReclaimSpace{}

	err = json.Unmarshal(data, &varReclaimSpace)

	if err != nil {
		return err
	}

	*o = ReclaimSpace(varReclaimSpace)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repo_hrefs")
		delete(additionalProperties, "repo_versions_keeplist")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReclaimSpace struct {
	value *ReclaimSpace
	isSet bool
}

func (v NullableReclaimSpace) Get() *ReclaimSpace {
	return v.value
}

func (v *NullableReclaimSpace) Set(val *ReclaimSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableReclaimSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableReclaimSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReclaimSpace(val *ReclaimSpace) *NullableReclaimSpace {
	return &NullableReclaimSpace{value: val, isSet: true}
}

func (v NullableReclaimSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReclaimSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


