# PythonPackageProvenanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**Package** | **string** | The package that the provenance is for. | 
**Provenance** | Pointer to **interface{}** |  | [optional] [readonly] 
**Sha256** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPythonPackageProvenanceResponse

`func NewPythonPackageProvenanceResponse(package_ string, ) *PythonPackageProvenanceResponse`

NewPythonPackageProvenanceResponse instantiates a new PythonPackageProvenanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPackageProvenanceResponseWithDefaults

`func NewPythonPackageProvenanceResponseWithDefaults() *PythonPackageProvenanceResponse`

NewPythonPackageProvenanceResponseWithDefaults instantiates a new PythonPackageProvenanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PythonPackageProvenanceResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PythonPackageProvenanceResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PythonPackageProvenanceResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PythonPackageProvenanceResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *PythonPackageProvenanceResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *PythonPackageProvenanceResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *PythonPackageProvenanceResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *PythonPackageProvenanceResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PythonPackageProvenanceResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PythonPackageProvenanceResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PythonPackageProvenanceResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PythonPackageProvenanceResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *PythonPackageProvenanceResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *PythonPackageProvenanceResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *PythonPackageProvenanceResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *PythonPackageProvenanceResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PythonPackageProvenanceResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPackageProvenanceResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPackageProvenanceResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPackageProvenanceResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *PythonPackageProvenanceResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *PythonPackageProvenanceResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *PythonPackageProvenanceResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *PythonPackageProvenanceResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetPackage

`func (o *PythonPackageProvenanceResponse) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PythonPackageProvenanceResponse) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PythonPackageProvenanceResponse) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetProvenance

`func (o *PythonPackageProvenanceResponse) GetProvenance() interface{}`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *PythonPackageProvenanceResponse) GetProvenanceOk() (*interface{}, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *PythonPackageProvenanceResponse) SetProvenance(v interface{})`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *PythonPackageProvenanceResponse) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### SetProvenanceNil

`func (o *PythonPackageProvenanceResponse) SetProvenanceNil(b bool)`

 SetProvenanceNil sets the value for Provenance to be an explicit nil

### UnsetProvenance
`func (o *PythonPackageProvenanceResponse) UnsetProvenance()`

UnsetProvenance ensures that no value is present for Provenance, not even an explicit nil
### GetSha256

`func (o *PythonPackageProvenanceResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PythonPackageProvenanceResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PythonPackageProvenanceResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *PythonPackageProvenanceResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


