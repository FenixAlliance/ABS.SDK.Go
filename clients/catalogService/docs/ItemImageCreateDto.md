# ItemImageCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemID** | Pointer to **NullableString** |  | [optional] 
**IsItemMozaicBG** | Pointer to **bool** |  | [optional] 
**MD5Hash** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**FileUploadURL** | Pointer to **NullableString** |  | [optional] 
**FileName** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentFileUploadID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemImageCreateDto

`func NewItemImageCreateDto(fileName string, ) *ItemImageCreateDto`

NewItemImageCreateDto instantiates a new ItemImageCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImageCreateDtoWithDefaults

`func NewItemImageCreateDtoWithDefaults() *ItemImageCreateDto`

NewItemImageCreateDtoWithDefaults instantiates a new ItemImageCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemImageCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemImageCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemImageCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemImageCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemImageCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemImageCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemImageCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemImageCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemID

`func (o *ItemImageCreateDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemImageCreateDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemImageCreateDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemImageCreateDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemImageCreateDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemImageCreateDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil
### GetIsItemMozaicBG

`func (o *ItemImageCreateDto) GetIsItemMozaicBG() bool`

GetIsItemMozaicBG returns the IsItemMozaicBG field if non-nil, zero value otherwise.

### GetIsItemMozaicBGOk

`func (o *ItemImageCreateDto) GetIsItemMozaicBGOk() (*bool, bool)`

GetIsItemMozaicBGOk returns a tuple with the IsItemMozaicBG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItemMozaicBG

`func (o *ItemImageCreateDto) SetIsItemMozaicBG(v bool)`

SetIsItemMozaicBG sets IsItemMozaicBG field to given value.

### HasIsItemMozaicBG

`func (o *ItemImageCreateDto) HasIsItemMozaicBG() bool`

HasIsItemMozaicBG returns a boolean if a field has been set.

### GetMD5Hash

`func (o *ItemImageCreateDto) GetMD5Hash() string`

GetMD5Hash returns the MD5Hash field if non-nil, zero value otherwise.

### GetMD5HashOk

`func (o *ItemImageCreateDto) GetMD5HashOk() (*string, bool)`

GetMD5HashOk returns a tuple with the MD5Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD5Hash

`func (o *ItemImageCreateDto) SetMD5Hash(v string)`

SetMD5Hash sets MD5Hash field to given value.

### HasMD5Hash

`func (o *ItemImageCreateDto) HasMD5Hash() bool`

HasMD5Hash returns a boolean if a field has been set.

### SetMD5HashNil

`func (o *ItemImageCreateDto) SetMD5HashNil(b bool)`

 SetMD5HashNil sets the value for MD5Hash to be an explicit nil

### UnsetMD5Hash
`func (o *ItemImageCreateDto) UnsetMD5Hash()`

UnsetMD5Hash ensures that no value is present for MD5Hash, not even an explicit nil
### GetMetadata

`func (o *ItemImageCreateDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemImageCreateDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemImageCreateDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemImageCreateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ItemImageCreateDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ItemImageCreateDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileUploadURL

`func (o *ItemImageCreateDto) GetFileUploadURL() string`

GetFileUploadURL returns the FileUploadURL field if non-nil, zero value otherwise.

### GetFileUploadURLOk

`func (o *ItemImageCreateDto) GetFileUploadURLOk() (*string, bool)`

GetFileUploadURLOk returns a tuple with the FileUploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadURL

`func (o *ItemImageCreateDto) SetFileUploadURL(v string)`

SetFileUploadURL sets FileUploadURL field to given value.

### HasFileUploadURL

`func (o *ItemImageCreateDto) HasFileUploadURL() bool`

HasFileUploadURL returns a boolean if a field has been set.

### SetFileUploadURLNil

`func (o *ItemImageCreateDto) SetFileUploadURLNil(b bool)`

 SetFileUploadURLNil sets the value for FileUploadURL to be an explicit nil

### UnsetFileUploadURL
`func (o *ItemImageCreateDto) UnsetFileUploadURL()`

UnsetFileUploadURL ensures that no value is present for FileUploadURL, not even an explicit nil
### GetFileName

`func (o *ItemImageCreateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemImageCreateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemImageCreateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTitle

`func (o *ItemImageCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemImageCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemImageCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemImageCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemImageCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemImageCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAbstract

`func (o *ItemImageCreateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemImageCreateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemImageCreateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemImageCreateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemImageCreateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemImageCreateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetAuthor

`func (o *ItemImageCreateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemImageCreateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemImageCreateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemImageCreateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemImageCreateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemImageCreateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetKeyWords

`func (o *ItemImageCreateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemImageCreateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemImageCreateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemImageCreateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemImageCreateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemImageCreateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetNotes

`func (o *ItemImageCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemImageCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemImageCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemImageCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemImageCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemImageCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetContentType

`func (o *ItemImageCreateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ItemImageCreateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ItemImageCreateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ItemImageCreateDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ItemImageCreateDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ItemImageCreateDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *ItemImageCreateDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ItemImageCreateDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ItemImageCreateDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ItemImageCreateDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetValidResponse

`func (o *ItemImageCreateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemImageCreateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemImageCreateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemImageCreateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetSocialProfileID

`func (o *ItemImageCreateDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemImageCreateDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemImageCreateDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemImageCreateDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemImageCreateDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemImageCreateDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetParentFileUploadID

`func (o *ItemImageCreateDto) GetParentFileUploadID() string`

GetParentFileUploadID returns the ParentFileUploadID field if non-nil, zero value otherwise.

### GetParentFileUploadIDOk

`func (o *ItemImageCreateDto) GetParentFileUploadIDOk() (*string, bool)`

GetParentFileUploadIDOk returns a tuple with the ParentFileUploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadID

`func (o *ItemImageCreateDto) SetParentFileUploadID(v string)`

SetParentFileUploadID sets ParentFileUploadID field to given value.

### HasParentFileUploadID

`func (o *ItemImageCreateDto) HasParentFileUploadID() bool`

HasParentFileUploadID returns a boolean if a field has been set.

### SetParentFileUploadIDNil

`func (o *ItemImageCreateDto) SetParentFileUploadIDNil(b bool)`

 SetParentFileUploadIDNil sets the value for ParentFileUploadID to be an explicit nil

### UnsetParentFileUploadID
`func (o *ItemImageCreateDto) UnsetParentFileUploadID()`

UnsetParentFileUploadID ensures that no value is present for ParentFileUploadID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


