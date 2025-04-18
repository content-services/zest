# ContentScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoVersion** | Pointer to **NullableString** | RepositoryVersion HREF with the packages to be checked. | [optional] 
**PackageJson** | Pointer to ***os.File** | package-lock.json file with the definition of dependencies to be checked. | [optional] 

## Methods

### NewContentScan

`func NewContentScan() *ContentScan`

NewContentScan instantiates a new ContentScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentScanWithDefaults

`func NewContentScanWithDefaults() *ContentScan`

NewContentScanWithDefaults instantiates a new ContentScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoVersion

`func (o *ContentScan) GetRepoVersion() string`

GetRepoVersion returns the RepoVersion field if non-nil, zero value otherwise.

### GetRepoVersionOk

`func (o *ContentScan) GetRepoVersionOk() (*string, bool)`

GetRepoVersionOk returns a tuple with the RepoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoVersion

`func (o *ContentScan) SetRepoVersion(v string)`

SetRepoVersion sets RepoVersion field to given value.

### HasRepoVersion

`func (o *ContentScan) HasRepoVersion() bool`

HasRepoVersion returns a boolean if a field has been set.

### SetRepoVersionNil

`func (o *ContentScan) SetRepoVersionNil(b bool)`

 SetRepoVersionNil sets the value for RepoVersion to be an explicit nil

### UnsetRepoVersion
`func (o *ContentScan) UnsetRepoVersion()`

UnsetRepoVersion ensures that no value is present for RepoVersion, not even an explicit nil
### GetPackageJson

`func (o *ContentScan) GetPackageJson() *os.File`

GetPackageJson returns the PackageJson field if non-nil, zero value otherwise.

### GetPackageJsonOk

`func (o *ContentScan) GetPackageJsonOk() (**os.File, bool)`

GetPackageJsonOk returns a tuple with the PackageJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageJson

`func (o *ContentScan) SetPackageJson(v *os.File)`

SetPackageJson sets PackageJson field to given value.

### HasPackageJson

`func (o *ContentScan) HasPackageJson() bool`

HasPackageJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


