/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
	"reflect"
)


// ContentPackagesAPIService ContentPackagesAPI service
type ContentPackagesAPIService service

type ContentPackagesAPIContentRpmPackagesCreateRequest struct {
	ctx context.Context
	ApiService *ContentPackagesAPIService
	pulpDomain string
	repository *string
	artifact *string
	relativePath *string
	file *os.File
	upload *string
}

// A URI of a repository the new content unit should be associated with.
func (r ContentPackagesAPIContentRpmPackagesCreateRequest) Repository(repository string) ContentPackagesAPIContentRpmPackagesCreateRequest {
	r.repository = &repository
	return r
}

// Artifact file representing the physical content
func (r ContentPackagesAPIContentRpmPackagesCreateRequest) Artifact(artifact string) ContentPackagesAPIContentRpmPackagesCreateRequest {
	r.artifact = &artifact
	return r
}

// Path where the artifact is located relative to distributions base_path
func (r ContentPackagesAPIContentRpmPackagesCreateRequest) RelativePath(relativePath string) ContentPackagesAPIContentRpmPackagesCreateRequest {
	r.relativePath = &relativePath
	return r
}

// An uploaded file that may be turned into the content unit.
func (r ContentPackagesAPIContentRpmPackagesCreateRequest) File(file *os.File) ContentPackagesAPIContentRpmPackagesCreateRequest {
	r.file = file
	return r
}

// An uncommitted upload that may be turned into the content unit.
func (r ContentPackagesAPIContentRpmPackagesCreateRequest) Upload(upload string) ContentPackagesAPIContentRpmPackagesCreateRequest {
	r.upload = &upload
	return r
}

func (r ContentPackagesAPIContentRpmPackagesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagesCreateExecute(r)
}

/*
ContentRpmPackagesCreate Create a package

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentPackagesAPIContentRpmPackagesCreateRequest
*/
func (a *ContentPackagesAPIService) ContentRpmPackagesCreate(ctx context.Context, pulpDomain string) ContentPackagesAPIContentRpmPackagesCreateRequest {
	return ContentPackagesAPIContentRpmPackagesCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentPackagesAPIService) ContentRpmPackagesCreateExecute(r ContentPackagesAPIContentRpmPackagesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagesAPIService.ContentRpmPackagesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/packages/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "")
	}
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "relative_path", r.relativePath, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentPackagesAPIContentRpmPackagesListRequest struct {
	ctx context.Context
	ApiService *ContentPackagesAPIService
	pulpDomain string
	arch *string
	archContains *string
	archIn *[]string
	archNe *string
	archStartswith *string
	checksumType *string
	checksumTypeIn *[]string
	checksumTypeNe *string
	epoch *string
	epochIn *[]string
	epochNe *string
	filename *string
	limit *int32
	name *string
	nameContains *string
	nameIn *[]string
	nameNe *string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pkgId *string
	pkgIdIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	release *string
	releaseContains *string
	releaseIn *[]string
	releaseNe *string
	releaseStartswith *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	version *string
	versionIn *[]string
	versionNe *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where arch matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) Arch(arch string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.arch = &arch
	return r
}

// Filter results where arch contains value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ArchContains(archContains string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.archContains = &archContains
	return r
}

// Filter results where arch is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) ArchIn(archIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.archIn = &archIn
	return r
}

// Filter results where arch not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ArchNe(archNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.archNe = &archNe
	return r
}

// Filter results where arch starts with value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ArchStartswith(archStartswith string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.archStartswith = &archStartswith
	return r
}

// Filter results where checksum_type matches value  * &#x60;unknown&#x60; - unknown * &#x60;md5&#x60; - md5 * &#x60;sha1&#x60; - sha1 * &#x60;sha1&#x60; - sha1 * &#x60;sha224&#x60; - sha224 * &#x60;sha256&#x60; - sha256 * &#x60;sha384&#x60; - sha384 * &#x60;sha512&#x60; - sha512
func (r ContentPackagesAPIContentRpmPackagesListRequest) ChecksumType(checksumType string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.checksumType = &checksumType
	return r
}

// Filter results where checksum_type is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) ChecksumTypeIn(checksumTypeIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.checksumTypeIn = &checksumTypeIn
	return r
}

// Filter results where checksum_type not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ChecksumTypeNe(checksumTypeNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.checksumTypeNe = &checksumTypeNe
	return r
}

// Filter results where epoch matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) Epoch(epoch string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.epoch = &epoch
	return r
}

// Filter results where epoch is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) EpochIn(epochIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.epochIn = &epochIn
	return r
}

// Filter results where epoch not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) EpochNe(epochNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.epochNe = &epochNe
	return r
}

func (r ContentPackagesAPIContentRpmPackagesListRequest) Filename(filename string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.filename = &filename
	return r
}

// Number of results to return per page.
func (r ContentPackagesAPIContentRpmPackagesListRequest) Limit(limit int32) ContentPackagesAPIContentRpmPackagesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) Name(name string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentPackagesAPIContentRpmPackagesListRequest) NameContains(nameContains string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) NameIn(nameIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) NameNe(nameNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.nameNe = &nameNe
	return r
}

