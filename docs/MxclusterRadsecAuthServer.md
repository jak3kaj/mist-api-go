# MxclusterRadsecAuthServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | ip / hostname of RADIUS server | [optional] [default to null]
**KeywrapEnabled** | **bool** | if used for Mist APs, enable keywrap algorithm. Default is false | [optional] [default to null]
**KeywrapFormat** | [***AllOfmxclusterRadsecAuthServerKeywrapFormat**](AllOfmxclusterRadsecAuthServerKeywrapFormat.md) |  | [optional] [default to null]
**KeywrapKek** | **string** | if used for Mist APs, encryption key | [optional] [default to null]
**KeywrapMack** | **string** | if used for Mist APs, Message Authentication Code Key | [optional] [default to null]
**Port** | **int32** | Auth port of RADIUS server | [optional] [default to 1812]
**Secret** | **string** | secret of RADIUS server | [optional] [default to null]
**Ssids** | **[]string** | list of ssids that will use this server if match_ssid is true and match is found | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

