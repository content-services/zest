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
	"time"
)

// checks if the GalaxyCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalaxyCollection{}

// GalaxyCollection A serializer for a Collection.
type GalaxyCollection struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// NewGalaxyCollection instantiates a new GalaxyCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalaxyCollection(id string, name string, created time.Time, modified time.Time) *GalaxyCollection {
	this := GalaxyCollection{}
	this.Id = id
	this.Name = name
	this.Created = created
	this.Modified = modified
	return &this
}

// NewGalaxyCollectionWithDefaults instantiates a new GalaxyCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalaxyCollectionWithDefaults() *GalaxyCollection {
	this := GalaxyCollection{}
	return &this
}

// GetId returns the Id field value
func (o *GalaxyCollection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GalaxyCollection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GalaxyCollection) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GalaxyCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GalaxyCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GalaxyCollection) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value
func (o *GalaxyCollection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GalaxyCollection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GalaxyCollection) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *GalaxyCollection) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *GalaxyCollection) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *GalaxyCollection) SetModified(v time.Time) {
	o.Modified = v
}

func (o GalaxyCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalaxyCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	return toSerialize, nil
}

type NullableGalaxyCollection struct {
	value *GalaxyCollection
	isSet bool
}

func (v NullableGalaxyCollection) Get() *GalaxyCollection {
	return v.value
}

func (v *NullableGalaxyCollection) Set(val *GalaxyCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGalaxyCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGalaxyCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalaxyCollection(val *GalaxyCollection) *NullableGalaxyCollection {
	return &NullableGalaxyCollection{value: val, isSet: true}
}

func (v NullableGalaxyCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalaxyCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


