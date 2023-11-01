# ExternalResourceID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSystem** | **string** | The name of the external system, e.g. GitHub | 
**ResourceId** | **string** | The resource within the external system, e.g. github_repo.id | 

## Methods

### NewExternalResourceID

`func NewExternalResourceID(externalSystem string, resourceId string, ) *ExternalResourceID`

NewExternalResourceID instantiates a new ExternalResourceID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalResourceIDWithDefaults

`func NewExternalResourceIDWithDefaults() *ExternalResourceID`

NewExternalResourceIDWithDefaults instantiates a new ExternalResourceID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSystem

`func (o *ExternalResourceID) GetExternalSystem() string`

GetExternalSystem returns the ExternalSystem field if non-nil, zero value otherwise.

### GetExternalSystemOk

`func (o *ExternalResourceID) GetExternalSystemOk() (*string, bool)`

GetExternalSystemOk returns a tuple with the ExternalSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSystem

`func (o *ExternalResourceID) SetExternalSystem(v string)`

SetExternalSystem sets ExternalSystem field to given value.


### GetResourceId

`func (o *ExternalResourceID) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExternalResourceID) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExternalResourceID) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


