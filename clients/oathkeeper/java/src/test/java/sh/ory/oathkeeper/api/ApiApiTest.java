/*
 * Ory Oathkeeper API
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * The version of the OpenAPI document: v0.38.20-beta.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.oathkeeper.api;

import sh.ory.oathkeeper.ApiException;
import sh.ory.oathkeeper.model.GenericError;
import sh.ory.oathkeeper.model.JsonWebKeySet;
import sh.ory.oathkeeper.model.Rule;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ApiApi
 */
@Ignore
public class ApiApiTest {

    private final ApiApi api = new ApiApi();

    
    /**
     * Access Control Decision API
     *
     * &gt; This endpoint works with all HTTP Methods (GET, POST, PUT, ...) and matches every path prefixed with /decision.  This endpoint mirrors the proxy capability of ORY Oathkeeper&#39;s proxy functionality but instead of forwarding the request to the upstream server, returns 200 (request should be allowed), 401 (unauthorized), or 403 (forbidden) status codes. This endpoint can be used to integrate with other API Proxies like Ambassador, Kong, Envoy, and many more.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void decisionsTest() throws ApiException {
                api.decisions();
        // TODO: test validations
    }
    
    /**
     * Retrieve a Rule
     *
     * Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getRuleTest() throws ApiException {
        String id = null;
                Rule response = api.getRule(id);
        // TODO: test validations
    }
    
    /**
     * Lists Cryptographic Keys
     *
     * This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWellKnownJSONWebKeysTest() throws ApiException {
                JsonWebKeySet response = api.getWellKnownJSONWebKeys();
        // TODO: test validations
    }
    
    /**
     * List All Rules
     *
     * This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full view of what rules you have currently in place.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listRulesTest() throws ApiException {
        Long limit = null;
        Long offset = null;
                List<Rule> response = api.listRules(limit, offset);
        // TODO: test validations
    }
    
}
