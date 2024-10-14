# Go API client for swagger

> Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2409.1.9
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

```bash
# Note it ignored the "packageName" option
curl -X POST -H "content-type:application/json" -d '{"options":{"packageName":"mistapi","packageVersion":"2409.1.9"}, "specURL":"https://www.juniper.net/documentation/us/en/software/mist/api/static/exports/mist-api-openapi3json.json", "lang":"go", "type":"CLIENT"}' https://generator3.swagger.io/api/generate -o go-client-generated-2409.1.9.zip
unzip go-client-generated-2409.1.9.zip
sed -E -i '' 's/([a-zA-Z0-9]*)__([a-zA-Z0-9]+)/\2_\1/' *.go
sed -i '' '/const (/,/)/d' swagger/model_*.go
sed -i '' 's/\*Object/*any/' swagger/model_*
sed -i '' 's/\]Object/]any/' swagger/model*

# modified:
# response.go
# swagger/api_constants_definitions.go
# swagger/api_installer.go
# swagger/api_msps_sso.go
# swagger/api_orgs_crl.go
# swagger/api_orgs_maps.go
# swagger/api_orgs_sdk_invites.go
# swagger/api_orgs_sites.go
# swagger/api_orgs_sso.go
# swagger/api_orgs_stats_devices.go
# swagger/api_sites_assets.go
# swagger/api_sites_maps.go
# swagger/api_sites_rfdiags.go
# swagger/api_sites_stats_assets.go
# swagger/api_sites_stats_devices.go
# swagger/api_sites_stats_devices.go
# swagger/client.go
# swagger/model_admins_admin_id_body.go
# swagger/model_ap_radio_antenna_mode.go
# swagger/model_ble_config_beacon_rate_mode.go
# swagger/model_capture_radiotap_band.go
# swagger/model_capture_radiotapwired_band.go
# swagger/model_capture_scan_aps_band.go swagger/model_capture_scan_band.go swagger/model_capture_wireless_band.go
# swagger/model_dot11_band.go
# swagger/model_dot11_bandwidth.go
# swagger/model_dot11_bandwidth24.go
# swagger/model_dot11_bandwidth5.go
# swagger/model_gateway_template_tunnel_ike_dh_group.go
# swagger/model_inline_response_200_16.go
# swagger/model_inline_response_200_17.go
# swagger/model_list_msp_logs_sort.go
# swagger/model_stats_ap.go
# swagger/model_stats_ap.go
# swagger/model_stats_ap_ble.go
# swagger/model_stats_beacon.go
# swagger/model_stats_device.go
# swagger/model_stats_devices.go
# swagger/model_stats_device.go

```
