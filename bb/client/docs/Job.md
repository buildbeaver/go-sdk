# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Job resource on the BuildBeaver server | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Etag** | **string** |  | 
**Name** | **string** | Job name, in URL format | 
**Workflow** | **string** | Workflow the job is a part of, or empty if the job is part of the default workflow. | 
**Description** | **string** | Optional human-readable description of the job. | 
**Type** | **string** | Type of the job (e.g. docker, exec etc.) | 
**RunsOn** | **[]string** | RunsOn contains a set of labels that this job requires runners to have. | 
**Docker** | Pointer to [**DockerConfig**](DockerConfig.md) |  | [optional] 
**StepExecution** | **string** | Determines how the runner will execute steps within this job | 
**Depends** | [**[]JobDependency**](JobDependency.md) | Dependencies on other jobs and their artifacts. Each JobDependency declares that this job depends on the successful execution of another, and optionally that this job consumes one or more artifacts from the other. | 
**Services** | [**[]Service**](Service.md) | Services to run in the background for the duration of the job; services are started before the first step is run, and stopped after the last step completes | 
**FingerprintCommands** | **[]string** | Shell commands to execute to generate a unique fingerprint for the jobs; two jobs in the same repo with the same name and fingerprint are considered identical | 
**Artifacts** | Pointer to [**[]ArtifactDefinition**](ArtifactDefinition.md) | A list of all artifacts the job is expected to produce that will be saved to the artifact store at the end of the job&#39;s execution | [optional] 
**Environment** | [**[]EnvVar**](EnvVar.md) | A list of environment variables to export prior to executing the job | 
**BuildId** | **string** | ID of the build this job forms a part of. | 
**RepoId** | **string** | RepoID that was committed to. | 
**CommitId** | **string** | CommitID that the job was generated from. | 
**LogDescriptorId** | **string** | LogDescriptorID points to the log for this job. | 
**RunnerId** | **string** | RunnerID is the id of the runner this job executed on, or empty if the job has not run yet (or did/will not run). | 
**IndirectToJobId** | **string** | IndirectToJobID records the ID of a job that previously ran successfully as part of another build and which is functionally identical to this job. If this is set it means this job did not actually run to avoid redundantly running the same thing more than once. | 
**Ref** | **string** | Ref is the git ref from the build that the job was generated from (e.g. branch or tag) | 
**Status** | **string** | Status reflects where the job is in the queue. | 
**Error** | Pointer to **string** | Error is set if the job finished with an error (or empty if the job succeeded). | [optional] 
**Timings** | [**WorkflowTimings**](WorkflowTimings.md) |  | 
**Fingerprint** | Pointer to **string** | Fingerprint contains the hashed output of FingerprintCommands, as well as any other inputs the agent added (such as artifact hashes). This is only available after the job has run successfully. | [optional] 
**FingerprintHashType** | Pointer to **string** | FingerprintHashType is the type of hashing algorithm used to produce the fingerprint. | [optional] 
**LogDescriptorUrl** | **string** | URL of the log for this job. | 
**IndirectJobUrl** | **string** | URL to the job that this job indirects to, if any. | 

## Methods

### NewJob

`func NewJob(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, workflow string, description string, type_ string, runsOn []string, stepExecution string, depends []JobDependency, services []Service, fingerprintCommands []string, environment []EnvVar, buildId string, repoId string, commitId string, logDescriptorId string, runnerId string, indirectToJobId string, ref string, status string, timings WorkflowTimings, logDescriptorUrl string, indirectJobUrl string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Job) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Job) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Job) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Job) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Job) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Job) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Job) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEtag

`func (o *Job) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Job) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Job) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetWorkflow

`func (o *Job) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *Job) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *Job) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.


### GetDescription

`func (o *Job) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Job) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Job) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.


### GetRunsOn

`func (o *Job) GetRunsOn() []string`

GetRunsOn returns the RunsOn field if non-nil, zero value otherwise.

### GetRunsOnOk

`func (o *Job) GetRunsOnOk() (*[]string, bool)`

GetRunsOnOk returns a tuple with the RunsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsOn

`func (o *Job) SetRunsOn(v []string)`

SetRunsOn sets RunsOn field to given value.


### GetDocker

`func (o *Job) GetDocker() DockerConfig`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *Job) GetDockerOk() (*DockerConfig, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *Job) SetDocker(v DockerConfig)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *Job) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetStepExecution

`func (o *Job) GetStepExecution() string`

GetStepExecution returns the StepExecution field if non-nil, zero value otherwise.

### GetStepExecutionOk

`func (o *Job) GetStepExecutionOk() (*string, bool)`

GetStepExecutionOk returns a tuple with the StepExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecution

`func (o *Job) SetStepExecution(v string)`

SetStepExecution sets StepExecution field to given value.


### GetDepends

