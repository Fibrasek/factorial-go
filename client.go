package factorial

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	authURL  = "https://api.factorialhr.com/oauth/authorize"
	tokenURL = "https://api.factorialhr.com/oauth/token"
)

// OAuthProvider keep the basic information
// needed for create and keep a connection
// using OAuth2
type OAuthProvider struct {
	conf *oauth2.Config
	ctx  context.Context
}

// OAuthProviderOption defines an option for a OAuthProvider.
type OAuthProviderOption func(*OAuthProvider)

// NewOAuthProvider will create a new OAuthProvider applying
// the given options
func NewOAuthProvider(opts ...OAuthProviderOption) *OAuthProvider {
	provider := &OAuthProvider{
		conf: &oauth2.Config{
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

// SetClientID will setup a new clientID for the OAuth2 config
func SetClientID(clientID string) OAuthProviderOption {
	return func(o *OAuthProvider) {
		o.conf.ClientID = clientID
	}
}

// SetClientSecret will setup a new client secret for the OAuth2 config
func SetClientSecret(clientSecret string) OAuthProviderOption {
	return func(o *OAuthProvider) {
		o.conf.ClientSecret = clientSecret
	}
}

// SetScopes will setup the scopes that the OAuth2 should ask
func SetScopes(scopes []string) OAuthProviderOption {
	return func(o *OAuthProvider) {
		o.conf.Scopes = scopes
	}
}

// SetRedirectURL will setup the redirect url needed for OAuth2
func SetRedirectURL(redirectURL string) OAuthProviderOption {
	return func(o *OAuthProvider) {
		o.conf.RedirectURL = redirectURL
	}
}

// GetAuthURL will return the return the url for redirect
// and start the OAuth2 process
func (o OAuthProvider) GetAuthURL(state string) string {
	return o.conf.AuthCodeURL(state)
}

// GetTokenFromCode method will find the token with the given code, this method
// should be called after a success callback received from auth process
func (o OAuthProvider) GetTokenFromCode(code string) (*oauth2.Token, error) {
	return o.conf.Exchange(o.ctx, code)
}

// RefreshToken method will refresh the given token
func (o OAuthProvider) RefreshToken(t *oauth2.Token) (*oauth2.Token, error) {
	return o.conf.TokenSource(o.ctx, t).Token()
}

// Client method will return a new http.Client for use in our calls, using
// the TokenSource will even refresh the token if needed
func (o OAuthProvider) Client(t *oauth2.Token) *http.Client {
	return o.conf.Client(o.ctx, t)
}
