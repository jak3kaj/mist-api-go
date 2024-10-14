# ConstDeviceAp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApType** | **string** |  | [optional] [default to null]
**Band24** | [***ConstDeviceApBand24**](const_device_ap_band24.md) |  | [optional] [default to null]
**Band5** | [***ConstDeviceApBand5**](const_device_ap_band5.md) |  | [optional] [default to null]
**Band6** | [***ConstDeviceApBand5**](const_device_ap_band5.md) |  | [optional] [default to null]
**Band24Usages** | [***AllOfconstDeviceApBand24Usages**](AllOfconstDeviceApBand24Usages.md) |  | [optional] [default to null]
**CeDfsOk** | **bool** |  | [optional] [default to null]
**CiscoPace** | **bool** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DisallowedChannels** | [**map[string]Object**](.md) | Property key is a list of country codes (e.g. \&quot;GB, DE\&quot;) | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**Extio** | [**map[string]ConstDeviceApExtios**](const_device_ap_extios.md) | Property key is the GPIO port name (e.g. \&quot;D0\&quot;, \&quot;A1\&quot;) | [optional] [default to null]
**FccDfsOk** | **bool** |  | [optional] [default to null]
**Has11ax** | **bool** |  | [optional] [default to null]
**HasCompass** | **bool** |  | [optional] [default to null]
**HasExtAnt** | **bool** |  | [optional] [default to null]
**HasExtio** | **bool** |  | [optional] [default to null]
**HasHeight** | **bool** |  | [optional] [default to null]
**HasModulePort** | **bool** |  | [optional] [default to null]
**HasPoeOut** | **bool** |  | [optional] [default to null]
**HasScanningRadio** | **bool** |  | [optional] [default to null]
**HasSelectableRadio** | **bool** |  | [optional] [default to null]
**HasUsb** | **bool** |  | [optional] [default to null]
**HasVble** | **bool** |  | [optional] [default to null]
**HasWifiBand24** | **bool** |  | [optional] [default to null]
**HasWifiBand5** | **bool** |  | [optional] [default to null]
**HasWifiBand6** | **bool** |  | [optional] [default to null]
**MaxPoeOut** | **int32** |  | [optional] [default to null]
**MaxWlans** | **int32** |  | [optional] [default to null]
**Model** | **string** |  | [optional] [default to null]
**OtherDfsOk** | **bool** |  | [optional] [default to null]
**Outdoor** | **bool** |  | [optional] [default to null]
**Radios** | **map[string]string** | Property key is the radio number (e.g. r0, r1, ...). Property value is the RF band (e.g. \&quot;24\&quot;, \&quot;5\&quot;, ...) | [optional] [default to null]
**SharedScanningRadio** | **bool** |  | [optional] [default to null]
**Type_** | **string** | Device Type. enum: &#x60;ap&#x60; | [default to null]
**Unmanaged** | **bool** |  | [optional] [default to null]
**Vble** | [***ConstDeviceApVble**](const_device_ap_vble.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

