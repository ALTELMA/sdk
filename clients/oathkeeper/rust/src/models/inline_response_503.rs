/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * The version of the OpenAPI document: v0.38.20-beta.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct InlineResponse503 {
    /// Errors contains a list of errors that caused the not ready status.
    #[serde(rename = "errors")]
    pub errors: ::std::collections::HashMap<String, String>,
}

impl InlineResponse503 {
    pub fn new(errors: ::std::collections::HashMap<String, String>) -> InlineResponse503 {
        InlineResponse503 {
            errors,
        }
    }
}


