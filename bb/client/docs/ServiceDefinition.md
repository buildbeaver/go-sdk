# ServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the service, within the parent job | 
**Image** | **string** | Name of the Docker image for the service to run | 
**BasicAuth** | Pointer to [**DockerBasicAuthDefinition**](DockerBasicAuthDefinition.md) |  | [optional] 
**AwsAuth** | Pointer to [**DockerAWSAuthDefinition**](DockerAWSAuthDefinition.md) |  | [optional] 
**Environment** | [**map[string]SecretStringDefinition**](SecretStringDefinition.md) | A list of environment variables to export prior to executing the service | 

## Methods

### NewServiceDefinition

`func NewServiceDefinition(name string, image string, environment map[string]SecretStringDefinition, ) *ServiceDefinition`

NewServiceDefinition instantiates a new ServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDefinitionWithDefaults

`func NewServiceDefinitionWithDefaults() *ServiceDefinition`

NewServiceDefinitionWithDefaults instantiates a new ServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *ServiceDefinition) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServiceDefinition) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServiceDefinition) SetImage(v string)`

SetImage sets Image field to given value.


### GetBasicAuth

`func (o *ServiceDefinition) GetBasicAuth() DockerBasicAuthDefinition`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *ServiceDefinition) GetBasicAuthOk() (*DockerBasicAuthDefinition, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *ServiceDefinition) SetBasicAuth(v DockerBasicAuthDefinition)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *ServiceDefinition) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetAwsAuth

`func (o *ServiceDefinition) GetAwsAuth() DockerAWSAuthDefinition`

GetAwsAuth returns the AwsAuth field if non-nil, zero value otherwise.

### GetAwsAuthOk

`func (o *ServiceDefinition) GetAwsAuthOk() (*DockerAWSAuthDefinition, bool)`

GetAwsAuthOk returns a tuple with the AwsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAuth

`func (o *ServiceDefinition) SetAwsAuth(v DockerAWSAuthDefinition)`

SetAwsAuth sets AwsAuth field to given value.

### HasAwsAuth

`func (o *ServiceDefinition) HasAwsAuth() bool`

HasAwsAuth returns a boolean if a field has been set.

### GetEnvironment

`func (o *ServiceDefinition) GetEnvironment() map[string]SecretStringDefinition`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceDefinition) GetEnvironmentOk() (*map[string]SecretStringDefinition, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceDefinition) SetEnvironment(v map[string]SecretStringDefinition)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


