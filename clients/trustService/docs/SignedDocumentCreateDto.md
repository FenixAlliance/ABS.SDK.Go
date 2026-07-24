# SignedDocumentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewSignedDocumentCreateDto

`func NewSignedDocumentCreateDto(title string, contactId string, ) *SignedDocumentCreateDto`

NewSignedDocumentCreateDto instantiates a new SignedDocumentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentCreateDtoWithDefaults

`func NewSignedDocumentCreateDtoWithDefaults() *SignedDocumentCreateDto`

NewSignedDocumentCreateDtoWithDefaults instantiates a new SignedDocumentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignedDocumentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignedDocumentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignedDocumentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignedDocumentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SignedDocumentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignedDocumentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignedDocumentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignedDocumentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUrl

`func (o *SignedDocumentCreateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SignedDocumentCreateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SignedDocumentCreateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SignedDocumentCreateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SignedDocumentCreateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SignedDocumentCreateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetType

`func (o *SignedDocumentCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignedDocumentCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignedDocumentCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignedDocumentCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SignedDocumentCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SignedDocumentCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *SignedDocumentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContentType

`func (o *SignedDocumentCreateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SignedDocumentCreateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SignedDocumentCreateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SignedDocumentCreateDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SignedDocumentCreateDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SignedDocumentCreateDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContactId

`func (o *SignedDocumentCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SignedDocumentCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SignedDocumentCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetDocumentStandard

`func (o *SignedDocumentCreateDto) GetDocumentStandard() string`

GetDocumentStandard returns the DocumentStandard field if non-nil, zero value otherwise.

### GetDocumentStandardOk

`func (o *SignedDocumentCreateDto) GetDocumentStandardOk() (*string, bool)`

GetDocumentStandardOk returns a tuple with the DocumentStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentStandard

`func (o *SignedDocumentCreateDto) SetDocumentStandard(v string)`

SetDocumentStandard sets DocumentStandard field to given value.

### HasDocumentStandard

`func (o *SignedDocumentCreateDto) HasDocumentStandard() bool`

HasDocumentStandard returns a boolean if a field has been set.

### SetDocumentStandardNil

`func (o *SignedDocumentCreateDto) SetDocumentStandardNil(b bool)`

 SetDocumentStandardNil sets the value for DocumentStandard to be an explicit nil

### UnsetDocumentStandard
`func (o *SignedDocumentCreateDto) UnsetDocumentStandard()`

UnsetDocumentStandard ensures that no value is present for DocumentStandard, not even an explicit nil
### GetTrustDocumentType

`func (o *SignedDocumentCreateDto) GetTrustDocumentType() string`

GetTrustDocumentType returns the TrustDocumentType field if non-nil, zero value otherwise.

### GetTrustDocumentTypeOk

`func (o *SignedDocumentCreateDto) GetTrustDocumentTypeOk() (*string, bool)`

GetTrustDocumentTypeOk returns a tuple with the TrustDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustDocumentType

`func (o *SignedDocumentCreateDto) SetTrustDocumentType(v string)`

SetTrustDocumentType sets TrustDocumentType field to given value.

### HasTrustDocumentType

`func (o *SignedDocumentCreateDto) HasTrustDocumentType() bool`

HasTrustDocumentType returns a boolean if a field has been set.

### SetTrustDocumentTypeNil

`func (o *SignedDocumentCreateDto) SetTrustDocumentTypeNil(b bool)`

 SetTrustDocumentTypeNil sets the value for TrustDocumentType to be an explicit nil

### UnsetTrustDocumentType
`func (o *SignedDocumentCreateDto) UnsetTrustDocumentType()`

UnsetTrustDocumentType ensures that no value is present for TrustDocumentType, not even an explicit nil
### GetCorrelationId

`func (o *SignedDocumentCreateDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignedDocumentCreateDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SignedDocumentCreateDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SignedDocumentCreateDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SignedDocumentCreateDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SignedDocumentCreateDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *SignedDocumentCreateDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *SignedDocumentCreateDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *SignedDocumentCreateDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *SignedDocumentCreateDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *SignedDocumentCreateDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *SignedDocumentCreateDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


