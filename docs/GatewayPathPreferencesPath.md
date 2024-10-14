# GatewayPathPreferencesPath

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **int32** |  | [optional] [default to null]
**Disabled** | **bool** | For SSR Only. &#x60;true&#x60;, if this specific path is undesired | [optional] [default to null]
**GatewayIp** | **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60;, if a different gateway is desired | [optional] [default to null]
**InternetAccess** | **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;vpn&#x60;, if this vpn path can be used for internet | [optional] [default to null]
**Name** | **string** | required when    * &#x60;type&#x60;&#x3D;&#x3D;&#x60;vpn&#x60;: the name of the VPN Path to use    * &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;: the name of the WAN interface to use&#x27; | [optional] [default to null]
**Networks** | **[]string** | required when &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] [default to null]
**TargetIps** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60;, if destination IP is to be replaced | [optional] [default to null]
**Type_** | [***AllOfgatewayPathPreferencesPathType_**](AllOfgatewayPathPreferencesPathType_.md) |  | [optional] [default to null]
**WanName** | **string** | optional if &#x60;type&#x60;&#x3D;&#x3D;&#x60;vpn&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

