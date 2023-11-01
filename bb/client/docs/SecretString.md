# SecretString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | An explicit, non-secret value for an environment variable or field. | [optional] 
**ValueFromSecret** | Pointer to **string** | The name of the secret that contains the value, for an environment variable or field. | [optional] 

## Methods

### NewSecretString

`func NewSecretString() *SecretString`

NewSecretString instantiates a new SecretString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStringWithDefaults

`func NewSecretStringWithDefaults() *SecretString`

NewSecretStringWithDefaults instantiates a new SecretString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SecretString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecretString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFromSecret

`func (o *SecretString) GetValueFromSecret() string`

GetValueFromSecret returns the ValueFromSecret field if non-nil, zero value otherwise.

### GetValueFromSecretOk

`func (o *SecretString) GetValueFromSecretOk() (*string, bool)`

GetValueFromSecretOk returns a tuple with the ValueFromSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFromSecret

`func (o *SecretString) SetValueFromSecret(v string)`

SetValueFromSecret sets ValueFromSecret field to given value.

### HasValueFromSecret

`func (o *SecretString) HasValueFromSecret() bool`

HasValueFromSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


