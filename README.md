# Go Auth proxy

### Create Proxy

```go
proxy := oauthProxy.Proxy{
  OAuthGrantType:        "password_domain",
  OAuthRefreshGrantType: "refresh_domain",
 
  Client    YourSpecificClient // client.OAuthClientInterface
  Converter YourSpecificConverter // converter.JwtConverterInterface
}
```

### Methods
```go
// Login
cryptedAuthData, userPayload, err := proxy.Login("username","password") // string, map[string]interface{}, error

// Logout
OAuthData := "crypted_and_converted_Access_token_array_to_string"
err := proxy.Logout(OAuthData) // error

// Logout All
OAuthData := "crypted_and_converted_Access_token_array_to_string"
err := proxy.LogoutAll(OAuthData) // error

// Check
OAuthData := "crypted_and_converted_Access_token_array_to_string"
err := proxy.Check(OAuthData) // error

// Refresh
OAuthData := "crypted_and_converted_Refresh_token_array_to_string"
cryptedAccessAndRefreshToken, err := proxy.Refres(OAuthData) // string, error

```
