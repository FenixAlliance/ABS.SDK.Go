# TruckDriverDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**NationalIdNumber** | Pointer to **NullableString** |  | [optional] 
**LicenseNumber** | Pointer to **NullableString** |  | [optional] 
**LicenseClass** | Pointer to **NullableString** |  | [optional] 
**LicenseExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**AdrCertified** | Pointer to **bool** |  | [optional] 
**AdrCertificateExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**MedicalExamExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckDriverDto

`func NewTruckDriverDto() *TruckDriverDto`

NewTruckDriverDto instantiates a new TruckDriverDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckDriverDtoWithDefaults

`func NewTruckDriverDtoWithDefaults() *TruckDriverDto`

NewTruckDriverDtoWithDefaults instantiates a new TruckDriverDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckDriverDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckDriverDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckDriverDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckDriverDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TruckDriverDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TruckDriverDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TruckDriverDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckDriverDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckDriverDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckDriverDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TruckDriverDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TruckDriverDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *TruckDriverDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TruckDriverDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TruckDriverDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TruckDriverDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TruckDriverDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TruckDriverDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *TruckDriverDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TruckDriverDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TruckDriverDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TruckDriverDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *TruckDriverDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *TruckDriverDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *TruckDriverDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TruckDriverDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TruckDriverDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TruckDriverDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TruckDriverDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TruckDriverDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNationalIdNumber

`func (o *TruckDriverDto) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *TruckDriverDto) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *TruckDriverDto) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *TruckDriverDto) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *TruckDriverDto) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *TruckDriverDto) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetLicenseNumber

`func (o *TruckDriverDto) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *TruckDriverDto) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *TruckDriverDto) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *TruckDriverDto) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *TruckDriverDto) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *TruckDriverDto) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseClass

`func (o *TruckDriverDto) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *TruckDriverDto) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *TruckDriverDto) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *TruckDriverDto) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *TruckDriverDto) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *TruckDriverDto) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetLicenseExpiryDate

`func (o *TruckDriverDto) GetLicenseExpiryDate() time.Time`

GetLicenseExpiryDate returns the LicenseExpiryDate field if non-nil, zero value otherwise.

### GetLicenseExpiryDateOk

`func (o *TruckDriverDto) GetLicenseExpiryDateOk() (*time.Time, bool)`

GetLicenseExpiryDateOk returns a tuple with the LicenseExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiryDate

`func (o *TruckDriverDto) SetLicenseExpiryDate(v time.Time)`

SetLicenseExpiryDate sets LicenseExpiryDate field to given value.

### HasLicenseExpiryDate

`func (o *TruckDriverDto) HasLicenseExpiryDate() bool`

HasLicenseExpiryDate returns a boolean if a field has been set.

### SetLicenseExpiryDateNil

`func (o *TruckDriverDto) SetLicenseExpiryDateNil(b bool)`

 SetLicenseExpiryDateNil sets the value for LicenseExpiryDate to be an explicit nil

### UnsetLicenseExpiryDate
`func (o *TruckDriverDto) UnsetLicenseExpiryDate()`

UnsetLicenseExpiryDate ensures that no value is present for LicenseExpiryDate, not even an explicit nil
### GetAdrCertified

`func (o *TruckDriverDto) GetAdrCertified() bool`

GetAdrCertified returns the AdrCertified field if non-nil, zero value otherwise.

### GetAdrCertifiedOk

`func (o *TruckDriverDto) GetAdrCertifiedOk() (*bool, bool)`

GetAdrCertifiedOk returns a tuple with the AdrCertified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrCertified

`func (o *TruckDriverDto) SetAdrCertified(v bool)`

SetAdrCertified sets AdrCertified field to given value.

### HasAdrCertified

`func (o *TruckDriverDto) HasAdrCertified() bool`

HasAdrCertified returns a boolean if a field has been set.

### GetAdrCertificateExpiryDate

`func (o *TruckDriverDto) GetAdrCertificateExpiryDate() time.Time`

GetAdrCertificateExpiryDate returns the AdrCertificateExpiryDate field if non-nil, zero value otherwise.

### GetAdrCertificateExpiryDateOk

`func (o *TruckDriverDto) GetAdrCertificateExpiryDateOk() (*time.Time, bool)`

GetAdrCertificateExpiryDateOk returns a tuple with the AdrCertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrCertificateExpiryDate

`func (o *TruckDriverDto) SetAdrCertificateExpiryDate(v time.Time)`

SetAdrCertificateExpiryDate sets AdrCertificateExpiryDate field to given value.

### HasAdrCertificateExpiryDate

`func (o *TruckDriverDto) HasAdrCertificateExpiryDate() bool`

HasAdrCertificateExpiryDate returns a boolean if a field has been set.

### SetAdrCertificateExpiryDateNil

`func (o *TruckDriverDto) SetAdrCertificateExpiryDateNil(b bool)`

 SetAdrCertificateExpiryDateNil sets the value for AdrCertificateExpiryDate to be an explicit nil

### UnsetAdrCertificateExpiryDate
`func (o *TruckDriverDto) UnsetAdrCertificateExpiryDate()`

UnsetAdrCertificateExpiryDate ensures that no value is present for AdrCertificateExpiryDate, not even an explicit nil
### GetMedicalExamExpiryDate

`func (o *TruckDriverDto) GetMedicalExamExpiryDate() time.Time`

GetMedicalExamExpiryDate returns the MedicalExamExpiryDate field if non-nil, zero value otherwise.

### GetMedicalExamExpiryDateOk

`func (o *TruckDriverDto) GetMedicalExamExpiryDateOk() (*time.Time, bool)`

GetMedicalExamExpiryDateOk returns a tuple with the MedicalExamExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalExamExpiryDate

`func (o *TruckDriverDto) SetMedicalExamExpiryDate(v time.Time)`

SetMedicalExamExpiryDate sets MedicalExamExpiryDate field to given value.

### HasMedicalExamExpiryDate

`func (o *TruckDriverDto) HasMedicalExamExpiryDate() bool`

HasMedicalExamExpiryDate returns a boolean if a field has been set.

### SetMedicalExamExpiryDateNil

`func (o *TruckDriverDto) SetMedicalExamExpiryDateNil(b bool)`

 SetMedicalExamExpiryDateNil sets the value for MedicalExamExpiryDate to be an explicit nil

### UnsetMedicalExamExpiryDate
`func (o *TruckDriverDto) UnsetMedicalExamExpiryDate()`

UnsetMedicalExamExpiryDate ensures that no value is present for MedicalExamExpiryDate, not even an explicit nil
### GetIsActive

`func (o *TruckDriverDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TruckDriverDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TruckDriverDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TruckDriverDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNotes

`func (o *TruckDriverDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TruckDriverDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TruckDriverDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TruckDriverDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TruckDriverDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TruckDriverDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetContactId

`func (o *TruckDriverDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TruckDriverDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TruckDriverDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TruckDriverDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *TruckDriverDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *TruckDriverDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetShippingCourierId

`func (o *TruckDriverDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *TruckDriverDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *TruckDriverDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *TruckDriverDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *TruckDriverDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *TruckDriverDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetTenantId

`func (o *TruckDriverDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TruckDriverDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TruckDriverDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TruckDriverDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TruckDriverDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TruckDriverDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TruckDriverDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TruckDriverDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TruckDriverDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TruckDriverDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TruckDriverDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TruckDriverDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


