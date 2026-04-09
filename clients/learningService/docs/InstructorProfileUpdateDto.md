# InstructorProfileUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Authorized** | Pointer to **bool** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Submitted** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**VerificationTimestamp** | Pointer to **time.Time** |  | [optional] 
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

## Methods

### NewInstructorProfileUpdateDto

`func NewInstructorProfileUpdateDto() *InstructorProfileUpdateDto`

NewInstructorProfileUpdateDto instantiates a new InstructorProfileUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructorProfileUpdateDtoWithDefaults

`func NewInstructorProfileUpdateDtoWithDefaults() *InstructorProfileUpdateDto`

NewInstructorProfileUpdateDtoWithDefaults instantiates a new InstructorProfileUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstructorProfileUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstructorProfileUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstructorProfileUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstructorProfileUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstructorProfileUpdateDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstructorProfileUpdateDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InstructorProfileUpdateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InstructorProfileUpdateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InstructorProfileUpdateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InstructorProfileUpdateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InstructorProfileUpdateDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InstructorProfileUpdateDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetAuthorized

`func (o *InstructorProfileUpdateDto) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *InstructorProfileUpdateDto) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *InstructorProfileUpdateDto) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *InstructorProfileUpdateDto) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetAbout

`func (o *InstructorProfileUpdateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *InstructorProfileUpdateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *InstructorProfileUpdateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *InstructorProfileUpdateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *InstructorProfileUpdateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *InstructorProfileUpdateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *InstructorProfileUpdateDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *InstructorProfileUpdateDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *InstructorProfileUpdateDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *InstructorProfileUpdateDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *InstructorProfileUpdateDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *InstructorProfileUpdateDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *InstructorProfileUpdateDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *InstructorProfileUpdateDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *InstructorProfileUpdateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *InstructorProfileUpdateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *InstructorProfileUpdateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *InstructorProfileUpdateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *InstructorProfileUpdateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *InstructorProfileUpdateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetQualifiedName

`func (o *InstructorProfileUpdateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *InstructorProfileUpdateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *InstructorProfileUpdateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *InstructorProfileUpdateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *InstructorProfileUpdateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *InstructorProfileUpdateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *InstructorProfileUpdateDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *InstructorProfileUpdateDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *InstructorProfileUpdateDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *InstructorProfileUpdateDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### GetData

`func (o *InstructorProfileUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstructorProfileUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstructorProfileUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *InstructorProfileUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *InstructorProfileUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *InstructorProfileUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *InstructorProfileUpdateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *InstructorProfileUpdateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *InstructorProfileUpdateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *InstructorProfileUpdateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *InstructorProfileUpdateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *InstructorProfileUpdateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *InstructorProfileUpdateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *InstructorProfileUpdateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *InstructorProfileUpdateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *InstructorProfileUpdateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *InstructorProfileUpdateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *InstructorProfileUpdateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *InstructorProfileUpdateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *InstructorProfileUpdateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *InstructorProfileUpdateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *InstructorProfileUpdateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *InstructorProfileUpdateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *InstructorProfileUpdateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *InstructorProfileUpdateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *InstructorProfileUpdateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *InstructorProfileUpdateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *InstructorProfileUpdateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *InstructorProfileUpdateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *InstructorProfileUpdateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *InstructorProfileUpdateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *InstructorProfileUpdateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *InstructorProfileUpdateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *InstructorProfileUpdateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *InstructorProfileUpdateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *InstructorProfileUpdateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *InstructorProfileUpdateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *InstructorProfileUpdateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *InstructorProfileUpdateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *InstructorProfileUpdateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *InstructorProfileUpdateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *InstructorProfileUpdateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *InstructorProfileUpdateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *InstructorProfileUpdateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *InstructorProfileUpdateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *InstructorProfileUpdateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *InstructorProfileUpdateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *InstructorProfileUpdateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *InstructorProfileUpdateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *InstructorProfileUpdateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *InstructorProfileUpdateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *InstructorProfileUpdateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *InstructorProfileUpdateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *InstructorProfileUpdateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *InstructorProfileUpdateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *InstructorProfileUpdateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *InstructorProfileUpdateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *InstructorProfileUpdateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *InstructorProfileUpdateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *InstructorProfileUpdateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *InstructorProfileUpdateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *InstructorProfileUpdateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *InstructorProfileUpdateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *InstructorProfileUpdateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *InstructorProfileUpdateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *InstructorProfileUpdateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *InstructorProfileUpdateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *InstructorProfileUpdateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *InstructorProfileUpdateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *InstructorProfileUpdateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *InstructorProfileUpdateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *InstructorProfileUpdateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *InstructorProfileUpdateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *InstructorProfileUpdateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *InstructorProfileUpdateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *InstructorProfileUpdateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *InstructorProfileUpdateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *InstructorProfileUpdateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *InstructorProfileUpdateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *InstructorProfileUpdateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *InstructorProfileUpdateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *InstructorProfileUpdateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *InstructorProfileUpdateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *InstructorProfileUpdateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *InstructorProfileUpdateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *InstructorProfileUpdateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *InstructorProfileUpdateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *InstructorProfileUpdateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *InstructorProfileUpdateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *InstructorProfileUpdateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *InstructorProfileUpdateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *InstructorProfileUpdateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *InstructorProfileUpdateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *InstructorProfileUpdateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *InstructorProfileUpdateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *InstructorProfileUpdateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *InstructorProfileUpdateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *InstructorProfileUpdateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *InstructorProfileUpdateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *InstructorProfileUpdateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *InstructorProfileUpdateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *InstructorProfileUpdateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *InstructorProfileUpdateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *InstructorProfileUpdateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *InstructorProfileUpdateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *InstructorProfileUpdateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *InstructorProfileUpdateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *InstructorProfileUpdateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *InstructorProfileUpdateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *InstructorProfileUpdateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *InstructorProfileUpdateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *InstructorProfileUpdateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *InstructorProfileUpdateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *InstructorProfileUpdateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *InstructorProfileUpdateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *InstructorProfileUpdateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *InstructorProfileUpdateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *InstructorProfileUpdateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *InstructorProfileUpdateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *InstructorProfileUpdateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


