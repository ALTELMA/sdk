/*
 * Ory Kratos
 *
 * Welcome to the ORY Kratos HTTP API documentation!
 *
 * The version of the OpenAPI document: v0.5.5-alpha.4
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateRecoveryLink {
    /// Link Expires In  The recovery link will expire at that point in time. Defaults to the configuration value of `selfservice.flows.recovery.request_lifespan`.
    #[serde(rename = "expires_in", skip_serializing_if = "Option::is_none")]
    pub expires_in: Option<String>,
    #[serde(rename = "identity_id")]
    pub identity_id: String,
}

impl CreateRecoveryLink {
    pub fn new(identity_id: String) -> CreateRecoveryLink {
        CreateRecoveryLink {
            expires_in: None,
            identity_id,
        }
    }
}

