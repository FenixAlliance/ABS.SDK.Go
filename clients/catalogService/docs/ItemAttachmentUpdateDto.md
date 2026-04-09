# ItemAttachmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadID** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemAttachmentUpdateDto

`func NewItemAttachmentUpdateDto() *ItemAttachmentUpdateDto`

NewItemAttachmentUpdateDto instantiates a new ItemAttachmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAttachmentUpdateDtoWithDefaults

`func NewItemAttachmentUpdateDtoWithDefaults() *ItemAttachmentUpdateDto`

NewItemAttachmentUpdateDtoWithDefaults instantiates a new ItemAttachmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *ItemAttachmentUpdateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemAttachmentUpdateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemAttachmentUpdateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemAttachmentUpdateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemAttachmentUpdateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemAttachmentUpdateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetMetadata

`func (o *ItemAttachmentUpdateDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemAttachmentUpdateDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemAttachmentUpdateDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemAttachmentUpdateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ItemAttachmentUpdateDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ItemAttachmentUpdateDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTitle

`func (o *ItemAttachmentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemAttachmentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemAttachmentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemAttachmentUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemAttachmentUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemAttachmentUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *ItemAttachmentUpdateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemAttachmentUpdateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemAttachmentUpdateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemAttachmentUpdateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemAttachmentUpdateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemAttachmentUpdateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *ItemAttachmentUpdateDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *ItemAttachmentUpdateDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *ItemAttachmentUpdateDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *ItemAttachmentUpdateDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetFileName

`func (o *ItemAttachmentUpdateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemAttachmentUpdateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemAttachmentUpdateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ItemAttachmentUpdateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ItemAttachmentUpdateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ItemAttachmentUpdateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *ItemAttachmentUpdateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemAttachmentUpdateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemAttachmentUpdateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemAttachmentUpdateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemAttachmentUpdateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemAttachmentUpdateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *ItemAttachmentUpdateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemAttachmentUpdateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemAttachmentUpdateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemAttachmentUpdateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemAttachmentUpdateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemAttachmentUpdateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetValidResponse

`func (o *ItemAttachmentUpdateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemAttachmentUpdateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemAttachmentUpdateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemAttachmentUpdateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadID

`func (o *ItemAttachmentUpdateDto) GetParentFileUploadID() string`

GetParentFileUploadID returns the ParentFileUploadID field if non-nil, zero value otherwise.

### GetParentFileUploadIDOk

`func (o *ItemAttachmentUpdateDto) GetParentFileUploadIDOk() (*string, bool)`

GetParentFileUploadIDOk returns a tuple with the ParentFileUploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadID

`func (o *ItemAttachmentUpdateDto) SetParentFileUploadID(v string)`

SetParentFileUploadID sets ParentFileUploadID field to given value.

### HasParentFileUploadID

`func (o *ItemAttachmentUpdateDto) HasParentFileUploadID() bool`

HasParentFileUploadID returns a boolean if a field has been set.

### SetParentFileUploadIDNil

`func (o *ItemAttachmentUpdateDto) SetParentFileUploadIDNil(b bool)`

 SetParentFileUploadIDNil sets the value for ParentFileUploadID to be an explicit nil

### UnsetParentFileUploadID
`func (o *ItemAttachmentUpdateDto) UnsetParentFileUploadID()`

UnsetParentFileUploadID ensures that no value is present for ParentFileUploadID, not even an explicit nil
### GetFilePath

`func (o *ItemAttachmentUpdateDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ItemAttachmentUpdateDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ItemAttachmentUpdateDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ItemAttachmentUpdateDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *ItemAttachmentUpdateDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *ItemAttachmentUpdateDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


