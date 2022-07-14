package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"oauth/service"
)

func tokenEndpoint(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)
	data, _ := io.ReadAll(req.Body)
	fmt.Println(string(data))
	ctx := req.Context()

	oauth2 := service.GetProvider()
	mySessionData := service.NewSession("", "")

	accessRequest, err := oauth2.NewAccessRequest(ctx, req, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewAccessRequest: %+v", err)
		oauth2.WriteAccessError(ctx, rw, accessRequest, err)
		return
	}

	if accessRequest.GetGrantTypes().ExactOne("client_credentials") {
		for _, scope := range accessRequest.GetRequestedScopes() {
			accessRequest.GrantScope(scope)
		}
	}

	response, err := oauth2.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		log.Printf("Error occurred in NewAccessResponse: %+v", err)
		oauth2.WriteAccessError(ctx, rw, accessRequest, err)
		return
	}

	oauth2.WriteAccessResponse(ctx, rw, accessRequest, response)
}
