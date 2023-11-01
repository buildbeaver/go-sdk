# ArtifactDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Uniquely identifies the group of artifacts specified in paths | 
**Paths** | **[]string** | One or more relative paths to artifacts that should be uploaded at the end of the build; these paths will be globbed, so that each path may identify one or more actual files | 

## Methods

### NewArtifactDefinition

`func NewArtifactDefinition(name string, paths []string, ) *ArtifactDefinition`

NewArtifactDefinition instantiates a new ArtifactDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDefinitionWithDefaults

`func NewArtifactDefinitionWithDefaults() *ArtifactDefinition`

NewArtifactDefinitionWithDefaults instantiates a new ArtifactDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArtifactDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetPaths

`func (o *ArtifactDefinition) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ArtifactDefinition) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ArtifactDefinition) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


