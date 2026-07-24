# SigningCertificateUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Type** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 
**ContactId** | **string** |  | 
**SecurityCertificateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningCertificateUpdateDto

`func NewSigningCertificateUpdateDto(title string, contactId string, ) *SigningCertificateUpdateDto`

NewSigningCertificateUpdateDto instantiates a new SigningCertificateUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningCertificateUpdateDtoWithDefaults

`func NewSigningCertificateUpdateDtoWithDefaults() *SigningCertificateUpdateDto`

NewSigningCertificateUpdateDtoWithDefaults instantiates a new SigningCertificateUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SigningCertificateUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SigningCertificateUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SigningCertificateUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *SigningCertificateUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningCertificateUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningCertificateUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningCertificateUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningCertificateUpdateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningCertificateUpdateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *SigningCertificateUpdateDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SigningCertificateUpdateDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SigningCertificateUpdateDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SigningCertificateUpdateDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SigningCertificateUpdateDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SigningCertificateUpdateDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCsr

`func (o *SigningCertificateUpdateDto) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *SigningCertificateUpdateDto) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *SigningCertificateUpdateDto) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *SigningCertificateUpdateDto) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *SigningCertificateUpdateDto) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *SigningCertificateUpdateDto) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetPublicKey

`func (o *SigningCertificateUpdateDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SigningCertificateUpdateDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SigningCertificateUpdateDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SigningCertificateUpdateDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *SigningCertificateUpdateDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *SigningCertificateUpdateDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetCertificateType

`func (o *SigningCertificateUpdateDto) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *SigningCertificateUpdateDto) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *SigningCertificateUpdateDto) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *SigningCertificateUpdateDto) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetContactId

`func (o *SigningCertificateUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningCertificateUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningCertificateUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetSecurityCertificateId

`func (o *SigningCertificateUpdateDto) GetSecurityCertificateId() string`

GetSecurityCertificateId returns the SecurityCertificateId field if non-nil, zero value otherwise.

### GetSecurityCertificateIdOk

`func (o *SigningCertificateUpdateDto) GetSecurityCertificateIdOk() (*string, bool)`

GetSecurityCertificateIdOk returns a tuple with the SecurityCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCertificateId

`func (o *SigningCertificateUpdateDto) SetSecurityCertificateId(v string)`

SetSecurityCertificateId sets SecurityCertificateId field to given value.

### HasSecurityCertificateId

`func (o *SigningCertificateUpdateDto) HasSecurityCertificateId() bool`

HasSecurityCertificateId returns a boolean if a field has been set.

### SetSecurityCertificateIdNil

`func (o *SigningCertificateUpdateDto) SetSecurityCertificateIdNil(b bool)`

 SetSecurityCertificateIdNil sets the value for SecurityCertificateId to be an explicit nil

### UnsetSecurityCertificateId
`func (o *SigningCertificateUpdateDto) UnsetSecurityCertificateId()`

UnsetSecurityCertificateId ensures that no value is present for SecurityCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