// Filter results where name starts with value
func (r ContentPackagesAPIContentRpmPackagesListRequest) NameStartswith(nameStartswith string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentPackagesAPIContentRpmPackagesListRequest) Offset(offset int32) ContentPackagesAPIContentRpmPackagesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;epoch&#x60; - Epoch * &#x60;-epoch&#x60; - Epoch (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;release&#x60; - Release * &#x60;-release&#x60; - Release (descending) * &#x60;arch&#x60; - Arch * &#x60;-arch&#x60; - Arch (descending) * &#x60;evr&#x60; - Evr * &#x60;-evr&#x60; - Evr (descending) * &#x60;pkgId&#x60; - Pkgid * &#x60;-pkgId&#x60; - Pkgid (descending) * &#x60;checksum_type&#x60; - Checksum type * &#x60;-checksum_type&#x60; - Checksum type (descending) * &#x60;summary&#x60; - Summary * &#x60;-summary&#x60; - Summary (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;changelogs&#x60; - Changelogs * &#x60;-changelogs&#x60; - Changelogs (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;requires&#x60; - Requires * &#x60;-requires&#x60; - Requires (descending) * &#x60;provides&#x60; - Provides * &#x60;-provides&#x60; - Provides (descending) * &#x60;conflicts&#x60; - Conflicts * &#x60;-conflicts&#x60; - Conflicts (descending) * &#x60;obsoletes&#x60; - Obsoletes * &#x60;-obsoletes&#x60; - Obsoletes (descending) * &#x60;suggests&#x60; - Suggests * &#x60;-suggests&#x60; - Suggests (descending) * &#x60;enhances&#x60; - Enhances * &#x60;-enhances&#x60; - Enhances (descending) * &#x60;recommends&#x60; - Recommends * &#x60;-recommends&#x60; - Recommends (descending) * &#x60;supplements&#x60; - Supplements * &#x60;-supplements&#x60; - Supplements (descending) * &#x60;location_base&#x60; - Location base * &#x60;-location_base&#x60; - Location base (descending) * &#x60;location_href&#x60; - Location href * &#x60;-location_href&#x60; - Location href (descending) * &#x60;rpm_buildhost&#x60; - Rpm buildhost * &#x60;-rpm_buildhost&#x60; - Rpm buildhost (descending) * &#x60;rpm_group&#x60; - Rpm group * &#x60;-rpm_group&#x60; - Rpm group (descending) * &#x60;rpm_license&#x60; - Rpm license * &#x60;-rpm_license&#x60; - Rpm license (descending) * &#x60;rpm_packager&#x60; - Rpm packager * &#x60;-rpm_packager&#x60; - Rpm packager (descending) * &#x60;rpm_sourcerpm&#x60; - Rpm sourcerpm * &#x60;-rpm_sourcerpm&#x60; - Rpm sourcerpm (descending) * &#x60;rpm_vendor&#x60; - Rpm vendor * &#x60;-rpm_vendor&#x60; - Rpm vendor (descending) * &#x60;rpm_header_start&#x60; - Rpm header start * &#x60;-rpm_header_start&#x60; - Rpm header start (descending) * &#x60;rpm_header_end&#x60; - Rpm header end * &#x60;-rpm_header_end&#x60; - Rpm header end (descending) * &#x60;size_archive&#x60; - Size archive * &#x60;-size_archive&#x60; - Size archive (descending) * &#x60;size_installed&#x60; - Size installed * &#x60;-size_installed&#x60; - Size installed (descending) * &#x60;size_package&#x60; - Size package * &#x60;-size_package&#x60; - Size package (descending) * &#x60;time_build&#x60; - Time build * &#x60;-time_build&#x60; - Time build (descending) * &#x60;time_file&#x60; - Time file * &#x60;-time_file&#x60; - Time file (descending) * &#x60;is_modular&#x60; - Is modular * &#x60;-is_modular&#x60; - Is modular (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackagesAPIContentRpmPackagesListRequest) Ordering(ordering []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pkgId matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) PkgId(pkgId string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.pkgId = &pkgId
	return r
}

// Filter results where pkgId is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) PkgIdIn(pkgIdIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.pkgIdIn = &pkgIdIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagesAPIContentRpmPackagesListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagesAPIContentRpmPackagesListRequest) PulpIdIn(pulpIdIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentPackagesAPIContentRpmPackagesListRequest) Q(q string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.q = &q
	return r
}

// Filter results where release matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) Release(release string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.release = &release
	return r
}

// Filter results where release contains value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ReleaseContains(releaseContains string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.releaseContains = &releaseContains
	return r
}

// Filter results where release is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) ReleaseIn(releaseIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.releaseIn = &releaseIn
	return r
}

// Filter results where release not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ReleaseNe(releaseNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.releaseNe = &releaseNe
	return r
}

