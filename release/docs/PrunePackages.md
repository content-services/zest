# PrunePackages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoHrefs** | **[]string** | Will prune old packages from the specified list of repos. Use [&#39;*&#39;] to specify all repos. Will prune based on the specified repositories&#39; latest_versions. | 
**KeepDays** | Pointer to **int64** | Prune packages introduced *prior-to* this many days ago. Default is 14. A value of 0 implies &#39;keep latest package only.&#39; | [optional] [default to 14]
**DryRun** | Pointer to **bool** | Determine what would-be-pruned and log the list of packages. Intended as a debugging aid. | [optional] [default to false]

## Methods

### NewPrunePackages

`func NewPrunePackages(repoHrefs []string, ) *PrunePackages`

NewPrunePackages instantiates a new PrunePackages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrunePackagesWithDefaults

`func NewPrunePackagesWithDefaults() *PrunePackages`

NewPrunePackagesWithDefaults instantiates a new PrunePackages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoHrefs

`func (o *PrunePackages) GetRepoHrefs() []string`

GetRepoHrefs returns the RepoHrefs field if non-nil, zero value otherwise.

### GetRepoHrefsOk

`func (o *PrunePackages) GetRepoHrefsOk() (*[]string, bool)`

GetRepoHrefsOk returns a tuple with the RepoHrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoHrefs

`func (o *PrunePackages) SetRepoHrefs(v []string)`

SetRepoHrefs sets RepoHrefs field to given value.


### GetKeepDays

`func (o *PrunePackages) GetKeepDays() int64`

GetKeepDays returns the KeepDays field if non-nil, zero value otherwise.

### GetKeepDaysOk

`func (o *PrunePackages) GetKeepDaysOk() (*int64, bool)`

GetKeepDaysOk returns a tuple with the KeepDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDays

`func (o *PrunePackages) SetKeepDays(v int64)`

SetKeepDays sets KeepDays field to given value.

### HasKeepDays

`func (o *PrunePackages) HasKeepDays() bool`

HasKeepDays returns a boolean if a field has been set.

### GetDryRun

`func (o *PrunePackages) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PrunePackages) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PrunePackages) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PrunePackages) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


