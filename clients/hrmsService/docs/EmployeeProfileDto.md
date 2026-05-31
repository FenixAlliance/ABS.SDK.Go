# EmployeeProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Submitted** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
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
**GrossPay** | Pointer to **NullableFloat64** |  | [optional] 
**NetSalary** | Pointer to **NullableFloat64** |  | [optional] 
**PayrollCurrency** | Pointer to **NullableString** |  | [optional] 
**MaxWorkHoursPerDay** | Pointer to **NullableInt32** |  | [optional] 
**JobTitleId** | Pointer to **NullableString** |  | [optional] 
**EmployeeTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmployeeProfileDto

`func NewEmployeeProfileDto() *EmployeeProfileDto`

NewEmployeeProfileDto instantiates a new EmployeeProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeProfileDtoWithDefaults

`func NewEmployeeProfileDtoWithDefaults() *EmployeeProfileDto`

NewEmployeeProfileDtoWithDefaults instantiates a new EmployeeProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeProfileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeProfileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeProfileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeProfileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EmployeeProfileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EmployeeProfileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *EmployeeProfileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmployeeProfileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmployeeProfileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmployeeProfileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *EmployeeProfileDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *EmployeeProfileDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetContactId

`func (o *EmployeeProfileDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *EmployeeProfileDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *EmployeeProfileDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *EmployeeProfileDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *EmployeeProfileDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *EmployeeProfileDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *EmployeeProfileDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmployeeProfileDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmployeeProfileDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmployeeProfileDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmployeeProfileDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmployeeProfileDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *EmployeeProfileDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmployeeProfileDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmployeeProfileDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmployeeProfileDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EmployeeProfileDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EmployeeProfileDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnrollmentId

`func (o *EmployeeProfileDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *EmployeeProfileDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *EmployeeProfileDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *EmployeeProfileDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *EmployeeProfileDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *EmployeeProfileDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAbout

`func (o *EmployeeProfileDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *EmployeeProfileDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *EmployeeProfileDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *EmployeeProfileDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *EmployeeProfileDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *EmployeeProfileDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *EmployeeProfileDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EmployeeProfileDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EmployeeProfileDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EmployeeProfileDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *EmployeeProfileDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *EmployeeProfileDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *EmployeeProfileDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *EmployeeProfileDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *EmployeeProfileDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *EmployeeProfileDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *EmployeeProfileDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *EmployeeProfileDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *EmployeeProfileDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *EmployeeProfileDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetContact

`func (o *EmployeeProfileDto) GetContact() ContactDto`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *EmployeeProfileDto) GetContactOk() (*ContactDto, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *EmployeeProfileDto) SetContact(v ContactDto)`

SetContact sets Contact field to given value.

### HasContact

`func (o *EmployeeProfileDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetQualifiedName

`func (o *EmployeeProfileDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *EmployeeProfileDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *EmployeeProfileDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *EmployeeProfileDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *EmployeeProfileDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *EmployeeProfileDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *EmployeeProfileDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *EmployeeProfileDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *EmployeeProfileDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *EmployeeProfileDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### SetVerificationTimestampNil

`func (o *EmployeeProfileDto) SetVerificationTimestampNil(b bool)`

 SetVerificationTimestampNil sets the value for VerificationTimestamp to be an explicit nil

### UnsetVerificationTimestamp
`func (o *EmployeeProfileDto) UnsetVerificationTimestamp()`

UnsetVerificationTimestamp ensures that no value is present for VerificationTimestamp, not even an explicit nil
### GetData

`func (o *EmployeeProfileDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmployeeProfileDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmployeeProfileDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EmployeeProfileDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EmployeeProfileDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EmployeeProfileDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *EmployeeProfileDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *EmployeeProfileDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *EmployeeProfileDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *EmployeeProfileDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *EmployeeProfileDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *EmployeeProfileDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *EmployeeProfileDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *EmployeeProfileDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *EmployeeProfileDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *EmployeeProfileDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *EmployeeProfileDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *EmployeeProfileDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *EmployeeProfileDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *EmployeeProfileDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *EmployeeProfileDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *EmployeeProfileDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *EmployeeProfileDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *EmployeeProfileDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *EmployeeProfileDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *EmployeeProfileDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *EmployeeProfileDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *EmployeeProfileDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *EmployeeProfileDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *EmployeeProfileDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *EmployeeProfileDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *EmployeeProfileDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *EmployeeProfileDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *EmployeeProfileDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *EmployeeProfileDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *EmployeeProfileDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *EmployeeProfileDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *EmployeeProfileDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *EmployeeProfileDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *EmployeeProfileDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *EmployeeProfileDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *EmployeeProfileDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *EmployeeProfileDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *EmployeeProfileDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *EmployeeProfileDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *EmployeeProfileDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *EmployeeProfileDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *EmployeeProfileDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *EmployeeProfileDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *EmployeeProfileDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *EmployeeProfileDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *EmployeeProfileDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *EmployeeProfileDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *EmployeeProfileDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *EmployeeProfileDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *EmployeeProfileDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *EmployeeProfileDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *EmployeeProfileDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *EmployeeProfileDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *EmployeeProfileDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *EmployeeProfileDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *EmployeeProfileDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *EmployeeProfileDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *EmployeeProfileDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *EmployeeProfileDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *EmployeeProfileDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *EmployeeProfileDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *EmployeeProfileDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *EmployeeProfileDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *EmployeeProfileDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *EmployeeProfileDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *EmployeeProfileDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *EmployeeProfileDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *EmployeeProfileDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *EmployeeProfileDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *EmployeeProfileDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *EmployeeProfileDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *EmployeeProfileDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *EmployeeProfileDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *EmployeeProfileDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *EmployeeProfileDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *EmployeeProfileDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *EmployeeProfileDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *EmployeeProfileDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *EmployeeProfileDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *EmployeeProfileDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *EmployeeProfileDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *EmployeeProfileDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *EmployeeProfileDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *EmployeeProfileDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *EmployeeProfileDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *EmployeeProfileDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *EmployeeProfileDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *EmployeeProfileDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *EmployeeProfileDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *EmployeeProfileDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *EmployeeProfileDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *EmployeeProfileDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *EmployeeProfileDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *EmployeeProfileDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *EmployeeProfileDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *EmployeeProfileDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *EmployeeProfileDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *EmployeeProfileDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *EmployeeProfileDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *EmployeeProfileDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *EmployeeProfileDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *EmployeeProfileDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *EmployeeProfileDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *EmployeeProfileDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *EmployeeProfileDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *EmployeeProfileDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *EmployeeProfileDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *EmployeeProfileDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *EmployeeProfileDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *EmployeeProfileDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *EmployeeProfileDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *EmployeeProfileDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *EmployeeProfileDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *EmployeeProfileDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetGrossPay

`func (o *EmployeeProfileDto) GetGrossPay() float64`

GetGrossPay returns the GrossPay field if non-nil, zero value otherwise.

### GetGrossPayOk

`func (o *EmployeeProfileDto) GetGrossPayOk() (*float64, bool)`

GetGrossPayOk returns a tuple with the GrossPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPay

`func (o *EmployeeProfileDto) SetGrossPay(v float64)`

SetGrossPay sets GrossPay field to given value.

### HasGrossPay

`func (o *EmployeeProfileDto) HasGrossPay() bool`

HasGrossPay returns a boolean if a field has been set.

### SetGrossPayNil

`func (o *EmployeeProfileDto) SetGrossPayNil(b bool)`

 SetGrossPayNil sets the value for GrossPay to be an explicit nil

### UnsetGrossPay
`func (o *EmployeeProfileDto) UnsetGrossPay()`

UnsetGrossPay ensures that no value is present for GrossPay, not even an explicit nil
### GetNetSalary

`func (o *EmployeeProfileDto) GetNetSalary() float64`

GetNetSalary returns the NetSalary field if non-nil, zero value otherwise.

### GetNetSalaryOk

`func (o *EmployeeProfileDto) GetNetSalaryOk() (*float64, bool)`

GetNetSalaryOk returns a tuple with the NetSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSalary

`func (o *EmployeeProfileDto) SetNetSalary(v float64)`

SetNetSalary sets NetSalary field to given value.

### HasNetSalary

`func (o *EmployeeProfileDto) HasNetSalary() bool`

HasNetSalary returns a boolean if a field has been set.

### SetNetSalaryNil

`func (o *EmployeeProfileDto) SetNetSalaryNil(b bool)`

 SetNetSalaryNil sets the value for NetSalary to be an explicit nil

### UnsetNetSalary
`func (o *EmployeeProfileDto) UnsetNetSalary()`

UnsetNetSalary ensures that no value is present for NetSalary, not even an explicit nil
### GetPayrollCurrency

`func (o *EmployeeProfileDto) GetPayrollCurrency() string`

GetPayrollCurrency returns the PayrollCurrency field if non-nil, zero value otherwise.

### GetPayrollCurrencyOk

`func (o *EmployeeProfileDto) GetPayrollCurrencyOk() (*string, bool)`

GetPayrollCurrencyOk returns a tuple with the PayrollCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollCurrency

`func (o *EmployeeProfileDto) SetPayrollCurrency(v string)`

SetPayrollCurrency sets PayrollCurrency field to given value.

### HasPayrollCurrency

`func (o *EmployeeProfileDto) HasPayrollCurrency() bool`

HasPayrollCurrency returns a boolean if a field has been set.

### SetPayrollCurrencyNil

`func (o *EmployeeProfileDto) SetPayrollCurrencyNil(b bool)`

 SetPayrollCurrencyNil sets the value for PayrollCurrency to be an explicit nil

### UnsetPayrollCurrency
`func (o *EmployeeProfileDto) UnsetPayrollCurrency()`

UnsetPayrollCurrency ensures that no value is present for PayrollCurrency, not even an explicit nil
### GetMaxWorkHoursPerDay

`func (o *EmployeeProfileDto) GetMaxWorkHoursPerDay() int32`

GetMaxWorkHoursPerDay returns the MaxWorkHoursPerDay field if non-nil, zero value otherwise.

### GetMaxWorkHoursPerDayOk

`func (o *EmployeeProfileDto) GetMaxWorkHoursPerDayOk() (*int32, bool)`

GetMaxWorkHoursPerDayOk returns a tuple with the MaxWorkHoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkHoursPerDay

`func (o *EmployeeProfileDto) SetMaxWorkHoursPerDay(v int32)`

SetMaxWorkHoursPerDay sets MaxWorkHoursPerDay field to given value.

### HasMaxWorkHoursPerDay

`func (o *EmployeeProfileDto) HasMaxWorkHoursPerDay() bool`

HasMaxWorkHoursPerDay returns a boolean if a field has been set.

### SetMaxWorkHoursPerDayNil

`func (o *EmployeeProfileDto) SetMaxWorkHoursPerDayNil(b bool)`

 SetMaxWorkHoursPerDayNil sets the value for MaxWorkHoursPerDay to be an explicit nil

### UnsetMaxWorkHoursPerDay
`func (o *EmployeeProfileDto) UnsetMaxWorkHoursPerDay()`

UnsetMaxWorkHoursPerDay ensures that no value is present for MaxWorkHoursPerDay, not even an explicit nil
### GetJobTitleId

`func (o *EmployeeProfileDto) GetJobTitleId() string`

GetJobTitleId returns the JobTitleId field if non-nil, zero value otherwise.

### GetJobTitleIdOk

`func (o *EmployeeProfileDto) GetJobTitleIdOk() (*string, bool)`

GetJobTitleIdOk returns a tuple with the JobTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitleId

`func (o *EmployeeProfileDto) SetJobTitleId(v string)`

SetJobTitleId sets JobTitleId field to given value.

### HasJobTitleId

`func (o *EmployeeProfileDto) HasJobTitleId() bool`

HasJobTitleId returns a boolean if a field has been set.

### SetJobTitleIdNil

`func (o *EmployeeProfileDto) SetJobTitleIdNil(b bool)`

 SetJobTitleIdNil sets the value for JobTitleId to be an explicit nil

### UnsetJobTitleId
`func (o *EmployeeProfileDto) UnsetJobTitleId()`

UnsetJobTitleId ensures that no value is present for JobTitleId, not even an explicit nil
### GetEmployeeTypeId

`func (o *EmployeeProfileDto) GetEmployeeTypeId() string`

GetEmployeeTypeId returns the EmployeeTypeId field if non-nil, zero value otherwise.

### GetEmployeeTypeIdOk

`func (o *EmployeeProfileDto) GetEmployeeTypeIdOk() (*string, bool)`

GetEmployeeTypeIdOk returns a tuple with the EmployeeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeTypeId

`func (o *EmployeeProfileDto) SetEmployeeTypeId(v string)`

SetEmployeeTypeId sets EmployeeTypeId field to given value.

### HasEmployeeTypeId

`func (o *EmployeeProfileDto) HasEmployeeTypeId() bool`

HasEmployeeTypeId returns a boolean if a field has been set.

### SetEmployeeTypeIdNil

`func (o *EmployeeProfileDto) SetEmployeeTypeIdNil(b bool)`

 SetEmployeeTypeIdNil sets the value for EmployeeTypeId to be an explicit nil

### UnsetEmployeeTypeId
`func (o *EmployeeProfileDto) UnsetEmployeeTypeId()`

UnsetEmployeeTypeId ensures that no value is present for EmployeeTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


