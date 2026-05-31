# SystemOverviewDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **string** |  | [optional] 
**OsDescription** | Pointer to **NullableString** |  | [optional] 
**MachineName** | Pointer to **NullableString** |  | [optional] 
**ProcessName** | Pointer to **NullableString** |  | [optional] 
**ProductVersion** | Pointer to **NullableString** |  | [optional] 
**PrivateMemoryMb** | Pointer to **int64** |  | [optional] 
**PagedMemoryMb** | Pointer to **int64** |  | [optional] 
**MemoryWorkingSetMb** | Pointer to **int64** |  | [optional] 
**IsDebugMode** | Pointer to **bool** |  | [optional] 
**IsDevMode** | Pointer to **bool** |  | [optional] 
**FrameworkDescription** | Pointer to **NullableString** |  | [optional] 
**RuntimeIdentifier** | Pointer to **NullableString** |  | [optional] 
**OsArchitecture** | Pointer to **NullableString** |  | [optional] 
**OsPlatform** | Pointer to **NullableString** |  | [optional] 
**ProcessArchitecture** | Pointer to **NullableString** |  | [optional] 
**UsersCount** | Pointer to **int32** |  | [optional] 
**OrdersCount** | Pointer to **int32** |  | [optional] 
**ContactsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSystemOverviewDto

`func NewSystemOverviewDto() *SystemOverviewDto`

NewSystemOverviewDto instantiates a new SystemOverviewDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemOverviewDtoWithDefaults

`func NewSystemOverviewDtoWithDefaults() *SystemOverviewDto`

NewSystemOverviewDtoWithDefaults instantiates a new SystemOverviewDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *SystemOverviewDto) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SystemOverviewDto) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SystemOverviewDto) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SystemOverviewDto) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetOsDescription

`func (o *SystemOverviewDto) GetOsDescription() string`

GetOsDescription returns the OsDescription field if non-nil, zero value otherwise.

### GetOsDescriptionOk

`func (o *SystemOverviewDto) GetOsDescriptionOk() (*string, bool)`

GetOsDescriptionOk returns a tuple with the OsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDescription

`func (o *SystemOverviewDto) SetOsDescription(v string)`

SetOsDescription sets OsDescription field to given value.

### HasOsDescription

`func (o *SystemOverviewDto) HasOsDescription() bool`

HasOsDescription returns a boolean if a field has been set.

### SetOsDescriptionNil

`func (o *SystemOverviewDto) SetOsDescriptionNil(b bool)`

 SetOsDescriptionNil sets the value for OsDescription to be an explicit nil

### UnsetOsDescription
`func (o *SystemOverviewDto) UnsetOsDescription()`

UnsetOsDescription ensures that no value is present for OsDescription, not even an explicit nil
### GetMachineName

`func (o *SystemOverviewDto) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *SystemOverviewDto) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *SystemOverviewDto) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *SystemOverviewDto) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### SetMachineNameNil

`func (o *SystemOverviewDto) SetMachineNameNil(b bool)`

 SetMachineNameNil sets the value for MachineName to be an explicit nil

### UnsetMachineName
`func (o *SystemOverviewDto) UnsetMachineName()`

UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
### GetProcessName

`func (o *SystemOverviewDto) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *SystemOverviewDto) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *SystemOverviewDto) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.

### HasProcessName

`func (o *SystemOverviewDto) HasProcessName() bool`

HasProcessName returns a boolean if a field has been set.

### SetProcessNameNil

`func (o *SystemOverviewDto) SetProcessNameNil(b bool)`

 SetProcessNameNil sets the value for ProcessName to be an explicit nil

### UnsetProcessName
`func (o *SystemOverviewDto) UnsetProcessName()`

UnsetProcessName ensures that no value is present for ProcessName, not even an explicit nil
### GetProductVersion

`func (o *SystemOverviewDto) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *SystemOverviewDto) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *SystemOverviewDto) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *SystemOverviewDto) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### SetProductVersionNil

`func (o *SystemOverviewDto) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *SystemOverviewDto) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetPrivateMemoryMb

`func (o *SystemOverviewDto) GetPrivateMemoryMb() int64`

GetPrivateMemoryMb returns the PrivateMemoryMb field if non-nil, zero value otherwise.

### GetPrivateMemoryMbOk

`func (o *SystemOverviewDto) GetPrivateMemoryMbOk() (*int64, bool)`

GetPrivateMemoryMbOk returns a tuple with the PrivateMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMemoryMb

`func (o *SystemOverviewDto) SetPrivateMemoryMb(v int64)`

SetPrivateMemoryMb sets PrivateMemoryMb field to given value.

### HasPrivateMemoryMb

