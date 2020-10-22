package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers myDID-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/myDID/did", createDidHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/myDID/did", listDidHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/did/{key}", getDidHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/did", setDidHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/myDID/did", deleteDidHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/myDID/user", createUserHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/myDID/user", listUserHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/user/{key}", getUserHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/user", setUserHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/myDID/user", deleteUserHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/myDID/post", createPostHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/myDID/post", listPostHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/post/{key}", getPostHandler(cliCtx, "myDID")).Methods("GET")
		r.HandleFunc("/myDID/post", setPostHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/myDID/post", deletePostHandler(cliCtx)).Methods("DELETE")

		
}
