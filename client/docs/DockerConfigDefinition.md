# DockerConfigDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | The default Docker image to run the job steps in, if the job is of type Docker | 
**Pull** | **string** | Determines if/when the Docker image is pulled during job execution, if the job is of type Docker | 
**BasicAuth** | Pointer to [**DockerBasicAuthDefinition**](DockerBasicAuthDefinition.md) |  | [optional] 
**AwsAuth** | Pointer to [**DockerAWSAuthDefinition**](DockerAWSAuthDefinition.md) |  | [optional] 
**Shell** | Pointer to **string** | Path to the shell to use to run build scripts with inside the container | [optional] 

## Methods

### NewDockerConfigDefinition

`func NewDockerConfigDefinition(image string, pull string, ) *DockerConfigDefinition`

NewDockerConfigDefinition instantiates a new DockerConfigDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerConfigDefinitionWithDefaults

`func NewDockerConfigDefinitionWithDefaults() *DockerConfigDefinition`

NewDockerConfigDefinitionWithDefaults instantiates a new DockerConfigDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *DockerConfigDefinition) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DockerConfigDefinition) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DockerConfigDefinition) SetImage(v string)`

SetImage sets Image field to given value.


### GetPull

`func (o *DockerConfigDefinition) GetPull() string`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *DockerConfigDefinition) GetPullOk() (*string, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *DockerConfigDefinition) SetPull(v string)`

SetPull sets Pull field to given value.


### GetBasicAuth

`func (o *DockerConfigDefinition) GetBasicAuth() DockerBasicAuthDefinition`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *DockerConfigDefinition) GetBasicAuthOk() (*DockerBasicAuthDefinition, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *DockerConfigDefinition) SetBasicAuth(v DockerBasicAuthDefinition)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *DockerConfigDefinition) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetAwsAuth

`func (o *DockerConfigDefinition) GetAwsAuth() DockerAWSAuthDefinition`

GetAwsAuth returns the AwsAuth field if non-nil, zero value otherwise.

### GetAwsAuthOk

`func (o *DockerConfigDefinition) GetAwsAuthOk() (*DockerAWSAuthDefinition, bool)`

GetAwsAuthOk returns a tuple with the AwsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAuth

`func (o *DockerConfigDefinition) SetAwsAuth(v DockerAWSAuthDefinition)`

SetAwsAuth sets AwsAuth field to given value.

### HasAwsAuth

`func (o *DockerConfigDefinition) HasAwsAuth() bool`

HasAwsAuth returns a boolean if a field has been set.

### GetShell

`func (o *DockerConfigDefinition) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *DockerConfigDefinition) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *DockerConfigDefinition) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *DockerConfigDefinition) HasShell() bool`

HasShell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


