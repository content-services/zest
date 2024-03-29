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

// checks if the AccessPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicy{}

// AccessPolicy Serializer for AccessPolicy.
type AccessPolicy struct {
	// List of callables that define the new permissions to be created for new objects.This is deprecated. Use `creation_hooks` instead.
	PermissionsAssignment []map[string]interface{} `json:"permissions_assignment,omitempty"`
	// List of callables that may associate user roles for new objects.
	CreationHooks []map[string]interface{} `json:"creation_hooks,omitempty"`
	// List of policy statements defining the policy.
	Statements []map[string]interface{} `json:"statements"`
	// A callable for performing queryset scoping. See plugin documentation for valid callables. Set to blank to turn off queryset scoping.
	QuerysetScoping map[string]interface{} `json:"queryset_scoping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicy AccessPolicy

// NewAccessPolicy instantiates a new AccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicy(statements []map[string]interface{}) *AccessPolicy {
	this := AccessPolicy{}
	this.Statements = statements
	return &this
}

// NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyWithDefaults() *AccessPolicy {
	this := AccessPolicy{}
	return &this
}

// GetPermissionsAssignment returns the PermissionsAssignment field value if set, zero value otherwise.
func (o *AccessPolicy) GetPermissionsAssignment() []map[string]interface{} {
	if o == nil || IsNil(o.PermissionsAssignment) {
		var ret []map[string]interface{}
		return ret
	}
	return o.PermissionsAssignment
}

// GetPermissionsAssignmentOk returns a tuple with the PermissionsAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetPermissionsAssignmentOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PermissionsAssignment) {
		return nil, false
	}
	return o.PermissionsAssignment, true
}

// HasPermissionsAssignment returns a boolean if a field has been set.
func (o *AccessPolicy) HasPermissionsAssignment() bool {
	if o != nil && !IsNil(o.PermissionsAssignment) {
		return true
	}

	return false
}

// SetPermissionsAssignment gets a reference to the given []map[string]interface{} and assigns it to the PermissionsAssignment field.
func (o *AccessPolicy) SetPermissionsAssignment(v []map[string]interface{}) {
	o.PermissionsAssignment = v
}

// GetCreationHooks returns the CreationHooks field value if set, zero value otherwise.
func (o *AccessPolicy) GetCreationHooks() []map[string]interface{} {
	if o == nil || IsNil(o.CreationHooks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.CreationHooks
}

// GetCreationHooksOk returns a tuple with the CreationHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetCreationHooksOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.CreationHooks) {
		return nil, false
	}
	return o.CreationHooks, true
}

// HasCreationHooks returns a boolean if a field has been set.
func (o *AccessPolicy) HasCreationHooks() bool {
	if o != nil && !IsNil(o.CreationHooks) {
		return true
	}

	return false
}

// SetCreationHooks gets a reference to the given []map[string]interface{} and assigns it to the CreationHooks field.
func (o *AccessPolicy) SetCreationHooks(v []map[string]interface{}) {
	o.CreationHooks = v
}

// GetStatements returns the Statements field value
func (o *AccessPolicy) GetStatements() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetStatementsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *AccessPolicy) SetStatements(v []map[string]interface{}) {
	o.Statements = v
}

// GetQuerysetScoping returns the QuerysetScoping field value if set, zero value otherwise.
func (o *AccessPolicy) GetQuerysetScoping() map[string]interface{} {
	if o == nil || IsNil(o.QuerysetScoping) {
		var ret map[string]interface{}
		return ret
	}
	return o.QuerysetScoping
}

// GetQuerysetScopingOk returns a tuple with the QuerysetScoping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetQuerysetScopingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.QuerysetScoping) {
		return map[string]interface{}{}, false
	}
	return o.QuerysetScoping, true
}

// HasQuerysetScoping returns a boolean if a field has been set.
func (o *AccessPolicy) HasQuerysetScoping() bool {
	if o != nil && !IsNil(o.QuerysetScoping) {
		return true
	}

	return false
}

// SetQuerysetScoping gets a reference to the given map[string]interface{} and assigns it to the QuerysetScoping field.
func (o *AccessPolicy) SetQuerysetScoping(v map[string]interface{}) {
	o.QuerysetScoping = v
}

func (o AccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionsAssignment) {
		toSerialize["permissions_assignment"] = o.PermissionsAssignment
	}
	if !IsNil(o.CreationHooks) {
		toSerialize["creation_hooks"] = o.CreationHooks
	}
	toSerialize["statements"] = o.Statements
	if !IsNil(o.QuerysetScoping) {
		toSerialize["queryset_scoping"] = o.QuerysetScoping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statements",
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

	varAccessPolicy := _AccessPolicy{}

	err = json.Unmarshal(data, &varAccessPolicy)

	if err != nil {
		return err
	}

	*o = AccessPolicy(varAccessPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions_assignment")
		delete(additionalProperties, "creation_hooks")
		delete(additionalProperties, "statements")
		delete(additionalProperties, "queryset_scoping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicy struct {
	value *AccessPolicy
	isSet bool
}

func (v NullableAccessPolicy) Get() *AccessPolicy {
	return v.value
}

func (v *NullableAccessPolicy) Set(val *AccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicy(val *AccessPolicy) *NullableAccessPolicy {
	return &NullableAccessPolicy{value: val, isSet: true}
}

func (v NullableAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


