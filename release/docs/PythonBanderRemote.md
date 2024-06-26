# PythonBanderRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | ***os.File** | A Bandersnatch config that may be used to construct a Python Remote. | 
**Name** | **string** | A unique name for this remote | 
**Policy** | Pointer to [**Policy762Enum**](Policy762Enum.md) | The policy to use when downloading content. The possible values include: &#39;immediate&#39;, &#39;on_demand&#39;, and &#39;streamed&#39;. &#39;on_demand&#39; is the default.* &#x60;immediate&#x60; - When syncing, download all metadata and content now.* &#x60;on_demand&#x60; - When syncing, download metadata, but do not download content now. Instead, download content as clients request it, and save it in Pulp to be served for future client requests.* &#x60;streamed&#x60; - When syncing, download metadata, but do not download content now. Instead,download content as clients request it, but never save it in Pulp. This causes future requests for that same content to have to be downloaded again. | [optional] [default to POLICY762ENUM_ON_DEMAND]

## Methods

### NewPythonBanderRemote

`func NewPythonBanderRemote(config *os.File, name string, ) *PythonBanderRemote`

NewPythonBanderRemote instantiates a new PythonBanderRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonBanderRemoteWithDefaults

`func NewPythonBanderRemoteWithDefaults() *PythonBanderRemote`

NewPythonBanderRemoteWithDefaults instantiates a new PythonBanderRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PythonBanderRemote) GetConfig() *os.File`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PythonBanderRemote) GetConfigOk() (**os.File, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PythonBanderRemote) SetConfig(v *os.File)`

SetConfig sets Config field to given value.


### GetName

`func (o *PythonBanderRemote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PythonBanderRemote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PythonBanderRemote) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *PythonBanderRemote) GetPolicy() Policy762Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PythonBanderRemote) GetPolicyOk() (*Policy762Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PythonBanderRemote) SetPolicy(v Policy762Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PythonBanderRemote) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


