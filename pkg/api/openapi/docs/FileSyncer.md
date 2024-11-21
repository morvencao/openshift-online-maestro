# FileSyncer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**ConsumerName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFileSyncer

`func NewFileSyncer() *FileSyncer`

NewFileSyncer instantiates a new FileSyncer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncerWithDefaults

`func NewFileSyncerWithDefaults() *FileSyncer`

NewFileSyncerWithDefaults instantiates a new FileSyncer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileSyncer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSyncer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSyncer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileSyncer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *FileSyncer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FileSyncer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FileSyncer) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FileSyncer) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *FileSyncer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FileSyncer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FileSyncer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FileSyncer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetConsumerName

`func (o *FileSyncer) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *FileSyncer) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *FileSyncer) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *FileSyncer) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetVersion

`func (o *FileSyncer) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileSyncer) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileSyncer) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FileSyncer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FileSyncer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileSyncer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileSyncer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileSyncer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileSyncer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileSyncer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileSyncer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileSyncer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *FileSyncer) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FileSyncer) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FileSyncer) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FileSyncer) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSpec

`func (o *FileSyncer) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FileSyncer) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FileSyncer) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FileSyncer) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *FileSyncer) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileSyncer) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileSyncer) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileSyncer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


