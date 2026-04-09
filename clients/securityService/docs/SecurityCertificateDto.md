# SecurityCertificateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ExpirePeriod** | Pointer to **NullableString** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityCertificateDto

`func NewSecurityCertificateDto() *SecurityCertificateDto`

NewSecurityCertificateDto instantiates a new SecurityCertificateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCertificateDtoWithDefaults

`func NewSecurityCertificateDtoWithDefaults() *SecurityCertificateDto`

NewSecurityCertificateDtoWithDefaults instantiates a new SecurityCertificateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityCertificateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityCertificateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityCertificateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityCertificateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SecurityCertificateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SecurityCertificateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SecurityCertificateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SecurityCertificateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SecurityCertificateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SecurityCertificateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SecurityCertificateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SecurityCertificateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetType

`func (o *SecurityCertificateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityCertificateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityCertificateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityCertificateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SecurityCertificateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SecurityCertificateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpirePeriod

`func (o *SecurityCertificateDto) GetExpirePeriod() string`

GetExpirePeriod returns the ExpirePeriod field if non-nil, zero value otherwise.

### GetExpirePeriodOk

`func (o *SecurityCertificateDto) GetExpirePeriodOk() (*string, bool)`

GetExpirePeriodOk returns a tuple with the ExpirePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePeriod

`func (o *SecurityCertificateDto) SetExpirePeriod(v string)`

SetExpirePeriod sets ExpirePeriod field to given value.

### HasExpirePeriod

`func (o *SecurityCertificateDto) HasExpirePeriod() bool`

HasExpirePeriod returns a boolean if a field has been set.

### SetExpirePeriodNil

`func (o *SecurityCertificateDto) SetExpirePeriodNil(b bool)`

 SetExpirePeriodNil sets the value for ExpirePeriod to be an explicit nil

### UnsetExpirePeriod
`func (o *SecurityCertificateDto) UnsetExpirePeriod()`

UnsetExpirePeriod ensures that no value is present for ExpirePeriod, not even an explicit nil
### GetExpired

`func (o *SecurityCertificateDto) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SecurityCertificateDto) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SecurityCertificateDto) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *SecurityCertificateDto) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetDisabled

`func (o *SecurityCertificateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SecurityCertificateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SecurityCertificateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SecurityCertificateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetBusinessID

`func (o *SecurityCertificateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SecurityCertificateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SecurityCertificateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SecurityCertificateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SecurityCertificateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SecurityCertificateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetCsr

`func (o *SecurityCertificateDto) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *SecurityCertificateDto) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *SecurityCertificateDto) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *SecurityCertificateDto) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *SecurityCertificateDto) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *SecurityCertificateDto) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


