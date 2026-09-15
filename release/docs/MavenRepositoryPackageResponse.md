# MavenRepositoryPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Maven groupId. Index rows are unique on GA. | 
**ArtifactId** | **string** | Maven artifactId. Index rows are unique on GA. | 
**LastUpdated** | **NullableTime** | When this package was last updated in the repository: the latest RepositoryContent.pulp_created among all MavenPackage units for this GA (any rebuild), falling back to the content unit&#39;s pulp_created. | 
**Versions** | **[]string** | Distinct logical version keys after rebuild-suffix strip, newest first. The set of values matches latest_releases[].version. | 
**LatestReleases** | [**[]MavenPackageReleaseResponse**](MavenPackageReleaseResponse.md) | Newest rebuild per logical version (latest pulp_created), newest version first. set(versions) &#x3D;&#x3D;&#x3D; set(latest_releases[].version). | 

## Methods

### NewMavenRepositoryPackageResponse

`func NewMavenRepositoryPackageResponse(groupId string, artifactId string, lastUpdated NullableTime, versions []string, latestReleases []MavenPackageReleaseResponse, ) *MavenRepositoryPackageResponse`

NewMavenRepositoryPackageResponse instantiates a new MavenRepositoryPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenRepositoryPackageResponseWithDefaults

`func NewMavenRepositoryPackageResponseWithDefaults() *MavenRepositoryPackageResponse`

NewMavenRepositoryPackageResponseWithDefaults instantiates a new MavenRepositoryPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *MavenRepositoryPackageResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenRepositoryPackageResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenRepositoryPackageResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetArtifactId

`func (o *MavenRepositoryPackageResponse) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenRepositoryPackageResponse) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenRepositoryPackageResponse) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.


### GetLastUpdated

`func (o *MavenRepositoryPackageResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MavenRepositoryPackageResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MavenRepositoryPackageResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *MavenRepositoryPackageResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *MavenRepositoryPackageResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetVersions

`func (o *MavenRepositoryPackageResponse) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MavenRepositoryPackageResponse) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MavenRepositoryPackageResponse) SetVersions(v []string)`

SetVersions sets Versions field to given value.


### GetLatestReleases

`func (o *MavenRepositoryPackageResponse) GetLatestReleases() []MavenPackageReleaseResponse`

GetLatestReleases returns the LatestReleases field if non-nil, zero value otherwise.

### GetLatestReleasesOk

`func (o *MavenRepositoryPackageResponse) GetLatestReleasesOk() (*[]MavenPackageReleaseResponse, bool)`

GetLatestReleasesOk returns a tuple with the LatestReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleases

`func (o *MavenRepositoryPackageResponse) SetLatestReleases(v []MavenPackageReleaseResponse)`

SetLatestReleases sets LatestReleases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


