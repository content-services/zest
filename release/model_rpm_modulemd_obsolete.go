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

// checks if the RpmModulemdObsolete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmModulemdObsolete{}

// RpmModulemdObsolete ModulemdObsolete serializer.
type RpmModulemdObsolete struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// Obsolete modified time.
	Modified string `json:"modified"`
	// Modulemd name.
	ModuleName string `json:"module_name"`
	// Modulemd's stream.
	ModuleStream string `json:"module_stream"`
	// Obsolete description.
	Message string `json:"message"`
	// Reset previous obsoletes.
	OverridePrevious NullableString `json:"override_previous"`
	// Modulemd's context.
	ModuleContext NullableString `json:"module_context"`
	// End of Life date.
	EolDate NullableString `json:"eol_date"`
	// Obsolete by module name.
	ObsoletedByModuleName NullableString `json:"obsoleted_by_module_name"`
	// Obsolete by module stream.
	ObsoletedByModuleStream NullableString `json:"obsoleted_by_module_stream"`
	// Module Obsolete snippet.
	Snippet string `json:"snippet"`
}

type _RpmModulemdObsolete RpmModulemdObsolete

// NewRpmModulemdObsolete instantiates a new RpmModulemdObsolete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmModulemdObsolete(modified string, moduleName string, moduleStream string, message string, overridePrevious NullableString, moduleContext NullableString, eolDate NullableString, obsoletedByModuleName NullableString, obsoletedByModuleStream NullableString, snippet string) *RpmModulemdObsolete {
	this := RpmModulemdObsolete{}
	this.Modified = modified
	this.ModuleName = moduleName
	this.ModuleStream = moduleStream
	this.Message = message
	this.OverridePrevious = overridePrevious
	this.ModuleContext = moduleContext
	this.EolDate = eolDate
	this.ObsoletedByModuleName = obsoletedByModuleName
	this.ObsoletedByModuleStream = obsoletedByModuleStream
	this.Snippet = snippet
	return &this
}

// NewRpmModulemdObsoleteWithDefaults instantiates a new RpmModulemdObsolete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmModulemdObsoleteWithDefaults() *RpmModulemdObsolete {
	this := RpmModulemdObsolete{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RpmModulemdObsolete) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RpmModulemdObsolete) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RpmModulemdObsolete) SetRepository(v string) {
	o.Repository = &v
}

// GetModified returns the Modified field value
func (o *RpmModulemdObsolete) GetModified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *RpmModulemdObsolete) SetModified(v string) {
	o.Modified = v
}

// GetModuleName returns the ModuleName field value
func (o *RpmModulemdObsolete) GetModuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetModuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleName, true
}

// SetModuleName sets field value
func (o *RpmModulemdObsolete) SetModuleName(v string) {
	o.ModuleName = v
}

// GetModuleStream returns the ModuleStream field value
func (o *RpmModulemdObsolete) GetModuleStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleStream
}

// GetModuleStreamOk returns a tuple with the ModuleStream field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetModuleStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleStream, true
}

// SetModuleStream sets field value
func (o *RpmModulemdObsolete) SetModuleStream(v string) {
	o.ModuleStream = v
}

// GetMessage returns the Message field value
func (o *RpmModulemdObsolete) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RpmModulemdObsolete) SetMessage(v string) {
	o.Message = v
}

// GetOverridePrevious returns the OverridePrevious field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmModulemdObsolete) GetOverridePrevious() string {
	if o == nil || o.OverridePrevious.Get() == nil {
		var ret string
		return ret
	}

	return *o.OverridePrevious.Get()
}

// GetOverridePreviousOk returns a tuple with the OverridePrevious field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdObsolete) GetOverridePreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverridePrevious.Get(), o.OverridePrevious.IsSet()
}

// SetOverridePrevious sets field value
func (o *RpmModulemdObsolete) SetOverridePrevious(v string) {
	o.OverridePrevious.Set(&v)
}

// GetModuleContext returns the ModuleContext field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmModulemdObsolete) GetModuleContext() string {
	if o == nil || o.ModuleContext.Get() == nil {
		var ret string
		return ret
	}

	return *o.ModuleContext.Get()
}