`func (o *SystemOverviewDto) HasPrivateMemoryMb() bool`

HasPrivateMemoryMb returns a boolean if a field has been set.

### GetPagedMemoryMb

`func (o *SystemOverviewDto) GetPagedMemoryMb() int64`

GetPagedMemoryMb returns the PagedMemoryMb field if non-nil, zero value otherwise.

### GetPagedMemoryMbOk

`func (o *SystemOverviewDto) GetPagedMemoryMbOk() (*int64, bool)`

GetPagedMemoryMbOk returns a tuple with the PagedMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagedMemoryMb

`func (o *SystemOverviewDto) SetPagedMemoryMb(v int64)`

SetPagedMemoryMb sets PagedMemoryMb field to given value.

### HasPagedMemoryMb

`func (o *SystemOverviewDto) HasPagedMemoryMb() bool`

HasPagedMemoryMb returns a boolean if a field has been set.

### GetMemoryWorkingSetMb

`func (o *SystemOverviewDto) GetMemoryWorkingSetMb() int64`

GetMemoryWorkingSetMb returns the MemoryWorkingSetMb field if non-nil, zero value otherwise.

### GetMemoryWorkingSetMbOk

`func (o *SystemOverviewDto) GetMemoryWorkingSetMbOk() (*int64, bool)`

GetMemoryWorkingSetMbOk returns a tuple with the MemoryWorkingSetMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryWorkingSetMb

`func (o *SystemOverviewDto) SetMemoryWorkingSetMb(v int64)`

SetMemoryWorkingSetMb sets MemoryWorkingSetMb field to given value.

### HasMemoryWorkingSetMb

`func (o *SystemOverviewDto) HasMemoryWorkingSetMb() bool`

HasMemoryWorkingSetMb returns a boolean if a field has been set.

### GetIsDebugMode

`func (o *SystemOverviewDto) GetIsDebugMode() bool`

GetIsDebugMode returns the IsDebugMode field if non-nil, zero value otherwise.

### GetIsDebugModeOk

`func (o *SystemOverviewDto) GetIsDebugModeOk() (*bool, bool)`

GetIsDebugModeOk returns a tuple with the IsDebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebugMode

`func (o *SystemOverviewDto) SetIsDebugMode(v bool)`

SetIsDebugMode sets IsDebugMode field to given value.

### HasIsDebugMode

`func (o *SystemOverviewDto) HasIsDebugMode() bool`

HasIsDebugMode returns a boolean if a field has been set.

### GetIsDevMode

`func (o *SystemOverviewDto) GetIsDevMode() bool`

GetIsDevMode returns the IsDevMode field if non-nil, zero value otherwise.

### GetIsDevModeOk

`func (o *SystemOverviewDto) GetIsDevModeOk() (*bool, bool)`

GetIsDevModeOk returns a tuple with the IsDevMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDevMode

`func (o *SystemOverviewDto) SetIsDevMode(v bool)`

SetIsDevMode sets IsDevMode field to given value.

### HasIsDevMode

`func (o *SystemOverviewDto) HasIsDevMode() bool`

HasIsDevMode returns a boolean if a field has been set.

### GetFrameworkDescription

`func (o *SystemOverviewDto) GetFrameworkDescription() string`

GetFrameworkDescription returns the FrameworkDescription field if non-nil, zero value otherwise.

### GetFrameworkDescriptionOk

`func (o *SystemOverviewDto) GetFrameworkDescriptionOk() (*string, bool)`

GetFrameworkDescriptionOk returns a tuple with the FrameworkDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkDescription

`func (o *SystemOverviewDto) SetFrameworkDescription(v string)`

SetFrameworkDescription sets FrameworkDescription field to given value.

### HasFrameworkDescription

`func (o *SystemOverviewDto) HasFrameworkDescription() bool`

HasFrameworkDescription returns a boolean if a field has been set.

### SetFrameworkDescriptionNil

`func (o *SystemOverviewDto) SetFrameworkDescriptionNil(b bool)`

 SetFrameworkDescriptionNil sets the value for FrameworkDescription to be an explicit nil

### UnsetFrameworkDescription
`func (o *SystemOverviewDto) UnsetFrameworkDescription()`

UnsetFrameworkDescription ensures that no value is present for FrameworkDescription, not even an explicit nil
### GetRuntimeIdentifier

`func (o *SystemOverviewDto) GetRuntimeIdentifier() string`

GetRuntimeIdentifier returns the RuntimeIdentifier field if non-nil, zero value otherwise.

### GetRuntimeIdentifierOk

`func (o *SystemOverviewDto) GetRuntimeIdentifierOk() (*string, bool)`

GetRuntimeIdentifierOk returns a tuple with the RuntimeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeIdentifier

