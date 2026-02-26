# PatchedcontainerContainerPullThroughDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | Pointer to **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**ContentGuard** | Pointer to **string** | An optional content-guard. If none is specified, a default one will be used. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**Name** | Pointer to **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | [optional] 
**Remote** | Pointer to **string** | Remote that can be used to fetch content when using pull-through caching. | [optional] 
**Distributions** | Pointer to **[]string** | Distributions created after pulling content through cache | [optional] 
**Private** | Pointer to **bool** | Restrict pull access to explicitly authorized users. Related distributions inherit this value. Defaults to unrestricted pull access. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 

## Methods

### NewPatchedcontainerContainerPullThroughDistribution

`func NewPatchedcontainerContainerPullThroughDistribution() *PatchedcontainerContainerPullThroughDistribution`

NewPatchedcontainerContainerPullThroughDistribution instantiates a new PatchedcontainerContainerPullThroughDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedcontainerContainerPullThroughDistributionWithDefaults

`func NewPatchedcontainerContainerPullThroughDistributionWithDefaults() *PatchedcontainerContainerPullThroughDistribution`

NewPatchedcontainerContainerPullThroughDistributionWithDefaults instantiates a new PatchedcontainerContainerPullThroughDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PatchedcontainerContainerPullThroughDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedcontainerContainerPullThroughDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedcontainerContainerPullThroughDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedcontainerContainerPullThroughDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedcontainerContainerPullThroughDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedcontainerContainerPullThroughDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedcontainerContainerPullThroughDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedcontainerContainerPullThroughDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetPulpLabels

`func (o *PatchedcontainerContainerPullThroughDistribution) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedcontainerContainerPullThroughDistribution) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedcontainerContainerPullThroughDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetContentGuard

`func (o *PatchedcontainerContainerPullThroughDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedcontainerContainerPullThroughDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedcontainerContainerPullThroughDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### GetHidden

`func (o *PatchedcontainerContainerPullThroughDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedcontainerContainerPullThroughDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedcontainerContainerPullThroughDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetName

`func (o *PatchedcontainerContainerPullThroughDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedcontainerContainerPullThroughDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedcontainerContainerPullThroughDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemote

`func (o *PatchedcontainerContainerPullThroughDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedcontainerContainerPullThroughDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedcontainerContainerPullThroughDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetDistributions

`func (o *PatchedcontainerContainerPullThroughDistribution) GetDistributions() []string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetDistributionsOk() (*[]string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *PatchedcontainerContainerPullThroughDistribution) SetDistributions(v []string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *PatchedcontainerContainerPullThroughDistribution) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### GetPrivate

`func (o *PatchedcontainerContainerPullThroughDistribution) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PatchedcontainerContainerPullThroughDistribution) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *PatchedcontainerContainerPullThroughDistribution) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedcontainerContainerPullThroughDistribution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedcontainerContainerPullThroughDistribution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedcontainerContainerPullThroughDistribution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedcontainerContainerPullThroughDistribution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedcontainerContainerPullThroughDistribution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedcontainerContainerPullThroughDistribution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