// Filter results where release starts with value
func (r ContentPackagesAPIContentRpmPackagesListRequest) ReleaseStartswith(releaseStartswith string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.releaseStartswith = &releaseStartswith
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagesAPIContentRpmPackagesListRequest) RepositoryVersion(repositoryVersion string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagesAPIContentRpmPackagesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagesAPIContentRpmPackagesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

func (r ContentPackagesAPIContentRpmPackagesListRequest) Sha256(sha256 string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where version matches value
func (r ContentPackagesAPIContentRpmPackagesListRequest) Version(version string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.version = &version
	return r
}

// Filter results where version is in a comma-separated list of values
func (r ContentPackagesAPIContentRpmPackagesListRequest) VersionIn(versionIn []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.versionIn = &versionIn
	return r
}

// Filter results where version not equal to value
func (r ContentPackagesAPIContentRpmPackagesListRequest) VersionNe(versionNe string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.versionNe = &versionNe
	return r
}

// A list of fields to include in the response.
func (r ContentPackagesAPIContentRpmPackagesListRequest) Fields(fields []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagesAPIContentRpmPackagesListRequest) ExcludeFields(excludeFields []string) ContentPackagesAPIContentRpmPackagesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagesAPIContentRpmPackagesListRequest) Execute() (*PaginatedrpmPackageResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagesListExecute(r)
}

/*
ContentRpmPackagesList List packages

A ViewSet for Package.

Define endpoint name which will appear in the API endpoint for this content type.
For example::
    http://pulp.example.com/pulp/api/v3/content/rpm/packages/

Also specify queryset and serializer for Package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentPackagesAPIContentRpmPackagesListRequest
*/
func (a *ContentPackagesAPIService) ContentRpmPackagesList(ctx context.Context, pulpDomain string) ContentPackagesAPIContentRpmPackagesListRequest {
	return ContentPackagesAPIContentRpmPackagesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageResponseList
func (a *ContentPackagesAPIService) ContentRpmPackagesListExecute(r ContentPackagesAPIContentRpmPackagesListRequest) (*PaginatedrpmPackageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagesAPIService.ContentRpmPackagesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/packages/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.arch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch", r.arch, "")
	}
	if r.archContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch__contains", r.archContains, "")
	}
	if r.archIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch__in", r.archIn, "csv")
	}
	if r.archNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch__ne", r.archNe, "")
	}
	if r.archStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch__startswith", r.archStartswith, "")
	}
	if r.checksumType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum_type", r.checksumType, "")
	}
	if r.checksumTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum_type__in", r.checksumTypeIn, "csv")
	}
	if r.checksumTypeNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum_type__ne", r.checksumTypeNe, "")
	}
	if r.epoch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "epoch", r.epoch, "")
	}
	if r.epochIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "epoch__in", r.epochIn, "csv")
	}
	if r.epochNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "epoch__ne", r.epochNe, "")
	}
	if r.filename != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filename", r.filename, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__ne", r.nameNe, "")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pkgId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pkgId", r.pkgId, "")
	}
	if r.pkgIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pkgId__in", r.pkgIdIn, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.release != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "release", r.release, "")
	}
	if r.releaseContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "release__contains", r.releaseContains, "")
	}
	if r.releaseIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "release__in", r.releaseIn, "csv")
	}
	if r.releaseNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "release__ne", r.releaseNe, "")
	}
	if r.releaseStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "release__startswith", r.releaseStartswith, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.versionIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version__in", r.versionIn, "csv")
	}
	if r.versionNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version__ne", r.versionNe, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentPackagesAPIContentRpmPackagesReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagesAPIService
	rpmPackageHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagesAPIContentRpmPackagesReadRequest) Fields(fields []string) ContentPackagesAPIContentRpmPackagesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagesAPIContentRpmPackagesReadRequest) ExcludeFields(excludeFields []string) ContentPackagesAPIContentRpmPackagesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagesAPIContentRpmPackagesReadRequest) Execute() (*RpmPackageResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagesReadExecute(r)
}

/*
ContentRpmPackagesRead Inspect a package

A ViewSet for Package.

Define endpoint name which will appear in the API endpoint for this content type.
For example::
    http://pulp.example.com/pulp/api/v3/content/rpm/packages/

Also specify queryset and serializer for Package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageHref
 @return ContentPackagesAPIContentRpmPackagesReadRequest
*/
func (a *ContentPackagesAPIService) ContentRpmPackagesRead(ctx context.Context, rpmPackageHref string) ContentPackagesAPIContentRpmPackagesReadRequest {
	return ContentPackagesAPIContentRpmPackagesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageHref: rpmPackageHref,
	}
}

// Execute executes the request
//  @return RpmPackageResponse
func (a *ContentPackagesAPIService) ContentRpmPackagesReadExecute(r ContentPackagesAPIContentRpmPackagesReadRequest) (*RpmPackageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagesAPIService.ContentRpmPackagesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageHref, "rpmPackageHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
