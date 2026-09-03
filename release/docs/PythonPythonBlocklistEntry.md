# PythonPythonBlocklistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Package name to block (for all versions). Compared after PEP 503 normalization. Required when &#39;filename&#39; is not provided. | [optional] 
**Version** | Pointer to **NullableString** | Exact version string to block (e.g. &#39;1.0&#39;). Only used when &#39;name&#39; is set. | [optional] 
**Filename** | Pointer to **NullableString** | Exact filename to block. Required when &#39;name&#39; is not provided. | [optional] 

## Methods

### NewPythonPythonBlocklistEntry

`func NewPythonPythonBlocklistEntry() *PythonPythonBlocklistEntry`

NewPythonPythonBlocklistEntry instantiates a new PythonPythonBlocklistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonBlocklistEntryWithDefaults

`func NewPythonPythonBlocklistEntryWithDefaults() *PythonPythonBlocklistEntry`

NewPythonPythonBlocklistEntryWithDefaults instantiates a new PythonPythonBlocklistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PythonPythonBlocklistEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonPythonBlocklistEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonPythonBlocklistEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PythonPythonBlocklistEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PythonPythonBlocklistEntry) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PythonPythonBlocklistEntry) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *PythonPythonBlocklistEntry) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PythonPythonBlocklistEntry) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PythonPythonBlocklistEntry) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PythonPythonBlocklistEntry) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PythonPythonBlocklistEntry) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PythonPythonBlocklistEntry) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFilename

`func (o *PythonPythonBlocklistEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PythonPythonBlocklistEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PythonPythonBlocklistEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PythonPythonBlocklistEntry) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *PythonPythonBlocklistEntry) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *PythonPythonBlocklistEntry) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


