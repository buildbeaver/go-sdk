# Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Build resource on the BuildBeaver server | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Etag** | **string** |  | 
**Name** | **string** | Name of the build, in URL format. Build names are typically numbers converted to strings. | 
**RepoId** | **string** | RepoID of the repo being built. | 
**CommitId** | **string** | CommitID that is being built. | 
**LogDescriptorId** | **string** | LogDescriptorID for the log for this build. | 
**Ref** | **string** | Ref is the git ref the build is for (e.g. branch or tag). | 
**Status** | **string** | Status reflects where the build is in the queue. | 
**Timings** | [**WorkflowTimings**](WorkflowTimings.md) |  | 
**Error** | Pointer to **string** | Error is set if the build finished with an error (or nil if the build succeeded). | [optional] 
**Opts** | [**BuildOptions**](BuildOptions.md) |  | 
**LogDescriptorUrl** | **string** | URL of the log for this build. | 
**ArtifactSearchUrl** | **string** | URL to use for searching for artifacts from this build. | 

## Methods

### NewBuild

`func NewBuild(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, repoId string, commitId string, logDescriptorId string, ref string, status string, timings WorkflowTimings, opts BuildOptions, logDescriptorUrl string, artifactSearchUrl string, ) *Build`

NewBuild instantiates a new Build object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildWithDefaults

`func NewBuildWithDefaults() *Build`

NewBuildWithDefaults instantiates a new Build object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Build) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Build) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Build) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Build) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Build) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Build) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Build) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Build) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Build) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Build) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Build) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Build) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Build) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Build) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Build) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Build) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEtag

`func (o *Build) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Build) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Build) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetName

`func (o *Build) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Build) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Build) SetName(v string)`

SetName sets Name field to given value.


### GetRepoId

`func (o *Build) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *Build) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *Build) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetCommitId

`func (o *Build) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *Build) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *Build) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetLogDescriptorId

`func (o *Build) GetLogDescriptorId() string`

GetLogDescriptorId returns the LogDescriptorId field if non-nil, zero value otherwise.

### GetLogDescriptorIdOk

`func (o *Build) GetLogDescriptorIdOk() (*string, bool)`

GetLogDescriptorIdOk returns a tuple with the LogDescriptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorId

`func (o *Build) SetLogDescriptorId(v string)`

SetLogDescriptorId sets LogDescriptorId field to given value.


### GetRef

`func (o *Build) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Build) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Build) SetRef(v string)`

SetRef sets Ref field to given value.


### GetStatus

`func (o *Build) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Build) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Build) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimings

`func (o *Build) GetTimings() WorkflowTimings`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *Build) GetTimingsOk() (*WorkflowTimings, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *Build) SetTimings(v WorkflowTimings)`

SetTimings sets Timings field to given value.


### GetError

`func (o *Build) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Build) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Build) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Build) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOpts

`func (o *Build) GetOpts() BuildOptions`

GetOpts returns the Opts field if non-nil, zero value otherwise.

### GetOptsOk

`func (o *Build) GetOptsOk() (*BuildOptions, bool)`

GetOptsOk returns a tuple with the Opts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpts

`func (o *Build) SetOpts(v BuildOptions)`

SetOpts sets Opts field to given value.


### GetLogDescriptorUrl

`func (o *Build) GetLogDescriptorUrl() string`

GetLogDescriptorUrl returns the LogDescriptorUrl field if non-nil, zero value otherwise.

### GetLogDescriptorUrlOk

`func (o *Build) GetLogDescriptorUrlOk() (*string, bool)`

GetLogDescriptorUrlOk returns a tuple with the LogDescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorUrl

`func (o *Build) SetLogDescriptorUrl(v string)`

SetLogDescriptorUrl sets LogDescriptorUrl field to given value.


### GetArtifactSearchUrl

`func (o *Build) GetArtifactSearchUrl() string`

GetArtifactSearchUrl returns the ArtifactSearchUrl field if non-nil, zero value otherwise.

### GetArtifactSearchUrlOk

`func (o *Build) GetArtifactSearchUrlOk() (*string, bool)`

GetArtifactSearchUrlOk returns a tuple with the ArtifactSearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSearchUrl

`func (o *Build) SetArtifactSearchUrl(v string)`

SetArtifactSearchUrl sets ArtifactSearchUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


