# SigningProfileUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSigningProfileUpdateDto

`func NewSigningProfileUpdateDto() *SigningProfileUpdateDto`

NewSigningProfileUpdateDto instantiates a new SigningProfileUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningProfileUpdateDtoWithDefaults

`func NewSigningProfileUpdateDtoWithDefaults() *SigningProfileUpdateDto`

NewSigningProfileUpdateDtoWithDefaults instantiates a new SigningProfileUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SigningProfileUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningProfileUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningProfileUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningProfileUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SigningProfileUpdateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SigningProfileUpdateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContactId

`func (o *SigningProfileUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningProfileUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningProfileUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SigningProfileUpdateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SigningProfileUpdateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SigningProfileUpdateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetAbout

`func (o *SigningProfileUpdateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *SigningProfileUpdateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *SigningProfileUpdateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *SigningProfileUpdateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *SigningProfileUpdateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *SigningProfileUpdateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetAvatarUrl

`func (o *SigningProfileUpdateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *SigningProfileUpdateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *SigningProfileUpdateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *SigningProfileUpdateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *SigningProfileUpdateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *SigningProfileUpdateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetData

`func (o *SigningProfileUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SigningProfileUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SigningProfileUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SigningProfileUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SigningProfileUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SigningProfileUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *SigningProfileUpdateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *SigningProfileUpdateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *SigningProfileUpdateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *SigningProfileUpdateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *SigningProfileUpdateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *SigningProfileUpdateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *SigningProfileUpdateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *SigningProfileUpdateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *SigningProfileUpdateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *SigningProfileUpdateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *SigningProfileUpdateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *SigningProfileUpdateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *SigningProfileUpdateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *SigningProfileUpdateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *SigningProfileUpdateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *SigningProfileUpdateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *SigningProfileUpdateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *SigningProfileUpdateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *SigningProfileUpdateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *SigningProfileUpdateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *SigningProfileUpdateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *SigningProfileUpdateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *SigningProfileUpdateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *SigningProfileUpdateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *SigningProfileUpdateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *SigningProfileUpdateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *SigningProfileUpdateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *SigningProfileUpdateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *SigningProfileUpdateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *SigningProfileUpdateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *SigningProfileUpdateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *SigningProfileUpdateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *SigningProfileUpdateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *SigningProfileUpdateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *SigningProfileUpdateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *SigningProfileUpdateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *SigningProfileUpdateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *SigningProfileUpdateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *SigningProfileUpdateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *SigningProfileUpdateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *SigningProfileUpdateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *SigningProfileUpdateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *SigningProfileUpdateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *SigningProfileUpdateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *SigningProfileUpdateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *SigningProfileUpdateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *SigningProfileUpdateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *SigningProfileUpdateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *SigningProfileUpdateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *SigningProfileUpdateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *SigningProfileUpdateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *SigningProfileUpdateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *SigningProfileUpdateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *SigningProfileUpdateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *SigningProfileUpdateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *SigningProfileUpdateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *SigningProfileUpdateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *SigningProfileUpdateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *SigningProfileUpdateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *SigningProfileUpdateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *SigningProfileUpdateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *SigningProfileUpdateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *SigningProfileUpdateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *SigningProfileUpdateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *SigningProfileUpdateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *SigningProfileUpdateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *SigningProfileUpdateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *SigningProfileUpdateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *SigningProfileUpdateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *SigningProfileUpdateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *SigningProfileUpdateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *SigningProfileUpdateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *SigningProfileUpdateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *SigningProfileUpdateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *SigningProfileUpdateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *SigningProfileUpdateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *SigningProfileUpdateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *SigningProfileUpdateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *SigningProfileUpdateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *SigningProfileUpdateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *SigningProfileUpdateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *SigningProfileUpdateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *SigningProfileUpdateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *SigningProfileUpdateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *SigningProfileUpdateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *SigningProfileUpdateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *SigningProfileUpdateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *SigningProfileUpdateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *SigningProfileUpdateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *SigningProfileUpdateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *SigningProfileUpdateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *SigningProfileUpdateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *SigningProfileUpdateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *SigningProfileUpdateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *SigningProfileUpdateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *SigningProfileUpdateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *SigningProfileUpdateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *SigningProfileUpdateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *SigningProfileUpdateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *SigningProfileUpdateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *SigningProfileUpdateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *SigningProfileUpdateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *SigningProfileUpdateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *SigningProfileUpdateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *SigningProfileUpdateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *SigningProfileUpdateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *SigningProfileUpdateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *SigningProfileUpdateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *SigningProfileUpdateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *SigningProfileUpdateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *SigningProfileUpdateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *SigningProfileUpdateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *SigningProfileUpdateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *SigningProfileUpdateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetSignatureFormat

`func (o *SigningProfileUpdateDto) GetSignatureFormat() string`

GetSignatureFormat returns the SignatureFormat field if non-nil, zero value otherwise.

### GetSignatureFormatOk

`func (o *SigningProfileUpdateDto) GetSignatureFormatOk() (*string, bool)`

GetSignatureFormatOk returns a tuple with the SignatureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFormat

`func (o *SigningProfileUpdateDto) SetSignatureFormat(v string)`

SetSignatureFormat sets SignatureFormat field to given value.

### HasSignatureFormat

`func (o *SigningProfileUpdateDto) HasSignatureFormat() bool`

HasSignatureFormat returns a boolean if a field has been set.

### SetSignatureFormatNil

`func (o *SigningProfileUpdateDto) SetSignatureFormatNil(b bool)`

 SetSignatureFormatNil sets the value for SignatureFormat to be an explicit nil

### UnsetSignatureFormat
`func (o *SigningProfileUpdateDto) UnsetSignatureFormat()`

UnsetSignatureFormat ensures that no value is present for SignatureFormat, not even an explicit nil
### GetSignaturePurpose

`func (o *SigningProfileUpdateDto) GetSignaturePurpose() string`

GetSignaturePurpose returns the SignaturePurpose field if non-nil, zero value otherwise.

### GetSignaturePurposeOk

`func (o *SigningProfileUpdateDto) GetSignaturePurposeOk() (*string, bool)`

GetSignaturePurposeOk returns a tuple with the SignaturePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturePurpose

`func (o *SigningProfileUpdateDto) SetSignaturePurpose(v string)`

SetSignaturePurpose sets SignaturePurpose field to given value.

### HasSignaturePurpose

`func (o *SigningProfileUpdateDto) HasSignaturePurpose() bool`

HasSignaturePurpose returns a boolean if a field has been set.

### SetSignaturePurposeNil

`func (o *SigningProfileUpdateDto) SetSignaturePurposeNil(b bool)`

 SetSignaturePurposeNil sets the value for SignaturePurpose to be an explicit nil

### UnsetSignaturePurpose
`func (o *SigningProfileUpdateDto) UnsetSignaturePurpose()`

UnsetSignaturePurpose ensures that no value is present for SignaturePurpose, not even an explicit nil
### GetDigestAlgorithm

`func (o *SigningProfileUpdateDto) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SigningProfileUpdateDto) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SigningProfileUpdateDto) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SigningProfileUpdateDto) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### SetDigestAlgorithmNil

`func (o *SigningProfileUpdateDto) SetDigestAlgorithmNil(b bool)`

 SetDigestAlgorithmNil sets the value for DigestAlgorithm to be an explicit nil

### UnsetDigestAlgorithm
`func (o *SigningProfileUpdateDto) UnsetDigestAlgorithm()`

UnsetDigestAlgorithm ensures that no value is present for DigestAlgorithm, not even an explicit nil
### GetSignatureAlgorithm

`func (o *SigningProfileUpdateDto) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SigningProfileUpdateDto) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SigningProfileUpdateDto) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SigningProfileUpdateDto) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### SetSignatureAlgorithmNil

