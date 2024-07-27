# RpmPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Md5** | Pointer to **string** | The MD5 checksum if available. | [optional] [readonly] 
**Sha1** | Pointer to **string** | The SHA-1 checksum if available. | [optional] [readonly] 
**Sha224** | Pointer to **string** | The SHA-224 checksum if available. | [optional] [readonly] 
**Sha256** | Pointer to **string** | The SHA-256 checksum if available. | [optional] [readonly] 
**Sha384** | Pointer to **string** | The SHA-384 checksum if available. | [optional] [readonly] 
**Sha512** | Pointer to **string** | The SHA-512 checksum if available. | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**Name** | Pointer to **string** | Name of the package | [optional] [readonly] 
**Epoch** | Pointer to **string** | The package&#39;s epoch | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the package. For example, &#39;2.8.0&#39; | [optional] [readonly] 
**Release** | Pointer to **string** | The release of a particular version of the package. e.g. &#39;1.el7&#39; or &#39;3.f24&#39; | [optional] [readonly] 
**Arch** | Pointer to **string** | The target architecture for a package.For example, &#39;x86_64&#39;, &#39;i686&#39;, or &#39;noarch&#39; | [optional] [readonly] 
**PkgId** | Pointer to **string** | Checksum of the package file | [optional] [readonly] 
**ChecksumType** | Pointer to **string** | Type of checksum, e.g. &#39;sha256&#39;, &#39;md5&#39; | [optional] [readonly] 
**Summary** | Pointer to **string** | Short description of the packaged software | [optional] [readonly] 
**Description** | Pointer to **string** | In-depth description of the packaged software | [optional] [readonly] 
**Url** | Pointer to **string** | URL with more information about the packaged software | [optional] [readonly] 
**Changelogs** | Pointer to **interface{}** | Changelogs that package contains | [optional] [readonly] [default to []]
**Files** | Pointer to **interface{}** | Files that package contains | [optional] [readonly] [default to []]
**Requires** | Pointer to **interface{}** | Capabilities the package requires | [optional] [readonly] [default to []]
**Provides** | Pointer to **interface{}** | Capabilities the package provides | [optional] [readonly] [default to []]
**Conflicts** | Pointer to **interface{}** | Capabilities the package conflicts | [optional] [readonly] [default to []]
**Obsoletes** | Pointer to **interface{}** | Capabilities the package obsoletes | [optional] [readonly] [default to []]
**Suggests** | Pointer to **interface{}** | Capabilities the package suggests | [optional] [readonly] [default to []]
**Enhances** | Pointer to **interface{}** | Capabilities the package enhances | [optional] [readonly] [default to []]
**Recommends** | Pointer to **interface{}** | Capabilities the package recommends | [optional] [readonly] [default to []]
**Supplements** | Pointer to **interface{}** | Capabilities the package supplements | [optional] [readonly] [default to []]
**LocationBase** | Pointer to **string** | Base location of this package | [optional] [readonly] 
**LocationHref** | Pointer to **string** | Relative location of package to the repodata | [optional] [readonly] 
**RpmBuildhost** | Pointer to **string** | Hostname of the system that built the package | [optional] [readonly] 
**RpmGroup** | Pointer to **string** | RPM group (See: http://fedoraproject.org/wiki/RPMGroups) | [optional] [readonly] 
**RpmLicense** | Pointer to **string** | License term applicable to the package software (GPLv2, etc.) | [optional] [readonly] 
**RpmPackager** | Pointer to **string** | Person or persons responsible for creating the package | [optional] [readonly] 
**RpmSourcerpm** | Pointer to **string** | Name of the source package (srpm) the package was built from | [optional] [readonly] 
**RpmVendor** | Pointer to **string** | Name of the organization that produced the package | [optional] [readonly] 
**RpmHeaderStart** | Pointer to **int64** | First byte of the header | [optional] [readonly] 
**RpmHeaderEnd** | Pointer to **int64** | Last byte of the header | [optional] [readonly] 
**IsModular** | Pointer to **bool** | Flag to identify if the package is modular | [optional] [readonly] 
**SizeArchive** | Pointer to **int64** | Size, in bytes, of the archive portion of the original package file | [optional] [readonly] 
**SizeInstalled** | Pointer to **int64** | Total size, in bytes, of every file installed by this package | [optional] [readonly] 
**SizePackage** | Pointer to **int64** | Size, in bytes, of the package | [optional] [readonly] 
**TimeBuild** | Pointer to **int64** | Time the package was built in seconds since the epoch | [optional] [readonly] 
**TimeFile** | Pointer to **int64** | The &#39;file&#39; time attribute in the primary XML - file mtime in seconds since the epoch. | [optional] [readonly] 

## Methods

### NewRpmPackageResponse

`func NewRpmPackageResponse() *RpmPackageResponse`

NewRpmPackageResponse instantiates a new RpmPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmPackageResponseWithDefaults

`func NewRpmPackageResponseWithDefaults() *RpmPackageResponse`

NewRpmPackageResponseWithDefaults instantiates a new RpmPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmPackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmPackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmPackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmPackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmPackageResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmPackageResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmPackageResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmPackageResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RpmPackageResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RpmPackageResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RpmPackageResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RpmPackageResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetMd5

`func (o *RpmPackageResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *RpmPackageResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *RpmPackageResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *RpmPackageResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *RpmPackageResponse) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *RpmPackageResponse) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *RpmPackageResponse) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *RpmPackageResponse) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha224

`func (o *RpmPackageResponse) GetSha224() string`

GetSha224 returns the Sha224 field if non-nil, zero value otherwise.

### GetSha224Ok

`func (o *RpmPackageResponse) GetSha224Ok() (*string, bool)`

GetSha224Ok returns a tuple with the Sha224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha224

`func (o *RpmPackageResponse) SetSha224(v string)`

SetSha224 sets Sha224 field to given value.

### HasSha224

`func (o *RpmPackageResponse) HasSha224() bool`

HasSha224 returns a boolean if a field has been set.

### GetSha256

`func (o *RpmPackageResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *RpmPackageResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *RpmPackageResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *RpmPackageResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSha384

`func (o *RpmPackageResponse) GetSha384() string`

GetSha384 returns the Sha384 field if non-nil, zero value otherwise.

### GetSha384Ok

`func (o *RpmPackageResponse) GetSha384Ok() (*string, bool)`

GetSha384Ok returns a tuple with the Sha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha384

`func (o *RpmPackageResponse) SetSha384(v string)`

SetSha384 sets Sha384 field to given value.

### HasSha384

`func (o *RpmPackageResponse) HasSha384() bool`

HasSha384 returns a boolean if a field has been set.

### GetSha512

`func (o *RpmPackageResponse) GetSha512() string`

GetSha512 returns the Sha512 field if non-nil, zero value otherwise.

### GetSha512Ok

`func (o *RpmPackageResponse) GetSha512Ok() (*string, bool)`

GetSha512Ok returns a tuple with the Sha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512

`func (o *RpmPackageResponse) SetSha512(v string)`

SetSha512 sets Sha512 field to given value.

### HasSha512

`func (o *RpmPackageResponse) HasSha512() bool`

HasSha512 returns a boolean if a field has been set.

### GetArtifact

`func (o *RpmPackageResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *RpmPackageResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *RpmPackageResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *RpmPackageResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetName

`func (o *RpmPackageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmPackageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmPackageResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RpmPackageResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpoch

`func (o *RpmPackageResponse) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *RpmPackageResponse) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *RpmPackageResponse) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *RpmPackageResponse) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetVersion

`func (o *RpmPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RpmPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RpmPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RpmPackageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *RpmPackageResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *RpmPackageResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *RpmPackageResponse) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *RpmPackageResponse) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetArch

`func (o *RpmPackageResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *RpmPackageResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *RpmPackageResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *RpmPackageResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetPkgId

`func (o *RpmPackageResponse) GetPkgId() string`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *RpmPackageResponse) GetPkgIdOk() (*string, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *RpmPackageResponse) SetPkgId(v string)`

SetPkgId sets PkgId field to given value.

### HasPkgId

`func (o *RpmPackageResponse) HasPkgId() bool`

HasPkgId returns a boolean if a field has been set.

### GetChecksumType

`func (o *RpmPackageResponse) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmPackageResponse) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmPackageResponse) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *RpmPackageResponse) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetSummary

`func (o *RpmPackageResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RpmPackageResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RpmPackageResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RpmPackageResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *RpmPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpmPackageResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *RpmPackageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RpmPackageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RpmPackageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RpmPackageResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetChangelogs

`func (o *RpmPackageResponse) GetChangelogs() interface{}`

GetChangelogs returns the Changelogs field if non-nil, zero value otherwise.

### GetChangelogsOk

`func (o *RpmPackageResponse) GetChangelogsOk() (*interface{}, bool)`

GetChangelogsOk returns a tuple with the Changelogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogs

`func (o *RpmPackageResponse) SetChangelogs(v interface{})`

SetChangelogs sets Changelogs field to given value.

### HasChangelogs

`func (o *RpmPackageResponse) HasChangelogs() bool`

HasChangelogs returns a boolean if a field has been set.

### SetChangelogsNil

`func (o *RpmPackageResponse) SetChangelogsNil(b bool)`

 SetChangelogsNil sets the value for Changelogs to be an explicit nil

### UnsetChangelogs
`func (o *RpmPackageResponse) UnsetChangelogs()`

UnsetChangelogs ensures that no value is present for Changelogs, not even an explicit nil
### GetFiles

`func (o *RpmPackageResponse) GetFiles() interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RpmPackageResponse) GetFilesOk() (*interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RpmPackageResponse) SetFiles(v interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RpmPackageResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *RpmPackageResponse) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *RpmPackageResponse) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetRequires

`func (o *RpmPackageResponse) GetRequires() interface{}`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *RpmPackageResponse) GetRequiresOk() (*interface{}, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *RpmPackageResponse) SetRequires(v interface{})`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *RpmPackageResponse) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### SetRequiresNil

`func (o *RpmPackageResponse) SetRequiresNil(b bool)`

 SetRequiresNil sets the value for Requires to be an explicit nil

### UnsetRequires
`func (o *RpmPackageResponse) UnsetRequires()`

UnsetRequires ensures that no value is present for Requires, not even an explicit nil
### GetProvides

`func (o *RpmPackageResponse) GetProvides() interface{}`

GetProvides returns the Provides field if non-nil, zero value otherwise.

### GetProvidesOk

`func (o *RpmPackageResponse) GetProvidesOk() (*interface{}, bool)`

GetProvidesOk returns a tuple with the Provides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvides

`func (o *RpmPackageResponse) SetProvides(v interface{})`

SetProvides sets Provides field to given value.

### HasProvides

`func (o *RpmPackageResponse) HasProvides() bool`

HasProvides returns a boolean if a field has been set.

### SetProvidesNil

`func (o *RpmPackageResponse) SetProvidesNil(b bool)`

 SetProvidesNil sets the value for Provides to be an explicit nil

### UnsetProvides
`func (o *RpmPackageResponse) UnsetProvides()`

UnsetProvides ensures that no value is present for Provides, not even an explicit nil
### GetConflicts

`func (o *RpmPackageResponse) GetConflicts() interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *RpmPackageResponse) GetConflictsOk() (*interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *RpmPackageResponse) SetConflicts(v interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *RpmPackageResponse) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### SetConflictsNil

`func (o *RpmPackageResponse) SetConflictsNil(b bool)`

 SetConflictsNil sets the value for Conflicts to be an explicit nil

### UnsetConflicts
`func (o *RpmPackageResponse) UnsetConflicts()`

UnsetConflicts ensures that no value is present for Conflicts, not even an explicit nil
### GetObsoletes

`func (o *RpmPackageResponse) GetObsoletes() interface{}`

GetObsoletes returns the Obsoletes field if non-nil, zero value otherwise.

### GetObsoletesOk

`func (o *RpmPackageResponse) GetObsoletesOk() (*interface{}, bool)`

GetObsoletesOk returns a tuple with the Obsoletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletes

`func (o *RpmPackageResponse) SetObsoletes(v interface{})`

SetObsoletes sets Obsoletes field to given value.

### HasObsoletes

`func (o *RpmPackageResponse) HasObsoletes() bool`

HasObsoletes returns a boolean if a field has been set.

### SetObsoletesNil

`func (o *RpmPackageResponse) SetObsoletesNil(b bool)`

 SetObsoletesNil sets the value for Obsoletes to be an explicit nil

### UnsetObsoletes
`func (o *RpmPackageResponse) UnsetObsoletes()`

UnsetObsoletes ensures that no value is present for Obsoletes, not even an explicit nil
### GetSuggests

`func (o *RpmPackageResponse) GetSuggests() interface{}`

GetSuggests returns the Suggests field if non-nil, zero value otherwise.

### GetSuggestsOk

`func (o *RpmPackageResponse) GetSuggestsOk() (*interface{}, bool)`

GetSuggestsOk returns a tuple with the Suggests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggests

`func (o *RpmPackageResponse) SetSuggests(v interface{})`

SetSuggests sets Suggests field to given value.

### HasSuggests

`func (o *RpmPackageResponse) HasSuggests() bool`

HasSuggests returns a boolean if a field has been set.

### SetSuggestsNil

`func (o *RpmPackageResponse) SetSuggestsNil(b bool)`

 SetSuggestsNil sets the value for Suggests to be an explicit nil

### UnsetSuggests
`func (o *RpmPackageResponse) UnsetSuggests()`

UnsetSuggests ensures that no value is present for Suggests, not even an explicit nil
### GetEnhances

`func (o *RpmPackageResponse) GetEnhances() interface{}`

GetEnhances returns the Enhances field if non-nil, zero value otherwise.

### GetEnhancesOk

`func (o *RpmPackageResponse) GetEnhancesOk() (*interface{}, bool)`

GetEnhancesOk returns a tuple with the Enhances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhances

`func (o *RpmPackageResponse) SetEnhances(v interface{})`

SetEnhances sets Enhances field to given value.

### HasEnhances

`func (o *RpmPackageResponse) HasEnhances() bool`

HasEnhances returns a boolean if a field has been set.

### SetEnhancesNil

`func (o *RpmPackageResponse) SetEnhancesNil(b bool)`

 SetEnhancesNil sets the value for Enhances to be an explicit nil

### UnsetEnhances
`func (o *RpmPackageResponse) UnsetEnhances()`

UnsetEnhances ensures that no value is present for Enhances, not even an explicit nil
### GetRecommends

`func (o *RpmPackageResponse) GetRecommends() interface{}`

GetRecommends returns the Recommends field if non-nil, zero value otherwise.

### GetRecommendsOk

`func (o *RpmPackageResponse) GetRecommendsOk() (*interface{}, bool)`

GetRecommendsOk returns a tuple with the Recommends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommends

`func (o *RpmPackageResponse) SetRecommends(v interface{})`

SetRecommends sets Recommends field to given value.

### HasRecommends

`func (o *RpmPackageResponse) HasRecommends() bool`

HasRecommends returns a boolean if a field has been set.

### SetRecommendsNil

`func (o *RpmPackageResponse) SetRecommendsNil(b bool)`

 SetRecommendsNil sets the value for Recommends to be an explicit nil

### UnsetRecommends
`func (o *RpmPackageResponse) UnsetRecommends()`

UnsetRecommends ensures that no value is present for Recommends, not even an explicit nil
### GetSupplements

`func (o *RpmPackageResponse) GetSupplements() interface{}`

GetSupplements returns the Supplements field if non-nil, zero value otherwise.

### GetSupplementsOk

`func (o *RpmPackageResponse) GetSupplementsOk() (*interface{}, bool)`

GetSupplementsOk returns a tuple with the Supplements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplements

`func (o *RpmPackageResponse) SetSupplements(v interface{})`

SetSupplements sets Supplements field to given value.

### HasSupplements

`func (o *RpmPackageResponse) HasSupplements() bool`

HasSupplements returns a boolean if a field has been set.

### SetSupplementsNil

`func (o *RpmPackageResponse) SetSupplementsNil(b bool)`

 SetSupplementsNil sets the value for Supplements to be an explicit nil

### UnsetSupplements
`func (o *RpmPackageResponse) UnsetSupplements()`

UnsetSupplements ensures that no value is present for Supplements, not even an explicit nil
### GetLocationBase

`func (o *RpmPackageResponse) GetLocationBase() string`

GetLocationBase returns the LocationBase field if non-nil, zero value otherwise.

### GetLocationBaseOk

`func (o *RpmPackageResponse) GetLocationBaseOk() (*string, bool)`

GetLocationBaseOk returns a tuple with the LocationBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationBase

`func (o *RpmPackageResponse) SetLocationBase(v string)`

SetLocationBase sets LocationBase field to given value.

### HasLocationBase

`func (o *RpmPackageResponse) HasLocationBase() bool`

HasLocationBase returns a boolean if a field has been set.

### GetLocationHref

`func (o *RpmPackageResponse) GetLocationHref() string`

GetLocationHref returns the LocationHref field if non-nil, zero value otherwise.

### GetLocationHrefOk

`func (o *RpmPackageResponse) GetLocationHrefOk() (*string, bool)`

GetLocationHrefOk returns a tuple with the LocationHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationHref

`func (o *RpmPackageResponse) SetLocationHref(v string)`

SetLocationHref sets LocationHref field to given value.

### HasLocationHref

`func (o *RpmPackageResponse) HasLocationHref() bool`

HasLocationHref returns a boolean if a field has been set.

### GetRpmBuildhost

`func (o *RpmPackageResponse) GetRpmBuildhost() string`

GetRpmBuildhost returns the RpmBuildhost field if non-nil, zero value otherwise.

### GetRpmBuildhostOk

`func (o *RpmPackageResponse) GetRpmBuildhostOk() (*string, bool)`

GetRpmBuildhostOk returns a tuple with the RpmBuildhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmBuildhost

`func (o *RpmPackageResponse) SetRpmBuildhost(v string)`

SetRpmBuildhost sets RpmBuildhost field to given value.

### HasRpmBuildhost

`func (o *RpmPackageResponse) HasRpmBuildhost() bool`

HasRpmBuildhost returns a boolean if a field has been set.

### GetRpmGroup

`func (o *RpmPackageResponse) GetRpmGroup() string`

GetRpmGroup returns the RpmGroup field if non-nil, zero value otherwise.

### GetRpmGroupOk

`func (o *RpmPackageResponse) GetRpmGroupOk() (*string, bool)`

GetRpmGroupOk returns a tuple with the RpmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmGroup

`func (o *RpmPackageResponse) SetRpmGroup(v string)`

SetRpmGroup sets RpmGroup field to given value.

### HasRpmGroup

`func (o *RpmPackageResponse) HasRpmGroup() bool`

HasRpmGroup returns a boolean if a field has been set.

### GetRpmLicense

`func (o *RpmPackageResponse) GetRpmLicense() string`

GetRpmLicense returns the RpmLicense field if non-nil, zero value otherwise.

### GetRpmLicenseOk

`func (o *RpmPackageResponse) GetRpmLicenseOk() (*string, bool)`

GetRpmLicenseOk returns a tuple with the RpmLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLicense

`func (o *RpmPackageResponse) SetRpmLicense(v string)`

SetRpmLicense sets RpmLicense field to given value.

### HasRpmLicense

`func (o *RpmPackageResponse) HasRpmLicense() bool`

HasRpmLicense returns a boolean if a field has been set.

### GetRpmPackager

`func (o *RpmPackageResponse) GetRpmPackager() string`

GetRpmPackager returns the RpmPackager field if non-nil, zero value otherwise.

### GetRpmPackagerOk

`func (o *RpmPackageResponse) GetRpmPackagerOk() (*string, bool)`

GetRpmPackagerOk returns a tuple with the RpmPackager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmPackager

`func (o *RpmPackageResponse) SetRpmPackager(v string)`

SetRpmPackager sets RpmPackager field to given value.

### HasRpmPackager

`func (o *RpmPackageResponse) HasRpmPackager() bool`

HasRpmPackager returns a boolean if a field has been set.

### GetRpmSourcerpm

`func (o *RpmPackageResponse) GetRpmSourcerpm() string`

GetRpmSourcerpm returns the RpmSourcerpm field if non-nil, zero value otherwise.

### GetRpmSourcerpmOk

`func (o *RpmPackageResponse) GetRpmSourcerpmOk() (*string, bool)`

GetRpmSourcerpmOk returns a tuple with the RpmSourcerpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmSourcerpm

`func (o *RpmPackageResponse) SetRpmSourcerpm(v string)`

SetRpmSourcerpm sets RpmSourcerpm field to given value.

### HasRpmSourcerpm

`func (o *RpmPackageResponse) HasRpmSourcerpm() bool`

HasRpmSourcerpm returns a boolean if a field has been set.

### GetRpmVendor

`func (o *RpmPackageResponse) GetRpmVendor() string`

GetRpmVendor returns the RpmVendor field if non-nil, zero value otherwise.

### GetRpmVendorOk

`func (o *RpmPackageResponse) GetRpmVendorOk() (*string, bool)`

GetRpmVendorOk returns a tuple with the RpmVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmVendor

`func (o *RpmPackageResponse) SetRpmVendor(v string)`

SetRpmVendor sets RpmVendor field to given value.

### HasRpmVendor

`func (o *RpmPackageResponse) HasRpmVendor() bool`

HasRpmVendor returns a boolean if a field has been set.

### GetRpmHeaderStart

`func (o *RpmPackageResponse) GetRpmHeaderStart() int64`

GetRpmHeaderStart returns the RpmHeaderStart field if non-nil, zero value otherwise.

### GetRpmHeaderStartOk

`func (o *RpmPackageResponse) GetRpmHeaderStartOk() (*int64, bool)`

GetRpmHeaderStartOk returns a tuple with the RpmHeaderStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmHeaderStart

`func (o *RpmPackageResponse) SetRpmHeaderStart(v int64)`

SetRpmHeaderStart sets RpmHeaderStart field to given value.

### HasRpmHeaderStart

`func (o *RpmPackageResponse) HasRpmHeaderStart() bool`

HasRpmHeaderStart returns a boolean if a field has been set.

### GetRpmHeaderEnd

`func (o *RpmPackageResponse) GetRpmHeaderEnd() int64`

GetRpmHeaderEnd returns the RpmHeaderEnd field if non-nil, zero value otherwise.

### GetRpmHeaderEndOk

`func (o *RpmPackageResponse) GetRpmHeaderEndOk() (*int64, bool)`

GetRpmHeaderEndOk returns a tuple with the RpmHeaderEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmHeaderEnd

`func (o *RpmPackageResponse) SetRpmHeaderEnd(v int64)`

SetRpmHeaderEnd sets RpmHeaderEnd field to given value.

### HasRpmHeaderEnd

`func (o *RpmPackageResponse) HasRpmHeaderEnd() bool`

HasRpmHeaderEnd returns a boolean if a field has been set.

### GetIsModular

`func (o *RpmPackageResponse) GetIsModular() bool`

GetIsModular returns the IsModular field if non-nil, zero value otherwise.

### GetIsModularOk

`func (o *RpmPackageResponse) GetIsModularOk() (*bool, bool)`

GetIsModularOk returns a tuple with the IsModular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModular

`func (o *RpmPackageResponse) SetIsModular(v bool)`

SetIsModular sets IsModular field to given value.

### HasIsModular

`func (o *RpmPackageResponse) HasIsModular() bool`

HasIsModular returns a boolean if a field has been set.

### GetSizeArchive

`func (o *RpmPackageResponse) GetSizeArchive() int64`

GetSizeArchive returns the SizeArchive field if non-nil, zero value otherwise.

### GetSizeArchiveOk

`func (o *RpmPackageResponse) GetSizeArchiveOk() (*int64, bool)`

GetSizeArchiveOk returns a tuple with the SizeArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeArchive

`func (o *RpmPackageResponse) SetSizeArchive(v int64)`

SetSizeArchive sets SizeArchive field to given value.

### HasSizeArchive

`func (o *RpmPackageResponse) HasSizeArchive() bool`

HasSizeArchive returns a boolean if a field has been set.

### GetSizeInstalled

`func (o *RpmPackageResponse) GetSizeInstalled() int64`

GetSizeInstalled returns the SizeInstalled field if non-nil, zero value otherwise.

### GetSizeInstalledOk

`func (o *RpmPackageResponse) GetSizeInstalledOk() (*int64, bool)`

GetSizeInstalledOk returns a tuple with the SizeInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInstalled

`func (o *RpmPackageResponse) SetSizeInstalled(v int64)`

SetSizeInstalled sets SizeInstalled field to given value.

### HasSizeInstalled

`func (o *RpmPackageResponse) HasSizeInstalled() bool`

HasSizeInstalled returns a boolean if a field has been set.

### GetSizePackage

`func (o *RpmPackageResponse) GetSizePackage() int64`

GetSizePackage returns the SizePackage field if non-nil, zero value otherwise.

### GetSizePackageOk

`func (o *RpmPackageResponse) GetSizePackageOk() (*int64, bool)`

GetSizePackageOk returns a tuple with the SizePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePackage

`func (o *RpmPackageResponse) SetSizePackage(v int64)`

SetSizePackage sets SizePackage field to given value.

### HasSizePackage

`func (o *RpmPackageResponse) HasSizePackage() bool`

HasSizePackage returns a boolean if a field has been set.

### GetTimeBuild

`func (o *RpmPackageResponse) GetTimeBuild() int64`

GetTimeBuild returns the TimeBuild field if non-nil, zero value otherwise.

### GetTimeBuildOk

`func (o *RpmPackageResponse) GetTimeBuildOk() (*int64, bool)`

GetTimeBuildOk returns a tuple with the TimeBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBuild

`func (o *RpmPackageResponse) SetTimeBuild(v int64)`

SetTimeBuild sets TimeBuild field to given value.

### HasTimeBuild

`func (o *RpmPackageResponse) HasTimeBuild() bool`

HasTimeBuild returns a boolean if a field has been set.

### GetTimeFile

`func (o *RpmPackageResponse) GetTimeFile() int64`

GetTimeFile returns the TimeFile field if non-nil, zero value otherwise.

### GetTimeFileOk

`func (o *RpmPackageResponse) GetTimeFileOk() (*int64, bool)`

GetTimeFileOk returns a tuple with the TimeFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFile

`func (o *RpmPackageResponse) SetTimeFile(v int64)`

SetTimeFile sets TimeFile field to given value.

### HasTimeFile

`func (o *RpmPackageResponse) HasTimeFile() bool`

HasTimeFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


