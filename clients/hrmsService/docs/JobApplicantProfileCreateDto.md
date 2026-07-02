# JobApplicantProfileCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to [**ContactCreateDto**](ContactCreateDto.md) |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
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

### NewJobApplicantProfileCreateDto

`func NewJobApplicantProfileCreateDto() *JobApplicantProfileCreateDto`

NewJobApplicantProfileCreateDto instantiates a new JobApplicantProfileCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobApplicantProfileCreateDtoWithDefaults

`func NewJobApplicantProfileCreateDtoWithDefaults() *JobApplicantProfileCreateDto`

NewJobApplicantProfileCreateDtoWithDefaults instantiates a new JobApplicantProfileCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobApplicantProfileCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobApplicantProfileCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobApplicantProfileCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobApplicantProfileCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JobApplicantProfileCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobApplicantProfileCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobApplicantProfileCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobApplicantProfileCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *JobApplicantProfileCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobApplicantProfileCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobApplicantProfileCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobApplicantProfileCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JobApplicantProfileCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JobApplicantProfileCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContactId

`func (o *JobApplicantProfileCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *JobApplicantProfileCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *JobApplicantProfileCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *JobApplicantProfileCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *JobApplicantProfileCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *JobApplicantProfileCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetContact

`func (o *JobApplicantProfileCreateDto) GetContact() ContactCreateDto`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *JobApplicantProfileCreateDto) GetContactOk() (*ContactCreateDto, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *JobApplicantProfileCreateDto) SetContact(v ContactCreateDto)`

SetContact sets Contact field to given value.

### HasContact

`func (o *JobApplicantProfileCreateDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAbout

`func (o *JobApplicantProfileCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *JobApplicantProfileCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *JobApplicantProfileCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *JobApplicantProfileCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *JobApplicantProfileCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *JobApplicantProfileCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetAvatarUrl

`func (o *JobApplicantProfileCreateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *JobApplicantProfileCreateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *JobApplicantProfileCreateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *JobApplicantProfileCreateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *JobApplicantProfileCreateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *JobApplicantProfileCreateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetData

`func (o *JobApplicantProfileCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobApplicantProfileCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobApplicantProfileCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JobApplicantProfileCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobApplicantProfileCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobApplicantProfileCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *JobApplicantProfileCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *JobApplicantProfileCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *JobApplicantProfileCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *JobApplicantProfileCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *JobApplicantProfileCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *JobApplicantProfileCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *JobApplicantProfileCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *JobApplicantProfileCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *JobApplicantProfileCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *JobApplicantProfileCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *JobApplicantProfileCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *JobApplicantProfileCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *JobApplicantProfileCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *JobApplicantProfileCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *JobApplicantProfileCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *JobApplicantProfileCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *JobApplicantProfileCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *JobApplicantProfileCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *JobApplicantProfileCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *JobApplicantProfileCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *JobApplicantProfileCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *JobApplicantProfileCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *JobApplicantProfileCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *JobApplicantProfileCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *JobApplicantProfileCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *JobApplicantProfileCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *JobApplicantProfileCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *JobApplicantProfileCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *JobApplicantProfileCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *JobApplicantProfileCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *JobApplicantProfileCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *JobApplicantProfileCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *JobApplicantProfileCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *JobApplicantProfileCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *JobApplicantProfileCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *JobApplicantProfileCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *JobApplicantProfileCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *JobApplicantProfileCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *JobApplicantProfileCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *JobApplicantProfileCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *JobApplicantProfileCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *JobApplicantProfileCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *JobApplicantProfileCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *JobApplicantProfileCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *JobApplicantProfileCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *JobApplicantProfileCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *JobApplicantProfileCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *JobApplicantProfileCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *JobApplicantProfileCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *JobApplicantProfileCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *JobApplicantProfileCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *JobApplicantProfileCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *JobApplicantProfileCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *JobApplicantProfileCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *JobApplicantProfileCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *JobApplicantProfileCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *JobApplicantProfileCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *JobApplicantProfileCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *JobApplicantProfileCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *JobApplicantProfileCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *JobApplicantProfileCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *JobApplicantProfileCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *JobApplicantProfileCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *JobApplicantProfileCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *JobApplicantProfileCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *JobApplicantProfileCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *JobApplicantProfileCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *JobApplicantProfileCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *JobApplicantProfileCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *JobApplicantProfileCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *JobApplicantProfileCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *JobApplicantProfileCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *JobApplicantProfileCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *JobApplicantProfileCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *JobApplicantProfileCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *JobApplicantProfileCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *JobApplicantProfileCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *JobApplicantProfileCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *JobApplicantProfileCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *JobApplicantProfileCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *JobApplicantProfileCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *JobApplicantProfileCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *JobApplicantProfileCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *JobApplicantProfileCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *JobApplicantProfileCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *JobApplicantProfileCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *JobApplicantProfileCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *JobApplicantProfileCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *JobApplicantProfileCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *JobApplicantProfileCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *JobApplicantProfileCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *JobApplicantProfileCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *JobApplicantProfileCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *JobApplicantProfileCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *JobApplicantProfileCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *JobApplicantProfileCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *JobApplicantProfileCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *JobApplicantProfileCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *JobApplicantProfileCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *JobApplicantProfileCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *JobApplicantProfileCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *JobApplicantProfileCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *JobApplicantProfileCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *JobApplicantProfileCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *JobApplicantProfileCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *JobApplicantProfileCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *JobApplicantProfileCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *JobApplicantProfileCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *JobApplicantProfileCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *JobApplicantProfileCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *JobApplicantProfileCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *JobApplicantProfileCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *JobApplicantProfileCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *JobApplicantProfileCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetAvailableForHire

`func (o *JobApplicantProfileCreateDto) GetAvailableForHire() bool`

GetAvailableForHire returns the AvailableForHire field if non-nil, zero value otherwise.

### GetAvailableForHireOk

`func (o *JobApplicantProfileCreateDto) GetAvailableForHireOk() (*bool, bool)`

GetAvailableForHireOk returns a tuple with the AvailableForHire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForHire

`func (o *JobApplicantProfileCreateDto) SetAvailableForHire(v bool)`

SetAvailableForHire sets AvailableForHire field to given value.

### HasAvailableForHire

`func (o *JobApplicantProfileCreateDto) HasAvailableForHire() bool`

HasAvailableForHire returns a boolean if a field has been set.

### GetCareerLevel

`func (o *JobApplicantProfileCreateDto) GetCareerLevel() string`

GetCareerLevel returns the CareerLevel field if non-nil, zero value otherwise.

### GetCareerLevelOk

`func (o *JobApplicantProfileCreateDto) GetCareerLevelOk() (*string, bool)`

GetCareerLevelOk returns a tuple with the CareerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCareerLevel

`func (o *JobApplicantProfileCreateDto) SetCareerLevel(v string)`

SetCareerLevel sets CareerLevel field to given value.

### HasCareerLevel

`func (o *JobApplicantProfileCreateDto) HasCareerLevel() bool`

HasCareerLevel returns a boolean if a field has been set.

### GetExperienceInYears

`func (o *JobApplicantProfileCreateDto) GetExperienceInYears() int32`

GetExperienceInYears returns the ExperienceInYears field if non-nil, zero value otherwise.

### GetExperienceInYearsOk

`func (o *JobApplicantProfileCreateDto) GetExperienceInYearsOk() (*int32, bool)`

GetExperienceInYearsOk returns a tuple with the ExperienceInYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInYears

`func (o *JobApplicantProfileCreateDto) SetExperienceInYears(v int32)`

SetExperienceInYears sets ExperienceInYears field to given value.

### HasExperienceInYears

`func (o *JobApplicantProfileCreateDto) HasExperienceInYears() bool`

HasExperienceInYears returns a boolean if a field has been set.

### GetCurrentSalary

`func (o *JobApplicantProfileCreateDto) GetCurrentSalary() float64`

GetCurrentSalary returns the CurrentSalary field if non-nil, zero value otherwise.

### GetCurrentSalaryOk

`func (o *JobApplicantProfileCreateDto) GetCurrentSalaryOk() (*float64, bool)`

GetCurrentSalaryOk returns a tuple with the CurrentSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSalary

`func (o *JobApplicantProfileCreateDto) SetCurrentSalary(v float64)`

SetCurrentSalary sets CurrentSalary field to given value.

### HasCurrentSalary

`func (o *JobApplicantProfileCreateDto) HasCurrentSalary() bool`

HasCurrentSalary returns a boolean if a field has been set.

### GetMinSalaryExpectation

`func (o *JobApplicantProfileCreateDto) GetMinSalaryExpectation() float64`

GetMinSalaryExpectation returns the MinSalaryExpectation field if non-nil, zero value otherwise.

### GetMinSalaryExpectationOk

`func (o *JobApplicantProfileCreateDto) GetMinSalaryExpectationOk() (*float64, bool)`

GetMinSalaryExpectationOk returns a tuple with the MinSalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSalaryExpectation

`func (o *JobApplicantProfileCreateDto) SetMinSalaryExpectation(v float64)`

SetMinSalaryExpectation sets MinSalaryExpectation field to given value.

### HasMinSalaryExpectation

`func (o *JobApplicantProfileCreateDto) HasMinSalaryExpectation() bool`

HasMinSalaryExpectation returns a boolean if a field has been set.

### GetMaxSalaryExpectation

`func (o *JobApplicantProfileCreateDto) GetMaxSalaryExpectation() float64`

GetMaxSalaryExpectation returns the MaxSalaryExpectation field if non-nil, zero value otherwise.

### GetMaxSalaryExpectationOk

`func (o *JobApplicantProfileCreateDto) GetMaxSalaryExpectationOk() (*float64, bool)`

GetMaxSalaryExpectationOk returns a tuple with the MaxSalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSalaryExpectation

`func (o *JobApplicantProfileCreateDto) SetMaxSalaryExpectation(v float64)`

SetMaxSalaryExpectation sets MaxSalaryExpectation field to given value.

### HasMaxSalaryExpectation

`func (o *JobApplicantProfileCreateDto) HasMaxSalaryExpectation() bool`

HasMaxSalaryExpectation returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobApplicantProfileCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobApplicantProfileCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobApplicantProfileCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobApplicantProfileCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobApplicantProfileCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobApplicantProfileCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


