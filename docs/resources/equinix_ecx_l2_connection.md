---
subcategory: "Fabric"
---

# equinix_ecx_l2_connection (Resource)

Resource `equinix_ecx_l2_connection` allows creation and management of Equinix Fabric
layer 2 connections.

## Example Usage

### Non-redundant Connection from own Equinix Port

```hcl
data "equinix_ecx_l2_sellerprofile" "aws" {
  name = "AWS Direct Connect"
}

data "equinix_ecx_port" "sv-qinq-pri" {
  name = "CX-SV5-NL-Dot1q-BO-10G-PRI"
}

resource "equinix_ecx_l2_connection" "port-2-aws" {
  name              = "tf-aws"
  profile_uuid      = data.equinix_ecx_l2_sellerprofile.aws.id
  speed             = 200
  speed_unit        = "MB"
  notifications     = ["marry@equinix.com", "john@equinix.com"]
  port_uuid         = data.equinix_ecx_port.sv-qinq-pri.id
  vlan_stag         = 777
  vlan_ctag         = 1000
  seller_region     = "us-west-1"
  seller_metro_code = "SV"
  authorization_key = "345742915919"
}
```

### Redundant Connection from own Equinix Ports

```hcl
data "equinix_ecx_l2_sellerprofile" "azure" {
  name = "Azure Express Route"
}

data "equinix_ecx_port" "sv-qinq-pri" {
  name = "CX-SV5-NL-Dot1q-BO-10G-PRI"
}

data "equinix_ecx_port" "sv-qinq-sec" {
  name = "CX-SV1-NL-Dot1q-BO-10G-SEC"
}

resource "equinix_ecx_l2_connection" "ports-2-azure" {
  name              = "tf-azure-pri"
  profile_uuid      = data.equinix_ecx_l2_sellerprofile.azure.id
  speed             = 50
  speed_unit        = "MB"
  notifications     = ["john@equinix.com", "marry@equinix.com"]
  port_uuid         = data.equinix_ecx_port.sv-qinq-pri.id
  vlan_stag         = 1482
  vlan_ctag         = 2512
  seller_metro_code = "SV"
  named_tag         = "PRIVATE"
  authorization_key = "c4dff8e8-b52f-4b34-b0d4-c4588f7338f3
  secondary_connection {
    name      = "tf-azure-sec"
    port_uuid = data.equinix_ecx_port.sv-qinq-sec.id
    vlan_stag = 1904
    vlan_ctag = 1631
  }
}
```

### Non-redundant Connection from Network Edge device

```hcl
data "equinix_ecx_l2_sellerprofile" "gcp-1" {
  name = "Google Cloud Partner Interconnect Zone 1"
}

resource "equinix_ecx_l2_connection" "router-to-gcp" {
  name                = "tf-gcp-pri"
  profile_uuid        = data.equinix_ecx_l2_sellerprofile.gcp-1.id
  device_uuid         = equinix_network_device.myrouter.id
  device_interface_id = 5
  speed               = 100
  speed_unit          = "MB"
  notifications       = ["john@equinix.com", "marry@equinix.com"]
  seller_metro_code   = "SV"
  seller_region       = "us-west1"
  authorization_key   = "4d335adc-00fd-4a41-c9f3-782ca31ab3f7/us-west1/1"
}
```

### Non-redundant Connection from an Equinix customer port using A-Side Service token

