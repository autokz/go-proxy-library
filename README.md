# go_proxy

Go proxy libriry

## Example

In oauthProxy/main.go

### Create Proxy

```go
proxy := oauthProxy.Proxy{
  BaseURL:               "http://localhost:8080",
  CheckURL:              "/oauth/check",
  OAuthURL:              "/oauth/auth",
  LogoutURL:             "/oauth/logout",
  LogoutAllURL:          "/oauth/logout/all",
  DomainURL:             "test.com",
  OAuthGrantType:        "password_domain",
  OAuthRefreshGrantType: "refresh_domain",

  // If You need override functions ... 
  Client    client.NetClientInterface
  Converter converter.JwtConverterInterface
}

proxy.Init() // If you want init default Converter and NetClient functions
```

### Methods
```go
// Login
cryptedAuthData, userPayload, err := proxy.Login("username","password") // string, map[string]interface{}, error

// Logout
OAuthData := "crypted_and_converted_Access_token_array_to_string"
ok, err := proxy.Logout(OAuthData) // bool, error

// Logout All
OAuthData := "crypted_and_converted_Access_token_array_to_string"
ok, err := proxy.LogoutAll(OAuthData) // bool, error

// Check
OAuthData := "crypted_and_converted_Access_token_array_to_string"
notModifiedSameOAuthData, err := proxy.Check(OAuthData) // string, error

// Refresh
OAuthData := "crypted_and_converted_Refresh_token_array_to_string"
cryptedAccessAndRefreshToken, err := proxy.Refres(OAuthData) // string, error

```
