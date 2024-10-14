# ModuleStatItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupVersion** | **string** |  | [optional] [default to null]
**BiosVersion** | **string** |  | [optional] [default to null]
**CpldVersion** | **string** |  | [optional] [default to null]
**Errors** | [**[]ModuleStatItemErrorsItems**](module_stat_item_errors_items.md) | used to report all error states the device node is running into. An error should always have &#x60;type&#x60; and &#x60;since&#x60; fields, and could have some other fields specific to that type. | [optional] [default to null]
**Fans** | [**[]ModuleStatItemFansItems**](module_stat_item_fans_items.md) |  | [optional] [default to null]
**FpcIdx** | **int32** |  | [optional] [default to null]
**FpgaVersion** | **string** |  | [optional] [default to null]
**LastSeen** | **float64** |  | [optional] [default to null]
**Model** | **string** |  | [optional] [default to null]
**OpticsCpldVersion** | **string** |  | [optional] [default to null]
**PendingVersion** | **string** |  | [optional] [default to null]
**Pics** | [**[]ModuleStatItemPicsItem**](module_stat_item_pics_item.md) |  | [optional] [default to null]
**Poe** | [***ModuleStatItemPoe**](module_stat_item_poe.md) |  | [optional] [default to null]
**PoeVersion** | **string** |  | [optional] [default to null]
**PowerCpldVersion** | **string** |  | [optional] [default to null]
**Psus** | [**[]ModuleStatItemPsusItem**](module_stat_item_psus_item.md) |  | [optional] [default to null]
**ReFpgaVersion** | **string** |  | [optional] [default to null]
**RecoveryVersion** | **string** |  | [optional] [default to null]
**Serial** | **string** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Temperatures** | [**[]ModuleStatItemTemperaturesItem**](module_stat_item_temperatures_item.md) |  | [optional] [default to null]
**TmcFpgaVersion** | **string** |  | [optional] [default to null]
**UbootVersion** | **string** |  | [optional] [default to null]
**Uptime** | **int32** |  | [optional] [default to null]
**VcLinks** | [**[]ModuleStatItemVcLinksItem**](module_stat_item_vc_links_item.md) |  | [optional] [default to null]
**VcMode** | **string** |  | [optional] [default to null]
**VcRole** | **string** | master / backup / linecard | [optional] [default to null]
**VcState** | **string** |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

