# MavenRepositoryFlatPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Maven groupId. | 
**ArtifactId** | **string** | Maven artifactId. | 
**Version** | **string** | Version stored on the MavenPackage, including any rebuild suffix. | 
**LastUpdated** | **time.Time** | When this GAV entered the repository version: RepositoryContent.pulp_created, falling back to the content unit&#39;s pulp_created. | 
**Description** | **string** | Description from the POM. Empty when the POM has none. | 
**Licenses** | [**[]MavenPackageLicenseResponse**](MavenPackageLicenseResponse.md) | Licenses from the POM. Empty when the POM has none. | 

## Methods

### NewMavenRepositoryFlatPackageResponse

`func NewMavenRepositoryFlatPackageResponse(groupId string, artifactId string, version string, lastUpdated time.Time, description string, licenses []MavenPackageLicenseResponse, ) *MavenRepositoryFlatPackageResponse`

NewMavenRepositoryFlatPackageResponse instantiates a new MavenRepositoryFlatPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenRepositoryFlatPackageResponseWithDefaults

`func NewMavenRepositoryFlatPackageResponseWithDefaults() *MavenRepositoryFlatPackageResponse`

NewMavenRepositoryFlatPackageResponseWithDefaults instantiates a new MavenRepositoryFlatPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *MavenRepositoryFlatPackageResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenRepositoryFlatPackageResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenRepositoryFlatPackageResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetArtifactId

`func (o *MavenRepositoryFlatPackageResponse) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenRepositoryFlatPackageResponse) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenRepositoryFlatPackageResponse) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.


### GetVersion

`func (o *MavenRepositoryFlatPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenRepositoryFlatPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenRepositoryFlatPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLastUpdated

`func (o *MavenRepositoryFlatPackageResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MavenRepositoryFlatPackageResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MavenRepositoryFlatPackageResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDescription

`func (o *MavenRepositoryFlatPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MavenRepositoryFlatPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MavenRepositoryFlatPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLicenses

`func (o *MavenRepositoryFlatPackageResponse) GetLicenses() []MavenPackageLicenseResponse`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *MavenRepositoryFlatPackageResponse) GetLicensesOk() (*[]MavenPackageLicenseResponse, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *MavenRepositoryFlatPackageResponse) SetLicenses(v []MavenPackageLicenseResponse)`

SetLicenses sets Licenses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


