# MavenMavenPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**GroupId** | Pointer to **string** | Group Id of the package. | [optional] [readonly] 
**ArtifactId** | Pointer to **string** | Artifact Id of the package. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the package. | [optional] [readonly] 
**BaseVersion** | Pointer to **string** | The package version with a trailing rebuild suffix stripped (matching \\.[a-zA-Z]+-\\d+$). Equal to version when no suffix is present. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Human-readable name from the POM. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description from the POM. | [optional] [readonly] 
**Packaging** | Pointer to **NullableString** | Packaging type (jar, war, pom, etc). | [optional] [readonly] 
**Url** | Pointer to **NullableString** | Project URL from the POM. | [optional] [readonly] 
**Licenses** | Pointer to **interface{}** | License information from the POM. | [optional] [readonly] 
**Dependencies** | Pointer to **interface{}** | Dependency list from the POM. | [optional] [readonly] 
**ScmUrl** | Pointer to **NullableString** | Source control URL from the POM. | [optional] [readonly] 

## Methods

### NewMavenMavenPackageResponse

`func NewMavenMavenPackageResponse() *MavenMavenPackageResponse`

NewMavenMavenPackageResponse instantiates a new MavenMavenPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenPackageResponseWithDefaults

`func NewMavenMavenPackageResponseWithDefaults() *MavenMavenPackageResponse`

NewMavenMavenPackageResponseWithDefaults instantiates a new MavenMavenPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *MavenMavenPackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *MavenMavenPackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *MavenMavenPackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *MavenMavenPackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *MavenMavenPackageResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *MavenMavenPackageResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *MavenMavenPackageResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *MavenMavenPackageResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *MavenMavenPackageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *MavenMavenPackageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *MavenMavenPackageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *MavenMavenPackageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *MavenMavenPackageResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *MavenMavenPackageResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *MavenMavenPackageResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *MavenMavenPackageResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenPackageResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenPackageResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenPackageResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenPackageResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *MavenMavenPackageResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *MavenMavenPackageResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *MavenMavenPackageResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *MavenMavenPackageResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetGroupId

`func (o *MavenMavenPackageResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenMavenPackageResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenMavenPackageResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MavenMavenPackageResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetArtifactId

`func (o *MavenMavenPackageResponse) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenMavenPackageResponse) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenMavenPackageResponse) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *MavenMavenPackageResponse) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetVersion

`func (o *MavenMavenPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenMavenPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenMavenPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MavenMavenPackageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBaseVersion

`func (o *MavenMavenPackageResponse) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *MavenMavenPackageResponse) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *MavenMavenPackageResponse) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.

### HasBaseVersion

`func (o *MavenMavenPackageResponse) HasBaseVersion() bool`

HasBaseVersion returns a boolean if a field has been set.

### GetName

`func (o *MavenMavenPackageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenMavenPackageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenMavenPackageResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MavenMavenPackageResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MavenMavenPackageResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MavenMavenPackageResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *MavenMavenPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MavenMavenPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MavenMavenPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MavenMavenPackageResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MavenMavenPackageResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MavenMavenPackageResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPackaging

`func (o *MavenMavenPackageResponse) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *MavenMavenPackageResponse) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *MavenMavenPackageResponse) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *MavenMavenPackageResponse) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### SetPackagingNil

`func (o *MavenMavenPackageResponse) SetPackagingNil(b bool)`

 SetPackagingNil sets the value for Packaging to be an explicit nil

### UnsetPackaging
`func (o *MavenMavenPackageResponse) UnsetPackaging()`

UnsetPackaging ensures that no value is present for Packaging, not even an explicit nil
### GetUrl

`func (o *MavenMavenPackageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MavenMavenPackageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MavenMavenPackageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MavenMavenPackageResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MavenMavenPackageResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MavenMavenPackageResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLicenses

`func (o *MavenMavenPackageResponse) GetLicenses() interface{}`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *MavenMavenPackageResponse) GetLicensesOk() (*interface{}, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *MavenMavenPackageResponse) SetLicenses(v interface{})`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *MavenMavenPackageResponse) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### SetLicensesNil

`func (o *MavenMavenPackageResponse) SetLicensesNil(b bool)`

 SetLicensesNil sets the value for Licenses to be an explicit nil

### UnsetLicenses
`func (o *MavenMavenPackageResponse) UnsetLicenses()`

UnsetLicenses ensures that no value is present for Licenses, not even an explicit nil
### GetDependencies

`func (o *MavenMavenPackageResponse) GetDependencies() interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *MavenMavenPackageResponse) GetDependenciesOk() (*interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *MavenMavenPackageResponse) SetDependencies(v interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *MavenMavenPackageResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *MavenMavenPackageResponse) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *MavenMavenPackageResponse) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetScmUrl

`func (o *MavenMavenPackageResponse) GetScmUrl() string`

GetScmUrl returns the ScmUrl field if non-nil, zero value otherwise.

### GetScmUrlOk

`func (o *MavenMavenPackageResponse) GetScmUrlOk() (*string, bool)`

GetScmUrlOk returns a tuple with the ScmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUrl

`func (o *MavenMavenPackageResponse) SetScmUrl(v string)`

SetScmUrl sets ScmUrl field to given value.

### HasScmUrl

`func (o *MavenMavenPackageResponse) HasScmUrl() bool`

HasScmUrl returns a boolean if a field has been set.

### SetScmUrlNil

`func (o *MavenMavenPackageResponse) SetScmUrlNil(b bool)`

 SetScmUrlNil sets the value for ScmUrl to be an explicit nil

### UnsetScmUrl
`func (o *MavenMavenPackageResponse) UnsetScmUrl()`

UnsetScmUrl ensures that no value is present for ScmUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