`func (o *SigningProfileUpdateDto) SetSignatureAlgorithmNil(b bool)`

 SetSignatureAlgorithmNil sets the value for SignatureAlgorithm to be an explicit nil

### UnsetSignatureAlgorithm
`func (o *SigningProfileUpdateDto) UnsetSignatureAlgorithm()`

UnsetSignatureAlgorithm ensures that no value is present for SignatureAlgorithm, not even an explicit nil
### GetCanonicalizationAlgorithm

`func (o *SigningProfileUpdateDto) GetCanonicalizationAlgorithm() string`

GetCanonicalizationAlgorithm returns the CanonicalizationAlgorithm field if non-nil, zero value otherwise.

### GetCanonicalizationAlgorithmOk

`func (o *SigningProfileUpdateDto) GetCanonicalizationAlgorithmOk() (*string, bool)`

GetCanonicalizationAlgorithmOk returns a tuple with the CanonicalizationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizationAlgorithm

`func (o *SigningProfileUpdateDto) SetCanonicalizationAlgorithm(v string)`

SetCanonicalizationAlgorithm sets CanonicalizationAlgorithm field to given value.

### HasCanonicalizationAlgorithm

`func (o *SigningProfileUpdateDto) HasCanonicalizationAlgorithm() bool`

HasCanonicalizationAlgorithm returns a boolean if a field has been set.

### SetCanonicalizationAlgorithmNil

`func (o *SigningProfileUpdateDto) SetCanonicalizationAlgorithmNil(b bool)`

 SetCanonicalizationAlgorithmNil sets the value for CanonicalizationAlgorithm to be an explicit nil