```hcl
data "equinix_ecx_l2_sellerprofile" "gcp" {
  name = "Google Cloud Partner Interconnect Zone 1"
}

resource "equinix_ecx_l2_connection" "token-to-gcp" {
  name                = "tf-gcp-pri"
  profile_uuid        = data.equinix_ecx_l2_sellerprofile.gcp-1.id
  service_token       = "e9c22453-d3a7-4d5d-9112-d50173531392"
  speed               = 100
  speed_unit          = "MB"
  notifications       = ["john@equinix.com", "marry@equinix.com"]
  seller_metro_code   = "SV"
  seller_region       = "us-west1"
  authorization_key   = "4d335adc-00fd-4a41-c9f3-782ca31ab3f7/us-west1/1"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Connection name. An alpha-numeric 24 characters string which can include only
hyphens and underscores
- `profile_uuid` - (Required) Unique identifier of the service provider's profile.
- `speed` - (Required) Speed/Bandwidth to be allocated to the connection.
- `speed_unit` - (Required) Unit of the speed/bandwidth to be allocated to the connection.
- `notifications` - (Required) A list of email addresses used for sending connection update
notifications.
- `purchase_order_number` - (Optional) Connection's purchase order number to reflect on the invoice
- `port_uuid` - (Required when `device_uuid` or `service_token` are not set) Unique identifier of
the Equinix port from which the connection would originate.
- `device_uuid` - (Required when `port_uuid` or `service_token` are not set) Unique identifier of
the Network Edge virtual device from which the connection would originate.
- `device_interface_id` - (Optional) Applicable with `device_uuid`, identifier of network interface
on a given device, used for a connection. If not specified then first available interface will be
selected.
- `service_token`- (Required when `port_uuid` or `device_uuid` are not set) - Unique Equinix Fabric
key given by a provider that grants you authorization to enable connectivity from a shared
multi-tenant Equinix port (a-side). More details in [Fabric Service Tokens](https://docs.equinix.com/en-us/Content/Interconnection/Fabric/service%20tokens/Fabric-Service-Tokens.htm).
- `vlan_stag` - (Required when port_uuid is set) S-Tag/Outer-Tag of the connection - a numeric
character ranging from 2 - 4094.
- `vlan_ctag` - (Optional) C-Tag/Inner-Tag of the connection - a numeric character ranging from 2
\- 4094.
- `named_tag` - (Optional) The type of peering to set up when connecting to Azure Express Route.
Valid values: `PRIVATE`, `MICROSOFT`, `MANUAL`, `PUBLIC`\*.
~> **NOTE:** _"PUBLIC"_ peering is deprecated. Use _"MICROSOFT"_ instead. More details in
[Microsoft public peering](https://docs.microsoft.com/en-us/azure/expressroute/about-public-peering) docs.
- `additional_info` - (Optional) one or more additional information key-value objects
  - `name` - (Required) additional information key
  - `value` - (Required) additional information value
- `zside_port_uuid` - (Optional) Unique identifier of the port on the remote side (z-side).
- `zside_vlan_stag` - (Optional) S-Tag/Outer-Tag of the connection on the remote side (z side) - a
numeric character ranging from 2 - 4094.
- `zside_vlan_ctag` - (Optional) C-Tag/Inner-Tag of the connectionon the remote side (z side). This
is only applicable if `named_tag` is 'MANUAL' - a numeric character ranging from 2 - 4094. If
`secondary_connection` is defined it will internally use same `zside_vlan_ctag` for the secondary
connection.
- `seller_region` - (Optional) The region in which the seller port resides.
- `seller_metro_code` - (Optional) The metro code that denotes the connection’s remote side
(z-side).
- `authorization_key` - (Optional) Text field used to authorize connection on the provider side.
Value depends on a provider service profile used for connection.
- `secondary_connection` - (Optional) Definition of secondary connection for redundant, HA
connectivity. See [Secondary Connection](#secondary-connection) below for more details.

### Secondary Connection

The `secondary_connection` block supports the following arguments:

- `name` - (Required) secondary connection name
- `speed` - (Optional) Speed/Bandwidth to be allocated to the secondary connection. If not
specified primary `speed` will be used.
- `speed_unit` - (Optional) Unit of the speed/bandwidth to be allocated to the secondary
connection. If not specified primary `speed_unit` will be used.
- `port_uuid` - (Optional) Applicable with primary `port_uuid`. Identifier of the Equinix port from
which the secondary connection would originate. If not specified primary `port_uuid` will be used.
- `device_uuid` - (Optional) Applicable with primary `device_uuid`. Identifier of the Network Edge
virtual device from which the secondary connection would originate. If not specified primary
`device_uuid` will be used.
- `device_interface_id` - (Optional) Applicable with `device_uuid`, identifier of network interface
on a given device. If not specified then first available interface will be selected.
- `service_token`- (Optional) Applicable with primary `service_token`. Unique Equinix Fabric key
given by a provider that grants you authorization to enable connectivity from a shared multi-tenant
Equinix port (a-side). More details in [Fabric Service Tokens](https://docs.equinix.com/en-us/Content/Interconnection/Fabric/service%20tokens/Fabric-Service-Tokens.htm). If not specified primary `service_token` will be used.
- `vlan_stag` - (Required when `port_uuid` is set) S-Tag/Outer-Tag of the secondary connection, a
numeric character ranging from 2 - 4094.
- `vlan_ctag` - (Optional) Applicable with `port_uuid`. C-Tag/Inner-Tag of the secondary
connection, a numeric character ranging from 2 - 4094.
- `seller_metro_code` - (Optional) The metro code that denotes the secondary connection’s
destination (Z side). If not specified primary `seller_metro_code` will be used.
- `seller_region` - (Optional) The region in which the seller port resides. If not specified
primary `seller_region` will be used.
- `authorization_key` - (Optional) Text field based on the service profile you want to connect to.
If not specified primary `authorization_key` will be used.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes
are exported:

- `uuid` - Unique identifier of the connection.
- `status` - Connection provisioning status on Equinix Fabric side.
- `provider_status` - Connection provisioning status on service provider's side.
- `redundant_uuid` - Unique identifier of the redundant connection, applicable for HA connections.
- `redundancy_type` - Connection redundancy type, applicable for HA connections. Valid values: `PRIMARY`,
`SECONDARY`.
- `redundancy_group` - Unique identifier of group containing a primary and secondary connection.
- `zside_port_uuid` - When not provided as an argument, it is identifier of the z-side port,
assigned by the Fabric.
- `zside_vlan_stag` - When not provided as an argument, it is S-Tag/Outer-Tag of the connection on
the Z side, assigned by the Fabric.
- `actions` - One or more pending actions to complete connection provisioning.
- `secondary_connection`:
  - `zside_port_uuid`
  - `zside_vlan_stag`
  - `zside_vlan_ctag`
  - `redundancy_type`
  - `redundancy_group`

## Update operation behavior

Update of most arguments will force replacement of a connection (including related redundant
connection in HA setup).

Following arguments can be updated. **NOTE** that Equinix Fabric may still forbid updates depending
on current connection state, used service provider or number of updates requested during the day.

- `name`
- `speed` and `speed_unit`

## Timeouts

This resource provides the following [Timeouts configuration](https://www.terraform.io/language/resources/syntax#operation-timeouts)
options:

- create - Default is 5 minutes
- delete - Default is 5 minutes

## Import

Equinix L2 connections can be imported using an existing `id`:

```sh
existing_connection_id='example-uuid-1'
terraform import equinix_ecx_l2_connection.example ${existing_connection_id}
```

To import a redundant connection it is required a single string with both connection `id` separated
by `:`, e.g.,

```sh
existing_primary_connection_id='example-uuid-1'
existing_secondary_connection_id='example-uuid-2'
terraform import equinix_ecx_l2_connection.example ${existing_primary_connection_id}:${existing_secondary_connection_id}
```
