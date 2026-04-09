# AzureStorageIntegrationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**ConnectionString** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureStorageIntegrationOptions

`func NewAzureStorageIntegrationOptions() *AzureStorageIntegrationOptions`

NewAzureStorageIntegrationOptions instantiates a new AzureStorageIntegrationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageIntegrationOptionsWithDefaults

`func NewAzureStorageIntegrationOptionsWithDefaults() *AzureStorageIntegrationOptions`

NewAzureStorageIntegrationOptionsWithDefaults instantiates a new AzureStorageIntegrationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AzureStorageIntegrationOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AzureStorageIntegrationOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AzureStorageIntegrationOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AzureStorageIntegrationOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetConnectionString

`func (o *AzureStorageIntegrationOptions) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *AzureStorageIntegrationOptions) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *AzureStorageIntegrationOptions) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *AzureStorageIntegrationOptions) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### SetConnectionStringNil

`func (o *AzureStorageIntegrationOptions) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *AzureStorageIntegrationOptions) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


