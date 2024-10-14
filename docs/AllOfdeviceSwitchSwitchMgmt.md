# AllOfdeviceSwitchSwitchMgmt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAffinityThreshold** | **int32** | ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used. | [optional] [default to 10]
**CliBanner** | **string** | Set Banners for switches. Allows markup formatting | [optional] [default to null]
**CliIdleTimeout** | **int32** | Sets timeout for switches | [optional] [default to null]
**ConfigRevertTimer** | **int32** | the rollback timer for commit confirmed | [optional] [default to 10]
**DhcpOptionFqdn** | **bool** | Enable to provide the FQDN with DHCP option 81 | [optional] [default to false]
**DisableOobDownAlarm** | **bool** |  | [optional] [default to null]
**LocalAccounts** | [**map[string]ConfigSwitchLocalAccountsUser**](config_switch_local_accounts_user.md) | Property key is the user name. For Local user authentication | [optional] [default to null]
**MxedgeProxyHost** | **string** |  | [optional] [default to null]
**MxedgeProxyPort** | **int32** |  | [optional] [default to 2222]
**ProtectRe** | [***Object**](.md) |  | [optional] [default to null]
**Radius** | [***Object**](.md) |  | [optional] [default to null]
**RootPassword** | **string** |  | [optional] [default to null]
**Tacacs** | [***Tacacs**](tacacs.md) |  | [optional] [default to null]
**UseMxedgeProxy** | **bool** | to use mxedge as proxy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

