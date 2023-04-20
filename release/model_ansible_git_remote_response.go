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

// checks if the AnsibleGitRemoteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleGitRemoteResponse{}

// AnsibleGitRemoteResponse A serializer for Git Collection Remotes.
type AnsibleGitRemoteResponse struct {
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// The URL of an external content source.
	Url string `json:"url"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A unique name for this remote.
	Name string `json:"name"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt64 `json:"rate_limit,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt64 `json:"max_retries,omitempty"`
	PulpHref *string `json:"pulp_href,omitempty"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// List of hidden (write only) fields
	HiddenFields []RemoteResponseHiddenFieldsInner `json:"hidden_fields,omitempty"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt64 `json:"download_concurrency,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	// Timestamp of the most recent update of the remote.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// If True, only metadata about the content will be stored in Pulp. Clients will retrieve content from the remote URL.
	MetadataOnly *bool `json:"metadata_only,omitempty"`
	// A git ref. e.g.: branch, tag, or commit sha.
	GitRef *string `json:"git_ref,omitempty"`
}

// NewAnsibleGitRemoteResponse instantiates a new AnsibleGitRemoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleGitRemoteResponse(url string, name string) *AnsibleGitRemoteResponse {
	this := AnsibleGitRemoteResponse{}
	this.Url = url
	this.Name = name
	return &this
}

// NewAnsibleGitRemoteResponseWithDefaults instantiates a new AnsibleGitRemoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleGitRemoteResponseWithDefaults() *AnsibleGitRemoteResponse {
	this := AnsibleGitRemoteResponse{}
	return &this
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *AnsibleGitRemoteResponse) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *AnsibleGitRemoteResponse) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetUrl returns the Url field value
func (o *AnsibleGitRemoteResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnsibleGitRemoteResponse) SetUrl(v string) {
	o.Url = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *AnsibleGitRemoteResponse) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetName returns the Name field value
func (o *AnsibleGitRemoteResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleGitRemoteResponse) SetName(v string) {
	o.Name = v
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *AnsibleGitRemoteResponse) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *AnsibleGitRemoteResponse) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetRateLimit() int64 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetRateLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt64 and assigns it to the RateLimit field.
func (o *AnsibleGitRemoteResponse) SetRateLimit(v int64) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleGitRemoteResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetMaxRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt64 and assigns it to the MaxRetries field.
func (o *AnsibleGitRemoteResponse) SetMaxRetries(v int64) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleGitRemoteResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *AnsibleGitRemoteResponse) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *AnsibleGitRemoteResponse) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleGitRemoteResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *AnsibleGitRemoteResponse) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetHiddenFields() []RemoteResponseHiddenFieldsInner {
	if o == nil || IsNil(o.HiddenFields) {
		var ret []RemoteResponseHiddenFieldsInner
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetHiddenFieldsOk() ([]RemoteResponseHiddenFieldsInner, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasHiddenFields() bool {
	if o != nil && !IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []RemoteResponseHiddenFieldsInner and assigns it to the HiddenFields field.
func (o *AnsibleGitRemoteResponse) SetHiddenFields(v []RemoteResponseHiddenFieldsInner) {
	o.HiddenFields = v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetDownloadConcurrency() int64 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int64
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetDownloadConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt64 and assigns it to the DownloadConcurrency field.
func (o *AnsibleGitRemoteResponse) SetDownloadConcurrency(v int64) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemoteResponse) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemoteResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *AnsibleGitRemoteResponse) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *AnsibleGitRemoteResponse) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *AnsibleGitRemoteResponse) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *AnsibleGitRemoteResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetMetadataOnly returns the MetadataOnly field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetMetadataOnly() bool {
	if o == nil || IsNil(o.MetadataOnly) {
		var ret bool
		return ret
	}
	return *o.MetadataOnly
}

// GetMetadataOnlyOk returns a tuple with the MetadataOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetMetadataOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MetadataOnly) {
		return nil, false
	}
	return o.MetadataOnly, true
}

// HasMetadataOnly returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasMetadataOnly() bool {
	if o != nil && !IsNil(o.MetadataOnly) {
		return true
	}

	return false
}

// SetMetadataOnly gets a reference to the given bool and assigns it to the MetadataOnly field.
func (o *AnsibleGitRemoteResponse) SetMetadataOnly(v bool) {
	o.MetadataOnly = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *AnsibleGitRemoteResponse) GetGitRef() string {
	if o == nil || IsNil(o.GitRef) {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemoteResponse) GetGitRefOk() (*string, bool) {
	if o == nil || IsNil(o.GitRef) {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *AnsibleGitRemoteResponse) HasGitRef() bool {
	if o != nil && !IsNil(o.GitRef) {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *AnsibleGitRemoteResponse) SetGitRef(v string) {
	o.GitRef = &v
}

func (o AnsibleGitRemoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleGitRemoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	toSerialize["url"] = o.Url
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	// skip: pulp_created is readOnly
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	// skip: pulp_href is readOnly
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	// skip: hidden_fields is readOnly
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	// skip: pulp_last_updated is readOnly
	if !IsNil(o.MetadataOnly) {
		toSerialize["metadata_only"] = o.MetadataOnly
	}
	if !IsNil(o.GitRef) {
		toSerialize["git_ref"] = o.GitRef
	}
	return toSerialize, nil
}

type NullableAnsibleGitRemoteResponse struct {
	value *AnsibleGitRemoteResponse
	isSet bool
}

func (v NullableAnsibleGitRemoteResponse) Get() *AnsibleGitRemoteResponse {
	return v.value
}

func (v *NullableAnsibleGitRemoteResponse) Set(val *AnsibleGitRemoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleGitRemoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleGitRemoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleGitRemoteResponse(val *AnsibleGitRemoteResponse) *NullableAnsibleGitRemoteResponse {
	return &NullableAnsibleGitRemoteResponse{value: val, isSet: true}
}

func (v NullableAnsibleGitRemoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleGitRemoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


