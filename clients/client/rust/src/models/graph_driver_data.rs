/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.7
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GraphDriverData {
    /// data
    #[serde(rename = "Data")]
    pub data: ::std::collections::HashMap<String, String>,
    /// name
    #[serde(rename = "Name")]
    pub name: String,
}

impl GraphDriverData {
    pub fn new(data: ::std::collections::HashMap<String, String>, name: String) -> GraphDriverData {
        GraphDriverData {
            data,
            name,
        }
    }
}


