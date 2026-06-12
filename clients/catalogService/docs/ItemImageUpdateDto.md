# ItemImageUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** |  | 
**IsItemMozaicBG** | Pointer to **bool** |  | [optional] 
**MD5Hash** | **string** |  | 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**FileUploadURL** | **string** |  | 
**FileName** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ContentType** | **string** |  | 
**FileLength** | Pointer to **int64** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemImageUpdateDto

`func NewItemImageUpdateDto(itemId string, mD5Hash string, fileUploadURL string, fileName string, contentType string, ) *ItemImageUpdateDto`

NewItemImageUpdateDto instantiates a new ItemImageUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImageUpdateDtoWithDefaults

`func NewItemImageUpdateDtoWithDefaults() *ItemImageUpdateDto`

NewItemImageUpdateDtoWithDefaults instantiates a new ItemImageUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *ItemImageUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemImageUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemImageUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetIsItemMozaicBG

`func (o *ItemImageUpdateDto) GetIsItemMozaicBG() bool`

GetIsItemMozaicBG returns the IsItemMozaicBG field if non-nil, zero value otherwise.

### GetIsItemMozaicBGOk

`func (o *ItemImageUpdateDto) GetIsItemMozaicBGOk() (*bool, bool)`

GetIsItemMozaicBGOk returns a tuple with the IsItemMozaicBG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItemMozaicBG

`func (o *ItemImageUpdateDto) SetIsItemMozaicBG(v bool)`

SetIsItemMozaicBG sets IsItemMozaicBG field to given value.

### HasIsItemMozaicBG

`func (o *ItemImageUpdateDto) HasIsItemMozaicBG() bool`

HasIsItemMozaicBG returns a boolean if a field has been set.

### GetMD5Hash

`func (o *ItemImageUpdateDto) GetMD5Hash() string`

GetMD5Hash returns the MD5Hash field if non-nil, zero value otherwise.

### GetMD5HashOk

`func (o *ItemImageUpdateDto) GetMD5HashOk() (*string, bool)`

GetMD5HashOk returns a tuple with the MD5Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD5Hash

`func (o *ItemImageUpdateDto) SetMD5Hash(v string)`

SetMD5Hash sets MD5Hash field to given value.


### GetMetadata

`func (o *ItemImageUpdateDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemImageUpdateDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemImageUpdateDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemImageUpdateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ItemImageUpdateDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ItemImageUpdateDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileUploadURL

`func (o *ItemImageUpdateDto) GetFileUploadURL() string`

GetFileUploadURL returns the FileUploadURL field if non-nil, zero value otherwise.

### GetFileUploadURLOk

`func (o *ItemImageUpdateDto) GetFileUploadURLOk() (*string, bool)`

GetFileUploadURLOk returns a tuple with the FileUploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadURL

`func (o *ItemImageUpdateDto) SetFileUploadURL(v string)`

SetFileUploadURL sets FileUploadURL field to given value.


### GetFileName

`func (o *ItemImageUpdateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemImageUpdateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemImageUpdateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTitle

`func (o *ItemImageUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemImageUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemImageUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemImageUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemImageUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemImageUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAbstract

`func (o *ItemImageUpdateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemImageUpdateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemImageUpdateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemImageUpdateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemImageUpdateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemImageUpdateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetAuthor

`func (o *ItemImageUpdateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemImageUpdateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemImageUpdateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemImageUpdateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemImageUpdateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemImageUpdateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetKeyWords

`func (o *ItemImageUpdateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemImageUpdateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemImageUpdateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemImageUpdateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemImageUpdateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemImageUpdateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetNotes

`func (o *ItemImageUpdateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemImageUpdateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemImageUpdateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemImageUpdateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemImageUpdateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemImageUpdateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetContentType

`func (o *ItemImageUpdateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ItemImageUpdateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ItemImageUpdateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetFileLength

`func (o *ItemImageUpdateDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ItemImageUpdateDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ItemImageUpdateDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ItemImageUpdateDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetValidResponse

`func (o *ItemImageUpdateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemImageUpdateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemImageUpdateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemImageUpdateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadId

`func (o *ItemImageUpdateDto) GetParentFileUploadId() string`

GetParentFileUploadId returns the ParentFileUploadId field if non-nil, zero value otherwise.

### GetParentFileUploadIdOk

`func (o *ItemImageUpdateDto) GetParentFileUploadIdOk() (*string, bool)`

GetParentFileUploadIdOk returns a tuple with the ParentFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadId

`func (o *ItemImageUpdateDto) SetParentFileUploadId(v string)`

SetParentFileUploadId sets ParentFileUploadId field to given value.

### HasParentFileUploadId

`func (o *ItemImageUpdateDto) HasParentFileUploadId() bool`

HasParentFileUploadId returns a boolean if a field has been set.

### SetParentFileUploadIdNil

`func (o *ItemImageUpdateDto) SetParentFileUploadIdNil(b bool)`

 SetParentFileUploadIdNil sets the value for ParentFileUploadId to be an explicit nil

### UnsetParentFileUploadId
`func (o *ItemImageUpdateDto) UnsetParentFileUploadId()`

UnsetParentFileUploadId ensures that no value is present for ParentFileUploadId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


