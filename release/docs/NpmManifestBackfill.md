# NpmManifestBackfill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | **string** | Npm repository href or UUID. | 
**Overwrite** | Pointer to **bool** | Replace existing manifests instead of only filling empty ones. | [optional] [default to false]

## Methods

### NewNpmManifestBackfill

`func NewNpmManifestBackfill(repository string, ) *NpmManifestBackfill`

NewNpmManifestBackfill instantiates a new NpmManifestBackfill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpmManifestBackfillWithDefaults

`func NewNpmManifestBackfillWithDefaults() *NpmManifestBackfill`

NewNpmManifestBackfillWithDefaults instantiates a new NpmManifestBackfill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *NpmManifestBackfill) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NpmManifestBackfill) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NpmManifestBackfill) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetOverwrite

`func (o *NpmManifestBackfill) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *NpmManifestBackfill) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *NpmManifestBackfill) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *NpmManifestBackfill) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


