package handlers

import (
	"encoding/json"
	"net/http"

	impl "github.com/Mahaveer86619/Everli/pkg/Implementations"
	resp "github.com/Mahaveer86619/Everli/pkg/Response"
)

func AuthenticateUserController(w http.ResponseWriter, r *http.Request) {
	var creds impl.AuthenticatingUser
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	returned_creds, statusCode, err := impl.AuthenticateUser(&creds)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(returned_creds)
	successResponse.SetMessage("User Authenticated successfully")
	successResponse.JSON(w)
}

func RegisterUserController(w http.ResponseWriter, r *http.Request) {
	var creds impl.RegisteringUser
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	returned_creds, statusCode, err := impl.RegisterUser(&creds)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(returned_creds)
	successResponse.SetMessage("User Registered successfully")
	successResponse.JSON(w)
}

func SendPassResetCodeController(w http.ResponseWriter, r *http.Request) {
	var reqBody impl.SendPassResetCodeBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	statusCode, err := impl.SendPassResetCode(reqBody.Email)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(nil)
	successResponse.SetMessage("Code sent successfully")
	successResponse.JSON(w)
}

func CheckResetPassCodeController(w http.ResponseWriter, r *http.Request) {
	var reqBody impl.CheckResetPassCodeBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	code, statusCode, err := impl.CheckResetPassCode(reqBody.Code, reqBody.Email)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(code)
	successResponse.SetMessage("Code is valid")
	successResponse.JSON(w)
}

func ResetPassController(w http.ResponseWriter, r *http.Request) {
	var reqBody impl.ResetPassBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	statusCode, err := impl.ResetPassword(reqBody.Email, reqBody.Password)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(nil)
	successResponse.SetMessage("Password is updated successfully")
	successResponse.JSON(w)
}

func RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	var refreshingToken impl.RefreshingToken
	err := json.NewDecoder(r.Body).Decode(&refreshingToken)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	returned_tokens, statusCode, err := impl.RefreshToken(&refreshingToken)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(returned_tokens)
	successResponse.SetMessage("Token refreshed successfully")
	successResponse.JSON(w)
}
