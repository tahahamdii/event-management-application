package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	impl "github.com/Mahaveer86619/Everli/pkg/Implementations"
	resp "github.com/Mahaveer86619/Everli/pkg/Response"
)

func CreateRoleController(w http.ResponseWriter, r *http.Request) {
	var my_role impl.Role
	err := json.NewDecoder(r.Body).Decode(&my_role)

	fmt.Print("Incoming role:")
	impl.PrintRole(&my_role)
	
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_role, statusCode, err := impl.CreateRole(&my_role)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusCreated)
	successResponse.SetData(pg_role)
	successResponse.SetMessage("Role created successfully")
	successResponse.JSON(w)
}

func GetRoleController(w http.ResponseWriter, r *http.Request) {
	role_id := r.URL.Query().Get("id")
	if role_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	pg_role, statusCode, err := impl.GetRoleById(role_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_role)
	successResponse.SetMessage("Role fetched successfully")
	successResponse.JSON(w)
}

func GetRolesByEventIdController(w http.ResponseWriter, r *http.Request) {
	event_id := r.URL.Query().Get("event_id")
	if event_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	pg_roles, statusCode, err := impl.GetRolesByEventId(event_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_roles)
	successResponse.SetMessage("Roles fetched successfully")
	successResponse.JSON(w)
}

func GetRolesByMemberIdController(w http.ResponseWriter, r *http.Request) {
	member_id := r.URL.Query().Get("member_id")
	if member_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	pg_roles, statusCode, err := impl.GetRolesByMemberId(member_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_roles)
	successResponse.SetMessage("Roles fetched successfully")
	successResponse.JSON(w)
}

func UpdateRoleController(w http.ResponseWriter, r *http.Request) {
	var my_role impl.Role
	err := json.NewDecoder(r.Body).Decode(&my_role)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_role, statusCode, err := impl.UpdateRole(&my_role)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_role)
	successResponse.SetMessage("Role updated successfully")
	successResponse.JSON(w)
}

func DeleteRoleController(w http.ResponseWriter, r *http.Request) {
	my_role := r.URL.Query().Get("id")
	if my_role == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	statusCode, err := impl.DeleteRole(my_role)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetMessage("Role deleted successfully")
	successResponse.JSON(w)
}
