## Usage

### Model

```go
type DefaultClient struct {
	ID             string   `json:"id"`
	Secret         []byte   `json:"client_secret,omitempty"`
	RotatedSecrets [][]byte `json:"rotated_secrets,omitempty"`
	RedirectURIs   []string `json:"redirect_uris"`
	GrantTypes     []string `json:"grant_types"`
	ResponseTypes  []string `json:"response_types"`
	Scopes         []string `json:"scopes"`
	Audience       []string `json:"audience"`
	Public         bool     `json:"public"`
}
```

### Start the server

```bash
#Start the server

go run main.go serve

```

### Register a client 

```bash
curl --location --request POST 'http://localhost:8080/api/v1/client/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ID":"my-client_1",
    "RedirectURIs":["http://localhost:3846/callback"],
    "ResponseTypes":["id_token", "code", "token", "id_token token", "code id_token", "code token", "code id_token token"],
    "GrantTypes": ["implicit", "refresh_token", "authorization_code", "password", "client_credentials"],
    "Scopes": ["fosite", "openid", "photos", "offline"]
}'

//get client secret
```

### Request authen code

```bash
http://localhost:8080/api/v1/auth

//query
client_id     = my-client
redirect_uri  = http://localhost:3846/callback
response_type = token id_token
scope         = fosite openid
state         = some-random-state-foobar
nonce         = some-random-nonce


curl http://localhost:8080/api/v1/auth?client_id=my-client&redirect_uri=http%3A%2F%2Flocalhost%3A3846%2Fcallback&response_type=token%20id_token&scope=fosite%20openid&state=some-random-state-foobar&nonce=some-random-nonce
```

### Request authen token

```go
   token, err := c.Exchange(context.Background(), req.URL.Query().Get("code"), opts...)
```
