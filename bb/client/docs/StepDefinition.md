# StepDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Step name, in URL format | 
**Description** | Pointer to **string** | Optional human-readable description of the job. | [optional] 
**Commands** | **[]string** | A list of one or more shell commands to execute during the step. | 
**Depends** | Pointer to **[]string** | Dependencies this step has on other steps within the job (see dependency syntax) | [optional] 

## Methods

### NewStepDefinition

`func NewStepDefinition(name string, commands []string, ) *StepDefinition`

NewStepDefinition instantiates a new StepDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepDefinitionWithDefaults

`func NewStepDefinitionWithDefaults() *StepDefinition`

NewStepDefinitionWithDefaults instantiates a new StepDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StepDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StepDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StepDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StepDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StepDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StepDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StepDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommands

`func (o *StepDefinition) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *StepDefinition) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *StepDefinition) SetCommands(v []string)`

SetCommands sets Commands field to given value.


### GetDepends

`func (o *StepDefinition) GetDepends() []string`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *StepDefinition) GetDependsOk() (*[]string, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *StepDefinition) SetDepends(v []string)`

SetDepends sets Depends field to given value.

### HasDepends

`func (o *StepDefinition) HasDepends() bool`

HasDepends returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


