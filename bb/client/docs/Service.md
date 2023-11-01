# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the service, within the parent job. | 
**Image** | **string** | Name of the Docker image for the service to run. | 
**BasicAuth** | Pointer to [**DockerBasicAuth**](DockerBasicAuth.md) |  | [optional] 
**AwsAuth** | Pointer to [**DockerAWSAuth**](DockerAWSAuth.md) |  | [optional] 
**Environment** | [**[]EnvVar**](EnvVar.md) | A list of environment variables to export prior to starting the service. | 

## Methods

### NewService

`func NewService(name string, image string, environment []EnvVar, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *Service) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Service) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Service) SetImage(v string)`

SetImage sets Image field to given value.


### GetBasicAuth

`func (o *Service) GetBasicAuth() DockerBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *Service) GetBasicAuthOk() (*DockerBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *Service) SetBasicAuth(v DockerBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *Service) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetAwsAuth

`func (o *Service) GetAwsAuth() DockerAWSAuth`

GetAwsAuth returns the AwsAuth field if non-nil, zero value otherwise.

### GetAwsAuthOk

`func (o *Service) GetAwsAuthOk() (*DockerAWSAuth, bool)`

GetAwsAuthOk returns a tuple with the AwsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAuth

`func (o *Service) SetAwsAuth(v DockerAWSAuth)`

SetAwsAuth sets AwsAuth field to given value.

### HasAwsAuth

`func (o *Service) HasAwsAuth() bool`

HasAwsAuth returns a boolean if a field has been set.

### GetEnvironment

`func (o *Service) GetEnvironment() []EnvVar`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Service) GetEnvironmentOk() (*[]EnvVar, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Service) SetEnvironment(v []EnvVar)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


