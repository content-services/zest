/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
	"time"
)

// checks if the AnsibleCollectionRemoteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleCollectionRemoteResponse{}

// AnsibleCollectionRemoteResponse A serializer for Collection Remotes.
type AnsibleCollectionRemoteResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A unique name for this remote.
	Name string `json:"name"`
	// The URL of an external content source.
	Url string `json:"url"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Timestamp of the most recent update of the remote.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt64 `json:"download_concurrency,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt64 `json:"max_retries,omitempty"`
	Policy *PolicyDb6Enum `json:"policy,omitempty"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt64 `json:"rate_limit,omitempty"`
	// List of hidden (write only) fields
	HiddenFields []RemoteResponseHiddenFieldsInner `json:"hidden_fields,omitempty"`
	// The string version of Collection requirements yaml.
	RequirementsFile NullableString `json:"requirements_file,omitempty"`
	// The URL to receive a session token from, e.g. used with Automation Hub.
	AuthUrl NullableString `json:"auth_url,omitempty"`
	// Sync dependencies for collections specified via requirements file
	SyncDependencies *bool `json:"sync_dependencies,omitempty"`
	// Sync only collections that have a signature
	SignedOnly *bool `json:"signed_only,omitempty"`
}

// NewAnsibleCollectionRemoteResponse instantiates a new AnsibleCollectionRemoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleCollectionRemoteResponse(name string, url string) *AnsibleCollectionRemoteResponse {
	this := AnsibleCollectionRemoteResponse{}
	this.Name = name
	this.Url = url
	var syncDependencies bool = true
	this.SyncDependencies = &syncDependencies
	var signedOnly bool = false
	this.SignedOnly = &signedOnly
	return &this
}

// NewAnsibleCollectionRemoteResponseWithDefaults instantiates a new AnsibleCollectionRemoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleCollectionRemoteResponseWithDefaults() *AnsibleCollectionRemoteResponse {
	this := AnsibleCollectionRemoteResponse{}
	var syncDependencies bool = true
	this.SyncDependencies = &syncDependencies
	var signedOnly bool = false
	this.SignedOnly = &signedOnly
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleCollectionRemoteResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleCollectionRemoteResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *AnsibleCollectionRemoteResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleCollectionRemoteResponse) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *AnsibleCollectionRemoteResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnsibleCollectionRemoteResponse) SetUrl(v string) {
	o.Url = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *AnsibleCollectionRemoteResponse) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *AnsibleCollectionRemoteResponse) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *AnsibleCollectionRemoteResponse) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *AnsibleCollectionRemoteResponse) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleCollectionRemoteResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *AnsibleCollectionRemoteResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetDownloadConcurrency() int64 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int64
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetDownloadConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt64 and assigns it to the DownloadConcurrency field.
func (o *AnsibleCollectionRemoteResponse) SetDownloadConcurrency(v int64) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetMaxRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt64 and assigns it to the MaxRetries field.
func (o *AnsibleCollectionRemoteResponse) SetMaxRetries(v int64) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetPolicy() PolicyDb6Enum {
	if o == nil || IsNil(o.Policy) {
		var ret PolicyDb6Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetPolicyOk() (*PolicyDb6Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PolicyDb6Enum and assigns it to the Policy field.
func (o *AnsibleCollectionRemoteResponse) SetPolicy(v PolicyDb6Enum) {
	o.Policy = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *AnsibleCollectionRemoteResponse) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *AnsibleCollectionRemoteResponse) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *AnsibleCollectionRemoteResponse) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *AnsibleCollectionRemoteResponse) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *AnsibleCollectionRemoteResponse) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetRateLimit() int64 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetRateLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt64 and assigns it to the RateLimit field.
func (o *AnsibleCollectionRemoteResponse) SetRateLimit(v int64) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetHiddenFields() []RemoteResponseHiddenFieldsInner {
	if o == nil || IsNil(o.HiddenFields) {
		var ret []RemoteResponseHiddenFieldsInner
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetHiddenFieldsOk() ([]RemoteResponseHiddenFieldsInner, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasHiddenFields() bool {
	if o != nil && !IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []RemoteResponseHiddenFieldsInner and assigns it to the HiddenFields field.
func (o *AnsibleCollectionRemoteResponse) SetHiddenFields(v []RemoteResponseHiddenFieldsInner) {
	o.HiddenFields = v
}

// GetRequirementsFile returns the RequirementsFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetRequirementsFile() string {
	if o == nil || IsNil(o.RequirementsFile.Get()) {
		var ret string
		return ret
	}
	return *o.RequirementsFile.Get()
}

// GetRequirementsFileOk returns a tuple with the RequirementsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetRequirementsFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequirementsFile.Get(), o.RequirementsFile.IsSet()
}

// HasRequirementsFile returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasRequirementsFile() bool {
	if o != nil && o.RequirementsFile.IsSet() {
		return true
	}

	return false
}

// SetRequirementsFile gets a reference to the given NullableString and assigns it to the RequirementsFile field.
func (o *AnsibleCollectionRemoteResponse) SetRequirementsFile(v string) {
	o.RequirementsFile.Set(&v)
}
// SetRequirementsFileNil sets the value for RequirementsFile to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetRequirementsFileNil() {
	o.RequirementsFile.Set(nil)
}

// UnsetRequirementsFile ensures that no value is present for RequirementsFile, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetRequirementsFile() {
	o.RequirementsFile.Unset()
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionRemoteResponse) GetAuthUrl() string {
	if o == nil || IsNil(o.AuthUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUrl.Get()
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionRemoteResponse) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUrl.Get(), o.AuthUrl.IsSet()
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasAuthUrl() bool {
	if o != nil && o.AuthUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given NullableString and assigns it to the AuthUrl field.
func (o *AnsibleCollectionRemoteResponse) SetAuthUrl(v string) {
	o.AuthUrl.Set(&v)
}
// SetAuthUrlNil sets the value for AuthUrl to be an explicit nil
func (o *AnsibleCollectionRemoteResponse) SetAuthUrlNil() {
	o.AuthUrl.Set(nil)
}

// UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil
func (o *AnsibleCollectionRemoteResponse) UnsetAuthUrl() {
	o.AuthUrl.Unset()
}

// GetSyncDependencies returns the SyncDependencies field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetSyncDependencies() bool {
	if o == nil || IsNil(o.SyncDependencies) {
		var ret bool
		return ret
	}
	return *o.SyncDependencies
}

// GetSyncDependenciesOk returns a tuple with the SyncDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetSyncDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncDependencies) {
		return nil, false
	}
	return o.SyncDependencies, true
}

// HasSyncDependencies returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasSyncDependencies() bool {
	if o != nil && !IsNil(o.SyncDependencies) {
		return true
	}

	return false
}

// SetSyncDependencies gets a reference to the given bool and assigns it to the SyncDependencies field.
func (o *AnsibleCollectionRemoteResponse) SetSyncDependencies(v bool) {
	o.SyncDependencies = &v
}

// GetSignedOnly returns the SignedOnly field value if set, zero value otherwise.
func (o *AnsibleCollectionRemoteResponse) GetSignedOnly() bool {
	if o == nil || IsNil(o.SignedOnly) {
		var ret bool
		return ret
	}
	return *o.SignedOnly
}

// GetSignedOnlyOk returns a tuple with the SignedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionRemoteResponse) GetSignedOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.SignedOnly) {
		return nil, false
	}
	return o.SignedOnly, true
}

// HasSignedOnly returns a boolean if a field has been set.
func (o *AnsibleCollectionRemoteResponse) HasSignedOnly() bool {
	if o != nil && !IsNil(o.SignedOnly) {
		return true
	}

	return false
}

// SetSignedOnly gets a reference to the given bool and assigns it to the SignedOnly field.
func (o *AnsibleCollectionRemoteResponse) SetSignedOnly(v bool) {
	o.SignedOnly = &v
}

func (o AnsibleCollectionRemoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleCollectionRemoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	// skip: pulp_last_updated is readOnly
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	// skip: hidden_fields is readOnly
	if o.RequirementsFile.IsSet() {
		toSerialize["requirements_file"] = o.RequirementsFile.Get()
	}
	if o.AuthUrl.IsSet() {
		toSerialize["auth_url"] = o.AuthUrl.Get()
	}
	if !IsNil(o.SyncDependencies) {
		toSerialize["sync_dependencies"] = o.SyncDependencies
	}
	if !IsNil(o.SignedOnly) {
		toSerialize["signed_only"] = o.SignedOnly
	}
	return toSerialize, nil
}

type NullableAnsibleCollectionRemoteResponse struct {
	value *AnsibleCollectionRemoteResponse
	isSet bool
}

func (v NullableAnsibleCollectionRemoteResponse) Get() *AnsibleCollectionRemoteResponse {
	return v.value
}

func (v *NullableAnsibleCollectionRemoteResponse) Set(val *AnsibleCollectionRemoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleCollectionRemoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleCollectionRemoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleCollectionRemoteResponse(val *AnsibleCollectionRemoteResponse) *NullableAnsibleCollectionRemoteResponse {
	return &NullableAnsibleCollectionRemoteResponse{value: val, isSet: true}
}

func (v NullableAnsibleCollectionRemoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleCollectionRemoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


