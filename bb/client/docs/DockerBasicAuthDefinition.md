# DockerBasicAuthDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | [**SecretStringDefinition**](SecretStringDefinition.md) |  | 
**Password** | [**SecretStringDefinition**](SecretStringDefinition.md) |  | 

## Methods

### NewDockerBasicAuthDefinition

`func NewDockerBasicAuthDefinition(username SecretStringDefinition, password SecretStringDefinition, ) *DockerBasicAuthDefinition`

NewDockerBasicAuthDefinition instantiates a new DockerBasicAuthDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerBasicAuthDefinitionWithDefaults

`func NewDockerBasicAuthDefinitionWithDefaults() *DockerBasicAuthDefinition`

NewDockerBasicAuthDefinitionWithDefaults instantiates a new DockerBasicAuthDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DockerBasicAuthDefinition) GetUsername() SecretStringDefinition`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DockerBasicAuthDefinition) GetUsernameOk() (*SecretStringDefinition, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DockerBasicAuthDefinition) SetUsername(v SecretStringDefinition)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *DockerBasicAuthDefinition) GetPassword() SecretStringDefinition`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DockerBasicAuthDefinition) GetPasswordOk() (*SecretStringDefinition, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DockerBasicAuthDefinition) SetPassword(v SecretStringDefinition)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


