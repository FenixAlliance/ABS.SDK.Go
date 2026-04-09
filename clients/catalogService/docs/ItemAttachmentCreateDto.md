# ItemAttachmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadId** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**ItemID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemAttachmentCreateDto

`func NewItemAttachmentCreateDto() *ItemAttachmentCreateDto`

NewItemAttachmentCreateDto instantiates a new ItemAttachmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAttachmentCreateDtoWithDefaults

`func NewItemAttachmentCreateDtoWithDefaults() *ItemAttachmentCreateDto`

NewItemAttachmentCreateDtoWithDefaults instantiates a new ItemAttachmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemAttachmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemAttachmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemAttachmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemAttachmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemAttachmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemAttachmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemAttachmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemAttachmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNotes

`func (o *ItemAttachmentCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemAttachmentCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemAttachmentCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemAttachmentCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemAttachmentCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemAttachmentCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *ItemAttachmentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemAttachmentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemAttachmentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemAttachmentCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemAttachmentCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemAttachmentCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *ItemAttachmentCreateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemAttachmentCreateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemAttachmentCreateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemAttachmentCreateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemAttachmentCreateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemAttachmentCreateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *ItemAttachmentCreateDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *ItemAttachmentCreateDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *ItemAttachmentCreateDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *ItemAttachmentCreateDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetFileName

`func (o *ItemAttachmentCreateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemAttachmentCreateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemAttachmentCreateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ItemAttachmentCreateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ItemAttachmentCreateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ItemAttachmentCreateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *ItemAttachmentCreateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemAttachmentCreateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemAttachmentCreateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemAttachmentCreateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemAttachmentCreateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemAttachmentCreateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *ItemAttachmentCreateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemAttachmentCreateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemAttachmentCreateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemAttachmentCreateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemAttachmentCreateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemAttachmentCreateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetValidResponse

`func (o *ItemAttachmentCreateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemAttachmentCreateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemAttachmentCreateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemAttachmentCreateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadId

`func (o *ItemAttachmentCreateDto) GetParentFileUploadId() string`

GetParentFileUploadId returns the ParentFileUploadId field if non-nil, zero value otherwise.

### GetParentFileUploadIdOk

`func (o *ItemAttachmentCreateDto) GetParentFileUploadIdOk() (*string, bool)`

GetParentFileUploadIdOk returns a tuple with the ParentFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadId

`func (o *ItemAttachmentCreateDto) SetParentFileUploadId(v string)`

SetParentFileUploadId sets ParentFileUploadId field to given value.

### HasParentFileUploadId

`func (o *ItemAttachmentCreateDto) HasParentFileUploadId() bool`

HasParentFileUploadId returns a boolean if a field has been set.

### SetParentFileUploadIdNil

`func (o *ItemAttachmentCreateDto) SetParentFileUploadIdNil(b bool)`

 SetParentFileUploadIdNil sets the value for ParentFileUploadId to be an explicit nil

### UnsetParentFileUploadId
`func (o *ItemAttachmentCreateDto) UnsetParentFileUploadId()`

UnsetParentFileUploadId ensures that no value is present for ParentFileUploadId, not even an explicit nil
### GetFilePath

`func (o *ItemAttachmentCreateDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ItemAttachmentCreateDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ItemAttachmentCreateDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ItemAttachmentCreateDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *ItemAttachmentCreateDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *ItemAttachmentCreateDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetItemID

`func (o *ItemAttachmentCreateDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemAttachmentCreateDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemAttachmentCreateDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemAttachmentCreateDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemAttachmentCreateDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemAttachmentCreateDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


