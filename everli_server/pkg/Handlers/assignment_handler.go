package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	impl "github.com/Mahaveer86619/Everli/pkg/Implementations"
	resp "github.com/Mahaveer86619/Everli/pkg/Response"
)

func CreateAssignmentController(w http.ResponseWriter, r *http.Request) {
	var my_assignment impl.Assignment
	err := json.NewDecoder(r.Body).Decode(&my_assignment)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_assignment, statusCode, err := impl.CreateAssignment(&my_assignment)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusCreated)
	successResponse.SetData(pg_assignment)
	successResponse.SetMessage("Assignment created successfully")
	successResponse.JSON(w)
}

func GetAssignmentController(w http.ResponseWriter, r *http.Request) {
	assignment_id := r.URL.Query().Get("id")
	if assignment_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	assignment, statusCode, err := impl.GetAssignment(assignment_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(assignment)
	successResponse.SetMessage("Assignment fetched successfully")
	successResponse.JSON(w)
}

func GetAssignmentsByMemberIdController(w http.ResponseWriter, r *http.Request) {
	member_id := r.URL.Query().Get("member_id")
	if member_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: member_id is required")
		failureResponse.JSON(w)
		return
	}

	assignments, statusCode, err := impl.GetAssignmentsByMemberId(member_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(assignments)
	successResponse.SetMessage("Assignments fetched successfully")
	successResponse.JSON(w)
}

func GetAssignmentsByEventIdController(w http.ResponseWriter, r *http.Request) {
	event_id := r.URL.Query().Get("event_id")
	if event_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: event_id is required")
		failureResponse.JSON(w)
		return
	}

	assignments, statusCode, err := impl.GetAssignmentsByEventId(event_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Print("No assignments found 1")
			successResponse := &resp.Success{}
			successResponse.SetStatusCode(http.StatusOK)
			successResponse.SetData([]impl.Assignment{})
			successResponse.SetMessage("Assignments fetched successfully")
			successResponse.JSON(w)
		}
		fmt.Print("No assignments found 2")
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(assignments)
	successResponse.SetMessage("Assignments fetched successfully")
	successResponse.JSON(w)
}

func UpdateAssignmentController(w http.ResponseWriter, r *http.Request) {
	var my_assignment impl.Assignment
	err := json.NewDecoder(r.Body).Decode(&my_assignment)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_assignment, statusCode, err := impl.UpdateAssignment(&my_assignment)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_assignment)
	successResponse.SetMessage("Assignment updated successfully")
	successResponse.JSON(w)
}

func DeleteAssignmentController(w http.ResponseWriter, r *http.Request) {
	assignment_id := r.URL.Query().Get("id")
	if assignment_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	statusCode, err := impl.DeleteAssignment(assignment_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetMessage("Assignment deleted successfully")
	successResponse.JSON(w)
}
