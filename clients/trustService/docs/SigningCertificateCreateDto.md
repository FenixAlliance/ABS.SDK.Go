# SigningCertificateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Type** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 
**ContactId** | **string** |  | 
**SecurityCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningCertificateCreateDto

`func NewSigningCertificateCreateDto(title string, contactId string, ) *SigningCertificateCreateDto`

NewSigningCertificateCreateDto instantiates a new SigningCertificateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningCertificateCreateDtoWithDefaults

`func NewSigningCertificateCreateDtoWithDefaults() *SigningCertificateCreateDto`

NewSigningCertificateCreateDtoWithDefaults instantiates a new SigningCertificateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningCertificateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningCertificateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningCertificateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningCertificateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SigningCertificateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningCertificateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningCertificateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningCertificateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SigningCertificateCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SigningCertificateCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SigningCertificateCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *SigningCertificateCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningCertificateCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningCertificateCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningCertificateCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningCertificateCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningCertificateCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *SigningCertificateCreateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SigningCertificateCreateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SigningCertificateCreateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SigningCertificateCreateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SigningCertificateCreateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SigningCertificateCreateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCsr

`func (o *SigningCertificateCreateDto) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *SigningCertificateCreateDto) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *SigningCertificateCreateDto) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *SigningCertificateCreateDto) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *SigningCertificateCreateDto) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *SigningCertificateCreateDto) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetPublicKey

`func (o *SigningCertificateCreateDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SigningCertificateCreateDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SigningCertificateCreateDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SigningCertificateCreateDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *SigningCertificateCreateDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *SigningCertificateCreateDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetCertificateType

`func (o *SigningCertificateCreateDto) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *SigningCertificateCreateDto) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *SigningCertificateCreateDto) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *SigningCertificateCreateDto) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetContactId

`func (o *SigningCertificateCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningCertificateCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningCertificateCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetSecurityCertificateId

`func (o *SigningCertificateCreateDto) GetSecurityCertificateId() string`

GetSecurityCertificateId returns the SecurityCertificateId field if non-nil, zero value otherwise.

### GetSecurityCertificateIdOk

`func (o *SigningCertificateCreateDto) GetSecurityCertificateIdOk() (*string, bool)`

GetSecurityCertificateIdOk returns a tuple with the SecurityCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateId

`func (o *SigningCertificateCreateDto) SetSecurityCertificateId(v string)`

SetSecurityCertificateId sets SecurityCertificateId field to given value.

### HasSecurityCertificateId

`func (o *SigningCertificateCreateDto) HasSecurityCertificateId() bool`

HasSecurityCertificateId returns a boolean if a field has been set.

### SetSecurityCertificateIdNil

`func (o *SigningCertificateCreateDto) SetSecurityCertificateIdNil(b bool)`

 SetSecurityCertificateIdNil sets the value for SecurityCertificateId to be an explicit nil

### UnsetSecurityCertificateId
`func (o *SigningCertificateCreateDto) UnsetSecurityCertificateId()`

UnsetSecurityCertificateId ensures that no value is present for SecurityCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


