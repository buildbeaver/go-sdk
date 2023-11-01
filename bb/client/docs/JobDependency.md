# JobDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | **string** | The name of the workflow containing the job that must complete, or an empty string for the default workflow. | 
**JobName** | **string** | The name of the job that must complete. | 
**ArtifactDependencies** | [**[]ArtifactDependency**](ArtifactDependency.md) | The set of artifacts required from the job that must complete. | 

## Methods

### NewJobDependency

`func NewJobDependency(workflow string, jobName string, artifactDependencies []ArtifactDependency, ) *JobDependency`

NewJobDependency instantiates a new JobDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDependencyWithDefaults

`func NewJobDependencyWithDefaults() *JobDependency`

NewJobDependencyWithDefaults instantiates a new JobDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *JobDependency) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *JobDependency) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *JobDependency) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetJobName

`func (o *JobDependency) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *JobDependency) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *JobDependency) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetArtifactDependencies

`func (o *JobDependency) GetArtifactDependencies() []ArtifactDependency`

GetArtifactDependencies returns the ArtifactDependencies field if non-nil, zero value otherwise.

### GetArtifactDependenciesOk

`func (o *JobDependency) GetArtifactDependenciesOk() (*[]ArtifactDependency, bool)`

GetArtifactDependenciesOk returns a tuple with the ArtifactDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactDependencies

`func (o *JobDependency) SetArtifactDependencies(v []ArtifactDependency)`

SetArtifactDependencies sets ArtifactDependencies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


