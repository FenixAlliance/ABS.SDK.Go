# TruckDriverCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**LicenseNumber** | Pointer to **NullableString** |  | [optional] 
**LicenseClass** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 
**AdrCertified** | Pointer to **bool** |  | [optional] 
**LicenseExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**MedicalExamExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**NationalIdNumber** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckDriverCreateDto

`func NewTruckDriverCreateDto() *TruckDriverCreateDto`

NewTruckDriverCreateDto instantiates a new TruckDriverCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckDriverCreateDtoWithDefaults

`func NewTruckDriverCreateDtoWithDefaults() *TruckDriverCreateDto`

NewTruckDriverCreateDtoWithDefaults instantiates a new TruckDriverCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckDriverCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckDriverCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckDriverCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckDriverCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TruckDriverCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckDriverCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckDriverCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckDriverCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TruckDriverCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TruckDriverCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TruckDriverCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TruckDriverCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TruckDriverCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TruckDriverCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLicenseNumber

`func (o *TruckDriverCreateDto) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *TruckDriverCreateDto) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *TruckDriverCreateDto) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *TruckDriverCreateDto) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *TruckDriverCreateDto) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *TruckDriverCreateDto) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseClass

`func (o *TruckDriverCreateDto) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *TruckDriverCreateDto) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *TruckDriverCreateDto) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *TruckDriverCreateDto) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *TruckDriverCreateDto) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *TruckDriverCreateDto) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetPhone

`func (o *TruckDriverCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TruckDriverCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TruckDriverCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TruckDriverCreateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *TruckDriverCreateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *TruckDriverCreateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *TruckDriverCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TruckDriverCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TruckDriverCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TruckDriverCreateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TruckDriverCreateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TruckDriverCreateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContactId

`func (o *TruckDriverCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TruckDriverCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TruckDriverCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TruckDriverCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *TruckDriverCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *TruckDriverCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetShippingCourierId

`func (o *TruckDriverCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *TruckDriverCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *TruckDriverCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *TruckDriverCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *TruckDriverCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *TruckDriverCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetAdrCertified

`func (o *TruckDriverCreateDto) GetAdrCertified() bool`

GetAdrCertified returns the AdrCertified field if non-nil, zero value otherwise.

### GetAdrCertifiedOk

`func (o *TruckDriverCreateDto) GetAdrCertifiedOk() (*bool, bool)`

GetAdrCertifiedOk returns a tuple with the AdrCertified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrCertified

`func (o *TruckDriverCreateDto) SetAdrCertified(v bool)`

SetAdrCertified sets AdrCertified field to given value.

### HasAdrCertified

`func (o *TruckDriverCreateDto) HasAdrCertified() bool`

HasAdrCertified returns a boolean if a field has been set.

### GetLicenseExpiryDate

`func (o *TruckDriverCreateDto) GetLicenseExpiryDate() time.Time`

GetLicenseExpiryDate returns the LicenseExpiryDate field if non-nil, zero value otherwise.

### GetLicenseExpiryDateOk

`func (o *TruckDriverCreateDto) GetLicenseExpiryDateOk() (*time.Time, bool)`

GetLicenseExpiryDateOk returns a tuple with the LicenseExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiryDate

`func (o *TruckDriverCreateDto) SetLicenseExpiryDate(v time.Time)`

SetLicenseExpiryDate sets LicenseExpiryDate field to given value.

### HasLicenseExpiryDate

`func (o *TruckDriverCreateDto) HasLicenseExpiryDate() bool`

HasLicenseExpiryDate returns a boolean if a field has been set.

### SetLicenseExpiryDateNil

`func (o *TruckDriverCreateDto) SetLicenseExpiryDateNil(b bool)`

 SetLicenseExpiryDateNil sets the value for LicenseExpiryDate to be an explicit nil

### UnsetLicenseExpiryDate
`func (o *TruckDriverCreateDto) UnsetLicenseExpiryDate()`

UnsetLicenseExpiryDate ensures that no value is present for LicenseExpiryDate, not even an explicit nil
### GetMedicalExamExpiryDate

`func (o *TruckDriverCreateDto) GetMedicalExamExpiryDate() time.Time`

GetMedicalExamExpiryDate returns the MedicalExamExpiryDate field if non-nil, zero value otherwise.

### GetMedicalExamExpiryDateOk

`func (o *TruckDriverCreateDto) GetMedicalExamExpiryDateOk() (*time.Time, bool)`

GetMedicalExamExpiryDateOk returns a tuple with the MedicalExamExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalExamExpiryDate

`func (o *TruckDriverCreateDto) SetMedicalExamExpiryDate(v time.Time)`

SetMedicalExamExpiryDate sets MedicalExamExpiryDate field to given value.

### HasMedicalExamExpiryDate

`func (o *TruckDriverCreateDto) HasMedicalExamExpiryDate() bool`

HasMedicalExamExpiryDate returns a boolean if a field has been set.

### SetMedicalExamExpiryDateNil

`func (o *TruckDriverCreateDto) SetMedicalExamExpiryDateNil(b bool)`

 SetMedicalExamExpiryDateNil sets the value for MedicalExamExpiryDate to be an explicit nil

### UnsetMedicalExamExpiryDate
`func (o *TruckDriverCreateDto) UnsetMedicalExamExpiryDate()`

UnsetMedicalExamExpiryDate ensures that no value is present for MedicalExamExpiryDate, not even an explicit nil
### GetNationalIdNumber

`func (o *TruckDriverCreateDto) GetNationalIdNumber() string`

GetNationalIdNumber returns the NationalIdNumber field if non-nil, zero value otherwise.

### GetNationalIdNumberOk

`func (o *TruckDriverCreateDto) GetNationalIdNumberOk() (*string, bool)`

GetNationalIdNumberOk returns a tuple with the NationalIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdNumber

`func (o *TruckDriverCreateDto) SetNationalIdNumber(v string)`

SetNationalIdNumber sets NationalIdNumber field to given value.

### HasNationalIdNumber

`func (o *TruckDriverCreateDto) HasNationalIdNumber() bool`

HasNationalIdNumber returns a boolean if a field has been set.

### SetNationalIdNumberNil

`func (o *TruckDriverCreateDto) SetNationalIdNumberNil(b bool)`

 SetNationalIdNumberNil sets the value for NationalIdNumber to be an explicit nil

### UnsetNationalIdNumber
`func (o *TruckDriverCreateDto) UnsetNationalIdNumber()`

UnsetNationalIdNumber ensures that no value is present for NationalIdNumber, not even an explicit nil
### GetNotes

`func (o *TruckDriverCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TruckDriverCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TruckDriverCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TruckDriverCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TruckDriverCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TruckDriverCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


