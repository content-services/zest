# MavenPackageReleaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Logical version key (rebuild suffix stripped). | 
**Release** | **string** | Rebuild/release qualifier within the version line (e.g. rhlw-00003). Empty when the selected unit has no rebuild suffix. | 
**CreatedAt** | **time.Time** | When this logical version entered the repository: RepositoryContent.pulp_created of the newest rebuild, falling back to the content unit&#39;s pulp_created. | 

## Methods

### NewMavenPackageReleaseResponse

`func NewMavenPackageReleaseResponse(version string, release string, createdAt time.Time, ) *MavenPackageReleaseResponse`

NewMavenPackageReleaseResponse instantiates a new MavenPackageReleaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenPackageReleaseResponseWithDefaults

`func NewMavenPackageReleaseResponseWithDefaults() *MavenPackageReleaseResponse`

NewMavenPackageReleaseResponseWithDefaults instantiates a new MavenPackageReleaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *MavenPackageReleaseResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenPackageReleaseResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenPackageReleaseResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRelease

`func (o *MavenPackageReleaseResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *MavenPackageReleaseResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *MavenPackageReleaseResponse) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetCreatedAt

`func (o *MavenPackageReleaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MavenPackageReleaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MavenPackageReleaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


