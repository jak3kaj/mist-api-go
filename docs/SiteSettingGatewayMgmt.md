# SiteSettingGatewayMgmt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSshkeys** | **[]string** | for SSR only, as direct root access is not allowed | [optional] [default to null]
**AppProbing** | [***AppProbing**](app_probing.md) |  | [optional] [default to null]
**AppUsage** | **bool** | consumes uplink bandwidth, requires WA license | [optional] [default to null]
**AutoSignatureUpdate** | [***SiteSettingGatewayMgmtAutoSignatureUpdate**](site_setting_gateway_mgmt_auto_signature_update.md) |  | [optional] [default to null]
**ConfigRevertTimer** | **int32** | he rollback timer for commit confirmed | [optional] [default to 10]
**DisableConsole** | **bool** | for both SSR and SRX disable console port | [optional] [default to false]
**DisableOob** | **bool** | for both SSR and SRX disable management interface | [optional] [default to false]
**ProbeHosts** | **[]string** |  | [optional] [default to null]
**ProtectRe** | [***AllOfsiteSettingGatewayMgmtProtectRe**](AllOfsiteSettingGatewayMgmtProtectRe.md) |  | [optional] [default to null]
**RootPassword** | **string** | for SRX only | [optional] [default to null]
**SecurityLogSourceAddress** | **string** |  | [optional] [default to null]
**SecurityLogSourceInterface** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

