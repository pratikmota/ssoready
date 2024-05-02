package idformat

import "github.com/ssoready/prettyuuid"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

var (
	AppUser        = prettyuuid.MustNewFormat("app_user_", alphabet)
	APIKey         = prettyuuid.MustNewFormat("api_key_", alphabet)
	Environment    = prettyuuid.MustNewFormat("env_", alphabet)
	Organization   = prettyuuid.MustNewFormat("org_", alphabet)
	SAMLConnection = prettyuuid.MustNewFormat("saml_conn_", alphabet)
	SAMLLoginEvent = prettyuuid.MustNewFormat("saml_login_event_", alphabet)
	SAMLAccessCode = prettyuuid.MustNewFormat("saml_access_code_", alphabet)
)
