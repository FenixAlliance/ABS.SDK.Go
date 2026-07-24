# QuickSignSignedDocumentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** |  | 
**SigningCertificateId** | **string** |  | 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQuickSignSignedDocumentDto

`func NewQuickSignSignedDocumentDto(providerName string, signingCertificateId string, ) *QuickSignSignedDocumentDto`

NewQuickSignSignedDocumentDto instantiates a new QuickSignSignedDocumentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickSignSignedDocumentDtoWithDefaults

`func NewQuickSignSignedDocumentDtoWithDefaults() *QuickSignSignedDocumentDto`

NewQuickSignSignedDocumentDtoWithDefaults instantiates a new QuickSignSignedDocumentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *QuickSignSignedDocumentDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *QuickSignSignedDocumentDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *QuickSignSignedDocumentDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSigningCertificateId

`func (o *QuickSignSignedDocumentDto) GetSigningCertificateId() string`

GetSigningCertificateId returns the SigningCertificateId field if non-nil, zero value otherwise.

### GetSigningCertificateIdOk

`func (o *QuickSignSignedDocumentDto) GetSigningCertificateIdOk() (*string, bool)`

GetSigningCertificateIdOk returns a tuple with the SigningCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificateId

`func (o *QuickSignSignedDocumentDto) SetSigningCertificateId(v string)`

SetSigningCertificateId sets SigningCertificateId field to given value.


### GetSigningProfileId

`func (o *QuickSignSignedDocumentDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *QuickSignSignedDocumentDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *QuickSignSignedDocumentDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *QuickSignSignedDocumentDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *QuickSignSignedDocumentDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *QuickSignSignedDocumentDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


