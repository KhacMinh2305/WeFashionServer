package authentication

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/service"
)

func RegisterAuthenticationRoute() {
	di.Router.POST("api/auth", service.FetchToken)
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
