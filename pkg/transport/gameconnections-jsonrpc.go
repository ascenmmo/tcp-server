// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strings"
	"sync"
)

func (http *httpGameConnections) serveSetSendMessage(ctx *fiber.Ctx) (err error) {
	return http.serveMethod(ctx, "setsendmessage", http.setSendMessage)
}
func (http *httpGameConnections) setSendMessage(ctx *fiber.Ctx, requestBase baseJsonRPC) (responseBase *baseJsonRPC) {

	var err error
	var request requestGameConnectionsSetSendMessage

	methodCtx := ctx.UserContext()
	if requestBase.Params != nil {
		if err = json.Unmarshal(requestBase.Params, &request); err != nil {
			return makeErrorResponseJsonRPC(requestBase.ID, parseError, "request body could not be decoded: "+err.Error(), nil)
		}
	}
	if requestBase.Version != Version {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "incorrect protocol version: "+requestBase.Version, nil)
	}

	if _token := string(ctx.Request().Header.Peek("Token")); _token != "" {
		var token string
		token = _token
		request.Token = token
	}

	var response responseGameConnectionsSetSendMessage
	err = http.svc.SetSendMessage(methodCtx, request.Token, request.Message)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		code := internalError
		if errCoder, ok := err.(withErrorCode); ok {
			code = errCoder.Code()
		}
		return makeErrorResponseJsonRPC(requestBase.ID, code, err.Error(), err)
	}
	responseBase = &baseJsonRPC{
		ID:      requestBase.ID,
		Version: Version,
	}
	if responseBase.Result, err = json.Marshal(response); err != nil {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "response body could not be encoded: "+err.Error(), nil)
	}
	return
}
func (http *httpGameConnections) serveGetMessage(ctx *fiber.Ctx) (err error) {
	return http.serveMethod(ctx, "getmessage", http.getMessage)
}
func (http *httpGameConnections) getMessage(ctx *fiber.Ctx, requestBase baseJsonRPC) (responseBase *baseJsonRPC) {

	var err error
	var request requestGameConnectionsGetMessage

	methodCtx := ctx.UserContext()
	if requestBase.Params != nil {
		if err = json.Unmarshal(requestBase.Params, &request); err != nil {
			return makeErrorResponseJsonRPC(requestBase.ID, parseError, "request body could not be decoded: "+err.Error(), nil)
		}
	}
	if requestBase.Version != Version {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "incorrect protocol version: "+requestBase.Version, nil)
	}

	if _token := string(ctx.Request().Header.Peek("Token")); _token != "" {
		var token string
		token = _token
		request.Token = token
	}

	var response responseGameConnectionsGetMessage
	response.Messages, err = http.svc.GetMessage(methodCtx, request.Token)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		code := internalError
		if errCoder, ok := err.(withErrorCode); ok {
			code = errCoder.Code()
		}
		return makeErrorResponseJsonRPC(requestBase.ID, code, err.Error(), err)
	}
	responseBase = &baseJsonRPC{
		ID:      requestBase.ID,
		Version: Version,
	}
	if responseBase.Result, err = json.Marshal(response); err != nil {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "response body could not be encoded: "+err.Error(), nil)
	}
	return
}
func (http *httpGameConnections) serveRemoveUser(ctx *fiber.Ctx) (err error) {
	return http.serveMethod(ctx, "removeuser", http.removeUser)
}
func (http *httpGameConnections) removeUser(ctx *fiber.Ctx, requestBase baseJsonRPC) (responseBase *baseJsonRPC) {

	var err error
	var request requestGameConnectionsRemoveUser

	methodCtx := ctx.UserContext()
	if requestBase.Params != nil {
		if err = json.Unmarshal(requestBase.Params, &request); err != nil {
			return makeErrorResponseJsonRPC(requestBase.ID, parseError, "request body could not be decoded: "+err.Error(), nil)
		}
	}
	if requestBase.Version != Version {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "incorrect protocol version: "+requestBase.Version, nil)
	}

	if _token := string(ctx.Request().Header.Peek("Token")); _token != "" {
		var token string
		token = _token
		request.Token = token
	}

	var response responseGameConnectionsRemoveUser
	err = http.svc.RemoveUser(methodCtx, request.Token, request.UserID)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		code := internalError
		if errCoder, ok := err.(withErrorCode); ok {
			code = errCoder.Code()
		}
		return makeErrorResponseJsonRPC(requestBase.ID, code, err.Error(), err)
	}
	responseBase = &baseJsonRPC{
		ID:      requestBase.ID,
		Version: Version,
	}
	if responseBase.Result, err = json.Marshal(response); err != nil {
		return makeErrorResponseJsonRPC(requestBase.ID, parseError, "response body could not be encoded: "+err.Error(), nil)
	}
	return
}
func (http *httpGameConnections) serveMethod(ctx *fiber.Ctx, methodName string, methodHandler methodJsonRPC) (err error) {

	methodHTTP := ctx.Method()
	if methodHTTP != fiber.MethodPost {
		ctx.Response().SetStatusCode(fiber.StatusMethodNotAllowed)
		if _, err = ctx.WriteString("only POST method supported"); err != nil {
			return
		}
	}
	var request baseJsonRPC
	var response *baseJsonRPC
	if err = json.Unmarshal(ctx.Body(), &request); err != nil {
		return sendResponse(ctx, makeErrorResponseJsonRPC([]byte("\"0\""), parseError, "request body could not be decoded: "+err.Error(), nil))
	}
	methodNameOrigin := request.Method
	method := strings.ToLower(request.Method)
	if method != "" && method != methodName {
		return sendResponse(ctx, makeErrorResponseJsonRPC(request.ID, methodNotFoundError, "invalid method "+methodNameOrigin, nil))
	}
	response = methodHandler(ctx, request)
	if response != nil {
		return sendResponse(ctx, response)
	}
	return
}
func (http *httpGameConnections) doBatch(ctx *fiber.Ctx, requests []baseJsonRPC) (responses jsonrpcResponses) {

	if len(requests) > http.maxBatchSize {
		responses.append(makeErrorResponseJsonRPC(nil, invalidRequestError, "batch size exceeded", nil))
		return
	}
	if strings.EqualFold(ctx.Get("X-Sync-On"), "true") {
		for _, request := range requests {
			response := http.doSingleBatch(ctx, request)
			if request.ID != nil {
				responses.append(response)
			}
		}
		return
	}
	var wg sync.WaitGroup
	batchSize := http.maxParallelBatch
	if len(requests) < batchSize {
		batchSize = len(requests)
	}
	callCh := make(chan baseJsonRPC, batchSize)
	responses = make(jsonrpcResponses, 0, len(requests))
	for i := 0; i < batchSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for request := range callCh {
				response := http.doSingleBatch(ctx, request)
				if request.ID != nil {
					responses.append(response)
				}
			}
		}()
	}
	for idx := range requests {
		callCh <- requests[idx]
	}
	close(callCh)
	wg.Wait()
	return
}
func (http *httpGameConnections) serveBatch(ctx *fiber.Ctx) (err error) {

	var single bool
	var requests []baseJsonRPC
	methodHTTP := ctx.Method()
	if methodHTTP != fiber.MethodPost {
		ctx.Response().SetStatusCode(fiber.StatusMethodNotAllowed)
		if _, err = ctx.WriteString("only POST method supported"); err != nil {
			return
		}
		return
	}
	if err = json.Unmarshal(ctx.Body(), &requests); err != nil {
		var request baseJsonRPC
		if err = json.Unmarshal(ctx.Body(), &request); err != nil {
			return sendResponse(ctx, makeErrorResponseJsonRPC([]byte("\"0\""), parseError, "request body could not be decoded: "+err.Error(), nil))
		}
		single = true
		requests = append(requests, request)
	}
	if single {
		return sendResponse(ctx, http.doSingleBatch(ctx, requests[0]))
	}
	return sendResponse(ctx, http.doBatch(ctx, requests))
}
func (http *httpGameConnections) doSingleBatch(ctx *fiber.Ctx, request baseJsonRPC) (response *baseJsonRPC) {

	methodNameOrigin := request.Method
	method := strings.ToLower(request.Method)
	switch method {
	case "setsendmessage":
		return http.setSendMessage(ctx, request)
	case "getmessage":
		return http.getMessage(ctx, request)
	case "removeuser":
		return http.removeUser(ctx, request)
	default:
		return makeErrorResponseJsonRPC(request.ID, methodNotFoundError, "invalid method '"+methodNameOrigin+"'", nil)
	}
}
