# PublicDebugAuthenticationHeadersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XRhIdentityPresent** | **bool** |  | 
**XPulpVpnVerified** | **bool** | Whether the trusted X-Pulp-VPN-Verified assertion reached Pulp. | 
**XPulpVpnAccessPresent** | **bool** |  | 

## Methods

### NewPublicDebugAuthenticationHeadersResponse

`func NewPublicDebugAuthenticationHeadersResponse(xRhIdentityPresent bool, xPulpVpnVerified bool, xPulpVpnAccessPresent bool, ) *PublicDebugAuthenticationHeadersResponse`

NewPublicDebugAuthenticationHeadersResponse instantiates a new PublicDebugAuthenticationHeadersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDebugAuthenticationHeadersResponseWithDefaults

`func NewPublicDebugAuthenticationHeadersResponseWithDefaults() *PublicDebugAuthenticationHeadersResponse`

NewPublicDebugAuthenticationHeadersResponseWithDefaults instantiates a new PublicDebugAuthenticationHeadersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXRhIdentityPresent

`func (o *PublicDebugAuthenticationHeadersResponse) GetXRhIdentityPresent() bool`

GetXRhIdentityPresent returns the XRhIdentityPresent field if non-nil, zero value otherwise.

### GetXRhIdentityPresentOk

`func (o *PublicDebugAuthenticationHeadersResponse) GetXRhIdentityPresentOk() (*bool, bool)`

GetXRhIdentityPresentOk returns a tuple with the XRhIdentityPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRhIdentityPresent

`func (o *PublicDebugAuthenticationHeadersResponse) SetXRhIdentityPresent(v bool)`

SetXRhIdentityPresent sets XRhIdentityPresent field to given value.


### GetXPulpVpnVerified

`func (o *PublicDebugAuthenticationHeadersResponse) GetXPulpVpnVerified() bool`

GetXPulpVpnVerified returns the XPulpVpnVerified field if non-nil, zero value otherwise.

### GetXPulpVpnVerifiedOk

`func (o *PublicDebugAuthenticationHeadersResponse) GetXPulpVpnVerifiedOk() (*bool, bool)`

GetXPulpVpnVerifiedOk returns a tuple with the XPulpVpnVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXPulpVpnVerified

`func (o *PublicDebugAuthenticationHeadersResponse) SetXPulpVpnVerified(v bool)`

SetXPulpVpnVerified sets XPulpVpnVerified field to given value.


### GetXPulpVpnAccessPresent

`func (o *PublicDebugAuthenticationHeadersResponse) GetXPulpVpnAccessPresent() bool`

GetXPulpVpnAccessPresent returns the XPulpVpnAccessPresent field if non-nil, zero value otherwise.

### GetXPulpVpnAccessPresentOk

`func (o *PublicDebugAuthenticationHeadersResponse) GetXPulpVpnAccessPresentOk() (*bool, bool)`

GetXPulpVpnAccessPresentOk returns a tuple with the XPulpVpnAccessPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXPulpVpnAccessPresent

`func (o *PublicDebugAuthenticationHeadersResponse) SetXPulpVpnAccessPresent(v bool)`

SetXPulpVpnAccessPresent sets XPulpVpnAccessPresent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


