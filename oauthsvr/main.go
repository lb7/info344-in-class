package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"fmt"
	"io"
)

const (
	headerContentType = "Content-Type"
	headerAccept      = "Accept"
)

const (
	//githubCurrentUserAPI is the URL for GitHub's current user API
	githubCurrentUserAPI = "https://api.github.com/user"

	//acceptGitHubV3JSON is the value you should include in
	//the Accept header when making requests to the GitHub API
	acceptGitHubV3JSON = "application/vnd.github.v3+json"
)

const (
	apiSignIn = "/oauth/signin"
	apiReply  = "/oauth/reply"
)

//HandlerContext is the receiver for our handler methods
//and contains various global values our handlers will need
type HandlerContext struct {
	//oauthConfig is the OAuth configuration for GitHub
	oauthConfig *oauth2.Config
	//stateCache is a cache of previously-generated OAuth state values
	stateCache *cache.Cache
}

//newStateValue returns a base64-encoded crypto-random value
//suitable for using as the `state` parameter in an OAuth2
//authorization request
func newStateValue() string {
	buf := make([]byte, 0, 32)
	if _, err := rand.Read(buf); err != nil {
		panic("error generating random bytes")
	}
	return base64.URLEncoding.EncodeToString(buf)
}

//OAuthSignInHandler handles requests for the oauth sign-on API
func (ctx *HandlerContext) OAuthSignInHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: implement this handler by:
	// - generating a new state value
	// - adding it to the cache (default timeout)
	// - redirecting the client to the authorization URL
	//   returned from the OAuth config
	state := newStateValue()
	log.Println(state)
	ctx.stateCache.Add(state, nil, cache.DefaultExpiration)

	redirURL := ctx.oauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}

//OAuthReplyHandler handles requests made after authenticating
//with the OAuth provider, and authorizing our application
func (ctx *HandlerContext) OAuthReplyHandler(w http.ResponseWriter, r *http.Request) {
	//This handler is called after the OAuth provider redirects the client
	//back to our server. The query string may contain either these parameters:
	// - code = authorization code
	// - state = state value we sent to the server
	//OR these params if there was an error:
	// - error = an error code: https://tools.ietf.org/html/rfc6749#section-4.1.2.1
	// - error_description (optional) = human-readable error message
	// - error_uri (optional) = human-readable web page

	//TODO: implement this handler by doing the following:
	// - if the query string contains an "error" parameter, handle the error
	// - if the "state" query string param is missing or is not found in
	//   the cache, respond with an error
	// - if it is found, delete it from the cache so that it can't be used again
	// - use the `.Exchange()` method on the OAuth config to get an access token
	// - use the token to get a new http.Client you can use to make requests on
	//   behalf of the authenticated user
	// - use that client to get the user's profile (see constants above)
	qsParams := r.URL.Query()
	if len(qsParams.Get("error")) > 0 {
		errorDesc := qsParams.Get("error_description")
		if len(errorDesc) == 0 {
			errorDesc = "Error signing in: " + qsParams.Get("error")
		}
		http.Error(w, fmt.Sprintf("Error signing in: %s", errorDesc), http.StatusInternalServerError)
		return
	}

	stateReturned := qsParams.Get("state")
	if _, found := ctx.stateCache.Get(stateReturned); !found {
		http.Error(w, "invalid state value returned from OAuth Provider", http.StatusBadRequest)
		return
	}
	ctx.stateCache.Delete(stateReturned)

	token, err := ctx.oauthConfig.Exchange(oauth2.NoContext, qsParams.Get("code"))
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting access token: %v", err), http.StatusInternalServerError)
		return
	}

	client := ctx.oauthConfig.Client(oauth2.NoContext, token)
	profileRequest, _ := http.NewRequest(http.MethodGet, githubCurrentUserAPI, nil)
	profileRequest.Header.Add(headerAccept, acceptGitHubV3JSON)

	profileResponse, err := client.Do(profileRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting profile: %v", err), http.StatusInternalServerError)
	}
	defer profileResponse.Body.Close()

	//After obtaining the current user's profile, this is where you
	//would typically create a new User record in your system,
	//and begin a new authenticated Session for that user.
	//For purposes of this demo, we will just stream the profile
	//to the client so that we can see what it contains
	w.Header().Add(headerContentType, profileResponse.Header.Get(headerContentType))
	io.Copy(w, profileResponse.Body)
}

func requireEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {
	addr := requireEnv("ADDR")
	clientID := requireEnv("CLIENT_ID")
	clientSecret := requireEnv("CLIENT_SECRET")

	ctx := &HandlerContext{
		oauthConfig: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"read:user"},
			RedirectURL:  "https://" + addr + apiReply,
			Endpoint:     github.Endpoint,
		},
		stateCache: cache.New(5*time.Minute, 10*time.Second),
	}

	mux := http.NewServeMux()
	mux.HandleFunc(apiSignIn, ctx.OAuthSignInHandler)
	mux.HandleFunc(apiReply, ctx.OAuthReplyHandler)

	log.Printf("API server is listening at https://%s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "./tls/fullchain.pem", "./tls/privkey.pem", mux))
}
