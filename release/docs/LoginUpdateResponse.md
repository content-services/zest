# LoginUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [optional] 
**FirstName** | Pointer to **string** | First name | [optional] 
**LastName** | Pointer to **string** | Last name | [optional] 
**Email** | Pointer to **string** | Email address | [optional] 

## Methods

### NewLoginUpdateResponse

`func NewLoginUpdateResponse() *LoginUpdateResponse`

NewLoginUpdateResponse instantiates a new LoginUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginUpdateResponseWithDefaults

`func NewLoginUpdateResponseWithDefaults() *LoginUpdateResponse`

NewLoginUpdateResponseWithDefaults instantiates a new LoginUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LoginUpdateResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoginUpdateResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoginUpdateResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoginUpdateResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *LoginUpdateResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LoginUpdateResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LoginUpdateResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LoginUpdateResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *LoginUpdateResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LoginUpdateResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LoginUpdateResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LoginUpdateResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *LoginUpdateResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginUpdateResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginUpdateResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoginUpdateResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


