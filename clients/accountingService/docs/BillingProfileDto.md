# BillingProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Submitted** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**VerificationTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**DataLabel** | Pointer to **NullableString** |  | [optional] 
**Data1** | Pointer to **NullableString** |  | [optional] 
**Data1Label** | Pointer to **NullableString** |  | [optional] 
**Data2** | Pointer to **NullableString** |  | [optional] 
**Data2Label** | Pointer to **NullableString** |  | [optional] 
**Data3** | Pointer to **NullableString** |  | [optional] 
**Data3Label** | Pointer to **NullableString** |  | [optional] 
**Data4** | Pointer to **NullableString** |  | [optional] 
**Data4Label** | Pointer to **NullableString** |  | [optional] 
**Data5** | Pointer to **NullableString** |  | [optional] 
**Data5Label** | Pointer to **NullableString** |  | [optional] 
**Data6** | Pointer to **NullableString** |  | [optional] 
**Data6Label** | Pointer to **NullableString** |  | [optional] 
**Data7** | Pointer to **NullableString** |  | [optional] 
**Data7Label** | Pointer to **NullableString** |  | [optional] 
**Data8** | Pointer to **NullableString** |  | [optional] 
**Data8Label** | Pointer to **NullableString** |  | [optional] 
**Data9** | Pointer to **NullableString** |  | [optional] 
**Data9Label** | Pointer to **NullableString** |  | [optional] 
**TaxId** | **string** |  | 
**Email** | **string** |  | 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Address1** | Pointer to **NullableString** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**BusinessName** | Pointer to **NullableString** |  | [optional] 
**CommercialName** | Pointer to **NullableString** |  | [optional] 
**Ticker** | Pointer to **NullableString** |  | [optional] 
**Duns** | Pointer to **NullableString** |  | [optional] 
**IsPublicCompany** | Pointer to **bool** |  | [optional] 
**IsFactaCustomer** | Pointer to **bool** |  | [optional] 
**TaxPayerType** | Pointer to **string** |  | [optional] 
**CountryId** | **string** |  | 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**FiscalIdentificationTypeId** | Pointer to **NullableString** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**FiscalRegimeId** | Pointer to **NullableString** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**FiscalAuthorityName** | Pointer to **NullableString** |  | [optional] 
**CountryName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillingProfileDto

`func NewBillingProfileDto(taxId string, email string, countryId string, ) *BillingProfileDto`

NewBillingProfileDto instantiates a new BillingProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileDtoWithDefaults

`func NewBillingProfileDtoWithDefaults() *BillingProfileDto`

NewBillingProfileDtoWithDefaults instantiates a new BillingProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingProfileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingProfileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingProfileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingProfileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BillingProfileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BillingProfileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BillingProfileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillingProfileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillingProfileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillingProfileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *BillingProfileDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingProfileDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingProfileDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingProfileDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BillingProfileDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BillingProfileDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetContactId

