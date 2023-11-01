# DockerAWSAuthDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRegion** | Pointer to **string** | The ECR AWS region to authenticate to. | [optional] 
**AwsAccessKeyId** | [**SecretStringDefinition**](SecretStringDefinition.md) |  | 
**AwsSecretAccessKey** | [**SecretStringDefinition**](SecretStringDefinition.md) |  | 

## Methods

### NewDockerAWSAuthDefinition

`func NewDockerAWSAuthDefinition(awsAccessKeyId SecretStringDefinition, awsSecretAccessKey SecretStringDefinition, ) *DockerAWSAuthDefinition`

NewDockerAWSAuthDefinition instantiates a new DockerAWSAuthDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerAWSAuthDefinitionWithDefaults

`func NewDockerAWSAuthDefinitionWithDefaults() *DockerAWSAuthDefinition`

NewDockerAWSAuthDefinitionWithDefaults instantiates a new DockerAWSAuthDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRegion

`func (o *DockerAWSAuthDefinition) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *DockerAWSAuthDefinition) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *DockerAWSAuthDefinition) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *DockerAWSAuthDefinition) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *DockerAWSAuthDefinition) GetAwsAccessKeyId() SecretStringDefinition`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *DockerAWSAuthDefinition) GetAwsAccessKeyIdOk() (*SecretStringDefinition, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *DockerAWSAuthDefinition) SetAwsAccessKeyId(v SecretStringDefinition)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### GetAwsSecretAccessKey

`func (o *DockerAWSAuthDefinition) GetAwsSecretAccessKey() SecretStringDefinition`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *DockerAWSAuthDefinition) GetAwsSecretAccessKeyOk() (*SecretStringDefinition, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *DockerAWSAuthDefinition) SetAwsSecretAccessKey(v SecretStringDefinition)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


