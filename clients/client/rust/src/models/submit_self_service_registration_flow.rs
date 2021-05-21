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
pub struct SubmitSelfServiceRegistrationFlow {
    /// The CSRF Token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method to use  This field must be set to `password` when using the password method.
    #[serde(rename = "method")]
    pub method: String,
    /// Password to sign the user up with
    #[serde(rename = "password", skip_serializing_if = "Option::is_none")]
    pub password: Option<String>,
    /// The identity's traits
    #[serde(rename = "traits", skip_serializing_if = "Option::is_none")]
    pub traits: Option<serde_json::Value>,
}

impl SubmitSelfServiceRegistrationFlow {
    pub fn new(method: String) -> SubmitSelfServiceRegistrationFlow {
        SubmitSelfServiceRegistrationFlow {
            csrf_token: None,
            method,
            password: None,
            traits: None,
        }
    }
}


