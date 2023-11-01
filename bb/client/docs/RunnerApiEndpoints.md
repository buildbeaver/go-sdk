# RunnerApiEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value of an environment variable, for variables that do not use secrets. | [optional] 
**FromSecret** | Pointer to **string** | The name of the secret that contains the value, for environment variables that use secrets. | [optional] 

## Methods

### NewRunnerApiEndpoints

`func NewRunnerApiEndpoints() *RunnerApiEndpoints`

NewRunnerApiEndpoints instantiates a new RunnerApiEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerApiEndpointsWithDefaults

`func NewRunnerApiEndpointsWithDefaults() *RunnerApiEndpoints`

NewRunnerApiEndpointsWithDefaults instantiates a new RunnerApiEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RunnerApiEndpoints) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RunnerApiEndpoints) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RunnerApiEndpoints) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RunnerApiEndpoints) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFromSecret

`func (o *RunnerApiEndpoints) GetFromSecret() string`

GetFromSecret returns the FromSecret field if non-nil, zero value otherwise.

### GetFromSecretOk

`func (o *RunnerApiEndpoints) GetFromSecretOk() (*string, bool)`

GetFromSecretOk returns a tuple with the FromSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSecret

`func (o *RunnerApiEndpoints) SetFromSecret(v string)`

SetFromSecret sets FromSecret field to given value.

### HasFromSecret

`func (o *RunnerApiEndpoints) HasFromSecret() bool`

HasFromSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


