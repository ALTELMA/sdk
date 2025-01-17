<?php
/**
 * OAuth2Client
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: latest
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.0.0-SNAPSHOT
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * OAuth2Client Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class OAuth2Client implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'oAuth2Client';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'allowed_cors_origins' => 'string[]',
        'audience' => 'string[]',
        'backchannel_logout_session_required' => 'bool',
        'backchannel_logout_uri' => 'string',
        'client_id' => 'string',
        'client_name' => 'string',
        'client_secret' => 'string',
        'client_secret_expires_at' => 'int',
        'client_uri' => 'string',
        'contacts' => 'string[]',
        'created_at' => '\DateTime',
        'frontchannel_logout_session_required' => 'bool',
        'frontchannel_logout_uri' => 'string',
        'grant_types' => 'string[]',
        'jwks' => 'object',
        'jwks_uri' => 'string',
        'logo_uri' => 'string',
        'metadata' => 'object',
        'owner' => 'string',
        'policy_uri' => 'string',
        'post_logout_redirect_uris' => 'string[]',
        'redirect_uris' => 'string[]',
        'registration_access_token' => 'string',
        'registration_client_uri' => 'string',
        'request_object_signing_alg' => 'string',
        'request_uris' => 'string[]',
        'response_types' => 'string[]',
        'scope' => 'string',
        'sector_identifier_uri' => 'string',
        'subject_type' => 'string',
        'token_endpoint_auth_method' => 'string',
        'token_endpoint_auth_signing_alg' => 'string',
        'tos_uri' => 'string',
        'updated_at' => '\DateTime',
        'userinfo_signed_response_alg' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'allowed_cors_origins' => null,
        'audience' => null,
        'backchannel_logout_session_required' => null,
        'backchannel_logout_uri' => null,
        'client_id' => null,
        'client_name' => null,
        'client_secret' => null,
        'client_secret_expires_at' => 'int64',
        'client_uri' => null,
        'contacts' => null,
        'created_at' => 'date-time',
        'frontchannel_logout_session_required' => null,
        'frontchannel_logout_uri' => null,
        'grant_types' => null,
        'jwks' => null,
        'jwks_uri' => null,
        'logo_uri' => null,
        'metadata' => null,
        'owner' => null,
        'policy_uri' => null,
        'post_logout_redirect_uris' => null,
        'redirect_uris' => null,
        'registration_access_token' => null,
        'registration_client_uri' => null,
        'request_object_signing_alg' => null,
        'request_uris' => null,
        'response_types' => null,
        'scope' => null,
        'sector_identifier_uri' => null,
        'subject_type' => null,
        'token_endpoint_auth_method' => null,
        'token_endpoint_auth_signing_alg' => null,
        'tos_uri' => null,
        'updated_at' => 'date-time',
        'userinfo_signed_response_alg' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'allowed_cors_origins' => 'allowed_cors_origins',
        'audience' => 'audience',
        'backchannel_logout_session_required' => 'backchannel_logout_session_required',
        'backchannel_logout_uri' => 'backchannel_logout_uri',
        'client_id' => 'client_id',
        'client_name' => 'client_name',
        'client_secret' => 'client_secret',
        'client_secret_expires_at' => 'client_secret_expires_at',
        'client_uri' => 'client_uri',
        'contacts' => 'contacts',
        'created_at' => 'created_at',
        'frontchannel_logout_session_required' => 'frontchannel_logout_session_required',
        'frontchannel_logout_uri' => 'frontchannel_logout_uri',
        'grant_types' => 'grant_types',
        'jwks' => 'jwks',
        'jwks_uri' => 'jwks_uri',
        'logo_uri' => 'logo_uri',
        'metadata' => 'metadata',
        'owner' => 'owner',
        'policy_uri' => 'policy_uri',
        'post_logout_redirect_uris' => 'post_logout_redirect_uris',
        'redirect_uris' => 'redirect_uris',
        'registration_access_token' => 'registration_access_token',
        'registration_client_uri' => 'registration_client_uri',
        'request_object_signing_alg' => 'request_object_signing_alg',
        'request_uris' => 'request_uris',
        'response_types' => 'response_types',
        'scope' => 'scope',
        'sector_identifier_uri' => 'sector_identifier_uri',
        'subject_type' => 'subject_type',
        'token_endpoint_auth_method' => 'token_endpoint_auth_method',
        'token_endpoint_auth_signing_alg' => 'token_endpoint_auth_signing_alg',
        'tos_uri' => 'tos_uri',
        'updated_at' => 'updated_at',
        'userinfo_signed_response_alg' => 'userinfo_signed_response_alg'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'allowed_cors_origins' => 'setAllowedCorsOrigins',
        'audience' => 'setAudience',
        'backchannel_logout_session_required' => 'setBackchannelLogoutSessionRequired',
        'backchannel_logout_uri' => 'setBackchannelLogoutUri',
        'client_id' => 'setClientId',
        'client_name' => 'setClientName',
        'client_secret' => 'setClientSecret',
        'client_secret_expires_at' => 'setClientSecretExpiresAt',
        'client_uri' => 'setClientUri',
        'contacts' => 'setContacts',
        'created_at' => 'setCreatedAt',
        'frontchannel_logout_session_required' => 'setFrontchannelLogoutSessionRequired',
        'frontchannel_logout_uri' => 'setFrontchannelLogoutUri',
        'grant_types' => 'setGrantTypes',
        'jwks' => 'setJwks',
        'jwks_uri' => 'setJwksUri',
        'logo_uri' => 'setLogoUri',
        'metadata' => 'setMetadata',
        'owner' => 'setOwner',
        'policy_uri' => 'setPolicyUri',
        'post_logout_redirect_uris' => 'setPostLogoutRedirectUris',
        'redirect_uris' => 'setRedirectUris',
        'registration_access_token' => 'setRegistrationAccessToken',
        'registration_client_uri' => 'setRegistrationClientUri',
        'request_object_signing_alg' => 'setRequestObjectSigningAlg',
        'request_uris' => 'setRequestUris',
        'response_types' => 'setResponseTypes',
        'scope' => 'setScope',
        'sector_identifier_uri' => 'setSectorIdentifierUri',
        'subject_type' => 'setSubjectType',
        'token_endpoint_auth_method' => 'setTokenEndpointAuthMethod',
        'token_endpoint_auth_signing_alg' => 'setTokenEndpointAuthSigningAlg',
        'tos_uri' => 'setTosUri',
        'updated_at' => 'setUpdatedAt',
        'userinfo_signed_response_alg' => 'setUserinfoSignedResponseAlg'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'allowed_cors_origins' => 'getAllowedCorsOrigins',
        'audience' => 'getAudience',
        'backchannel_logout_session_required' => 'getBackchannelLogoutSessionRequired',
        'backchannel_logout_uri' => 'getBackchannelLogoutUri',
        'client_id' => 'getClientId',
        'client_name' => 'getClientName',
        'client_secret' => 'getClientSecret',
        'client_secret_expires_at' => 'getClientSecretExpiresAt',
        'client_uri' => 'getClientUri',
        'contacts' => 'getContacts',
        'created_at' => 'getCreatedAt',
        'frontchannel_logout_session_required' => 'getFrontchannelLogoutSessionRequired',
        'frontchannel_logout_uri' => 'getFrontchannelLogoutUri',
        'grant_types' => 'getGrantTypes',
        'jwks' => 'getJwks',
        'jwks_uri' => 'getJwksUri',
        'logo_uri' => 'getLogoUri',
        'metadata' => 'getMetadata',
        'owner' => 'getOwner',
        'policy_uri' => 'getPolicyUri',
        'post_logout_redirect_uris' => 'getPostLogoutRedirectUris',
        'redirect_uris' => 'getRedirectUris',
        'registration_access_token' => 'getRegistrationAccessToken',
        'registration_client_uri' => 'getRegistrationClientUri',
        'request_object_signing_alg' => 'getRequestObjectSigningAlg',
        'request_uris' => 'getRequestUris',
        'response_types' => 'getResponseTypes',
        'scope' => 'getScope',
        'sector_identifier_uri' => 'getSectorIdentifierUri',
        'subject_type' => 'getSubjectType',
        'token_endpoint_auth_method' => 'getTokenEndpointAuthMethod',
        'token_endpoint_auth_signing_alg' => 'getTokenEndpointAuthSigningAlg',
        'tos_uri' => 'getTosUri',
        'updated_at' => 'getUpdatedAt',
        'userinfo_signed_response_alg' => 'getUserinfoSignedResponseAlg'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['allowed_cors_origins'] = $data['allowed_cors_origins'] ?? null;
        $this->container['audience'] = $data['audience'] ?? null;
        $this->container['backchannel_logout_session_required'] = $data['backchannel_logout_session_required'] ?? null;
        $this->container['backchannel_logout_uri'] = $data['backchannel_logout_uri'] ?? null;
        $this->container['client_id'] = $data['client_id'] ?? null;
        $this->container['client_name'] = $data['client_name'] ?? null;
        $this->container['client_secret'] = $data['client_secret'] ?? null;
        $this->container['client_secret_expires_at'] = $data['client_secret_expires_at'] ?? null;
        $this->container['client_uri'] = $data['client_uri'] ?? null;
        $this->container['contacts'] = $data['contacts'] ?? null;
        $this->container['created_at'] = $data['created_at'] ?? null;
        $this->container['frontchannel_logout_session_required'] = $data['frontchannel_logout_session_required'] ?? null;
        $this->container['frontchannel_logout_uri'] = $data['frontchannel_logout_uri'] ?? null;
        $this->container['grant_types'] = $data['grant_types'] ?? null;
        $this->container['jwks'] = $data['jwks'] ?? null;
        $this->container['jwks_uri'] = $data['jwks_uri'] ?? null;
        $this->container['logo_uri'] = $data['logo_uri'] ?? null;
        $this->container['metadata'] = $data['metadata'] ?? null;
        $this->container['owner'] = $data['owner'] ?? null;
        $this->container['policy_uri'] = $data['policy_uri'] ?? null;
        $this->container['post_logout_redirect_uris'] = $data['post_logout_redirect_uris'] ?? null;
        $this->container['redirect_uris'] = $data['redirect_uris'] ?? null;
        $this->container['registration_access_token'] = $data['registration_access_token'] ?? null;
        $this->container['registration_client_uri'] = $data['registration_client_uri'] ?? null;
        $this->container['request_object_signing_alg'] = $data['request_object_signing_alg'] ?? null;
        $this->container['request_uris'] = $data['request_uris'] ?? null;
        $this->container['response_types'] = $data['response_types'] ?? null;
        $this->container['scope'] = $data['scope'] ?? null;
        $this->container['sector_identifier_uri'] = $data['sector_identifier_uri'] ?? null;
        $this->container['subject_type'] = $data['subject_type'] ?? null;
        $this->container['token_endpoint_auth_method'] = $data['token_endpoint_auth_method'] ?? null;
        $this->container['token_endpoint_auth_signing_alg'] = $data['token_endpoint_auth_signing_alg'] ?? null;
        $this->container['tos_uri'] = $data['tos_uri'] ?? null;
        $this->container['updated_at'] = $data['updated_at'] ?? null;
        $this->container['userinfo_signed_response_alg'] = $data['userinfo_signed_response_alg'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (!is_null($this->container['scope']) && !preg_match("/([a-zA-Z0-9\\.\\*]+\\s?)+/", $this->container['scope'])) {
            $invalidProperties[] = "invalid value for 'scope', must be conform to the pattern /([a-zA-Z0-9\\.\\*]+\\s?)+/.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets allowed_cors_origins
     *
     * @return string[]|null
     */
    public function getAllowedCorsOrigins()
    {
        return $this->container['allowed_cors_origins'];
    }

    /**
     * Sets allowed_cors_origins
     *
     * @param string[]|null $allowed_cors_origins allowed_cors_origins
     *
     * @return self
     */
    public function setAllowedCorsOrigins($allowed_cors_origins)
    {
        $this->container['allowed_cors_origins'] = $allowed_cors_origins;

        return $this;
    }

    /**
     * Gets audience
     *
     * @return string[]|null
     */
    public function getAudience()
    {
        return $this->container['audience'];
    }

    /**
     * Sets audience
     *
     * @param string[]|null $audience audience
     *
     * @return self
     */
    public function setAudience($audience)
    {
        $this->container['audience'] = $audience;

        return $this;
    }

    /**
     * Gets backchannel_logout_session_required
     *
     * @return bool|null
     */
    public function getBackchannelLogoutSessionRequired()
    {
        return $this->container['backchannel_logout_session_required'];
    }

    /**
     * Sets backchannel_logout_session_required
     *
     * @param bool|null $backchannel_logout_session_required Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.
     *
     * @return self
     */
    public function setBackchannelLogoutSessionRequired($backchannel_logout_session_required)
    {
        $this->container['backchannel_logout_session_required'] = $backchannel_logout_session_required;

        return $this;
    }

    /**
     * Gets backchannel_logout_uri
     *
     * @return string|null
     */
    public function getBackchannelLogoutUri()
    {
        return $this->container['backchannel_logout_uri'];
    }

    /**
     * Sets backchannel_logout_uri
     *
     * @param string|null $backchannel_logout_uri RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.
     *
     * @return self
     */
    public function setBackchannelLogoutUri($backchannel_logout_uri)
    {
        $this->container['backchannel_logout_uri'] = $backchannel_logout_uri;

        return $this;
    }

    /**
     * Gets client_id
     *
     * @return string|null
     */
    public function getClientId()
    {
        return $this->container['client_id'];
    }

    /**
     * Sets client_id
     *
     * @param string|null $client_id ID  is the id for this client.
     *
     * @return self
     */
    public function setClientId($client_id)
    {
        $this->container['client_id'] = $client_id;

        return $this;
    }

    /**
     * Gets client_name
     *
     * @return string|null
     */
    public function getClientName()
    {
        return $this->container['client_name'];
    }

    /**
     * Sets client_name
     *
     * @param string|null $client_name Name is the human-readable string name of the client to be presented to the end-user during authorization.
     *
     * @return self
     */
    public function setClientName($client_name)
    {
        $this->container['client_name'] = $client_name;

        return $this;
    }

    /**
     * Gets client_secret
     *
     * @return string|null
     */
    public function getClientSecret()
    {
        return $this->container['client_secret'];
    }

    /**
     * Sets client_secret
     *
     * @param string|null $client_secret Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
     *
     * @return self
     */
    public function setClientSecret($client_secret)
    {
        $this->container['client_secret'] = $client_secret;

        return $this;
    }

    /**
     * Gets client_secret_expires_at
     *
     * @return int|null
     */
    public function getClientSecretExpiresAt()
    {
        return $this->container['client_secret_expires_at'];
    }

    /**
     * Sets client_secret_expires_at
     *
     * @param int|null $client_secret_expires_at SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire. The time is represented as the number of seconds from 1970-01-01T00:00:00Z as measured in UTC until the date/time of expiration.  This feature is currently not supported and it's value will always be set to 0.
     *
     * @return self
     */
    public function setClientSecretExpiresAt($client_secret_expires_at)
    {
        $this->container['client_secret_expires_at'] = $client_secret_expires_at;

        return $this;
    }

    /**
     * Gets client_uri
     *
     * @return string|null
     */
    public function getClientUri()
    {
        return $this->container['client_uri'];
    }

    /**
     * Sets client_uri
     *
     * @param string|null $client_uri ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
     *
     * @return self
     */
    public function setClientUri($client_uri)
    {
        $this->container['client_uri'] = $client_uri;

        return $this;
    }

    /**
     * Gets contacts
     *
     * @return string[]|null
     */
    public function getContacts()
    {
        return $this->container['contacts'];
    }

    /**
     * Sets contacts
     *
     * @param string[]|null $contacts contacts
     *
     * @return self
     */
    public function setContacts($contacts)
    {
        $this->container['contacts'] = $contacts;

        return $this;
    }

    /**
     * Gets created_at
     *
     * @return \DateTime|null
     */
    public function getCreatedAt()
    {
        return $this->container['created_at'];
    }

    /**
     * Sets created_at
     *
     * @param \DateTime|null $created_at CreatedAt returns the timestamp of the client's creation.
     *
     * @return self
     */
    public function setCreatedAt($created_at)
    {
        $this->container['created_at'] = $created_at;

        return $this;
    }

    /**
     * Gets frontchannel_logout_session_required
     *
     * @return bool|null
     */
    public function getFrontchannelLogoutSessionRequired()
    {
        return $this->container['frontchannel_logout_session_required'];
    }

    /**
     * Sets frontchannel_logout_session_required
     *
     * @param bool|null $frontchannel_logout_session_required Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.
     *
     * @return self
     */
    public function setFrontchannelLogoutSessionRequired($frontchannel_logout_session_required)
    {
        $this->container['frontchannel_logout_session_required'] = $frontchannel_logout_session_required;

        return $this;
    }

    /**
     * Gets frontchannel_logout_uri
     *
     * @return string|null
     */
    public function getFrontchannelLogoutUri()
    {
        return $this->container['frontchannel_logout_uri'];
    }

    /**
     * Sets frontchannel_logout_uri
     *
     * @param string|null $frontchannel_logout_uri RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.
     *
     * @return self
     */
    public function setFrontchannelLogoutUri($frontchannel_logout_uri)
    {
        $this->container['frontchannel_logout_uri'] = $frontchannel_logout_uri;

        return $this;
    }

    /**
     * Gets grant_types
     *
     * @return string[]|null
     */
    public function getGrantTypes()
    {
        return $this->container['grant_types'];
    }

    /**
     * Sets grant_types
     *
     * @param string[]|null $grant_types grant_types
     *
     * @return self
     */
    public function setGrantTypes($grant_types)
    {
        $this->container['grant_types'] = $grant_types;

        return $this;
    }

    /**
     * Gets jwks
     *
     * @return object|null
     */
    public function getJwks()
    {
        return $this->container['jwks'];
    }

    /**
     * Sets jwks
     *
     * @param object|null $jwks jwks
     *
     * @return self
     */
    public function setJwks($jwks)
    {
        $this->container['jwks'] = $jwks;

        return $this;
    }

    /**
     * Gets jwks_uri
     *
     * @return string|null
     */
    public function getJwksUri()
    {
        return $this->container['jwks_uri'];
    }

    /**
     * Sets jwks_uri
     *
     * @param string|null $jwks_uri URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
     *
     * @return self
     */
    public function setJwksUri($jwks_uri)
    {
        $this->container['jwks_uri'] = $jwks_uri;

        return $this;
    }

    /**
     * Gets logo_uri
     *
     * @return string|null
     */
    public function getLogoUri()
    {
        return $this->container['logo_uri'];
    }

    /**
     * Sets logo_uri
     *
     * @param string|null $logo_uri LogoURI is an URL string that references a logo for the client.
     *
     * @return self
     */
    public function setLogoUri($logo_uri)
    {
        $this->container['logo_uri'] = $logo_uri;

        return $this;
    }

    /**
     * Gets metadata
     *
     * @return object|null
     */
    public function getMetadata()
    {
        return $this->container['metadata'];
    }

    /**
     * Sets metadata
     *
     * @param object|null $metadata metadata
     *
     * @return self
     */
    public function setMetadata($metadata)
    {
        $this->container['metadata'] = $metadata;

        return $this;
    }

    /**
     * Gets owner
     *
     * @return string|null
     */
    public function getOwner()
    {
        return $this->container['owner'];
    }

    /**
     * Sets owner
     *
     * @param string|null $owner Owner is a string identifying the owner of the OAuth 2.0 Client.
     *
     * @return self
     */
    public function setOwner($owner)
    {
        $this->container['owner'] = $owner;

        return $this;
    }

    /**
     * Gets policy_uri
     *
     * @return string|null
     */
    public function getPolicyUri()
    {
        return $this->container['policy_uri'];
    }

    /**
     * Sets policy_uri
     *
     * @param string|null $policy_uri PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
     *
     * @return self
     */
    public function setPolicyUri($policy_uri)
    {
        $this->container['policy_uri'] = $policy_uri;

        return $this;
    }

    /**
     * Gets post_logout_redirect_uris
     *
     * @return string[]|null
     */
    public function getPostLogoutRedirectUris()
    {
        return $this->container['post_logout_redirect_uris'];
    }

    /**
     * Sets post_logout_redirect_uris
     *
     * @param string[]|null $post_logout_redirect_uris post_logout_redirect_uris
     *
     * @return self
     */
    public function setPostLogoutRedirectUris($post_logout_redirect_uris)
    {
        $this->container['post_logout_redirect_uris'] = $post_logout_redirect_uris;

        return $this;
    }

    /**
     * Gets redirect_uris
     *
     * @return string[]|null
     */
    public function getRedirectUris()
    {
        return $this->container['redirect_uris'];
    }

    /**
     * Sets redirect_uris
     *
     * @param string[]|null $redirect_uris redirect_uris
     *
     * @return self
     */
    public function setRedirectUris($redirect_uris)
    {
        $this->container['redirect_uris'] = $redirect_uris;

        return $this;
    }

    /**
     * Gets registration_access_token
     *
     * @return string|null
     */
    public function getRegistrationAccessToken()
    {
        return $this->container['registration_access_token'];
    }

    /**
     * Sets registration_access_token
     *
     * @param string|null $registration_access_token RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client.
     *
     * @return self
     */
    public function setRegistrationAccessToken($registration_access_token)
    {
        $this->container['registration_access_token'] = $registration_access_token;

        return $this;
    }

    /**
     * Gets registration_client_uri
     *
     * @return string|null
     */
    public function getRegistrationClientUri()
    {
        return $this->container['registration_client_uri'];
    }

    /**
     * Sets registration_client_uri
     *
     * @param string|null $registration_client_uri RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.
     *
     * @return self
     */
    public function setRegistrationClientUri($registration_client_uri)
    {
        $this->container['registration_client_uri'] = $registration_client_uri;

        return $this;
    }

    /**
     * Gets request_object_signing_alg
     *
     * @return string|null
     */
    public function getRequestObjectSigningAlg()
    {
        return $this->container['request_object_signing_alg'];
    }

    /**
     * Sets request_object_signing_alg
     *
     * @param string|null $request_object_signing_alg JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.
     *
     * @return self
     */
    public function setRequestObjectSigningAlg($request_object_signing_alg)
    {
        $this->container['request_object_signing_alg'] = $request_object_signing_alg;

        return $this;
    }

    /**
     * Gets request_uris
     *
     * @return string[]|null
     */
    public function getRequestUris()
    {
        return $this->container['request_uris'];
    }

    /**
     * Sets request_uris
     *
     * @param string[]|null $request_uris request_uris
     *
     * @return self
     */
    public function setRequestUris($request_uris)
    {
        $this->container['request_uris'] = $request_uris;

        return $this;
    }

    /**
     * Gets response_types
     *
     * @return string[]|null
     */
    public function getResponseTypes()
    {
        return $this->container['response_types'];
    }

    /**
     * Sets response_types
     *
     * @param string[]|null $response_types response_types
     *
     * @return self
     */
    public function setResponseTypes($response_types)
    {
        $this->container['response_types'] = $response_types;

        return $this;
    }

    /**
     * Gets scope
     *
     * @return string|null
     */
    public function getScope()
    {
        return $this->container['scope'];
    }

    /**
     * Sets scope
     *
     * @param string|null $scope Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
     *
     * @return self
     */
    public function setScope($scope)
    {

        if (!is_null($scope) && (!preg_match("/([a-zA-Z0-9\\.\\*]+\\s?)+/", $scope))) {
            throw new \InvalidArgumentException("invalid value for $scope when calling OAuth2Client., must conform to the pattern /([a-zA-Z0-9\\.\\*]+\\s?)+/.");
        }

        $this->container['scope'] = $scope;

        return $this;
    }

    /**
     * Gets sector_identifier_uri
     *
     * @return string|null
     */
    public function getSectorIdentifierUri()
    {
        return $this->container['sector_identifier_uri'];
    }

    /**
     * Sets sector_identifier_uri
     *
     * @param string|null $sector_identifier_uri URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.
     *
     * @return self
     */
    public function setSectorIdentifierUri($sector_identifier_uri)
    {
        $this->container['sector_identifier_uri'] = $sector_identifier_uri;

        return $this;
    }

    /**
     * Gets subject_type
     *
     * @return string|null
     */
    public function getSubjectType()
    {
        return $this->container['subject_type'];
    }

    /**
     * Sets subject_type
     *
     * @param string|null $subject_type SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
     *
     * @return self
     */
    public function setSubjectType($subject_type)
    {
        $this->container['subject_type'] = $subject_type;

        return $this;
    }

    /**
     * Gets token_endpoint_auth_method
     *
     * @return string|null
     */
    public function getTokenEndpointAuthMethod()
    {
        return $this->container['token_endpoint_auth_method'];
    }

    /**
     * Sets token_endpoint_auth_method
     *
     * @param string|null $token_endpoint_auth_method Requested Client Authentication method for the Token Endpoint. The options are client_secret_post, client_secret_basic, private_key_jwt, and none.
     *
     * @return self
     */
    public function setTokenEndpointAuthMethod($token_endpoint_auth_method)
    {
        $this->container['token_endpoint_auth_method'] = $token_endpoint_auth_method;

        return $this;
    }

    /**
     * Gets token_endpoint_auth_signing_alg
     *
     * @return string|null
     */
    public function getTokenEndpointAuthSigningAlg()
    {
        return $this->container['token_endpoint_auth_signing_alg'];
    }

    /**
     * Sets token_endpoint_auth_signing_alg
     *
     * @param string|null $token_endpoint_auth_signing_alg Requested Client Authentication signing algorithm for the Token Endpoint.
     *
     * @return self
     */
    public function setTokenEndpointAuthSigningAlg($token_endpoint_auth_signing_alg)
    {
        $this->container['token_endpoint_auth_signing_alg'] = $token_endpoint_auth_signing_alg;

        return $this;
    }

    /**
     * Gets tos_uri
     *
     * @return string|null
     */
    public function getTosUri()
    {
        return $this->container['tos_uri'];
    }

    /**
     * Sets tos_uri
     *
     * @param string|null $tos_uri TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
     *
     * @return self
     */
    public function setTosUri($tos_uri)
    {
        $this->container['tos_uri'] = $tos_uri;

        return $this;
    }

    /**
     * Gets updated_at
     *
     * @return \DateTime|null
     */
    public function getUpdatedAt()
    {
        return $this->container['updated_at'];
    }

    /**
     * Sets updated_at
     *
     * @param \DateTime|null $updated_at UpdatedAt returns the timestamp of the last update.
     *
     * @return self
     */
    public function setUpdatedAt($updated_at)
    {
        $this->container['updated_at'] = $updated_at;

        return $this;
    }

    /**
     * Gets userinfo_signed_response_alg
     *
     * @return string|null
     */
    public function getUserinfoSignedResponseAlg()
    {
        return $this->container['userinfo_signed_response_alg'];
    }

    /**
     * Sets userinfo_signed_response_alg
     *
     * @param string|null $userinfo_signed_response_alg JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.
     *
     * @return self
     */
    public function setUserinfoSignedResponseAlg($userinfo_signed_response_alg)
    {
        $this->container['userinfo_signed_response_alg'] = $userinfo_signed_response_alg;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


