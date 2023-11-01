# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**RepoId** | **string** | ID of the repo that the commit was made against. | 
**ConfigType** | **string** | The type of the config file found in the commit (Config itself is not included in the document). | 
**Sha** | **string** | The unique SHA hash of the commit. | 
**Message** | **string** | The commit message. | 
**AuthorId** | **string** | The ID of the legal entity that authored the commit, if known. | 
**AuthorName** | **string** | The author name recorded on the commit. | 
**AuthorEmail** | **string** | The author email address recorded on the commit. | 
**CommitterId** | **string** | The ID of the legal entity that committed the commit, if known. | 
**CommitterName** | **string** | The committer name recorded on the commit, if any. | 
**CommitterEmail** | **string** | The committer email address recorded on the commit, if any. | 
**Link** | **string** | A link (URL) to the Commit in the SCM, for users to browse to. | 
**CommitterUrl** | **string** | URL to fetch additional information about the committer of this commit. | 
**AuthorUrl** | **string** | URL to fetch additional information about the author of this commit. | 

## Methods

### NewCommit

`func NewCommit(id string, createdAt time.Time, repoId string, configType string, sha string, message string, authorId string, authorName string, authorEmail string, committerId string, committerName string, committerEmail string, link string, committerUrl string, authorUrl string, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Commit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Commit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Commit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRepoId

`func (o *Commit) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *Commit) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *Commit) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetConfigType

`func (o *Commit) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *Commit) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *Commit) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetSha

`func (o *Commit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Commit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Commit) SetSha(v string)`

SetSha sets Sha field to given value.


### GetMessage

`func (o *Commit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthorId

`func (o *Commit) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Commit) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Commit) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *Commit) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Commit) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Commit) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetAuthorEmail

`func (o *Commit) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *Commit) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *Commit) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.


### GetCommitterId

`func (o *Commit) GetCommitterId() string`

GetCommitterId returns the CommitterId field if non-nil, zero value otherwise.

### GetCommitterIdOk

`func (o *Commit) GetCommitterIdOk() (*string, bool)`

GetCommitterIdOk returns a tuple with the CommitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterId

`func (o *Commit) SetCommitterId(v string)`

SetCommitterId sets CommitterId field to given value.


### GetCommitterName

`func (o *Commit) GetCommitterName() string`

GetCommitterName returns the CommitterName field if non-nil, zero value otherwise.

### GetCommitterNameOk

`func (o *Commit) GetCommitterNameOk() (*string, bool)`

GetCommitterNameOk returns a tuple with the CommitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterName

`func (o *Commit) SetCommitterName(v string)`

SetCommitterName sets CommitterName field to given value.


### GetCommitterEmail

`func (o *Commit) GetCommitterEmail() string`

GetCommitterEmail returns the CommitterEmail field if non-nil, zero value otherwise.

### GetCommitterEmailOk

`func (o *Commit) GetCommitterEmailOk() (*string, bool)`

GetCommitterEmailOk returns a tuple with the CommitterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterEmail

`func (o *Commit) SetCommitterEmail(v string)`

SetCommitterEmail sets CommitterEmail field to given value.


### GetLink

`func (o *Commit) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Commit) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Commit) SetLink(v string)`

SetLink sets Link field to given value.


### GetCommitterUrl

`func (o *Commit) GetCommitterUrl() string`

GetCommitterUrl returns the CommitterUrl field if non-nil, zero value otherwise.

### GetCommitterUrlOk

`func (o *Commit) GetCommitterUrlOk() (*string, bool)`

GetCommitterUrlOk returns a tuple with the CommitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterUrl

`func (o *Commit) SetCommitterUrl(v string)`

SetCommitterUrl sets CommitterUrl field to given value.


### GetAuthorUrl

`func (o *Commit) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *Commit) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *Commit) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


