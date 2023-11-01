# BuildGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | [**Build**](Build.md) |  | 
**Jobs** | [**[]JobGraph**](JobGraph.md) | The current set of jobs making up the build, including the full job graph/steps for each. | 
**Repo** | [**Repo**](Repo.md) |  | 
**Commit** | [**Commit**](Commit.md) |  | 

## Methods

### NewBuildGraph

`func NewBuildGraph(build Build, jobs []JobGraph, repo Repo, commit Commit, ) *BuildGraph`

NewBuildGraph instantiates a new BuildGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildGraphWithDefaults

`func NewBuildGraphWithDefaults() *BuildGraph`

NewBuildGraphWithDefaults instantiates a new BuildGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *BuildGraph) GetBuild() Build`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *BuildGraph) GetBuildOk() (*Build, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *BuildGraph) SetBuild(v Build)`

SetBuild sets Build field to given value.


### GetJobs

`func (o *BuildGraph) GetJobs() []JobGraph`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BuildGraph) GetJobsOk() (*[]JobGraph, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BuildGraph) SetJobs(v []JobGraph)`

SetJobs sets Jobs field to given value.


### GetRepo

`func (o *BuildGraph) GetRepo() Repo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *BuildGraph) GetRepoOk() (*Repo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *BuildGraph) SetRepo(v Repo)`

SetRepo sets Repo field to given value.


### GetCommit

`func (o *BuildGraph) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BuildGraph) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BuildGraph) SetCommit(v Commit)`

SetCommit sets Commit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