`func (o *SystemOverviewDto) SetRuntimeIdentifier(v string)`

SetRuntimeIdentifier sets RuntimeIdentifier field to given value.

### HasRuntimeIdentifier

`func (o *SystemOverviewDto) HasRuntimeIdentifier() bool`

HasRuntimeIdentifier returns a boolean if a field has been set.

### SetRuntimeIdentifierNil

`func (o *SystemOverviewDto) SetRuntimeIdentifierNil(b bool)`

 SetRuntimeIdentifierNil sets the value for RuntimeIdentifier to be an explicit nil

### UnsetRuntimeIdentifier
`func (o *SystemOverviewDto) UnsetRuntimeIdentifier()`

UnsetRuntimeIdentifier ensures that no value is present for RuntimeIdentifier, not even an explicit nil
### GetOsArchitecture

`func (o *SystemOverviewDto) GetOsArchitecture() string`

GetOsArchitecture returns the OsArchitecture field if non-nil, zero value otherwise.

### GetOsArchitectureOk

`func (o *SystemOverviewDto) GetOsArchitectureOk() (*string, bool)`

GetOsArchitectureOk returns a tuple with the OsArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArchitecture

`func (o *SystemOverviewDto) SetOsArchitecture(v string)`

SetOsArchitecture sets OsArchitecture field to given value.

### HasOsArchitecture

`func (o *SystemOverviewDto) HasOsArchitecture() bool`

HasOsArchitecture returns a boolean if a field has been set.

### SetOsArchitectureNil

`func (o *SystemOverviewDto) SetOsArchitectureNil(b bool)`

 SetOsArchitectureNil sets the value for OsArchitecture to be an explicit nil

### UnsetOsArchitecture
`func (o *SystemOverviewDto) UnsetOsArchitecture()`

UnsetOsArchitecture ensures that no value is present for OsArchitecture, not even an explicit nil
### GetOsPlatform

`func (o *SystemOverviewDto) GetOsPlatform() string`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *SystemOverviewDto) GetOsPlatformOk() (*string, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *SystemOverviewDto) SetOsPlatform(v string)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *SystemOverviewDto) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### SetOsPlatformNil

`func (o *SystemOverviewDto) SetOsPlatformNil(b bool)`

 SetOsPlatformNil sets the value for OsPlatform to be an explicit nil

### UnsetOsPlatform
`func (o *SystemOverviewDto) UnsetOsPlatform()`

UnsetOsPlatform ensures that no value is present for OsPlatform, not even an explicit nil
### GetProcessArchitecture

`func (o *SystemOverviewDto) GetProcessArchitecture() string`

GetProcessArchitecture returns the ProcessArchitecture field if non-nil, zero value otherwise.

### GetProcessArchitectureOk

`func (o *SystemOverviewDto) GetProcessArchitectureOk() (*string, bool)`

GetProcessArchitectureOk returns a tuple with the ProcessArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessArchitecture

`func (o *SystemOverviewDto) SetProcessArchitecture(v string)`

SetProcessArchitecture sets ProcessArchitecture field to given value.

### HasProcessArchitecture

`func (o *SystemOverviewDto) HasProcessArchitecture() bool`

HasProcessArchitecture returns a boolean if a field has been set.

### SetProcessArchitectureNil

`func (o *SystemOverviewDto) SetProcessArchitectureNil(b bool)`

 SetProcessArchitectureNil sets the value for ProcessArchitecture to be an explicit nil

### UnsetProcessArchitecture
`func (o *SystemOverviewDto) UnsetProcessArchitecture()`

UnsetProcessArchitecture ensures that no value is present for ProcessArchitecture, not even an explicit nil
### GetUsersCount

`func (o *SystemOverviewDto) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *SystemOverviewDto) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *SystemOverviewDto) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.

### HasUsersCount

`func (o *SystemOverviewDto) HasUsersCount() bool`

HasUsersCount returns a boolean if a field has been set.

### GetOrdersCount

`func (o *SystemOverviewDto) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *SystemOverviewDto) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *SystemOverviewDto) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *SystemOverviewDto) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### GetContactsCount

`func (o *SystemOverviewDto) GetContactsCount() int32`

GetContactsCount returns the ContactsCount field if non-nil, zero value otherwise.

### GetContactsCountOk

`func (o *SystemOverviewDto) GetContactsCountOk() (*int32, bool)`

GetContactsCountOk returns a tuple with the ContactsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactsCount

`func (o *SystemOverviewDto) SetContactsCount(v int32)`

SetContactsCount sets ContactsCount field to given value.

### HasContactsCount

`func (o *SystemOverviewDto) HasContactsCount() bool`

HasContactsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


