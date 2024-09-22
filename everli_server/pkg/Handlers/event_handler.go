package handlers

import (
	"encoding/json"
	"net/http"

	impl "github.com/Mahaveer86619/Everli/pkg/Implementations"
	resp "github.com/Mahaveer86619/Everli/pkg/Response"
)

func CreateEventController(w http.ResponseWriter, r *http.Request) {
	var my_event impl.Event
	err := json.NewDecoder(r.Body).Decode(&my_event)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_event, statusCode, err := impl.CreateEvent(&my_event)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusCreated)
	successResponse.SetData(pg_event)
	successResponse.SetMessage("Event created successfully")
	successResponse.JSON(w)
}

func GetEventController(w http.ResponseWriter, r *http.Request) {
	event_id := r.URL.Query().Get("id")
	if event_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	pg_event, statusCode, err := impl.GetEvent(event_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_event)
	successResponse.SetMessage("Event fetched successfully")
	successResponse.JSON(w)
}

func GetAllEventsController(w http.ResponseWriter, r *http.Request) {
	events, statusCode, err := impl.GetAllEvents()
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(statusCode)
	successResponse.SetData(events)
	successResponse.SetMessage("Events fetched successfully")
	successResponse.JSON(w)
}

func UpdateEventController(w http.ResponseWriter, r *http.Request) {
	var my_event impl.Event
	err := json.NewDecoder(r.Body).Decode(&my_event)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body")
		failureResponse.JSON(w)
		return
	}

	pg_event, statusCode, err := impl.UpdateEvent(&my_event)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetData(pg_event)
	successResponse.SetMessage("Event updated successfully")
	successResponse.JSON(w)
}

func DeleteEventController(w http.ResponseWriter, r *http.Request) {
	event_id := r.URL.Query().Get("id")
	if event_id == "" {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(http.StatusBadRequest)
		failureResponse.SetMessage("Invalid request body: id is required")
		failureResponse.JSON(w)
		return
	}

	statusCode, err := impl.DeleteEvent(event_id)
	if err != nil {
		failureResponse := resp.Failure{}
		failureResponse.SetStatusCode(statusCode)
		failureResponse.SetMessage(err.Error())
		failureResponse.JSON(w)
		return
	}

	successResponse := &resp.Success{}
	successResponse.SetStatusCode(http.StatusOK)
	successResponse.SetMessage("Event deleted successfully")
	successResponse.JSON(w)
}
