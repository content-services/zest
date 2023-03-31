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

// checks if the PatchedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUser{}

// PatchedUser Serializer for User.
type PatchedUser struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username *string `json:"username,omitempty"`
	// Users password. Set to ``null`` to disable password authentication.
	Password NullableString `json:"password,omitempty"`
	// First name
	FirstName *string `json:"first_name,omitempty"`
	// Last name
	LastName *string `json:"last_name,omitempty"`
	// Email address
	Email *string `json:"email,omitempty"`
	// Designates whether the user can log into this admin site.
	IsStaff *bool `json:"is_staff,omitempty"`
	// Designates whether this user should be treated as active.
	IsActive *bool `json:"is_active,omitempty"`
}

// NewPatchedUser instantiates a new PatchedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUser() *PatchedUser {
	this := PatchedUser{}
	var isStaff bool = false
	this.IsStaff = &isStaff
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewPatchedUserWithDefaults instantiates a new PatchedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserWithDefaults() *PatchedUser {
	this := PatchedUser{}
	var isStaff bool = false
	this.IsStaff = &isStaff
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedUser) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUser) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUser) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedUser) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PatchedUser) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PatchedUser) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PatchedUser) UnsetPassword() {
	o.Password.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PatchedUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PatchedUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PatchedUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PatchedUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PatchedUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PatchedUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedUser) SetEmail(v string) {
	o.Email = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *PatchedUser) GetIsStaff() bool {
	if o == nil || IsNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetIsStaffOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaff) {
		return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *PatchedUser) HasIsStaff() bool {
	if o != nil && !IsNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *PatchedUser) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedUser) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUser) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedUser) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedUser) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o PatchedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsStaff) {
		toSerialize["is_staff"] = o.IsStaff
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	return toSerialize, nil
}

type NullablePatchedUser struct {
	value *PatchedUser
	isSet bool
}

func (v NullablePatchedUser) Get() *PatchedUser {
	return v.value
}

func (v *NullablePatchedUser) Set(val *PatchedUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUser(val *PatchedUser) *NullablePatchedUser {
	return &NullablePatchedUser{value: val, isSet: true}
}

func (v NullablePatchedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


