# AllOftunnelProviderOptionsZscaler

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AupAcceptanceRequired** | **bool** |  | [optional] [default to true]
**AupExpire** | **int32** | days before AUP is requested again | [optional] [default to 1]
**AupSslProxy** | **bool** | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser | [optional] [default to false]
**DownloadMbps** | **int32** | the download bandwidth cap of the link, in Mbps | [optional] [default to null]
**EnableAup** | **bool** | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60;, display Acceptable Use Policy (AUP) | [optional] [default to false]
**EnableCaution** | **bool** | when &#x60;enforce_authentication&#x60;&#x3D;&#x3D;&#x60;false&#x60;, display caution notification for non-authenticated users | [optional] [default to false]
**EnforceAuthentication** | **bool** |  | [optional] [default to false]
**Name** | **string** |  | [optional] [default to null]
**SubLocations** | [**[]TunnelProviderOptionsZscalerSubLocation**](tunnel_provider_options_zscaler_sub_location.md) | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] [default to null]
**UploadMbps** | **int32** | the download bandwidth cap of the link, in Mbps | [optional] [default to null]
**UseXff** | **bool** | location uses proxy chaining to forward traffic | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

