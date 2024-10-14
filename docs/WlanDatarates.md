# WlanDatarates

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ht** | **string** | MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 00ff 00f0 001f limits HT rates to MCS 0-7 for 1 stream, MCS 4-7 for 2 stream (i.e. MCS 12-15), MCS 1-5 for 3 stream (i.e. MCS 16-20) | [optional] [default to null]
**Legacy** | [**[]WlanDataratesLegacyItem**](wlan_datarates_legacy_item.md) | list of supported rates (IE&#x3D;1) and extended supported rates (IE&#x3D;50) for custom template, append ‘b’ at the end to indicate a rate being basic/mandatory. If &#x60;template&#x60;&#x3D;&#x3D;&#x60;custom&#x60; is configured and legacy does not define at least one basic rate, it will use &#x60;no-legacy&#x60; default values | [optional] [default to null]
**MinRssi** | **int32** | Minimum RSSI for client to connect, 0 means not enforcing | [optional] [default to null]
**Template** | **string** | * &#x60;no-legacy&#x60;: no 11b * &#x60;compatible&#x60;: all, like before, default setting that Broadcom/Atheros used * &#x60;legacy-only&#x60;: disable 802.11n and 802.11ac * &#x60;high-density&#x60;: no 11b, no low rates * &#x60;custom&#x60;: user defined | [optional] [default to null]
**Vht** | **string** | MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 03ff 01ff 00ff limits VHT rates to MCS 0-9 for 1 stream, MCS 0-8 for 2 streams, and MCS 0-7 for 3 streams. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

