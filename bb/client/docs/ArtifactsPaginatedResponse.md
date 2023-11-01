# ArtifactsPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of objects contained in the results, in this case &#39;artifact&#39; | 
**Results** | [**[]Artifact**](Artifact.md) |  | 
**PrevUrl** | Pointer to **string** | A URL to fetch to obtain the previous page of results before this one. | [optional] 
**PrevCursor** | Pointer to **string** | A cursor that can be used as a query parameter to obtain the previous page of results before this one. | [optional] 
**NextUrl** | Pointer to **string** | A URL to fetch to obtain the next page of results after this one. | [optional] 
**NextCursor** | Pointer to **string** | A cursor that can be used as a query parameter to obtain the next page of results after this one. | [optional] 

## Methods

### NewArtifactsPaginatedResponse

`func NewArtifactsPaginatedResponse(kind string, results []Artifact, ) *ArtifactsPaginatedResponse`

NewArtifactsPaginatedResponse instantiates a new ArtifactsPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsPaginatedResponseWithDefaults

`func NewArtifactsPaginatedResponseWithDefaults() *ArtifactsPaginatedResponse`

NewArtifactsPaginatedResponseWithDefaults instantiates a new ArtifactsPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ArtifactsPaginatedResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactsPaginatedResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactsPaginatedResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetResults

`func (o *ArtifactsPaginatedResponse) GetResults() []Artifact`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ArtifactsPaginatedResponse) GetResultsOk() (*[]Artifact, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ArtifactsPaginatedResponse) SetResults(v []Artifact)`

SetResults sets Results field to given value.


### GetPrevUrl

`func (o *ArtifactsPaginatedResponse) GetPrevUrl() string`

GetPrevUrl returns the PrevUrl field if non-nil, zero value otherwise.

### GetPrevUrlOk

`func (o *ArtifactsPaginatedResponse) GetPrevUrlOk() (*string, bool)`

GetPrevUrlOk returns a tuple with the PrevUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevUrl

`func (o *ArtifactsPaginatedResponse) SetPrevUrl(v string)`

SetPrevUrl sets PrevUrl field to given value.

### HasPrevUrl

`func (o *ArtifactsPaginatedResponse) HasPrevUrl() bool`

HasPrevUrl returns a boolean if a field has been set.

### GetPrevCursor

`func (o *ArtifactsPaginatedResponse) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *ArtifactsPaginatedResponse) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *ArtifactsPaginatedResponse) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *ArtifactsPaginatedResponse) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.

### GetNextUrl

`func (o *ArtifactsPaginatedResponse) GetNextUrl() string`

GetNextUrl returns the NextUrl field if non-nil, zero value otherwise.

### GetNextUrlOk

`func (o *ArtifactsPaginatedResponse) GetNextUrlOk() (*string, bool)`

GetNextUrlOk returns a tuple with the NextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUrl

`func (o *ArtifactsPaginatedResponse) SetNextUrl(v string)`

SetNextUrl sets NextUrl field to given value.

### HasNextUrl

`func (o *ArtifactsPaginatedResponse) HasNextUrl() bool`

HasNextUrl returns a boolean if a field has been set.

### GetNextCursor

`func (o *ArtifactsPaginatedResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ArtifactsPaginatedResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ArtifactsPaginatedResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ArtifactsPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


