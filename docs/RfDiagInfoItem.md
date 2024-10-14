# RfDiagInfoItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;asset&#x60;, id of the asset | [optional] [default to null]
**AssetName** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;asset&#x60;, name of the asset | [optional] [default to null]
**ClientName** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;, hostname of the client | [optional] [default to null]
**Duration** | **int32** | recording length in seconds, max is 120 | [default to null]
**EndTime** | **int32** | timestamp of end of recording | [default to null]
**FrameCount** | **int32** | Number of frames in the output | [default to null]
**Id** | **string** |  | [optional] [default to null]
**Mac** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60; or &#x60;asset&#x60;, mac of the device | [optional] [default to null]
**MapId** | **string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Next** | **string** | Optional. id of the next recoding if present. Only valid for site survey. | [optional] [default to null]
**RawEvents** | **string** | URL to a JSON file that contains array of raw location diag events | [default to null]
**Ready** | **bool** | whether it’s ready for playback | [default to null]
**SdkclientId** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, sdkclient_id of this recording | [optional] [default to null]
**SdkclientName** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, name of the sdkclient | [optional] [default to null]
**SdkclientUuid** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, device_id of sdkclient | [optional] [default to null]
**StartTime** | **int32** | timestamp of the recording (the start) | [default to null]
**Type_** | [***AllOfrfDiagInfoItemType_**](AllOfrfDiagInfoItemType_.md) |  | [default to null]
**Url** | **string** | URL to a JSON file that contains an array of frames, each frame is the same format | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

