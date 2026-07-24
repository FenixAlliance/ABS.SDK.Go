# SigningProfileDto

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
**ContactName** | Pointer to **NullableString** |  | [optional] 
**SignatureFormat** | Pointer to **NullableString** |  | [optional] 
**SignaturePurpose** | Pointer to **NullableString** |  | [optional] 
**DigestAlgorithm** | Pointer to **NullableString** |  | [optional] 
**SignatureAlgorithm** | Pointer to **NullableString** |  | [optional] 
**CanonicalizationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**PolicyIdentifier** | Pointer to **NullableString** |  | [optional] 
**PolicyUri** | Pointer to **NullableString** |  | [optional] 
**AuthorityProfile** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**DefaultForDocumentType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningProfileDto

`func NewSigningProfileDto() *SigningProfileDto`

NewSigningProfileDto instantiates a new SigningProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningProfileDtoWithDefaults

`func NewSigningProfileDtoWithDefaults() *SigningProfileDto`

NewSigningProfileDtoWithDefaults instantiates a new SigningProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningProfileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningProfileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningProfileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningProfileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningProfileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningProfileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SigningProfileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningProfileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningProfileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningProfileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SigningProfileDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SigningProfileDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetContactId

`func (o *SigningProfileDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningProfileDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningProfileDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SigningProfileDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SigningProfileDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SigningProfileDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *SigningProfileDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningProfileDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningProfileDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningProfileDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningProfileDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningProfileDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *SigningProfileDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningProfileDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningProfileDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningProfileDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningProfileDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningProfileDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnrollmentId

`func (o *SigningProfileDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SigningProfileDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SigningProfileDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SigningProfileDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SigningProfileDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SigningProfileDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAbout

`func (o *SigningProfileDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *SigningProfileDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *SigningProfileDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *SigningProfileDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *SigningProfileDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *SigningProfileDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *SigningProfileDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SigningProfileDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SigningProfileDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SigningProfileDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *SigningProfileDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *SigningProfileDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *SigningProfileDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *SigningProfileDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *SigningProfileDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *SigningProfileDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *SigningProfileDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *SigningProfileDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *SigningProfileDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *SigningProfileDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetContact

`func (o *SigningProfileDto) GetContact() ContactDto`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SigningProfileDto) GetContactOk() (*ContactDto, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SigningProfileDto) SetContact(v ContactDto)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SigningProfileDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetQualifiedName

`func (o *SigningProfileDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *SigningProfileDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *SigningProfileDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *SigningProfileDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *SigningProfileDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *SigningProfileDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *SigningProfileDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *SigningProfileDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *SigningProfileDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *SigningProfileDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### SetVerificationTimestampNil

`func (o *SigningProfileDto) SetVerificationTimestampNil(b bool)`

 SetVerificationTimestampNil sets the value for VerificationTimestamp to be an explicit nil

### UnsetVerificationTimestamp
`func (o *SigningProfileDto) UnsetVerificationTimestamp()`

UnsetVerificationTimestamp ensures that no value is present for VerificationTimestamp, not even an explicit nil
### GetData

`func (o *SigningProfileDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SigningProfileDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SigningProfileDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SigningProfileDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SigningProfileDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SigningProfileDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *SigningProfileDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *SigningProfileDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *SigningProfileDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *SigningProfileDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *SigningProfileDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *SigningProfileDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *SigningProfileDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *SigningProfileDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *SigningProfileDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *SigningProfileDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *SigningProfileDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *SigningProfileDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *SigningProfileDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *SigningProfileDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *SigningProfileDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *SigningProfileDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *SigningProfileDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *SigningProfileDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *SigningProfileDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *SigningProfileDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *SigningProfileDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *SigningProfileDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *SigningProfileDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *SigningProfileDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *SigningProfileDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *SigningProfileDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *SigningProfileDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *SigningProfileDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *SigningProfileDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *SigningProfileDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *SigningProfileDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *SigningProfileDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *SigningProfileDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *SigningProfileDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *SigningProfileDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *SigningProfileDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *SigningProfileDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *SigningProfileDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *SigningProfileDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *SigningProfileDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *SigningProfileDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *SigningProfileDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *SigningProfileDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *SigningProfileDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *SigningProfileDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *SigningProfileDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *SigningProfileDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *SigningProfileDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *SigningProfileDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *SigningProfileDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *SigningProfileDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *SigningProfileDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *SigningProfileDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *SigningProfileDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *SigningProfileDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *SigningProfileDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *SigningProfileDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *SigningProfileDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *SigningProfileDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *SigningProfileDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *SigningProfileDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *SigningProfileDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *SigningProfileDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *SigningProfileDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *SigningProfileDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *SigningProfileDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *SigningProfileDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *SigningProfileDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *SigningProfileDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *SigningProfileDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *SigningProfileDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *SigningProfileDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *SigningProfileDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *SigningProfileDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *SigningProfileDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *SigningProfileDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *SigningProfileDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *SigningProfileDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *SigningProfileDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *SigningProfileDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *SigningProfileDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *SigningProfileDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *SigningProfileDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *SigningProfileDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *SigningProfileDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *SigningProfileDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *SigningProfileDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *SigningProfileDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *SigningProfileDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *SigningProfileDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *SigningProfileDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *SigningProfileDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *SigningProfileDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *SigningProfileDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *SigningProfileDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *SigningProfileDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *SigningProfileDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *SigningProfileDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *SigningProfileDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *SigningProfileDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *SigningProfileDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *SigningProfileDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *SigningProfileDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *SigningProfileDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *SigningProfileDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *SigningProfileDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *SigningProfileDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *SigningProfileDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *SigningProfileDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *SigningProfileDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *SigningProfileDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *SigningProfileDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *SigningProfileDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *SigningProfileDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetContactName

`func (o *SigningProfileDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SigningProfileDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SigningProfileDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SigningProfileDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *SigningProfileDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *SigningProfileDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetSignatureFormat

`func (o *SigningProfileDto) GetSignatureFormat() string`

GetSignatureFormat returns the SignatureFormat field if non-nil, zero value otherwise.

### GetSignatureFormatOk

`func (o *SigningProfileDto) GetSignatureFormatOk() (*string, bool)`

GetSignatureFormatOk returns a tuple with the SignatureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFormat

`func (o *SigningProfileDto) SetSignatureFormat(v string)`

SetSignatureFormat sets SignatureFormat field to given value.

### HasSignatureFormat

`func (o *SigningProfileDto) HasSignatureFormat() bool`

HasSignatureFormat returns a boolean if a field has been set.

### SetSignatureFormatNil

`func (o *SigningProfileDto) SetSignatureFormatNil(b bool)`

 SetSignatureFormatNil sets the value for SignatureFormat to be an explicit nil

### UnsetSignatureFormat
`func (o *SigningProfileDto) UnsetSignatureFormat()`

UnsetSignatureFormat ensures that no value is present for SignatureFormat, not even an explicit nil
### GetSignaturePurpose

`func (o *SigningProfileDto) GetSignaturePurpose() string`

GetSignaturePurpose returns the SignaturePurpose field if non-nil, zero value otherwise.

### GetSignaturePurposeOk

`func (o *SigningProfileDto) GetSignaturePurposeOk() (*string, bool)`

GetSignaturePurposeOk returns a tuple with the SignaturePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturePurpose

`func (o *SigningProfileDto) SetSignaturePurpose(v string)`

SetSignaturePurpose sets SignaturePurpose field to given value.

### HasSignaturePurpose

`func (o *SigningProfileDto) HasSignaturePurpose() bool`

HasSignaturePurpose returns a boolean if a field has been set.

### SetSignaturePurposeNil

`func (o *SigningProfileDto) SetSignaturePurposeNil(b bool)`

 SetSignaturePurposeNil sets the value for SignaturePurpose to be an explicit nil

### UnsetSignaturePurpose
`func (o *SigningProfileDto) UnsetSignaturePurpose()`

UnsetSignaturePurpose ensures that no value is present for SignaturePurpose, not even an explicit nil
### GetDigestAlgorithm

`func (o *SigningProfileDto) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SigningProfileDto) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SigningProfileDto) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SigningProfileDto) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### SetDigestAlgorithmNil

`func (o *SigningProfileDto) SetDigestAlgorithmNil(b bool)`

 SetDigestAlgorithmNil sets the value for DigestAlgorithm to be an explicit nil

### UnsetDigestAlgorithm
`func (o *SigningProfileDto) UnsetDigestAlgorithm()`

UnsetDigestAlgorithm ensures that no value is present for DigestAlgorithm, not even an explicit nil
### GetSignatureAlgorithm

`func (o *SigningProfileDto) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SigningProfileDto) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SigningProfileDto) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SigningProfileDto) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### SetSignatureAlgorithmNil

