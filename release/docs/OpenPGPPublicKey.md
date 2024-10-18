# OpenPGPPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 

## Methods

### NewOpenPGPPublicKey

`func NewOpenPGPPublicKey() *OpenPGPPublicKey`

NewOpenPGPPublicKey instantiates a new OpenPGPPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPGPPublicKeyWithDefaults

`func NewOpenPGPPublicKeyWithDefaults() *OpenPGPPublicKey`

NewOpenPGPPublicKeyWithDefaults instantiates a new OpenPGPPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *OpenPGPPublicKey) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OpenPGPPublicKey) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OpenPGPPublicKey) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *OpenPGPPublicKey) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetFile

`func (o *OpenPGPPublicKey) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *OpenPGPPublicKey) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *OpenPGPPublicKey) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *OpenPGPPublicKey) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *OpenPGPPublicKey) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *OpenPGPPublicKey) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *OpenPGPPublicKey) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *OpenPGPPublicKey) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *OpenPGPPublicKey) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *OpenPGPPublicKey) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *OpenPGPPublicKey) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *OpenPGPPublicKey) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


