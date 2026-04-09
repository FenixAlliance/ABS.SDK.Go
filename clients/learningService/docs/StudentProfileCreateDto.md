# StudentProfileCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStudentProfileCreateDto

`func NewStudentProfileCreateDto() *StudentProfileCreateDto`

NewStudentProfileCreateDto instantiates a new StudentProfileCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentProfileCreateDtoWithDefaults

`func NewStudentProfileCreateDtoWithDefaults() *StudentProfileCreateDto`

NewStudentProfileCreateDtoWithDefaults instantiates a new StudentProfileCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StudentProfileCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudentProfileCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudentProfileCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudentProfileCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *StudentProfileCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StudentProfileCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StudentProfileCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StudentProfileCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAbout

`func (o *StudentProfileCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *StudentProfileCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *StudentProfileCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *StudentProfileCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *StudentProfileCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *StudentProfileCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetVerified

`func (o *StudentProfileCreateDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *StudentProfileCreateDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *StudentProfileCreateDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *StudentProfileCreateDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetSubmitted

`func (o *StudentProfileCreateDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *StudentProfileCreateDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *StudentProfileCreateDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *StudentProfileCreateDto) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *StudentProfileCreateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *StudentProfileCreateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *StudentProfileCreateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *StudentProfileCreateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *StudentProfileCreateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *StudentProfileCreateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetQualifiedName

`func (o *StudentProfileCreateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *StudentProfileCreateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *StudentProfileCreateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *StudentProfileCreateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *StudentProfileCreateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *StudentProfileCreateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetVerificationTimestamp

`func (o *StudentProfileCreateDto) GetVerificationTimestamp() time.Time`

GetVerificationTimestamp returns the VerificationTimestamp field if non-nil, zero value otherwise.

### GetVerificationTimestampOk

`func (o *StudentProfileCreateDto) GetVerificationTimestampOk() (*time.Time, bool)`

GetVerificationTimestampOk returns a tuple with the VerificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTimestamp

`func (o *StudentProfileCreateDto) SetVerificationTimestamp(v time.Time)`

SetVerificationTimestamp sets VerificationTimestamp field to given value.

### HasVerificationTimestamp

`func (o *StudentProfileCreateDto) HasVerificationTimestamp() bool`

HasVerificationTimestamp returns a boolean if a field has been set.

### GetData

`func (o *StudentProfileCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StudentProfileCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StudentProfileCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *StudentProfileCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *StudentProfileCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *StudentProfileCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *StudentProfileCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *StudentProfileCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *StudentProfileCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *StudentProfileCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *StudentProfileCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *StudentProfileCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *StudentProfileCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *StudentProfileCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *StudentProfileCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *StudentProfileCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *StudentProfileCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *StudentProfileCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *StudentProfileCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *StudentProfileCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *StudentProfileCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *StudentProfileCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *StudentProfileCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *StudentProfileCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *StudentProfileCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *StudentProfileCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *StudentProfileCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *StudentProfileCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *StudentProfileCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *StudentProfileCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *StudentProfileCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *StudentProfileCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *StudentProfileCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *StudentProfileCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *StudentProfileCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *StudentProfileCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *StudentProfileCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *StudentProfileCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *StudentProfileCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *StudentProfileCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *StudentProfileCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *StudentProfileCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *StudentProfileCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *StudentProfileCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *StudentProfileCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *StudentProfileCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *StudentProfileCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *StudentProfileCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *StudentProfileCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *StudentProfileCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *StudentProfileCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *StudentProfileCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *StudentProfileCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *StudentProfileCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *StudentProfileCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *StudentProfileCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *StudentProfileCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *StudentProfileCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *StudentProfileCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *StudentProfileCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *StudentProfileCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *StudentProfileCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *StudentProfileCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *StudentProfileCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *StudentProfileCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *StudentProfileCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *StudentProfileCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *StudentProfileCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *StudentProfileCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *StudentProfileCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *StudentProfileCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *StudentProfileCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *StudentProfileCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *StudentProfileCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *StudentProfileCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *StudentProfileCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *StudentProfileCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *StudentProfileCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *StudentProfileCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *StudentProfileCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *StudentProfileCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *StudentProfileCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *StudentProfileCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *StudentProfileCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *StudentProfileCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *StudentProfileCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *StudentProfileCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *StudentProfileCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *StudentProfileCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *StudentProfileCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *StudentProfileCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *StudentProfileCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *StudentProfileCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *StudentProfileCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *StudentProfileCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *StudentProfileCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *StudentProfileCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *StudentProfileCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *StudentProfileCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *StudentProfileCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *StudentProfileCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *StudentProfileCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *StudentProfileCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *StudentProfileCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *StudentProfileCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *StudentProfileCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *StudentProfileCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *StudentProfileCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *StudentProfileCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *StudentProfileCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *StudentProfileCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *StudentProfileCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *StudentProfileCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *StudentProfileCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *StudentProfileCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *StudentProfileCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *StudentProfileCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *StudentProfileCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *StudentProfileCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *StudentProfileCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetBusinessID

`func (o *StudentProfileCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *StudentProfileCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *StudentProfileCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *StudentProfileCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *StudentProfileCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *StudentProfileCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetContactID

`func (o *StudentProfileCreateDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *StudentProfileCreateDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *StudentProfileCreateDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *StudentProfileCreateDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *StudentProfileCreateDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *StudentProfileCreateDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *StudentProfileCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *StudentProfileCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *StudentProfileCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *StudentProfileCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *StudentProfileCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *StudentProfileCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


