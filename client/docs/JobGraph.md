# JobGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**Job**](Job.md) |  | 
**Steps** | [**[]Step**](Step.md) | The set of steps within the job | 

## Methods

### NewJobGraph

`func NewJobGraph(job Job, steps []Step, ) *JobGraph`

NewJobGraph instantiates a new JobGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobGraphWithDefaults

`func NewJobGraphWithDefaults() *JobGraph`

NewJobGraphWithDefaults instantiates a new JobGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *JobGraph) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobGraph) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobGraph) SetJob(v Job)`

SetJob sets Job field to given value.


### GetSteps

`func (o *JobGraph) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *JobGraph) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *JobGraph) SetSteps(v []Step)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


