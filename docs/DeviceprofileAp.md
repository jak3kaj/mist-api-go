# DeviceprofileAp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aeroscout** | [***AllOfdeviceprofileApAeroscout**](AllOfdeviceprofileApAeroscout.md) |  | [optional] [default to null]
**BleConfig** | [***AllOfdeviceprofileApBleConfig**](AllOfdeviceprofileApBleConfig.md) |  | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**DisableEth1** | **bool** | whether to disable eth1 port | [optional] [default to false]
**DisableEth2** | **bool** | whether to disable eth2 port | [optional] [default to false]
**DisableEth3** | **bool** | whether to disable eth3 port | [optional] [default to false]
**DisableModule** | **bool** | whether to disable module port | [optional] [default to false]
**EslConfig** | [***ApEslConfig**](ap_esl_config.md) |  | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IotConfig** | [***AllOfdeviceprofileApIotConfig**](AllOfdeviceprofileApIotConfig.md) |  | [optional] [default to null]
**IpConfig** | [***AllOfdeviceprofileApIpConfig**](AllOfdeviceprofileApIpConfig.md) |  | [optional] [default to null]
**Led** | [***AllOfdeviceprofileApLed**](AllOfdeviceprofileApLed.md) |  | [optional] [default to null]
**Mesh** | [***AllOfdeviceprofileApMesh**](AllOfdeviceprofileApMesh.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**NtpServers** | **[]string** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**PoePassthrough** | **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | [**map[string]ApPortConfig**](ap_port_config.md) | Property key is the interface(s) name (e.g. \&quot;eth1,eth2\&quot;) | [optional] [default to null]
**PwrConfig** | [***AllOfdeviceprofileApPwrConfig**](AllOfdeviceprofileApPwrConfig.md) |  | [optional] [default to null]
**RadioConfig** | [***AllOfdeviceprofileApRadioConfig**](AllOfdeviceprofileApRadioConfig.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SwitchConfig** | [***AllOfdeviceprofileApSwitchConfig**](AllOfdeviceprofileApSwitchConfig.md) |  | [optional] [default to null]
**Type_** | **string** | Device Type. enum: &#x60;ap&#x60; | [default to null]
**UplinkPortConfig** | [***ApUplinkPortConfig**](ap_uplink_port_config.md) |  | [optional] [default to null]
**UsbConfig** | [***AllOfdeviceprofileApUsbConfig**](AllOfdeviceprofileApUsbConfig.md) |  | [optional] [default to null]
**Vars** | **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

