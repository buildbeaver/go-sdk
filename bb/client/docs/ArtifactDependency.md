# ArtifactDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The name of the workflow containing the job that will produce the artifact(s), or an empty string for the default workflow. | 
**JobName** | **string** | The name of the job that that will produce the artifact(s). | 
**GroupName** | **string** | The name of the group of artifacts. | 

## Methods

### NewArtifactDependency

`func NewArtifactDependency(workflow string, jobName string, groupName string, ) *ArtifactDependency`

NewArtifactDependency instantiates a new ArtifactDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDependencyWithDefaults

`func NewArtifactDependencyWithDefaults() *ArtifactDependency`

NewArtifactDependencyWithDefaults instantiates a new ArtifactDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *ArtifactDependency) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ArtifactDependency) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ArtifactDependency) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetJobName

`func (o *ArtifactDependency) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ArtifactDependency) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ArtifactDependency) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetGroupName

`func (o *ArtifactDependency) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ArtifactDependency) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ArtifactDependency) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


