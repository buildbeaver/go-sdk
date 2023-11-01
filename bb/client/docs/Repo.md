# Repo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A link to the Repo resource on the BuildBeaver server. | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Etag** | **string** |  | 
**Name** | **string** | Name of the repo. | 
**Description** | **string** | A description of the repo, taken from the Source Control Management system (e.g. GitHub). | 
**LegalEntityId** | **string** | The ID of the Legal Entity (person or company) that owns this repo. | 
**SshUrl** | **string** | The URL of this repo on the Source Control Management system (e.g. GitHub) for SSH access to the repo. | 
**HttpUrl** | **string** | The URL of this repo on the Source Control Management system (e.g. GitHub) for HTTP access to the repo. | 
**Link** | **string** | A link (URL) to the Repo on the Web, for users to browse to. | 
**DefaultBranch** | **string** | The name of the default branch on this repo (typically &#39;main&#39; or &#39;master&#39;) | 
**Private** | **bool** | True if this is a private repo, false if it is public. | 
**Enabled** | **bool** | True if builds are enabled for this repo. | 
**SshKeySecretId** | Pointer to **string** | The ID of the secret containing the SSH key to use to check out code from the repo. | [optional] 
**ExternalId** | Pointer to [**ExternalResourceID**](ExternalResourceID.md) |  | [optional] 
**ExternalMetadata** | **string** | Extra information relating to the repo in the Source Control Management system (e.g. GitHub). The exact information stored here will depend on which SCM contains the repo. | 
**BuildsUrl** | **string** | URL to fetch a list of builds for this repo. | 
**BuildSearchUrl** | **string** | URL to search for specific builds for this repo. | 
**SecretsUrl** | **string** | URL to fetch secrets for this repo (subject to access control). | 

## Methods

### NewRepo

`func NewRepo(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, description string, legalEntityId string, sshUrl string, httpUrl string, link string, defaultBranch string, private bool, enabled bool, externalMetadata string, buildsUrl string, buildSearchUrl string, secretsUrl string, ) *Repo`

NewRepo instantiates a new Repo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoWithDefaults

`func NewRepoWithDefaults() *Repo`

NewRepoWithDefaults instantiates a new Repo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Repo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Repo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repo) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Repo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Repo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Repo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Repo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Repo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Repo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Repo) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Repo) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Repo) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Repo) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEtag

`func (o *Repo) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *Repo) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *Repo) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetName

`func (o *Repo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Repo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Repo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Repo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLegalEntityId

`func (o *Repo) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *Repo) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *Repo) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetSshUrl

`func (o *Repo) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *Repo) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *Repo) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetHttpUrl

`func (o *Repo) GetHttpUrl() string`

GetHttpUrl returns the HttpUrl field if non-nil, zero value otherwise.

### GetHttpUrlOk

`func (o *Repo) GetHttpUrlOk() (*string, bool)`

GetHttpUrlOk returns a tuple with the HttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUrl

`func (o *Repo) SetHttpUrl(v string)`

SetHttpUrl sets HttpUrl field to given value.


### GetLink

`func (o *Repo) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Repo) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Repo) SetLink(v string)`

SetLink sets Link field to given value.


### GetDefaultBranch

`func (o *Repo) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *Repo) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *Repo) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetPrivate

`func (o *Repo) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Repo) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Repo) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetEnabled

`func (o *Repo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Repo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Repo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSshKeySecretId

`func (o *Repo) GetSshKeySecretId() string`

GetSshKeySecretId returns the SshKeySecretId field if non-nil, zero value otherwise.

### GetSshKeySecretIdOk

`func (o *Repo) GetSshKeySecretIdOk() (*string, bool)`

GetSshKeySecretIdOk returns a tuple with the SshKeySecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeySecretId

`func (o *Repo) SetSshKeySecretId(v string)`

SetSshKeySecretId sets SshKeySecretId field to given value.

### HasSshKeySecretId

`func (o *Repo) HasSshKeySecretId() bool`

HasSshKeySecretId returns a boolean if a field has been set.

### GetExternalId

`func (o *Repo) GetExternalId() ExternalResourceID`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Repo) GetExternalIdOk() (*ExternalResourceID, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Repo) SetExternalId(v ExternalResourceID)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Repo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalMetadata

`func (o *Repo) GetExternalMetadata() string`

GetExternalMetadata returns the ExternalMetadata field if non-nil, zero value otherwise.

### GetExternalMetadataOk

`func (o *Repo) GetExternalMetadataOk() (*string, bool)`

GetExternalMetadataOk returns a tuple with the ExternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMetadata

`func (o *Repo) SetExternalMetadata(v string)`

SetExternalMetadata sets ExternalMetadata field to given value.


### GetBuildsUrl

`func (o *Repo) GetBuildsUrl() string`

GetBuildsUrl returns the BuildsUrl field if non-nil, zero value otherwise.

### GetBuildsUrlOk

`func (o *Repo) GetBuildsUrlOk() (*string, bool)`

GetBuildsUrlOk returns a tuple with the BuildsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildsUrl

`func (o *Repo) SetBuildsUrl(v string)`

SetBuildsUrl sets BuildsUrl field to given value.


### GetBuildSearchUrl

`func (o *Repo) GetBuildSearchUrl() string`

GetBuildSearchUrl returns the BuildSearchUrl field if non-nil, zero value otherwise.

### GetBuildSearchUrlOk

`func (o *Repo) GetBuildSearchUrlOk() (*string, bool)`

GetBuildSearchUrlOk returns a tuple with the BuildSearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSearchUrl

`func (o *Repo) SetBuildSearchUrl(v string)`

SetBuildSearchUrl sets BuildSearchUrl field to given value.


### GetSecretsUrl

`func (o *Repo) GetSecretsUrl() string`

GetSecretsUrl returns the SecretsUrl field if non-nil, zero value otherwise.

### GetSecretsUrlOk

`func (o *Repo) GetSecretsUrlOk() (*string, bool)`

GetSecretsUrlOk returns a tuple with the SecretsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsUrl

`func (o *Repo) SetSecretsUrl(v string)`

SetSecretsUrl sets SecretsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


