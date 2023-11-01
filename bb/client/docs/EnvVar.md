# EnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the environment variable. | 
**Value** | Pointer to **string** | Value of the variable, if the variable is set explicitly. | [optional] 
**ValueFromSecret** | Pointer to **string** | ValueFromSecret is the name of the secret to set this variable to, if setting the variable to a secret. | [optional] 

## Methods

### NewEnvVar

`func NewEnvVar(name string, ) *EnvVar`

NewEnvVar instantiates a new EnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarWithDefaults

`func NewEnvVarWithDefaults() *EnvVar`

NewEnvVarWithDefaults instantiates a new EnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFromSecret

`func (o *EnvVar) GetValueFromSecret() string`

GetValueFromSecret returns the ValueFromSecret field if non-nil, zero value otherwise.

### GetValueFromSecretOk

`func (o *EnvVar) GetValueFromSecretOk() (*string, bool)`

GetValueFromSecretOk returns a tuple with the ValueFromSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFromSecret

`func (o *EnvVar) SetValueFromSecret(v string)`

SetValueFromSecret sets ValueFromSecret field to given value.

### HasValueFromSecret

`func (o *EnvVar) HasValueFromSecret() bool`

HasValueFromSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


