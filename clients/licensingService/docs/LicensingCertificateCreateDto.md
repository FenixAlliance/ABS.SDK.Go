# LicensingCertificateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ContactId** | **string** |  | 
**Csr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicensingCertificateCreateDto

`func NewLicensingCertificateCreateDto(contactId string, ) *LicensingCertificateCreateDto`

NewLicensingCertificateCreateDto instantiates a new LicensingCertificateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingCertificateCreateDtoWithDefaults

`func NewLicensingCertificateCreateDtoWithDefaults() *LicensingCertificateCreateDto`

NewLicensingCertificateCreateDtoWithDefaults instantiates a new LicensingCertificateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicensingCertificateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensingCertificateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensingCertificateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicensingCertificateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LicensingCertificateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicensingCertificateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicensingCertificateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicensingCertificateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *LicensingCertificateCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicensingCertificateCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicensingCertificateCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LicensingCertificateCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LicensingCertificateCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LicensingCertificateCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContactId

`func (o *LicensingCertificateCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *LicensingCertificateCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *LicensingCertificateCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetCsr

`func (o *LicensingCertificateCreateDto) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *LicensingCertificateCreateDto) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *LicensingCertificateCreateDto) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *LicensingCertificateCreateDto) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *LicensingCertificateCreateDto) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *LicensingCertificateCreateDto) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