`func (o *Job) GetDepends() []JobDependency`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *Job) GetDependsOk() (*[]JobDependency, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *Job) SetDepends(v []JobDependency)`

SetDepends sets Depends field to given value.


### GetServices

`func (o *Job) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Job) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Job) SetServices(v []Service)`

SetServices sets Services field to given value.


### GetFingerprintCommands

`func (o *Job) GetFingerprintCommands() []string`

GetFingerprintCommands returns the FingerprintCommands field if non-nil, zero value otherwise.

### GetFingerprintCommandsOk

`func (o *Job) GetFingerprintCommandsOk() (*[]string, bool)`

GetFingerprintCommandsOk returns a tuple with the FingerprintCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintCommands

`func (o *Job) SetFingerprintCommands(v []string)`

SetFingerprintCommands sets FingerprintCommands field to given value.


### GetArtifacts

`func (o *Job) GetArtifacts() []ArtifactDefinition`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Job) GetArtifactsOk() (*[]ArtifactDefinition, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Job) SetArtifacts(v []ArtifactDefinition)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *Job) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetEnvironment

`func (o *Job) GetEnvironment() []EnvVar`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Job) GetEnvironmentOk() (*[]EnvVar, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Job) SetEnvironment(v []EnvVar)`

SetEnvironment sets Environment field to given value.


### GetBuildId

`func (o *Job) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Job) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Job) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetRepoId

`func (o *Job) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *Job) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *Job) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetCommitId

`func (o *Job) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *Job) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *Job) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetLogDescriptorId

`func (o *Job) GetLogDescriptorId() string`

GetLogDescriptorId returns the LogDescriptorId field if non-nil, zero value otherwise.

### GetLogDescriptorIdOk

`func (o *Job) GetLogDescriptorIdOk() (*string, bool)`

GetLogDescriptorIdOk returns a tuple with the LogDescriptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorId

`func (o *Job) SetLogDescriptorId(v string)`

SetLogDescriptorId sets LogDescriptorId field to given value.


### GetRunnerId

`func (o *Job) GetRunnerId() string`

GetRunnerId returns the RunnerId field if non-nil, zero value otherwise.

### GetRunnerIdOk

`func (o *Job) GetRunnerIdOk() (*string, bool)`

GetRunnerIdOk returns a tuple with the RunnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerId

`func (o *Job) SetRunnerId(v string)`

SetRunnerId sets RunnerId field to given value.


### GetIndirectToJobId

`func (o *Job) GetIndirectToJobId() string`

GetIndirectToJobId returns the IndirectToJobId field if non-nil, zero value otherwise.

### GetIndirectToJobIdOk

`func (o *Job) GetIndirectToJobIdOk() (*string, bool)`

GetIndirectToJobIdOk returns a tuple with the IndirectToJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectToJobId

`func (o *Job) SetIndirectToJobId(v string)`

SetIndirectToJobId sets IndirectToJobId field to given value.


### GetRef

`func (o *Job) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Job) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Job) SetRef(v string)`

SetRef sets Ref field to given value.


### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *Job) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Job) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Job) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Job) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTimings

`func (o *Job) GetTimings() WorkflowTimings`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *Job) GetTimingsOk() (*WorkflowTimings, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *Job) SetTimings(v WorkflowTimings)`

SetTimings sets Timings field to given value.


### GetFingerprint

`func (o *Job) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Job) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Job) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Job) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprintHashType

`func (o *Job) GetFingerprintHashType() string`

GetFingerprintHashType returns the FingerprintHashType field if non-nil, zero value otherwise.

### GetFingerprintHashTypeOk

`func (o *Job) GetFingerprintHashTypeOk() (*string, bool)`

GetFingerprintHashTypeOk returns a tuple with the FingerprintHashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintHashType

`func (o *Job) SetFingerprintHashType(v string)`

SetFingerprintHashType sets FingerprintHashType field to given value.

### HasFingerprintHashType

`func (o *Job) HasFingerprintHashType() bool`

HasFingerprintHashType returns a boolean if a field has been set.

### GetLogDescriptorUrl

`func (o *Job) GetLogDescriptorUrl() string`

GetLogDescriptorUrl returns the LogDescriptorUrl field if non-nil, zero value otherwise.

### GetLogDescriptorUrlOk

`func (o *Job) GetLogDescriptorUrlOk() (*string, bool)`

GetLogDescriptorUrlOk returns a tuple with the LogDescriptorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDescriptorUrl

`func (o *Job) SetLogDescriptorUrl(v string)`

SetLogDescriptorUrl sets LogDescriptorUrl field to given value.


### GetIndirectJobUrl

`func (o *Job) GetIndirectJobUrl() string`

GetIndirectJobUrl returns the IndirectJobUrl field if non-nil, zero value otherwise.

### GetIndirectJobUrlOk

`func (o *Job) GetIndirectJobUrlOk() (*string, bool)`

GetIndirectJobUrlOk returns a tuple with the IndirectJobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectJobUrl

`func (o *Job) SetIndirectJobUrl(v string)`

SetIndirectJobUrl sets IndirectJobUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


