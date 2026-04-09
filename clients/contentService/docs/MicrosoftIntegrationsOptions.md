# MicrosoftIntegrationsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Azure** | Pointer to [**MicrosoftAzureIntegrationOptions**](MicrosoftAzureIntegrationOptions.md) |  | [optional] 

## Methods

### NewMicrosoftIntegrationsOptions

`func NewMicrosoftIntegrationsOptions() *MicrosoftIntegrationsOptions`

NewMicrosoftIntegrationsOptions instantiates a new MicrosoftIntegrationsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftIntegrationsOptionsWithDefaults

`func NewMicrosoftIntegrationsOptionsWithDefaults() *MicrosoftIntegrationsOptions`

NewMicrosoftIntegrationsOptionsWithDefaults instantiates a new MicrosoftIntegrationsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *MicrosoftIntegrationsOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MicrosoftIntegrationsOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MicrosoftIntegrationsOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MicrosoftIntegrationsOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAzure

`func (o *MicrosoftIntegrationsOptions) GetAzure() MicrosoftAzureIntegrationOptions`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *MicrosoftIntegrationsOptions) GetAzureOk() (*MicrosoftAzureIntegrationOptions, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *MicrosoftIntegrationsOptions) SetAzure(v MicrosoftAzureIntegrationOptions)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *MicrosoftIntegrationsOptions) HasAzure() bool`

HasAzure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


