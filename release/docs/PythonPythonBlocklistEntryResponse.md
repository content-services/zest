# PythonPythonBlocklistEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** | The URL of this blocklist entry. | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Repository** | Pointer to **string** | Repository this blocklist entry belongs to. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Package name to block (for all versions). Compared after PEP 503 normalization. Required when &#39;filename&#39; is not provided. | [optional] 
**Version** | Pointer to **NullableString** | Exact version string to block (e.g. &#39;1.0&#39;). Only used when &#39;name&#39; is set. | [optional] 
**Filename** | Pointer to **NullableString** | Exact filename to block. Required when &#39;name&#39; is not provided. | [optional] 
**AddedBy** | Pointer to **string** | PRN of the user who added this blocklist entry. | [optional] [readonly] 

## Methods

### NewPythonPythonBlocklistEntryResponse

`func NewPythonPythonBlocklistEntryResponse() *PythonPythonBlocklistEntryResponse`

NewPythonPythonBlocklistEntryResponse instantiates a new PythonPythonBlocklistEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonBlocklistEntryResponseWithDefaults

`func NewPythonPythonBlocklistEntryResponseWithDefaults() *PythonPythonBlocklistEntryResponse`

NewPythonPythonBlocklistEntryResponseWithDefaults instantiates a new PythonPythonBlocklistEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PythonPythonBlocklistEntryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PythonPythonBlocklistEntryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PythonPythonBlocklistEntryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PythonPythonBlocklistEntryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *PythonPythonBlocklistEntryResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *PythonPythonBlocklistEntryResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *PythonPythonBlocklistEntryResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *PythonPythonBlocklistEntryResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PythonPythonBlocklistEntryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PythonPythonBlocklistEntryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PythonPythonBlocklistEntryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PythonPythonBlocklistEntryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *PythonPythonBlocklistEntryResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *PythonPythonBlocklistEntryResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *PythonPythonBlocklistEntryResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *PythonPythonBlocklistEntryResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetRepository

`func (o *PythonPythonBlocklistEntryResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PythonPythonBlocklistEntryResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PythonPythonBlocklistEntryResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PythonPythonBlocklistEntryResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetName

`func (o *PythonPythonBlocklistEntryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonPythonBlocklistEntryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonPythonBlocklistEntryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PythonPythonBlocklistEntryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PythonPythonBlocklistEntryResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PythonPythonBlocklistEntryResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *PythonPythonBlocklistEntryResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PythonPythonBlocklistEntryResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PythonPythonBlocklistEntryResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PythonPythonBlocklistEntryResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PythonPythonBlocklistEntryResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PythonPythonBlocklistEntryResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFilename

`func (o *PythonPythonBlocklistEntryResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PythonPythonBlocklistEntryResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PythonPythonBlocklistEntryResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PythonPythonBlocklistEntryResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *PythonPythonBlocklistEntryResponse) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *PythonPythonBlocklistEntryResponse) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetAddedBy

`func (o *PythonPythonBlocklistEntryResponse) GetAddedBy() string`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *PythonPythonBlocklistEntryResponse) GetAddedByOk() (*string, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *PythonPythonBlocklistEntryResponse) SetAddedBy(v string)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *PythonPythonBlocklistEntryResponse) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


