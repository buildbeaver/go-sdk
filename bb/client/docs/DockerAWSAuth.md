# DockerAWSAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRegion** | Pointer to **string** | The ECR AWS region to authenticate to. | [optional] 
**AwsAccessKeyId** | [**SecretString**](SecretString.md) |  | 
**AwsSecretAccessKey** | [**SecretString**](SecretString.md) |  | 

## Methods

### NewDockerAWSAuth

`func NewDockerAWSAuth(awsAccessKeyId SecretString, awsSecretAccessKey SecretString, ) *DockerAWSAuth`

NewDockerAWSAuth instantiates a new DockerAWSAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerAWSAuthWithDefaults

`func NewDockerAWSAuthWithDefaults() *DockerAWSAuth`

NewDockerAWSAuthWithDefaults instantiates a new DockerAWSAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRegion

`func (o *DockerAWSAuth) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *DockerAWSAuth) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *DockerAWSAuth) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *DockerAWSAuth) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *DockerAWSAuth) GetAwsAccessKeyId() SecretString`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *DockerAWSAuth) GetAwsAccessKeyIdOk() (*SecretString, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *DockerAWSAuth) SetAwsAccessKeyId(v SecretString)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### GetAwsSecretAccessKey

`func (o *DockerAWSAuth) GetAwsSecretAccessKey() SecretString`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *DockerAWSAuth) GetAwsSecretAccessKeyOk() (*SecretString, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *DockerAWSAuth) SetAwsSecretAccessKey(v SecretString)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


