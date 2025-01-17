/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * The version of the OpenAPI document: v0.38.20-beta.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// GenericError : GenericError The standard error format



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GenericError {
    /// code
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<i64>,
    /// details
    #[serde(rename = "details", skip_serializing_if = "Option::is_none")]
    pub details: Option<Vec<::std::collections::HashMap<String, serde_json::Value>>>,
    /// message
    #[serde(rename = "message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
    /// reason
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<String>,
    /// request
    #[serde(rename = "request", skip_serializing_if = "Option::is_none")]
    pub request: Option<String>,
    /// status
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
}

impl GenericError {
    /// GenericError The standard error format
    pub fn new() -> GenericError {
        GenericError {
            code: None,
            details: None,
            message: None,
            reason: None,
            request: None,
            status: None,
        }
    }
}


