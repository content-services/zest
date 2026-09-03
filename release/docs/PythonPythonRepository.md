# PythonPythonRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**RetainCheckpoints** | Pointer to **NullableInt64** | Retain X checkpoint publications for the repository. Default is null which retains all checkpoints. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. [Deprecated] | [optional] [default to false]
**AllowPackageSubstitution** | Pointer to **bool** | Whether to allow package substitution (replacing existing packages with packages that have the same filename but a different checksum). When False, any new repository version that would cause such a substitution will be rejected. This applies to all repository version creation paths including uploads, modify, and sync. When True (the default), package substitution is allowed. | [optional] [default to true]
**ErrorOnReject** | Pointer to **bool** | Whether to fail the entire repository version when packages are rejected by the package substitution or blocklist policies. When True (the default), a ValidationError is raised and no packages from the request are added. When False, rejected packages are skipped and remaining packages are added; skipped packages are recorded in a task progress report. | [optional] [default to true]

## Methods

### NewPythonPythonRepository

`func NewPythonPythonRepository(name string, ) *PythonPythonRepository`

NewPythonPythonRepository instantiates a new PythonPythonRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonRepositoryWithDefaults

`func NewPythonPythonRepositoryWithDefaults() *PythonPythonRepository`

NewPythonPythonRepositoryWithDefaults instantiates a new PythonPythonRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *PythonPythonRepository) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPythonRepository) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPythonRepository) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPythonRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PythonPythonRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonPythonRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonPythonRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PythonPythonRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PythonPythonRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PythonPythonRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PythonPythonRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PythonPythonRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PythonPythonRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *PythonPythonRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PythonPythonRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PythonPythonRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PythonPythonRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PythonPythonRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PythonPythonRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRetainCheckpoints

`func (o *PythonPythonRepository) GetRetainCheckpoints() int64`

GetRetainCheckpoints returns the RetainCheckpoints field if non-nil, zero value otherwise.

### GetRetainCheckpointsOk

`func (o *PythonPythonRepository) GetRetainCheckpointsOk() (*int64, bool)`

GetRetainCheckpointsOk returns a tuple with the RetainCheckpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainCheckpoints

`func (o *PythonPythonRepository) SetRetainCheckpoints(v int64)`

SetRetainCheckpoints sets RetainCheckpoints field to given value.

### HasRetainCheckpoints

`func (o *PythonPythonRepository) HasRetainCheckpoints() bool`

HasRetainCheckpoints returns a boolean if a field has been set.

### SetRetainCheckpointsNil

`func (o *PythonPythonRepository) SetRetainCheckpointsNil(b bool)`

 SetRetainCheckpointsNil sets the value for RetainCheckpoints to be an explicit nil

### UnsetRetainCheckpoints
`func (o *PythonPythonRepository) UnsetRetainCheckpoints()`

UnsetRetainCheckpoints ensures that no value is present for RetainCheckpoints, not even an explicit nil
### GetRemote

`func (o *PythonPythonRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PythonPythonRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PythonPythonRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PythonPythonRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PythonPythonRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PythonPythonRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *PythonPythonRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *PythonPythonRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *PythonPythonRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *PythonPythonRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetAllowPackageSubstitution

`func (o *PythonPythonRepository) GetAllowPackageSubstitution() bool`

GetAllowPackageSubstitution returns the AllowPackageSubstitution field if non-nil, zero value otherwise.

### GetAllowPackageSubstitutionOk

`func (o *PythonPythonRepository) GetAllowPackageSubstitutionOk() (*bool, bool)`

GetAllowPackageSubstitutionOk returns a tuple with the AllowPackageSubstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPackageSubstitution

`func (o *PythonPythonRepository) SetAllowPackageSubstitution(v bool)`

SetAllowPackageSubstitution sets AllowPackageSubstitution field to given value.

### HasAllowPackageSubstitution

`func (o *PythonPythonRepository) HasAllowPackageSubstitution() bool`

HasAllowPackageSubstitution returns a boolean if a field has been set.

### GetErrorOnReject

`func (o *PythonPythonRepository) GetErrorOnReject() bool`

GetErrorOnReject returns the ErrorOnReject field if non-nil, zero value otherwise.

### GetErrorOnRejectOk

`func (o *PythonPythonRepository) GetErrorOnRejectOk() (*bool, bool)`

GetErrorOnRejectOk returns a tuple with the ErrorOnReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnReject

`func (o *PythonPythonRepository) SetErrorOnReject(v bool)`

SetErrorOnReject sets ErrorOnReject field to given value.

### HasErrorOnReject

`func (o *PythonPythonRepository) HasErrorOnReject() bool`

HasErrorOnReject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


