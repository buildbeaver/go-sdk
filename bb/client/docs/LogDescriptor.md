# LogDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Log Descriptor on the BuildBeaver server | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Etag** | **string** |  | 
**ParentLogId** | **string** | The ID of the log descriptor that this one is nested within. | 
**ResourceId** | **string** | The ID of the resource that this log belongs to. | 
**Sealed** | **bool** | Sealed is set to true when the log is completed and has become immutable. | 
**SizeBytes** | **int64** | The size of the log, in bytes, calculated and set at the time the log is sealed. | 
**DataUrl** | **string** | URL to use for fetching the log data. | 

## Methods

### NewLogDescriptor

`func NewLogDescriptor(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, parentLogId string, resourceId string, sealed bool, sizeBytes int64, dataUrl string, ) *LogDescriptor`

NewLogDescriptor instantiates a new LogDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDescriptorWithDefaults

`func NewLogDescriptorWithDefaults() *LogDescriptor`

NewLogDescriptorWithDefaults instantiates a new LogDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LogDescriptor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogDescriptor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogDescriptor) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *LogDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogDescriptor) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *LogDescriptor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogDescriptor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogDescriptor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogDescriptor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogDescriptor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogDescriptor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEtag

`func (o *LogDescriptor) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *LogDescriptor) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *LogDescriptor) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetParentLogId

`func (o *LogDescriptor) GetParentLogId() string`

GetParentLogId returns the ParentLogId field if non-nil, zero value otherwise.

### GetParentLogIdOk

`func (o *LogDescriptor) GetParentLogIdOk() (*string, bool)`

GetParentLogIdOk returns a tuple with the ParentLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogId

`func (o *LogDescriptor) SetParentLogId(v string)`

SetParentLogId sets ParentLogId field to given value.


### GetResourceId

`func (o *LogDescriptor) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *LogDescriptor) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *LogDescriptor) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetSealed

`func (o *LogDescriptor) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *LogDescriptor) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *LogDescriptor) SetSealed(v bool)`

SetSealed sets Sealed field to given value.


### GetSizeBytes

`func (o *LogDescriptor) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *LogDescriptor) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *LogDescriptor) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.


### GetDataUrl

`func (o *LogDescriptor) GetDataUrl() string`

GetDataUrl returns the DataUrl field if non-nil, zero value otherwise.

### GetDataUrlOk

`func (o *LogDescriptor) GetDataUrlOk() (*string, bool)`

GetDataUrlOk returns a tuple with the DataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUrl

`func (o *LogDescriptor) SetDataUrl(v string)`

SetDataUrl sets DataUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


