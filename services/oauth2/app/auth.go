package app

import (
	"fmt"
	"net/http"
	"oauth/service"
)

func AuthEndpoint(rw http.ResponseWriter, req *http.Request) {

	oauth2 := service.GetProvider()
	ctx := req.Context()

	ar, err := oauth2.NewAuthorizeRequest(ctx, req)
	if err != nil {
		fmt.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		oauth2.WriteAuthorizeError(ctx, rw, ar, err)
		return
	}

	var requestedScopes string
	for _, this := range ar.GetRequestedScopes() {
		requestedScopes += fmt.Sprintf(`<li><input type="checkbox" name="scopes" value="%s">%s</li>`, this, this)
	}

	req.ParseForm()
	if req.PostForm.Get("username") == "" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write([]byte(`<h1>Login page</h1>`))
		rw.Write([]byte(fmt.Sprintf(`
			<p>Howdy! This is the log in page. For this example, it is enough to supply the username.</p>
			<form method="post">
				<p>
					By logging in, you consent to grant these scopes:
					<ul>%s</ul>
				</p>
				<input type="text" name="username" /> <small>try peter</small><br>
				<input type="submit">
			</form>
		`, requestedScopes)))
		return
	}

	for _, scope := range req.PostForm["scopes"] {
		ar.GrantScope(scope)
	}

	mySessionData := service.NewSession(req.PostForm.Get("username"), req.PostForm.Get("nonce"))
	response, err := oauth2.NewAuthorizeResponse(ctx, ar, mySessionData)
	if err != nil {
		fmt.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
		oauth2.WriteAuthorizeError(ctx, rw, ar, err)
		return
	}

	// Last but not least, send the response!
	oauth2.WriteAuthorizeResponse(ctx, rw, ar, response)
}
