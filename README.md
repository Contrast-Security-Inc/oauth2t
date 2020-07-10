![Go](https://github.com/Contrast-Security-Inc/oauth2t/workflows/Go/badge.svg) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Overview
A simple cmdline tool that will fetch an OAuth2 token and print it to the
cmdline.  This is useful in scripts such as:

    curl -v --request GET --url http://my.api -H "Authorization: Bearer `./oauth2t`"

It will only use the client-credentials OAuth2 flow since its designed to be
used in automated scripts and pipelines. It can take configuration for the
issuer, client_id, and client_secret via config file, environment variables, or
cmdline variables.

# Configuration
The `oauth2t` will need to know the client_id and client_secret specific to
your project when performing the OAuth2 client-credentials authentication flow.
These values would be available to you when you created your service account.

Configuration can be specific on the cmdline, in environment variables, and in
a config file, in highest precedence order respectively. `oauth2t` will look in all locations
and build the configuration after evaluating config sources.

## File
The config file can be specified as `oauth2t.json`, `oauth2.yaml`,
`oauth2.properties` to describe configuration in json, yaml, or java properties
respectively. `oauth2t` will look for the file in your current working
directory.

YAML example:
```
client_id: "YOUR CLIENT ID"
client_secret: "YOUR CLIENT SECRET"
```

ENV Variable example:
```
OAUTH2T_CLIENT_ID="YOUR CLIENT ID"
OAUTH2T_CLIENT_SECRET="YOUR CLIENT SECRET"
```
