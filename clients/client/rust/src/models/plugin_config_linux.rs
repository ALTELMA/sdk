/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.7
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginConfigLinux : PluginConfigLinux plugin config linux



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginConfigLinux {
    /// allow all devices
    #[serde(rename = "AllowAllDevices")]
    pub allow_all_devices: bool,
    /// capabilities
    #[serde(rename = "Capabilities")]
    pub capabilities: Vec<String>,
    /// devices
    #[serde(rename = "Devices")]
    pub devices: Vec<crate::models::PluginDevice>,
}

impl PluginConfigLinux {
    /// PluginConfigLinux plugin config linux
    pub fn new(allow_all_devices: bool, capabilities: Vec<String>, devices: Vec<crate::models::PluginDevice>) -> PluginConfigLinux {
        PluginConfigLinux {
            allow_all_devices,
            capabilities,
            devices,
        }
    }
}


