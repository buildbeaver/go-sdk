# JobDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Job name. This can optionally include a workflow as a prefix (dot-separated), as an alternative to specifying an explicit &#39;workflow&#39; element. | 
**Workflow** | Pointer to **string** | Workflow the job is a part of. If not specified then the job is part of the default workflow. | [optional] 
**Description** | Pointer to **string** | Optional human-readable description of the job. | [optional] 
**Type** | Pointer to **string** | Type of the job (e.g. docker, exec etc.) | [optional] 
**RunsOn** | Pointer to **[]string** | RunsOn contains a set of labels that this job requires runners to have. | [optional] 
**Docker** | Pointer to [**DockerConfigDefinition**](DockerConfigDefinition.md) |  | [optional] 
**StepExecution** | **string** | Determines how the runner will execute steps within this job | 
**Depends** | Pointer to **[]string** | Dependencies on other jobs and their artifacts (see dependency syntax) | [optional] 
**Services** | Pointer to [**[]ServiceDefinition**](ServiceDefinition.md) | Services to run in the background for the duration of the job; services are started before the first step is run, and stopped after the last step completes | [optional] 
**Fingerprint** | Pointer to **[]string** | Shell commands to execute to generate a unique fingerprint for the jobs; two jobs in the same repo with the same name and fingerprint are considered identical | [optional] 
**Artifacts** | Pointer to [**[]ArtifactDefinition**](ArtifactDefinition.md) | A list of all artifacts the job is expected to produce that will be saved to the artifact store at the end of the job&#39;s execution | [optional] 
**Environment** | [**map[string]SecretStringDefinition**](SecretStringDefinition.md) | A list of environment variables to export prior to executing the job | 
**Steps** | [**[]StepDefinition**](StepDefinition.md) | The set of steps within the job | 

## Methods

### NewJobDefinition

`func NewJobDefinition(name string, stepExecution string, environment map[string]SecretStringDefinition, steps []StepDefinition, ) *JobDefinition`

NewJobDefinition instantiates a new JobDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDefinitionWithDefaults

`func NewJobDefinitionWithDefaults() *JobDefinition`

NewJobDefinitionWithDefaults instantiates a new JobDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetWorkflow

`func (o *JobDefinition) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *JobDefinition) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *JobDefinition) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *JobDefinition) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetDescription

`func (o *JobDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *JobDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRunsOn

`func (o *JobDefinition) GetRunsOn() []string`

GetRunsOn returns the RunsOn field if non-nil, zero value otherwise.

### GetRunsOnOk

`func (o *JobDefinition) GetRunsOnOk() (*[]string, bool)`

GetRunsOnOk returns a tuple with the RunsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsOn

`func (o *JobDefinition) SetRunsOn(v []string)`

SetRunsOn sets RunsOn field to given value.

### HasRunsOn

`func (o *JobDefinition) HasRunsOn() bool`

HasRunsOn returns a boolean if a field has been set.

### GetDocker

`func (o *JobDefinition) GetDocker() DockerConfigDefinition`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *JobDefinition) GetDockerOk() (*DockerConfigDefinition, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *JobDefinition) SetDocker(v DockerConfigDefinition)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *JobDefinition) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetStepExecution

`func (o *JobDefinition) GetStepExecution() string`

GetStepExecution returns the StepExecution field if non-nil, zero value otherwise.

### GetStepExecutionOk

`func (o *JobDefinition) GetStepExecutionOk() (*string, bool)`

GetStepExecutionOk returns a tuple with the StepExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecution

`func (o *JobDefinition) SetStepExecution(v string)`

SetStepExecution sets StepExecution field to given value.


### GetDepends

`func (o *JobDefinition) GetDepends() []string`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *JobDefinition) GetDependsOk() (*[]string, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *JobDefinition) SetDepends(v []string)`

SetDepends sets Depends field to given value.

### HasDepends

`func (o *JobDefinition) HasDepends() bool`

HasDepends returns a boolean if a field has been set.

### GetServices

`func (o *JobDefinition) GetServices() []ServiceDefinition`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *JobDefinition) GetServicesOk() (*[]ServiceDefinition, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *JobDefinition) SetServices(v []ServiceDefinition)`

SetServices sets Services field to given value.

### HasServices

`func (o *JobDefinition) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetFingerprint

`func (o *JobDefinition) GetFingerprint() []string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *JobDefinition) GetFingerprintOk() (*[]string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *JobDefinition) SetFingerprint(v []string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *JobDefinition) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetArtifacts

`func (o *JobDefinition) GetArtifacts() []ArtifactDefinition`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *JobDefinition) GetArtifactsOk() (*[]ArtifactDefinition, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *JobDefinition) SetArtifacts(v []ArtifactDefinition)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *JobDefinition) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetEnvironment

`func (o *JobDefinition) GetEnvironment() map[string]SecretStringDefinition`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *JobDefinition) GetEnvironmentOk() (*map[string]SecretStringDefinition, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *JobDefinition) SetEnvironment(v map[string]SecretStringDefinition)`

SetEnvironment sets Environment field to given value.


### GetSteps

`func (o *JobDefinition) GetSteps() []StepDefinition`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *JobDefinition) GetStepsOk() (*[]StepDefinition, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *JobDefinition) SetSteps(v []StepDefinition)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


