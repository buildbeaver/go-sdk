# WorkflowTimings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueuedAt** | Pointer to **time.Time** |  | [optional] 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**RunningAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**CanceledAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWorkflowTimings

`func NewWorkflowTimings() *WorkflowTimings`

NewWorkflowTimings instantiates a new WorkflowTimings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTimingsWithDefaults

`func NewWorkflowTimingsWithDefaults() *WorkflowTimings`

NewWorkflowTimingsWithDefaults instantiates a new WorkflowTimings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuedAt

`func (o *WorkflowTimings) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *WorkflowTimings) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *WorkflowTimings) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *WorkflowTimings) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *WorkflowTimings) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *WorkflowTimings) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *WorkflowTimings) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *WorkflowTimings) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetRunningAt

`func (o *WorkflowTimings) GetRunningAt() time.Time`

GetRunningAt returns the RunningAt field if non-nil, zero value otherwise.

### GetRunningAtOk

`func (o *WorkflowTimings) GetRunningAtOk() (*time.Time, bool)`

GetRunningAtOk returns a tuple with the RunningAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningAt

`func (o *WorkflowTimings) SetRunningAt(v time.Time)`

SetRunningAt sets RunningAt field to given value.

### HasRunningAt

`func (o *WorkflowTimings) HasRunningAt() bool`

HasRunningAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *WorkflowTimings) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *WorkflowTimings) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *WorkflowTimings) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *WorkflowTimings) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetCanceledAt

`func (o *WorkflowTimings) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *WorkflowTimings) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *WorkflowTimings) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *WorkflowTimings) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


