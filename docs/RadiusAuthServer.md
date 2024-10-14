# RadiusAuthServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | ip / hostname of RADIUS server | [default to null]
**KeywrapEnabled** | **bool** |  | [optional] [default to null]
**KeywrapFormat** | [***AllOfradiusAuthServerKeywrapFormat**](AllOfradiusAuthServerKeywrapFormat.md) |  | [optional] [default to null]
**KeywrapKek** | **string** |  | [optional] [default to null]
**KeywrapMack** | **string** |  | [optional] [default to null]
**Port** | **int32** | Auth port of RADIUS server | [optional] [default to 1812]
**RequireMessageAuthenticator** | **bool** | whether to require Message-Authenticator in requests | [optional] [default to false]
**Secret** | **string** | secret of RADIUS server | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

