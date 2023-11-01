# BuildDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version number of the build definition format used | 
**Jobs** | [**[]JobDefinition**](JobDefinition.md) |  | 

## Methods

### NewBuildDefinition

`func NewBuildDefinition(version string, jobs []JobDefinition, ) *BuildDefinition`

NewBuildDefinition instantiates a new BuildDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildDefinitionWithDefaults

`func NewBuildDefinitionWithDefaults() *BuildDefinition`

NewBuildDefinitionWithDefaults instantiates a new BuildDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *BuildDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BuildDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BuildDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetJobs

`func (o *BuildDefinition) GetJobs() []JobDefinition`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BuildDefinition) GetJobsOk() (*[]JobDefinition, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BuildDefinition) SetJobs(v []JobDefinition)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


