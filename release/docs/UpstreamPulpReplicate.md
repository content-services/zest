# UpstreamPulpReplicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QSelect** | Pointer to **NullableString** | Filter distributions on the upstream Pulp using complex filtering. When specified, overrides the stored q_select for this replication run only. E.g. pulp_label_select&#x3D;\&quot;foo\&quot; OR pulp_label_select&#x3D;\&quot;key&#x3D;val\&quot; | [optional] 

## Methods

### NewUpstreamPulpReplicate

`func NewUpstreamPulpReplicate() *UpstreamPulpReplicate`

NewUpstreamPulpReplicate instantiates a new UpstreamPulpReplicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamPulpReplicateWithDefaults

`func NewUpstreamPulpReplicateWithDefaults() *UpstreamPulpReplicate`

NewUpstreamPulpReplicateWithDefaults instantiates a new UpstreamPulpReplicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQSelect

`func (o *UpstreamPulpReplicate) GetQSelect() string`

GetQSelect returns the QSelect field if non-nil, zero value otherwise.

### GetQSelectOk

`func (o *UpstreamPulpReplicate) GetQSelectOk() (*string, bool)`

GetQSelectOk returns a tuple with the QSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQSelect

`func (o *UpstreamPulpReplicate) SetQSelect(v string)`

SetQSelect sets QSelect field to given value.

### HasQSelect

`func (o *UpstreamPulpReplicate) HasQSelect() bool`

HasQSelect returns a boolean if a field has been set.

### SetQSelectNil

`func (o *UpstreamPulpReplicate) SetQSelectNil(b bool)`

 SetQSelectNil sets the value for QSelect to be an explicit nil

### UnsetQSelect
`func (o *UpstreamPulpReplicate) UnsetQSelect()`

UnsetQSelect ensures that no value is present for QSelect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


