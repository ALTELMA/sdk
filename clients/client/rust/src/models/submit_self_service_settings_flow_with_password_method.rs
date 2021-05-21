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
pub struct SubmitSelfServiceSettingsFlowWithPasswordMethod {
    /// CSRFToken is the anti-CSRF token  type: string
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to password when trying to update a password.  type: string
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<String>,
    /// Password is the updated password  type: string
    #[serde(rename = "password")]
    pub password: String,
}

impl SubmitSelfServiceSettingsFlowWithPasswordMethod {
    pub fn new(password: String) -> SubmitSelfServiceSettingsFlowWithPasswordMethod {
        SubmitSelfServiceSettingsFlowWithPasswordMethod {
            csrf_token: None,
            method: None,
            password,
        }
    }
}


