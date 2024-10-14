# AllOfstatsMxedgeMemoryStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int32** | The amount of memory, in kibibytes, that has been used more recently and is usually not reclaimed unless absolutely necessary. | [optional] [default to null]
**Available** | **int64** | An estimate of how much memory is available for starting new applications, without swapping. | [optional] [default to null]
**Buffers** | **int32** | The amount, in kibibytes, of temporary storage for raw disk blocks. | [optional] [default to null]
**Cached** | **int32** | The amount of physical RAM, in kibibytes, used as cache memory. | [optional] [default to null]
**Free** | **int64** | The amount of physical RAM, in kibibytes, left unused by the system | [optional] [default to null]
**Inactive** | **int32** | The amount of memory, in kibibytes, that has been used less recently and is more eligible to be reclaimed for other purposes. | [optional] [default to null]
**SwapCached** | **int32** | The amount of memory, in kibibytes, that has once been moved into swap, then back into the main memory, but still also remains in the swapfile. | [optional] [default to null]
**SwapFree** | **int32** | The total amount of swap free, in kibibytes. | [optional] [default to null]
**SwapTotal** | **int32** | The total amount of swap available, in kibibytes. | [optional] [default to null]
**Total** | **int64** | Total amount of usable RAM, in kibibytes, which is physical RAM minus a number of reserved bits and the kernel binary code | [optional] [default to null]
**Usage** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