// GetModuleContextOk returns a tuple with the ModuleContext field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdObsolete) GetModuleContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleContext.Get(), o.ModuleContext.IsSet()
}

// SetModuleContext sets field value
func (o *RpmModulemdObsolete) SetModuleContext(v string) {
	o.ModuleContext.Set(&v)
}

// GetEolDate returns the EolDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmModulemdObsolete) GetEolDate() string {
	if o == nil || o.EolDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EolDate.Get()
}

// GetEolDateOk returns a tuple with the EolDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdObsolete) GetEolDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EolDate.Get(), o.EolDate.IsSet()
}

// SetEolDate sets field value
func (o *RpmModulemdObsolete) SetEolDate(v string) {
	o.EolDate.Set(&v)
}

// GetObsoletedByModuleName returns the ObsoletedByModuleName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmModulemdObsolete) GetObsoletedByModuleName() string {
	if o == nil || o.ObsoletedByModuleName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObsoletedByModuleName.Get()
}

// GetObsoletedByModuleNameOk returns a tuple with the ObsoletedByModuleName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdObsolete) GetObsoletedByModuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObsoletedByModuleName.Get(), o.ObsoletedByModuleName.IsSet()
}

// SetObsoletedByModuleName sets field value
func (o *RpmModulemdObsolete) SetObsoletedByModuleName(v string) {
	o.ObsoletedByModuleName.Set(&v)
}

// GetObsoletedByModuleStream returns the ObsoletedByModuleStream field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmModulemdObsolete) GetObsoletedByModuleStream() string {
	if o == nil || o.ObsoletedByModuleStream.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObsoletedByModuleStream.Get()
}

// GetObsoletedByModuleStreamOk returns a tuple with the ObsoletedByModuleStream field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdObsolete) GetObsoletedByModuleStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObsoletedByModuleStream.Get(), o.ObsoletedByModuleStream.IsSet()
}

// SetObsoletedByModuleStream sets field value
func (o *RpmModulemdObsolete) SetObsoletedByModuleStream(v string) {
	o.ObsoletedByModuleStream.Set(&v)
}

// GetSnippet returns the Snippet field value
func (o *RpmModulemdObsolete) GetSnippet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdObsolete) GetSnippetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snippet, true
}

// SetSnippet sets field value
func (o *RpmModulemdObsolete) SetSnippet(v string) {
	o.Snippet = v
}

func (o RpmModulemdObsolete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmModulemdObsolete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	toSerialize["modified"] = o.Modified
	toSerialize["module_name"] = o.ModuleName
	toSerialize["module_stream"] = o.ModuleStream
	toSerialize["message"] = o.Message
	toSerialize["override_previous"] = o.OverridePrevious.Get()
	toSerialize["module_context"] = o.ModuleContext.Get()
	toSerialize["eol_date"] = o.EolDate.Get()
	toSerialize["obsoleted_by_module_name"] = o.ObsoletedByModuleName.Get()
	toSerialize["obsoleted_by_module_stream"] = o.ObsoletedByModuleStream.Get()
	toSerialize["snippet"] = o.Snippet
	return toSerialize, nil
}

func (o *RpmModulemdObsolete) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modified",
		"module_name",
		"module_stream",
		"message",
		"override_previous",
		"module_context",
		"eol_date",
		"obsoleted_by_module_name",
		"obsoleted_by_module_stream",
		"snippet",
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

	varRpmModulemdObsolete := _RpmModulemdObsolete{}

	err = json.Unmarshal(bytes, &varRpmModulemdObsolete)

	if err != nil {
		return err
	}

	*o = RpmModulemdObsolete(varRpmModulemdObsolete)

	return err
}

type NullableRpmModulemdObsolete struct {
	value *RpmModulemdObsolete
	isSet bool
}

func (v NullableRpmModulemdObsolete) Get() *RpmModulemdObsolete {
	return v.value
}

func (v *NullableRpmModulemdObsolete) Set(val *RpmModulemdObsolete) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmModulemdObsolete) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmModulemdObsolete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmModulemdObsolete(val *RpmModulemdObsolete) *NullableRpmModulemdObsolete {
	return &NullableRpmModulemdObsolete{value: val, isSet: true}
}

func (v NullableRpmModulemdObsolete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmModulemdObsolete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