### UnsetCanonicalizationAlgorithm
`func (o *SigningProfileUpdateDto) UnsetCanonicalizationAlgorithm()`

UnsetCanonicalizationAlgorithm ensures that no value is present for CanonicalizationAlgorithm, not even an explicit nil
### GetPolicyIdentifier

`func (o *SigningProfileUpdateDto) GetPolicyIdentifier() string`

GetPolicyIdentifier returns the PolicyIdentifier field if non-nil, zero value otherwise.

### GetPolicyIdentifierOk

`func (o *SigningProfileUpdateDto) GetPolicyIdentifierOk() (*string, bool)`

GetPolicyIdentifierOk returns a tuple with the PolicyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifier

`func (o *SigningProfileUpdateDto) SetPolicyIdentifier(v string)`

SetPolicyIdentifier sets PolicyIdentifier field to given value.

### HasPolicyIdentifier

`func (o *SigningProfileUpdateDto) HasPolicyIdentifier() bool`

HasPolicyIdentifier returns a boolean if a field has been set.

### SetPolicyIdentifierNil

`func (o *SigningProfileUpdateDto) SetPolicyIdentifierNil(b bool)`

 SetPolicyIdentifierNil sets the value for PolicyIdentifier to be an explicit nil

### UnsetPolicyIdentifier
`func (o *SigningProfileUpdateDto) UnsetPolicyIdentifier()`

UnsetPolicyIdentifier ensures that no value is present for PolicyIdentifier, not even an explicit nil
### GetPolicyUri

`func (o *SigningProfileUpdateDto) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *SigningProfileUpdateDto) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *SigningProfileUpdateDto) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *SigningProfileUpdateDto) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### SetPolicyUriNil

`func (o *SigningProfileUpdateDto) SetPolicyUriNil(b bool)`

 SetPolicyUriNil sets the value for PolicyUri to be an explicit nil

### UnsetPolicyUri
`func (o *SigningProfileUpdateDto) UnsetPolicyUri()`

UnsetPolicyUri ensures that no value is present for PolicyUri, not even an explicit nil
### GetAuthorityProfile

`func (o *SigningProfileUpdateDto) GetAuthorityProfile() string`

GetAuthorityProfile returns the AuthorityProfile field if non-nil, zero value otherwise.

### GetAuthorityProfileOk

`func (o *SigningProfileUpdateDto) GetAuthorityProfileOk() (*string, bool)`

GetAuthorityProfileOk returns a tuple with the AuthorityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityProfile

`func (o *SigningProfileUpdateDto) SetAuthorityProfile(v string)`

SetAuthorityProfile sets AuthorityProfile field to given value.

### HasAuthorityProfile

`func (o *SigningProfileUpdateDto) HasAuthorityProfile() bool`

HasAuthorityProfile returns a boolean if a field has been set.

### SetAuthorityProfileNil

`func (o *SigningProfileUpdateDto) SetAuthorityProfileNil(b bool)`

 SetAuthorityProfileNil sets the value for AuthorityProfile to be an explicit nil

### UnsetAuthorityProfile
`func (o *SigningProfileUpdateDto) UnsetAuthorityProfile()`

UnsetAuthorityProfile ensures that no value is present for AuthorityProfile, not even an explicit nil
### GetIsActive

`func (o *SigningProfileUpdateDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SigningProfileUpdateDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SigningProfileUpdateDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SigningProfileUpdateDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *SigningProfileUpdateDto) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *SigningProfileUpdateDto) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetDefaultForDocumentType

`func (o *SigningProfileUpdateDto) GetDefaultForDocumentType() string`

GetDefaultForDocumentType returns the DefaultForDocumentType field if non-nil, zero value otherwise.

### GetDefaultForDocumentTypeOk

`func (o *SigningProfileUpdateDto) GetDefaultForDocumentTypeOk() (*string, bool)`

GetDefaultForDocumentTypeOk returns a tuple with the DefaultForDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForDocumentType

`func (o *SigningProfileUpdateDto) SetDefaultForDocumentType(v string)`

SetDefaultForDocumentType sets DefaultForDocumentType field to given value.

### HasDefaultForDocumentType

`func (o *SigningProfileUpdateDto) HasDefaultForDocumentType() bool`

HasDefaultForDocumentType returns a boolean if a field has been set.

### SetDefaultForDocumentTypeNil

`func (o *SigningProfileUpdateDto) SetDefaultForDocumentTypeNil(b bool)`

 SetDefaultForDocumentTypeNil sets the value for DefaultForDocumentType to be an explicit nil

### UnsetDefaultForDocumentType
`func (o *SigningProfileUpdateDto) UnsetDefaultForDocumentType()`

UnsetDefaultForDocumentType ensures that no value is present for DefaultForDocumentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


