# Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Step resource on the BuildBeaver server | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Etag** | **string** |  | 
**Name** | **string** | Step name, in URL format | 
**Description** | **string** | Optional human-readable description of the job. | 
**Commands** | **[]string** | A list of one or more shell commands to execute during the step. | 
**Depends** | [**[]StepDependency**](StepDependency.md) | Dependencies this step has on other steps within the job (see dependency syntax) | 
**JobId** | **string** | ID of the job this step forms a part of | 
**RepoId** | **string** | RepoID that the step is building from. | 
**RunnerId** | **string** | RunnerID that ran the step (or empty if the step has not run yet). | 
**LogDescriptorId** | **string** | LogDescriptorID points to the log for this step. | 
**Status** | **string** | Status reflects where the step is in processing. | 
**Error** | Pointer to **string** | Error is set if the step finished with an error (or empty if the step succeeded). | [optional] 
**Timings** | [**WorkflowTimings**](WorkflowTimings.md) |  | 
**RunnerUrl** | **string** | URL of the runner that ran the step (or empty if the step has not run yet). | 
**LogDescriptorUrl** | **string** | URL of the log for this step. | 

## Methods

### NewStep

`func NewStep(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, description string, commands []string, depends []StepDependency, jobId string, repoId string, runnerId string, logDescriptorId string, status string, timings WorkflowTimings, runnerUrl string, logDescriptorUrl string, ) *Step`

NewStep instantiates a new Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepWithDefaults

`func NewStepWithDefaults() *Step`

NewStepWithDefaults instantiates a new Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Step) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Step) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Step) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Step) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Step) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Step) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Step) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Step) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Step) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Step) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Step) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Step) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Step) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Step) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Step) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Step) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEtag

`func (o *Step) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Step) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Step) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetName

`func (o *Step) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Step) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Step) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Step) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Step) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Step) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCommands

`func (o *Step) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *Step) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *Step) SetCommands(v []string)`

SetCommands sets Commands field to given value.


### GetDepends

`func (o *Step) GetDepends() []StepDependency`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *Step) GetDependsOk() (*[]StepDependency, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *Step) SetDepends(v []StepDependency)`

SetDepends sets Depends field to given value.


### GetJobId

`func (o *Step) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Step) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Step) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetRepoId

`func (o *Step) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *Step) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *Step) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetRunnerId

`func (o *Step) GetRunnerId() string`

GetRunnerId returns the RunnerId field if non-nil, zero value otherwise.

### GetRunnerIdOk

`func (o *Step) GetRunnerIdOk() (*string, bool)`

GetRunnerIdOk returns a tuple with the RunnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerId

`func (o *Step) SetRunnerId(v string)`

SetRunnerId sets RunnerId field to given value.


### GetLogDescriptorId

`func (o *Step) GetLogDescriptorId() string`

GetLogDescriptorId returns the LogDescriptorId field if non-nil, zero value otherwise.

### GetLogDescriptorIdOk

`func (o *Step) GetLogDescriptorIdOk() (*string, bool)`

GetLogDescriptorIdOk returns a tuple with the LogDescriptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorId

`func (o *Step) SetLogDescriptorId(v string)`

SetLogDescriptorId sets LogDescriptorId field to given value.


### GetStatus

`func (o *Step) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Step) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Step) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *Step) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Step) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Step) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Step) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTimings

`func (o *Step) GetTimings() WorkflowTimings`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *Step) GetTimingsOk() (*WorkflowTimings, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *Step) SetTimings(v WorkflowTimings)`

SetTimings sets Timings field to given value.


### GetRunnerUrl

`func (o *Step) GetRunnerUrl() string`

GetRunnerUrl returns the RunnerUrl field if non-nil, zero value otherwise.

### GetRunnerUrlOk

`func (o *Step) GetRunnerUrlOk() (*string, bool)`

GetRunnerUrlOk returns a tuple with the RunnerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerUrl

`func (o *Step) SetRunnerUrl(v string)`

SetRunnerUrl sets RunnerUrl field to given value.


### GetLogDescriptorUrl

`func (o *Step) GetLogDescriptorUrl() string`

GetLogDescriptorUrl returns the LogDescriptorUrl field if non-nil, zero value otherwise.

### GetLogDescriptorUrlOk

`func (o *Step) GetLogDescriptorUrlOk() (*string, bool)`

GetLogDescriptorUrlOk returns a tuple with the LogDescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorUrl

`func (o *Step) SetLogDescriptorUrl(v string)`

SetLogDescriptorUrl sets LogDescriptorUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


