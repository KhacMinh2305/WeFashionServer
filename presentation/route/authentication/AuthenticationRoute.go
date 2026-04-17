package authentication

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/authentication"
)

func RegisterAuthenticationRoute() {
	di.Router.POST("api/auth", authentication.FetchToken)
}

/***********************info***********************/
/*
RegisterAuthenticationRoute:
+) method : POST
+) endpoint: api/auth
+) body:
{
"id" : // user id
"token" : // old token or empty
}
*/