`func (o *SigningProfileDto) SetSignatureAlgorithmNil(b bool)`

 SetSignatureAlgorithmNil sets the value for SignatureAlgorithm to be an explicit nil

### UnsetSignatureAlgorithm
`func (o *SigningProfileDto) UnsetSignatureAlgorithm()`

UnsetSignatureAlgorithm ensures that no value is present for SignatureAlgorithm, not even an explicit nil
### GetCanonicalizationAlgorithm

`func (o *SigningProfileDto) GetCanonicalizationAlgorithm() string`

GetCanonicalizationAlgorithm returns the CanonicalizationAlgorithm field if non-nil, zero value otherwise.

### GetCanonicalizationAlgorithmOk

`func (o *SigningProfileDto) GetCanonicalizationAlgorithmOk() (*string, bool)`

GetCanonicalizationAlgorithmOk returns a tuple with the CanonicalizationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizationAlgorithm

`func (o *SigningProfileDto) SetCanonicalizationAlgorithm(v string)`

SetCanonicalizationAlgorithm sets CanonicalizationAlgorithm field to given value.

### HasCanonicalizationAlgorithm

`func (o *SigningProfileDto) HasCanonicalizationAlgorithm() bool`

HasCanonicalizationAlgorithm returns a boolean if a field has been set.

