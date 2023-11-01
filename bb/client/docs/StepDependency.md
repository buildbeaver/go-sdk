# StepDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | **string** | The name of the step that must complete before this step. | 

## Methods

### NewStepDependency

`func NewStepDependency(stepName string, ) *StepDependency`

NewStepDependency instantiates a new StepDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepDependencyWithDefaults

`func NewStepDependencyWithDefaults() *StepDependency`

NewStepDependencyWithDefaults instantiates a new StepDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepName

`func (o *StepDependency) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *StepDependency) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *StepDependency) SetStepName(v string)`

SetStepName sets StepName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