`func (o *BillingProfileDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *BillingProfileDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *BillingProfileDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *BillingProfileDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *BillingProfileDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *BillingProfileDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetEnrollmentId

`func (o *BillingProfileDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BillingProfileDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BillingProfileDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BillingProfileDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BillingProfileDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BillingProfileDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAbout

`func (o *BillingProfileDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *BillingProfileDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *BillingProfileDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *BillingProfileDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *BillingProfileDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *BillingProfileDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *BillingProfileDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *BillingProfileDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *BillingProfileDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *BillingProfileDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *BillingProfileDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *BillingProfileDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *BillingProfileDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *BillingProfileDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *BillingProfileDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *BillingProfileDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *BillingProfileDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *BillingProfileDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *BillingProfileDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *BillingProfileDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetQualifiedName

`func (o *BillingProfileDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *BillingProfileDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *BillingProfileDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *BillingProfileDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *BillingProfileDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *BillingProfileDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *BillingProfileDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *BillingProfileDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *BillingProfileDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *BillingProfileDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### SetVerificationTimestampNil

`func (o *BillingProfileDto) SetVerificationTimestampNil(b bool)`

 SetVerificationTimestampNil sets the value for VerificationTimestamp to be an explicit nil

### UnsetVerificationTimestamp
`func (o *BillingProfileDto) UnsetVerificationTimestamp()`

UnsetVerificationTimestamp ensures that no value is present for VerificationTimestamp, not even an explicit nil
### GetData

`func (o *BillingProfileDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingProfileDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingProfileDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BillingProfileDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BillingProfileDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BillingProfileDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *BillingProfileDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *BillingProfileDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *BillingProfileDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *BillingProfileDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *BillingProfileDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *BillingProfileDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *BillingProfileDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *BillingProfileDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *BillingProfileDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *BillingProfileDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *BillingProfileDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *BillingProfileDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *BillingProfileDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *BillingProfileDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *BillingProfileDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *BillingProfileDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *BillingProfileDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *BillingProfileDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *BillingProfileDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *BillingProfileDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *BillingProfileDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *BillingProfileDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *BillingProfileDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *BillingProfileDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *BillingProfileDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *BillingProfileDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *BillingProfileDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *BillingProfileDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *BillingProfileDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *BillingProfileDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *BillingProfileDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *BillingProfileDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *BillingProfileDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *BillingProfileDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *BillingProfileDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *BillingProfileDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *BillingProfileDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *BillingProfileDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *BillingProfileDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *BillingProfileDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *BillingProfileDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *BillingProfileDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *BillingProfileDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *BillingProfileDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *BillingProfileDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *BillingProfileDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *BillingProfileDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *BillingProfileDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *BillingProfileDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *BillingProfileDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *BillingProfileDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *BillingProfileDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *BillingProfileDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *BillingProfileDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *BillingProfileDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *BillingProfileDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *BillingProfileDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *BillingProfileDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *BillingProfileDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *BillingProfileDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *BillingProfileDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *BillingProfileDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *BillingProfileDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *BillingProfileDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *BillingProfileDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *BillingProfileDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *BillingProfileDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *BillingProfileDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *BillingProfileDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *BillingProfileDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *BillingProfileDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *BillingProfileDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *BillingProfileDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *BillingProfileDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *BillingProfileDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *BillingProfileDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *BillingProfileDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *BillingProfileDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *BillingProfileDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *BillingProfileDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *BillingProfileDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *BillingProfileDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *BillingProfileDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *BillingProfileDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *BillingProfileDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *BillingProfileDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *BillingProfileDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *BillingProfileDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *BillingProfileDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *BillingProfileDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *BillingProfileDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *BillingProfileDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *BillingProfileDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *BillingProfileDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *BillingProfileDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *BillingProfileDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *BillingProfileDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *BillingProfileDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *BillingProfileDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *BillingProfileDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *BillingProfileDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *BillingProfileDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *BillingProfileDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *BillingProfileDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *BillingProfileDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *BillingProfileDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *BillingProfileDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *BillingProfileDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *BillingProfileDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *BillingProfileDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *BillingProfileDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *BillingProfileDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *BillingProfileDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *BillingProfileDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetTaxId

`func (o *BillingProfileDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingProfileDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingProfileDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### GetEmail

`func (o *BillingProfileDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingProfileDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingProfileDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *BillingProfileDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingProfileDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingProfileDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BillingProfileDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *BillingProfileDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *BillingProfileDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAddress

`func (o *BillingProfileDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingProfileDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingProfileDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BillingProfileDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *BillingProfileDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *BillingProfileDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddress1

`func (o *BillingProfileDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *BillingProfileDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *BillingProfileDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *BillingProfileDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *BillingProfileDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *BillingProfileDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *BillingProfileDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *BillingProfileDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *BillingProfileDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *BillingProfileDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *BillingProfileDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *BillingProfileDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetPostalCode

`func (o *BillingProfileDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingProfileDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingProfileDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingProfileDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *BillingProfileDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *BillingProfileDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetBusinessName

`func (o *BillingProfileDto) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *BillingProfileDto) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *BillingProfileDto) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *BillingProfileDto) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *BillingProfileDto) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *BillingProfileDto) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetCommercialName

`func (o *BillingProfileDto) GetCommercialName() string`

GetCommercialName returns the CommercialName field if non-nil, zero value otherwise.

### GetCommercialNameOk

`func (o *BillingProfileDto) GetCommercialNameOk() (*string, bool)`

GetCommercialNameOk returns a tuple with the CommercialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialName

`func (o *BillingProfileDto) SetCommercialName(v string)`

SetCommercialName sets CommercialName field to given value.

### HasCommercialName

`func (o *BillingProfileDto) HasCommercialName() bool`

HasCommercialName returns a boolean if a field has been set.

### SetCommercialNameNil

`func (o *BillingProfileDto) SetCommercialNameNil(b bool)`

 SetCommercialNameNil sets the value for CommercialName to be an explicit nil

### UnsetCommercialName
`func (o *BillingProfileDto) UnsetCommercialName()`

UnsetCommercialName ensures that no value is present for CommercialName, not even an explicit nil
### GetTicker

`func (o *BillingProfileDto) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *BillingProfileDto) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *BillingProfileDto) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *BillingProfileDto) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *BillingProfileDto) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *BillingProfileDto) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetDuns

