# MavenRepositoryMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageCount** | **int64** | Distinct (group_id, artifact_id) pairs among MavenPackage units. | 
**VersionCount** | **int64** | Distinct (group_id, artifact_id, base_version) triples after rebuild-suffix strip. | 
**BuildCount** | **int64** | Distinct (group_id, artifact_id, full version) GAVs among MavenPackage units. | 

## Methods

### NewMavenRepositoryMetricsResponse

`func NewMavenRepositoryMetricsResponse(packageCount int64, versionCount int64, buildCount int64, ) *MavenRepositoryMetricsResponse`

NewMavenRepositoryMetricsResponse instantiates a new MavenRepositoryMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenRepositoryMetricsResponseWithDefaults

`func NewMavenRepositoryMetricsResponseWithDefaults() *MavenRepositoryMetricsResponse`

NewMavenRepositoryMetricsResponseWithDefaults instantiates a new MavenRepositoryMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageCount

`func (o *MavenRepositoryMetricsResponse) GetPackageCount() int64`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *MavenRepositoryMetricsResponse) GetPackageCountOk() (*int64, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *MavenRepositoryMetricsResponse) SetPackageCount(v int64)`

SetPackageCount sets PackageCount field to given value.


### GetVersionCount

`func (o *MavenRepositoryMetricsResponse) GetVersionCount() int64`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *MavenRepositoryMetricsResponse) GetVersionCountOk() (*int64, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *MavenRepositoryMetricsResponse) SetVersionCount(v int64)`

SetVersionCount sets VersionCount field to given value.


### GetBuildCount

`func (o *MavenRepositoryMetricsResponse) GetBuildCount() int64`

GetBuildCount returns the BuildCount field if non-nil, zero value otherwise.

### GetBuildCountOk

`func (o *MavenRepositoryMetricsResponse) GetBuildCountOk() (*int64, bool)`

GetBuildCountOk returns a tuple with the BuildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCount

`func (o *MavenRepositoryMetricsResponse) SetBuildCount(v int64)`

SetBuildCount sets BuildCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


