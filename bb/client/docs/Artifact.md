# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Artifact on the BuildBeaver server | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Etag** | **string** |  | 
**Name** | **string** | Name of the artifact. | 
**StepId** | **string** | The ID of the step (within a job) that created this artifact. | 
**GroupName** | **string** | The name associated with the one or more artifacts identified by an ArtifactDefinition in the build config. | 
**Path** | **string** | The filesystem path that the artifact was found at, relative to the job workspace. | 
**HashType** | **string** | tThe type of hashing algorithm used to hash the data. | 
**Hash** | **string** | The hex-encoded hash of the artifact data. This This may be set later if the hash is not known yet. | 
**Size** | **int32** | Size of the artifact file in bytes. | 
**Mime** | **string** | Mime type of the artifact, or empty if not known. | 
**Sealed** | **bool** | Sealed is true once the data for the artifact has successfully been uploaded and the file contents are now locked. Until Sealed is true various pieces of metadata such as the file size and hash etc. will be unset. Note that if sealed is false it doesn&#39;t necessarily mean no data has been uploaded to the blob store yet, and so we must still verify that the backing data is deleted before garbage collecting unsealed artifact files. | 
**DataUrl** | **string** | URL to use for fetching the bytes of data making up the artifact. | 

## Methods

### NewArtifact

`func NewArtifact(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, stepId string, groupName string, path string, hashType string, hash string, size int32, mime string, sealed bool, dataUrl string, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Artifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Artifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Artifact) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Artifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artifact) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Artifact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Artifact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Artifact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Artifact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Artifact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Artifact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEtag

`func (o *Artifact) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Artifact) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Artifact) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetName

`func (o *Artifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artifact) SetName(v string)`

SetName sets Name field to given value.


### GetStepId

`func (o *Artifact) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *Artifact) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *Artifact) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetGroupName

`func (o *Artifact) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Artifact) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Artifact) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetPath

`func (o *Artifact) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Artifact) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Artifact) SetPath(v string)`

SetPath sets Path field to given value.


### GetHashType

`func (o *Artifact) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *Artifact) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *Artifact) SetHashType(v string)`

SetHashType sets HashType field to given value.


### GetHash

`func (o *Artifact) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Artifact) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Artifact) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSize

`func (o *Artifact) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Artifact) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Artifact) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMime

`func (o *Artifact) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *Artifact) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *Artifact) SetMime(v string)`

SetMime sets Mime field to given value.


### GetSealed

`func (o *Artifact) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *Artifact) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *Artifact) SetSealed(v bool)`

SetSealed sets Sealed field to given value.


### GetDataUrl

`func (o *Artifact) GetDataUrl() string`

GetDataUrl returns the DataUrl field if non-nil, zero value otherwise.

### GetDataUrlOk

`func (o *Artifact) GetDataUrlOk() (*string, bool)`

GetDataUrlOk returns a tuple with the DataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrl

`func (o *Artifact) SetDataUrl(v string)`

SetDataUrl sets DataUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


