# ory-oathkeeper-client

OryOathkeeperClient - the Ruby gem for the Ory Oathkeeper API

Documentation for all of Ory Oathkeeper's APIs.


This SDK is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: v0.38.20-beta.1
- Package version: v0.38.20-beta.1
- Build package: org.openapitools.codegen.languages.RubyClientCodegen

## Installation

### Build a gem

To build the Ruby code into a gem:

```shell
gem build ory-oathkeeper-client.gemspec
```

Then either install the gem locally:

```shell
gem install ./ory-oathkeeper-client-v0.38.20-beta.1.gem
```

(for development, run `gem install --dev ./ory-oathkeeper-client-v0.38.20-beta.1.gem` to install the development dependencies)

or publish the gem to a gem hosting service, e.g. [RubyGems](https://rubygems.org/).

Finally add this to the Gemfile:

    gem 'ory-oathkeeper-client', '~> v0.38.20-beta.1'

### Install from Git

If the Ruby gem is hosted at a git repository: https://github.com/ory/sdk, then add the following in the Gemfile:

    gem 'ory-oathkeeper-client', :git => 'https://github.com/ory/sdk.git'

### Include the Ruby code directly

Include the Ruby code directly using `-I` as follows:

```shell
ruby -Ilib script.rb
```

## Getting Started

Please follow the [installation](#installation) procedure and then run the following code:

```ruby
# Load the gem
require 'ory-oathkeeper-client'

api_instance = OryOathkeeperClient::ApiApi.new

begin
  #Access Control Decision API
  api_instance.decisions
rescue OryOathkeeperClient::ApiError => e
  puts "Exception when calling ApiApi->decisions: #{e}"
end

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OryOathkeeperClient::ApiApi* | [**decisions**](docs/ApiApi.md#decisions) | **GET** /decisions | Access Control Decision API
*OryOathkeeperClient::ApiApi* | [**get_rule**](docs/ApiApi.md#get_rule) | **GET** /rules/{id} | Retrieve a Rule
*OryOathkeeperClient::ApiApi* | [**get_well_known_json_web_keys**](docs/ApiApi.md#get_well_known_json_web_keys) | **GET** /.well-known/jwks.json | Lists Cryptographic Keys
*OryOathkeeperClient::ApiApi* | [**list_rules**](docs/ApiApi.md#list_rules) | **GET** /rules | List All Rules
*OryOathkeeperClient::MetadataApi* | [**get_version**](docs/MetadataApi.md#get_version) | **GET** /version | Return Running Software Version.
*OryOathkeeperClient::MetadataApi* | [**is_alive**](docs/MetadataApi.md#is_alive) | **GET** /health/alive | Check HTTP Server Status
*OryOathkeeperClient::MetadataApi* | [**is_ready**](docs/MetadataApi.md#is_ready) | **GET** /health/ready | Check HTTP Server and Database Status


## Documentation for Models

 - [OryOathkeeperClient::GenericError](docs/GenericError.md)
 - [OryOathkeeperClient::HealthNotReadyStatus](docs/HealthNotReadyStatus.md)
 - [OryOathkeeperClient::HealthStatus](docs/HealthStatus.md)
 - [OryOathkeeperClient::InlineResponse200](docs/InlineResponse200.md)
 - [OryOathkeeperClient::InlineResponse2001](docs/InlineResponse2001.md)
 - [OryOathkeeperClient::InlineResponse503](docs/InlineResponse503.md)
 - [OryOathkeeperClient::JsonWebKey](docs/JsonWebKey.md)
 - [OryOathkeeperClient::JsonWebKeySet](docs/JsonWebKeySet.md)
 - [OryOathkeeperClient::Rule](docs/Rule.md)
 - [OryOathkeeperClient::RuleHandler](docs/RuleHandler.md)
 - [OryOathkeeperClient::RuleMatch](docs/RuleMatch.md)
 - [OryOathkeeperClient::Upstream](docs/Upstream.md)
 - [OryOathkeeperClient::Version](docs/Version.md)


## Documentation for Authorization

 All endpoints do not require authorization.