`func (o *BillingProfileDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *BillingProfileDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *BillingProfileDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *BillingProfileDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *BillingProfileDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *BillingProfileDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetIsPublicCompany

`func (o *BillingProfileDto) GetIsPublicCompany() bool`

GetIsPublicCompany returns the IsPublicCompany field if non-nil, zero value otherwise.

### GetIsPublicCompanyOk

`func (o *BillingProfileDto) GetIsPublicCompanyOk() (*bool, bool)`

GetIsPublicCompanyOk returns a tuple with the IsPublicCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicCompany

`func (o *BillingProfileDto) SetIsPublicCompany(v bool)`

SetIsPublicCompany sets IsPublicCompany field to given value.

### HasIsPublicCompany

`func (o *BillingProfileDto) HasIsPublicCompany() bool`

HasIsPublicCompany returns a boolean if a field has been set.

### GetIsFactaCustomer

`func (o *BillingProfileDto) GetIsFactaCustomer() bool`

GetIsFactaCustomer returns the IsFactaCustomer field if non-nil, zero value otherwise.

### GetIsFactaCustomerOk

`func (o *BillingProfileDto) GetIsFactaCustomerOk() (*bool, bool)`

GetIsFactaCustomerOk returns a tuple with the IsFactaCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFactaCustomer

`func (o *BillingProfileDto) SetIsFactaCustomer(v bool)`

SetIsFactaCustomer sets IsFactaCustomer field to given value.

### HasIsFactaCustomer

`func (o *BillingProfileDto) HasIsFactaCustomer() bool`

HasIsFactaCustomer returns a boolean if a field has been set.

### GetTaxPayerType

`func (o *BillingProfileDto) GetTaxPayerType() string`

GetTaxPayerType returns the TaxPayerType field if non-nil, zero value otherwise.

### GetTaxPayerTypeOk

`func (o *BillingProfileDto) GetTaxPayerTypeOk() (*string, bool)`

GetTaxPayerTypeOk returns a tuple with the TaxPayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPayerType

`func (o *BillingProfileDto) SetTaxPayerType(v string)`

SetTaxPayerType sets TaxPayerType field to given value.

### HasTaxPayerType

`func (o *BillingProfileDto) HasTaxPayerType() bool`

HasTaxPayerType returns a boolean if a field has been set.

### GetCountryId

`func (o *BillingProfileDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *BillingProfileDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *BillingProfileDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.


### GetStateId

`func (o *BillingProfileDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *BillingProfileDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *BillingProfileDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *BillingProfileDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *BillingProfileDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *BillingProfileDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *BillingProfileDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *BillingProfileDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *BillingProfileDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *BillingProfileDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *BillingProfileDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *BillingProfileDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetFiscalIdentificationTypeId

`func (o *BillingProfileDto) GetFiscalIdentificationTypeId() string`

GetFiscalIdentificationTypeId returns the FiscalIdentificationTypeId field if non-nil, zero value otherwise.

### GetFiscalIdentificationTypeIdOk

`func (o *BillingProfileDto) GetFiscalIdentificationTypeIdOk() (*string, bool)`

GetFiscalIdentificationTypeIdOk returns a tuple with the FiscalIdentificationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalIdentificationTypeId

`func (o *BillingProfileDto) SetFiscalIdentificationTypeId(v string)`

SetFiscalIdentificationTypeId sets FiscalIdentificationTypeId field to given value.

### HasFiscalIdentificationTypeId

`func (o *BillingProfileDto) HasFiscalIdentificationTypeId() bool`

HasFiscalIdentificationTypeId returns a boolean if a field has been set.

### SetFiscalIdentificationTypeIdNil

`func (o *BillingProfileDto) SetFiscalIdentificationTypeIdNil(b bool)`

 SetFiscalIdentificationTypeIdNil sets the value for FiscalIdentificationTypeId to be an explicit nil

### UnsetFiscalIdentificationTypeId
`func (o *BillingProfileDto) UnsetFiscalIdentificationTypeId()`

UnsetFiscalIdentificationTypeId ensures that no value is present for FiscalIdentificationTypeId, not even an explicit nil
### GetFiscalAuthorityId

`func (o *BillingProfileDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *BillingProfileDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *BillingProfileDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *BillingProfileDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *BillingProfileDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *BillingProfileDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetFiscalRegimeId

`func (o *BillingProfileDto) GetFiscalRegimeId() string`

GetFiscalRegimeId returns the FiscalRegimeId field if non-nil, zero value otherwise.

### GetFiscalRegimeIdOk

`func (o *BillingProfileDto) GetFiscalRegimeIdOk() (*string, bool)`

GetFiscalRegimeIdOk returns a tuple with the FiscalRegimeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalRegimeId

`func (o *BillingProfileDto) SetFiscalRegimeId(v string)`

SetFiscalRegimeId sets FiscalRegimeId field to given value.

### HasFiscalRegimeId

`func (o *BillingProfileDto) HasFiscalRegimeId() bool`

HasFiscalRegimeId returns a boolean if a field has been set.

### SetFiscalRegimeIdNil

`func (o *BillingProfileDto) SetFiscalRegimeIdNil(b bool)`

 SetFiscalRegimeIdNil sets the value for FiscalRegimeId to be an explicit nil

### UnsetFiscalRegimeId
`func (o *BillingProfileDto) UnsetFiscalRegimeId()`

UnsetFiscalRegimeId ensures that no value is present for FiscalRegimeId, not even an explicit nil
### GetContactName

`func (o *BillingProfileDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *BillingProfileDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *BillingProfileDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *BillingProfileDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *BillingProfileDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *BillingProfileDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetFiscalAuthorityName

`func (o *BillingProfileDto) GetFiscalAuthorityName() string`

GetFiscalAuthorityName returns the FiscalAuthorityName field if non-nil, zero value otherwise.

### GetFiscalAuthorityNameOk

`func (o *BillingProfileDto) GetFiscalAuthorityNameOk() (*string, bool)`

GetFiscalAuthorityNameOk returns a tuple with the FiscalAuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityName

`func (o *BillingProfileDto) SetFiscalAuthorityName(v string)`

SetFiscalAuthorityName sets FiscalAuthorityName field to given value.

### HasFiscalAuthorityName

`func (o *BillingProfileDto) HasFiscalAuthorityName() bool`

HasFiscalAuthorityName returns a boolean if a field has been set.

### SetFiscalAuthorityNameNil

`func (o *BillingProfileDto) SetFiscalAuthorityNameNil(b bool)`

 SetFiscalAuthorityNameNil sets the value for FiscalAuthorityName to be an explicit nil

### UnsetFiscalAuthorityName
`func (o *BillingProfileDto) UnsetFiscalAuthorityName()`

UnsetFiscalAuthorityName ensures that no value is present for FiscalAuthorityName, not even an explicit nil
### GetCountryName

`func (o *BillingProfileDto) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *BillingProfileDto) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *BillingProfileDto) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *BillingProfileDto) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### SetCountryNameNil

`func (o *BillingProfileDto) SetCountryNameNil(b bool)`

 SetCountryNameNil sets the value for CountryName to be an explicit nil

### UnsetCountryName
`func (o *BillingProfileDto) UnsetCountryName()`

UnsetCountryName ensures that no value is present for CountryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


