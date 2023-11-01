# BuildOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | **bool** | True to force all jobs in the build to run by ignoring fingerprints. | 
**NodesToRun** | [**[]NodeFQN**](NodeFQN.md) | Contains zero or more workflows, jobs and steps to run. If no nodes are specified then all workflows, jobs and steps will be run. | 

## Methods

### NewBuildOptions

`func NewBuildOptions(force bool, nodesToRun []NodeFQN, ) *BuildOptions`

NewBuildOptions instantiates a new BuildOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildOptionsWithDefaults

`func NewBuildOptionsWithDefaults() *BuildOptions`

NewBuildOptionsWithDefaults instantiates a new BuildOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *BuildOptions) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *BuildOptions) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *BuildOptions) SetForce(v bool)`

SetForce sets Force field to given value.


### GetNodesToRun

`func (o *BuildOptions) GetNodesToRun() []NodeFQN`

GetNodesToRun returns the NodesToRun field if non-nil, zero value otherwise.

### GetNodesToRunOk

`func (o *BuildOptions) GetNodesToRunOk() (*[]NodeFQN, bool)`

GetNodesToRunOk returns a tuple with the NodesToRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesToRun

`func (o *BuildOptions) SetNodesToRun(v []NodeFQN)`

SetNodesToRun sets NodesToRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


