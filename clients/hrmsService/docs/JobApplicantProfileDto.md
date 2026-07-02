# JobApplicantProfileDto

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
**AvailableForHire** | Pointer to **bool** |  | [optional] 
**CareerLevel** | Pointer to **string** |  | [optional] 
**ExperienceInYears** | Pointer to **int32** |  | [optional] 
**CurrentSalary** | Pointer to **float64** |  | [optional] 
**MinSalaryExpectation** | Pointer to **float64** |  | [optional] 
**MaxSalaryExpectation** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobApplicantProfileDto

`func NewJobApplicantProfileDto() *JobApplicantProfileDto`

NewJobApplicantProfileDto instantiates a new JobApplicantProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobApplicantProfileDtoWithDefaults

`func NewJobApplicantProfileDtoWithDefaults() *JobApplicantProfileDto`

NewJobApplicantProfileDtoWithDefaults instantiates a new JobApplicantProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobApplicantProfileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobApplicantProfileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobApplicantProfileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobApplicantProfileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobApplicantProfileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobApplicantProfileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JobApplicantProfileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobApplicantProfileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobApplicantProfileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobApplicantProfileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JobApplicantProfileDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JobApplicantProfileDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetContactId

`func (o *JobApplicantProfileDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *JobApplicantProfileDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *JobApplicantProfileDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *JobApplicantProfileDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *JobApplicantProfileDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *JobApplicantProfileDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *JobApplicantProfileDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JobApplicantProfileDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JobApplicantProfileDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JobApplicantProfileDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JobApplicantProfileDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JobApplicantProfileDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *JobApplicantProfileDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobApplicantProfileDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobApplicantProfileDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobApplicantProfileDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JobApplicantProfileDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JobApplicantProfileDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnrollmentId

`func (o *JobApplicantProfileDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *JobApplicantProfileDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *JobApplicantProfileDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *JobApplicantProfileDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *JobApplicantProfileDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *JobApplicantProfileDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAbout

`func (o *JobApplicantProfileDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *JobApplicantProfileDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *JobApplicantProfileDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *JobApplicantProfileDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *JobApplicantProfileDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *JobApplicantProfileDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *JobApplicantProfileDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *JobApplicantProfileDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *JobApplicantProfileDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *JobApplicantProfileDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *JobApplicantProfileDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *JobApplicantProfileDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *JobApplicantProfileDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *JobApplicantProfileDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *JobApplicantProfileDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *JobApplicantProfileDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *JobApplicantProfileDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *JobApplicantProfileDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *JobApplicantProfileDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *JobApplicantProfileDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetContact

`func (o *JobApplicantProfileDto) GetContact() ContactDto`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *JobApplicantProfileDto) GetContactOk() (*ContactDto, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *JobApplicantProfileDto) SetContact(v ContactDto)`

SetContact sets Contact field to given value.

### HasContact

`func (o *JobApplicantProfileDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetQualifiedName

`func (o *JobApplicantProfileDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *JobApplicantProfileDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *JobApplicantProfileDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *JobApplicantProfileDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *JobApplicantProfileDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *JobApplicantProfileDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *JobApplicantProfileDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *JobApplicantProfileDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *JobApplicantProfileDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *JobApplicantProfileDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### SetVerificationTimestampNil

`func (o *JobApplicantProfileDto) SetVerificationTimestampNil(b bool)`

 SetVerificationTimestampNil sets the value for VerificationTimestamp to be an explicit nil

### UnsetVerificationTimestamp
`func (o *JobApplicantProfileDto) UnsetVerificationTimestamp()`

UnsetVerificationTimestamp ensures that no value is present for VerificationTimestamp, not even an explicit nil
### GetData

`func (o *JobApplicantProfileDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobApplicantProfileDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobApplicantProfileDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JobApplicantProfileDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobApplicantProfileDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobApplicantProfileDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *JobApplicantProfileDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *JobApplicantProfileDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *JobApplicantProfileDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *JobApplicantProfileDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *JobApplicantProfileDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *JobApplicantProfileDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *JobApplicantProfileDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *JobApplicantProfileDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *JobApplicantProfileDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *JobApplicantProfileDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *JobApplicantProfileDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *JobApplicantProfileDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *JobApplicantProfileDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *JobApplicantProfileDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *JobApplicantProfileDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *JobApplicantProfileDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *JobApplicantProfileDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *JobApplicantProfileDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *JobApplicantProfileDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *JobApplicantProfileDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *JobApplicantProfileDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *JobApplicantProfileDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *JobApplicantProfileDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *JobApplicantProfileDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *JobApplicantProfileDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *JobApplicantProfileDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *JobApplicantProfileDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *JobApplicantProfileDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *JobApplicantProfileDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *JobApplicantProfileDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *JobApplicantProfileDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *JobApplicantProfileDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *JobApplicantProfileDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *JobApplicantProfileDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *JobApplicantProfileDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *JobApplicantProfileDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *JobApplicantProfileDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *JobApplicantProfileDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *JobApplicantProfileDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *JobApplicantProfileDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *JobApplicantProfileDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *JobApplicantProfileDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *JobApplicantProfileDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *JobApplicantProfileDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *JobApplicantProfileDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *JobApplicantProfileDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *JobApplicantProfileDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *JobApplicantProfileDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *JobApplicantProfileDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *JobApplicantProfileDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *JobApplicantProfileDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *JobApplicantProfileDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *JobApplicantProfileDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *JobApplicantProfileDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *JobApplicantProfileDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *JobApplicantProfileDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *JobApplicantProfileDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *JobApplicantProfileDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *JobApplicantProfileDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *JobApplicantProfileDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *JobApplicantProfileDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *JobApplicantProfileDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *JobApplicantProfileDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *JobApplicantProfileDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *JobApplicantProfileDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *JobApplicantProfileDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *JobApplicantProfileDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *JobApplicantProfileDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *JobApplicantProfileDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *JobApplicantProfileDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *JobApplicantProfileDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *JobApplicantProfileDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *JobApplicantProfileDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *JobApplicantProfileDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *JobApplicantProfileDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *JobApplicantProfileDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *JobApplicantProfileDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *JobApplicantProfileDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *JobApplicantProfileDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *JobApplicantProfileDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *JobApplicantProfileDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *JobApplicantProfileDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *JobApplicantProfileDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *JobApplicantProfileDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *JobApplicantProfileDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *JobApplicantProfileDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *JobApplicantProfileDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *JobApplicantProfileDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *JobApplicantProfileDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *JobApplicantProfileDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *JobApplicantProfileDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *JobApplicantProfileDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *JobApplicantProfileDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *JobApplicantProfileDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *JobApplicantProfileDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *JobApplicantProfileDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *JobApplicantProfileDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *JobApplicantProfileDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *JobApplicantProfileDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *JobApplicantProfileDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *JobApplicantProfileDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *JobApplicantProfileDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *JobApplicantProfileDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *JobApplicantProfileDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *JobApplicantProfileDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *JobApplicantProfileDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *JobApplicantProfileDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *JobApplicantProfileDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *JobApplicantProfileDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *JobApplicantProfileDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *JobApplicantProfileDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *JobApplicantProfileDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *JobApplicantProfileDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *JobApplicantProfileDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetAvailableForHire

`func (o *JobApplicantProfileDto) GetAvailableForHire() bool`

GetAvailableForHire returns the AvailableForHire field if non-nil, zero value otherwise.

### GetAvailableForHireOk

`func (o *JobApplicantProfileDto) GetAvailableForHireOk() (*bool, bool)`

GetAvailableForHireOk returns a tuple with the AvailableForHire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForHire

`func (o *JobApplicantProfileDto) SetAvailableForHire(v bool)`

SetAvailableForHire sets AvailableForHire field to given value.

### HasAvailableForHire

`func (o *JobApplicantProfileDto) HasAvailableForHire() bool`

HasAvailableForHire returns a boolean if a field has been set.

### GetCareerLevel

`func (o *JobApplicantProfileDto) GetCareerLevel() string`

GetCareerLevel returns the CareerLevel field if non-nil, zero value otherwise.

### GetCareerLevelOk

`func (o *JobApplicantProfileDto) GetCareerLevelOk() (*string, bool)`

GetCareerLevelOk returns a tuple with the CareerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareerLevel

`func (o *JobApplicantProfileDto) SetCareerLevel(v string)`

SetCareerLevel sets CareerLevel field to given value.

### HasCareerLevel

`func (o *JobApplicantProfileDto) HasCareerLevel() bool`

HasCareerLevel returns a boolean if a field has been set.

### GetExperienceInYears

`func (o *JobApplicantProfileDto) GetExperienceInYears() int32`

GetExperienceInYears returns the ExperienceInYears field if non-nil, zero value otherwise.

### GetExperienceInYearsOk

`func (o *JobApplicantProfileDto) GetExperienceInYearsOk() (*int32, bool)`

GetExperienceInYearsOk returns a tuple with the ExperienceInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInYears

`func (o *JobApplicantProfileDto) SetExperienceInYears(v int32)`

SetExperienceInYears sets ExperienceInYears field to given value.

### HasExperienceInYears

`func (o *JobApplicantProfileDto) HasExperienceInYears() bool`

HasExperienceInYears returns a boolean if a field has been set.

### GetCurrentSalary

`func (o *JobApplicantProfileDto) GetCurrentSalary() float64`

GetCurrentSalary returns the CurrentSalary field if non-nil, zero value otherwise.

### GetCurrentSalaryOk

`func (o *JobApplicantProfileDto) GetCurrentSalaryOk() (*float64, bool)`

GetCurrentSalaryOk returns a tuple with the CurrentSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSalary

`func (o *JobApplicantProfileDto) SetCurrentSalary(v float64)`

SetCurrentSalary sets CurrentSalary field to given value.

### HasCurrentSalary

`func (o *JobApplicantProfileDto) HasCurrentSalary() bool`

HasCurrentSalary returns a boolean if a field has been set.

### GetMinSalaryExpectation

`func (o *JobApplicantProfileDto) GetMinSalaryExpectation() float64`

GetMinSalaryExpectation returns the MinSalaryExpectation field if non-nil, zero value otherwise.

### GetMinSalaryExpectationOk

`func (o *JobApplicantProfileDto) GetMinSalaryExpectationOk() (*float64, bool)`

GetMinSalaryExpectationOk returns a tuple with the MinSalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSalaryExpectation

`func (o *JobApplicantProfileDto) SetMinSalaryExpectation(v float64)`

SetMinSalaryExpectation sets MinSalaryExpectation field to given value.

### HasMinSalaryExpectation

`func (o *JobApplicantProfileDto) HasMinSalaryExpectation() bool`

HasMinSalaryExpectation returns a boolean if a field has been set.

### GetMaxSalaryExpectation

`func (o *JobApplicantProfileDto) GetMaxSalaryExpectation() float64`

GetMaxSalaryExpectation returns the MaxSalaryExpectation field if non-nil, zero value otherwise.

### GetMaxSalaryExpectationOk

`func (o *JobApplicantProfileDto) GetMaxSalaryExpectationOk() (*float64, bool)`

GetMaxSalaryExpectationOk returns a tuple with the MaxSalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSalaryExpectation

`func (o *JobApplicantProfileDto) SetMaxSalaryExpectation(v float64)`

SetMaxSalaryExpectation sets MaxSalaryExpectation field to given value.

### HasMaxSalaryExpectation

`func (o *JobApplicantProfileDto) HasMaxSalaryExpectation() bool`

HasMaxSalaryExpectation returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobApplicantProfileDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobApplicantProfileDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobApplicantProfileDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobApplicantProfileDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobApplicantProfileDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobApplicantProfileDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


