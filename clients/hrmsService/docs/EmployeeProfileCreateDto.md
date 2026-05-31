# EmployeeProfileCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
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
**GrossPay** | Pointer to **float64** |  | [optional] 
**NetSalary** | Pointer to **float64** |  | [optional] 
**PayrollCurrency** | Pointer to **NullableString** |  | [optional] 
**MaxWorkHoursPerDay** | Pointer to **int32** |  | [optional] 
**JobTitleId** | Pointer to **NullableString** |  | [optional] 
**EmployeeTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmployeeProfileCreateDto

`func NewEmployeeProfileCreateDto() *EmployeeProfileCreateDto`

NewEmployeeProfileCreateDto instantiates a new EmployeeProfileCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeProfileCreateDtoWithDefaults

`func NewEmployeeProfileCreateDtoWithDefaults() *EmployeeProfileCreateDto`

NewEmployeeProfileCreateDtoWithDefaults instantiates a new EmployeeProfileCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeProfileCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeProfileCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeProfileCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeProfileCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EmployeeProfileCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmployeeProfileCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmployeeProfileCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EmployeeProfileCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *EmployeeProfileCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmployeeProfileCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmployeeProfileCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmployeeProfileCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EmployeeProfileCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EmployeeProfileCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContactId

`func (o *EmployeeProfileCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *EmployeeProfileCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *EmployeeProfileCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *EmployeeProfileCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *EmployeeProfileCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *EmployeeProfileCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetAbout

`func (o *EmployeeProfileCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *EmployeeProfileCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *EmployeeProfileCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *EmployeeProfileCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *EmployeeProfileCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *EmployeeProfileCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetAvatarUrl

`func (o *EmployeeProfileCreateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *EmployeeProfileCreateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *EmployeeProfileCreateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *EmployeeProfileCreateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *EmployeeProfileCreateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *EmployeeProfileCreateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetData

`func (o *EmployeeProfileCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmployeeProfileCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmployeeProfileCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EmployeeProfileCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EmployeeProfileCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EmployeeProfileCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *EmployeeProfileCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *EmployeeProfileCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *EmployeeProfileCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *EmployeeProfileCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *EmployeeProfileCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *EmployeeProfileCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *EmployeeProfileCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *EmployeeProfileCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *EmployeeProfileCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *EmployeeProfileCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *EmployeeProfileCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *EmployeeProfileCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *EmployeeProfileCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *EmployeeProfileCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *EmployeeProfileCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *EmployeeProfileCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *EmployeeProfileCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *EmployeeProfileCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *EmployeeProfileCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *EmployeeProfileCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *EmployeeProfileCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *EmployeeProfileCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *EmployeeProfileCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *EmployeeProfileCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *EmployeeProfileCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *EmployeeProfileCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *EmployeeProfileCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *EmployeeProfileCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *EmployeeProfileCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *EmployeeProfileCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *EmployeeProfileCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *EmployeeProfileCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *EmployeeProfileCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *EmployeeProfileCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *EmployeeProfileCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *EmployeeProfileCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *EmployeeProfileCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *EmployeeProfileCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *EmployeeProfileCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *EmployeeProfileCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *EmployeeProfileCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *EmployeeProfileCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *EmployeeProfileCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *EmployeeProfileCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *EmployeeProfileCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *EmployeeProfileCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *EmployeeProfileCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *EmployeeProfileCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *EmployeeProfileCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *EmployeeProfileCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *EmployeeProfileCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *EmployeeProfileCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *EmployeeProfileCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *EmployeeProfileCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *EmployeeProfileCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *EmployeeProfileCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *EmployeeProfileCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *EmployeeProfileCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *EmployeeProfileCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *EmployeeProfileCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *EmployeeProfileCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *EmployeeProfileCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *EmployeeProfileCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *EmployeeProfileCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *EmployeeProfileCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *EmployeeProfileCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *EmployeeProfileCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *EmployeeProfileCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *EmployeeProfileCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *EmployeeProfileCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *EmployeeProfileCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *EmployeeProfileCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *EmployeeProfileCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *EmployeeProfileCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *EmployeeProfileCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *EmployeeProfileCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *EmployeeProfileCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *EmployeeProfileCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *EmployeeProfileCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *EmployeeProfileCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *EmployeeProfileCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *EmployeeProfileCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *EmployeeProfileCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *EmployeeProfileCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *EmployeeProfileCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *EmployeeProfileCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *EmployeeProfileCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *EmployeeProfileCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *EmployeeProfileCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *EmployeeProfileCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *EmployeeProfileCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *EmployeeProfileCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *EmployeeProfileCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *EmployeeProfileCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *EmployeeProfileCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *EmployeeProfileCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *EmployeeProfileCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *EmployeeProfileCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *EmployeeProfileCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *EmployeeProfileCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *EmployeeProfileCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *EmployeeProfileCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *EmployeeProfileCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *EmployeeProfileCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *EmployeeProfileCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *EmployeeProfileCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *EmployeeProfileCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *EmployeeProfileCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *EmployeeProfileCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *EmployeeProfileCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *EmployeeProfileCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *EmployeeProfileCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *EmployeeProfileCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *EmployeeProfileCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetGrossPay

`func (o *EmployeeProfileCreateDto) GetGrossPay() float64`

GetGrossPay returns the GrossPay field if non-nil, zero value otherwise.

### GetGrossPayOk

`func (o *EmployeeProfileCreateDto) GetGrossPayOk() (*float64, bool)`

GetGrossPayOk returns a tuple with the GrossPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPay

`func (o *EmployeeProfileCreateDto) SetGrossPay(v float64)`

SetGrossPay sets GrossPay field to given value.

### HasGrossPay

`func (o *EmployeeProfileCreateDto) HasGrossPay() bool`

HasGrossPay returns a boolean if a field has been set.

### GetNetSalary

`func (o *EmployeeProfileCreateDto) GetNetSalary() float64`

GetNetSalary returns the NetSalary field if non-nil, zero value otherwise.

### GetNetSalaryOk

`func (o *EmployeeProfileCreateDto) GetNetSalaryOk() (*float64, bool)`

GetNetSalaryOk returns a tuple with the NetSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSalary

`func (o *EmployeeProfileCreateDto) SetNetSalary(v float64)`

SetNetSalary sets NetSalary field to given value.

### HasNetSalary

`func (o *EmployeeProfileCreateDto) HasNetSalary() bool`

HasNetSalary returns a boolean if a field has been set.

### GetPayrollCurrency

`func (o *EmployeeProfileCreateDto) GetPayrollCurrency() string`

GetPayrollCurrency returns the PayrollCurrency field if non-nil, zero value otherwise.

### GetPayrollCurrencyOk

`func (o *EmployeeProfileCreateDto) GetPayrollCurrencyOk() (*string, bool)`

GetPayrollCurrencyOk returns a tuple with the PayrollCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollCurrency

`func (o *EmployeeProfileCreateDto) SetPayrollCurrency(v string)`

SetPayrollCurrency sets PayrollCurrency field to given value.

### HasPayrollCurrency

`func (o *EmployeeProfileCreateDto) HasPayrollCurrency() bool`

HasPayrollCurrency returns a boolean if a field has been set.

### SetPayrollCurrencyNil

`func (o *EmployeeProfileCreateDto) SetPayrollCurrencyNil(b bool)`

 SetPayrollCurrencyNil sets the value for PayrollCurrency to be an explicit nil

### UnsetPayrollCurrency
`func (o *EmployeeProfileCreateDto) UnsetPayrollCurrency()`

UnsetPayrollCurrency ensures that no value is present for PayrollCurrency, not even an explicit nil
### GetMaxWorkHoursPerDay

`func (o *EmployeeProfileCreateDto) GetMaxWorkHoursPerDay() int32`

GetMaxWorkHoursPerDay returns the MaxWorkHoursPerDay field if non-nil, zero value otherwise.

### GetMaxWorkHoursPerDayOk

`func (o *EmployeeProfileCreateDto) GetMaxWorkHoursPerDayOk() (*int32, bool)`

GetMaxWorkHoursPerDayOk returns a tuple with the MaxWorkHoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkHoursPerDay

`func (o *EmployeeProfileCreateDto) SetMaxWorkHoursPerDay(v int32)`

SetMaxWorkHoursPerDay sets MaxWorkHoursPerDay field to given value.

### HasMaxWorkHoursPerDay

`func (o *EmployeeProfileCreateDto) HasMaxWorkHoursPerDay() bool`

HasMaxWorkHoursPerDay returns a boolean if a field has been set.

### GetJobTitleId

`func (o *EmployeeProfileCreateDto) GetJobTitleId() string`

GetJobTitleId returns the JobTitleId field if non-nil, zero value otherwise.

### GetJobTitleIdOk

`func (o *EmployeeProfileCreateDto) GetJobTitleIdOk() (*string, bool)`

GetJobTitleIdOk returns a tuple with the JobTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitleId

`func (o *EmployeeProfileCreateDto) SetJobTitleId(v string)`

SetJobTitleId sets JobTitleId field to given value.

### HasJobTitleId

`func (o *EmployeeProfileCreateDto) HasJobTitleId() bool`

HasJobTitleId returns a boolean if a field has been set.

### SetJobTitleIdNil

`func (o *EmployeeProfileCreateDto) SetJobTitleIdNil(b bool)`

 SetJobTitleIdNil sets the value for JobTitleId to be an explicit nil

### UnsetJobTitleId
`func (o *EmployeeProfileCreateDto) UnsetJobTitleId()`

UnsetJobTitleId ensures that no value is present for JobTitleId, not even an explicit nil
### GetEmployeeTypeId

`func (o *EmployeeProfileCreateDto) GetEmployeeTypeId() string`

GetEmployeeTypeId returns the EmployeeTypeId field if non-nil, zero value otherwise.

### GetEmployeeTypeIdOk

`func (o *EmployeeProfileCreateDto) GetEmployeeTypeIdOk() (*string, bool)`

GetEmployeeTypeIdOk returns a tuple with the EmployeeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeTypeId

`func (o *EmployeeProfileCreateDto) SetEmployeeTypeId(v string)`

SetEmployeeTypeId sets EmployeeTypeId field to given value.

### HasEmployeeTypeId

`func (o *EmployeeProfileCreateDto) HasEmployeeTypeId() bool`

HasEmployeeTypeId returns a boolean if a field has been set.

### SetEmployeeTypeIdNil

`func (o *EmployeeProfileCreateDto) SetEmployeeTypeIdNil(b bool)`

 SetEmployeeTypeIdNil sets the value for EmployeeTypeId to be an explicit nil

### UnsetEmployeeTypeId
`func (o *EmployeeProfileCreateDto) UnsetEmployeeTypeId()`

UnsetEmployeeTypeId ensures that no value is present for EmployeeTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


