/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.93
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct InlineObject {
    /// Project ID  The Project ID you want to set active.  format: uuid
    #[serde(rename = "project_id")]
    pub project_id: String,
}

impl InlineObject {
    pub fn new(project_id: String) -> InlineObject {
        InlineObject {
            project_id,
        }
    }
}