### SetCanonicalizationAlgorithmNil

`func (o *SigningProfileDto) SetCanonicalizationAlgorithmNil(b bool)`

 SetCanonicalizationAlgorithmNil sets the value for CanonicalizationAlgorithm to be an explicit nil

### UnsetCanonicalizationAlgorithm
`func (o *SigningProfileDto) UnsetCanonicalizationAlgorithm()`

UnsetCanonicalizationAlgorithm ensures that no value is present for CanonicalizationAlgorithm, not even an explicit nil
### GetPolicyIdentifier

`func (o *SigningProfileDto) GetPolicyIdentifier() string`

GetPolicyIdentifier returns the PolicyIdentifier field if non-nil, zero value otherwise.

### GetPolicyIdentifierOk

`func (o *SigningProfileDto) GetPolicyIdentifierOk() (*string, bool)`

GetPolicyIdentifierOk returns a tuple with the PolicyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifier

`func (o *SigningProfileDto) SetPolicyIdentifier(v string)`

SetPolicyIdentifier sets PolicyIdentifier field to given value.

### HasPolicyIdentifier

`func (o *SigningProfileDto) HasPolicyIdentifier() bool`

HasPolicyIdentifier returns a boolean if a field has been set.

### SetPolicyIdentifierNil

`func (o *SigningProfileDto) SetPolicyIdentifierNil(b bool)`

 SetPolicyIdentifierNil sets the value for PolicyIdentifier to be an explicit nil

### UnsetPolicyIdentifier
`func (o *SigningProfileDto) UnsetPolicyIdentifier()`

UnsetPolicyIdentifier ensures that no value is present for PolicyIdentifier, not even an explicit nil
### GetPolicyUri

`func (o *SigningProfileDto) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *SigningProfileDto) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *SigningProfileDto) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *SigningProfileDto) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### SetPolicyUriNil

`func (o *SigningProfileDto) SetPolicyUriNil(b bool)`

 SetPolicyUriNil sets the value for PolicyUri to be an explicit nil

### UnsetPolicyUri
`func (o *SigningProfileDto) UnsetPolicyUri()`

UnsetPolicyUri ensures that no value is present for PolicyUri, not even an explicit nil
### GetAuthorityProfile

`func (o *SigningProfileDto) GetAuthorityProfile() string`

GetAuthorityProfile returns the AuthorityProfile field if non-nil, zero value otherwise.

### GetAuthorityProfileOk

`func (o *SigningProfileDto) GetAuthorityProfileOk() (*string, bool)`

GetAuthorityProfileOk returns a tuple with the AuthorityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityProfile

`func (o *SigningProfileDto) SetAuthorityProfile(v string)`

SetAuthorityProfile sets AuthorityProfile field to given value.

### HasAuthorityProfile

`func (o *SigningProfileDto) HasAuthorityProfile() bool`

HasAuthorityProfile returns a boolean if a field has been set.

### SetAuthorityProfileNil

`func (o *SigningProfileDto) SetAuthorityProfileNil(b bool)`

 SetAuthorityProfileNil sets the value for AuthorityProfile to be an explicit nil

### UnsetAuthorityProfile
`func (o *SigningProfileDto) UnsetAuthorityProfile()`

UnsetAuthorityProfile ensures that no value is present for AuthorityProfile, not even an explicit nil
### GetIsActive

`func (o *SigningProfileDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SigningProfileDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SigningProfileDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SigningProfileDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *SigningProfileDto) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *SigningProfileDto) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetDefaultForDocumentType

`func (o *SigningProfileDto) GetDefaultForDocumentType() string`

GetDefaultForDocumentType returns the DefaultForDocumentType field if non-nil, zero value otherwise.

### GetDefaultForDocumentTypeOk

`func (o *SigningProfileDto) GetDefaultForDocumentTypeOk() (*string, bool)`

GetDefaultForDocumentTypeOk returns a tuple with the DefaultForDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForDocumentType

`func (o *SigningProfileDto) SetDefaultForDocumentType(v string)`

SetDefaultForDocumentType sets DefaultForDocumentType field to given value.

### HasDefaultForDocumentType

`func (o *SigningProfileDto) HasDefaultForDocumentType() bool`

HasDefaultForDocumentType returns a boolean if a field has been set.

### SetDefaultForDocumentTypeNil

`func (o *SigningProfileDto) SetDefaultForDocumentTypeNil(b bool)`

 SetDefaultForDocumentTypeNil sets the value for DefaultForDocumentType to be an explicit nil

### UnsetDefaultForDocumentType
`func (o *SigningProfileDto) UnsetDefaultForDocumentType()`

UnsetDefaultForDocumentType ensures that no value is present for DefaultForDocumentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


