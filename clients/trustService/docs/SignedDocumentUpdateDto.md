# SignedDocumentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ContactId** | **string** |  | 
**DocumentStandard** | Pointer to **NullableString** |  | [optional] 
**TrustDocumentType** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignedDocumentUpdateDto

`func NewSignedDocumentUpdateDto(title string, contactId string, ) *SignedDocumentUpdateDto`

NewSignedDocumentUpdateDto instantiates a new SignedDocumentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentUpdateDtoWithDefaults

`func NewSignedDocumentUpdateDtoWithDefaults() *SignedDocumentUpdateDto`

NewSignedDocumentUpdateDtoWithDefaults instantiates a new SignedDocumentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SignedDocumentUpdateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SignedDocumentUpdateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SignedDocumentUpdateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SignedDocumentUpdateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SignedDocumentUpdateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SignedDocumentUpdateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetType

`func (o *SignedDocumentUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignedDocumentUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignedDocumentUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignedDocumentUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SignedDocumentUpdateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SignedDocumentUpdateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SignedDocumentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContentType

`func (o *SignedDocumentUpdateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SignedDocumentUpdateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SignedDocumentUpdateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SignedDocumentUpdateDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SignedDocumentUpdateDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SignedDocumentUpdateDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContactId

`func (o *SignedDocumentUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SignedDocumentUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SignedDocumentUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetDocumentStandard

`func (o *SignedDocumentUpdateDto) GetDocumentStandard() string`

GetDocumentStandard returns the DocumentStandard field if non-nil, zero value otherwise.

### GetDocumentStandardOk

`func (o *SignedDocumentUpdateDto) GetDocumentStandardOk() (*string, bool)`

GetDocumentStandardOk returns a tuple with the DocumentStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentStandard

`func (o *SignedDocumentUpdateDto) SetDocumentStandard(v string)`

SetDocumentStandard sets DocumentStandard field to given value.

### HasDocumentStandard

`func (o *SignedDocumentUpdateDto) HasDocumentStandard() bool`

HasDocumentStandard returns a boolean if a field has been set.

### SetDocumentStandardNil

`func (o *SignedDocumentUpdateDto) SetDocumentStandardNil(b bool)`

 SetDocumentStandardNil sets the value for DocumentStandard to be an explicit nil

### UnsetDocumentStandard
`func (o *SignedDocumentUpdateDto) UnsetDocumentStandard()`

UnsetDocumentStandard ensures that no value is present for DocumentStandard, not even an explicit nil
### GetTrustDocumentType

`func (o *SignedDocumentUpdateDto) GetTrustDocumentType() string`

GetTrustDocumentType returns the TrustDocumentType field if non-nil, zero value otherwise.

### GetTrustDocumentTypeOk

`func (o *SignedDocumentUpdateDto) GetTrustDocumentTypeOk() (*string, bool)`

GetTrustDocumentTypeOk returns a tuple with the TrustDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustDocumentType

`func (o *SignedDocumentUpdateDto) SetTrustDocumentType(v string)`

SetTrustDocumentType sets TrustDocumentType field to given value.

### HasTrustDocumentType

`func (o *SignedDocumentUpdateDto) HasTrustDocumentType() bool`

HasTrustDocumentType returns a boolean if a field has been set.

### SetTrustDocumentTypeNil

`func (o *SignedDocumentUpdateDto) SetTrustDocumentTypeNil(b bool)`

 SetTrustDocumentTypeNil sets the value for TrustDocumentType to be an explicit nil

### UnsetTrustDocumentType
`func (o *SignedDocumentUpdateDto) UnsetTrustDocumentType()`

UnsetTrustDocumentType ensures that no value is present for TrustDocumentType, not even an explicit nil
### GetCorrelationId

`func (o *SignedDocumentUpdateDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignedDocumentUpdateDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SignedDocumentUpdateDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SignedDocumentUpdateDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SignedDocumentUpdateDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SignedDocumentUpdateDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *SignedDocumentUpdateDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *SignedDocumentUpdateDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *SignedDocumentUpdateDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *SignedDocumentUpdateDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *SignedDocumentUpdateDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *SignedDocumentUpdateDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


