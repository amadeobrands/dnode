// Implements REST API for multisig modules.
package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/dfinance/dnode/x/multisig/types"
)

// Registering routes in the REST API.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/call/{id}", types.ModuleName), getCall(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/calls", types.ModuleName), getCalls(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/unique/{unique}", types.ModuleName), getCallByUnique(cliCtx)).Methods("GET")
}

// GetUniqueCall godoc
// @Tags multisig
// @Summary Get call
// @Description Get call object by it's uniqueID
// @ID multisigGetUniqueCall
// @Accept  json
// @Produce json
// @Param uniqueID path string true "call uniqueID"
// @Success 200 {object} MSRespGetCall
// @Failure 400 {object} rest.ErrorResponse "Returned if the request doesn't have valid query params"
// @Failure 500 {object} rest.ErrorResponse "Returned on server error"
// @Router /multisig/unique/{uniqueID} [get]
func getCallByUnique(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		req := types.UniqueReq{UniqueId: vars["unique"]}

		bz, err := cliCtx.Codec.MarshalJSON(req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/unique", types.ModuleName), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// GetCall godoc
// @Tags multisig
// @Summary Get call
// @Description Get call object by it's ID
// @ID multisigGetCall
// @Accept  json
// @Produce json
// @Param id path uint true "call ID"
// @Success 200 {object} MSRespGetCall
// @Failure 400 {object} rest.ErrorResponse "Returned if the request doesn't have valid query params"
// @Failure 500 {object} rest.ErrorResponse "Returned on server error"
// @Router /multisig/call/{id} [get]
func getCall(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		req := types.CallReq{CallId: id}
		bz, err := cliCtx.Codec.MarshalJSON(req)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/call", types.ModuleName), bz)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// GetCalls godoc
// @Tags multisig
// @Summary Get calls
// @Description Get call objects
// @ID multisigGetCalls
// @Accept  json
// @Produce json
// @Success 200 {object} MSRespGetCalls
// @Failure 500 {object} rest.ErrorResponse "Returned on server error"
// @Router /multisig/calls [get]
func getCalls(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/calls", types.ModuleName), nil)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
