package Implementations

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	db "github.com/Mahaveer86619/Everli/pkg/DB"
	"github.com/google/uuid"
)

type Invitation struct {
	Id         string `json:"id"`
	EventId    string `json:"event_id"`
	Code       string `json:"code"`
	Role       string `json:"role"`
	ExpiryDate string `json:"expiry"`
	CreatedAt  string `json:"created_at"`
}

func CreateInvitation(invitation *Invitation) (*Invitation, int, error) {
	conn := db.GetDBConnection()

	query := `
		INSERT INTO invitations (id, event_id, code, role, expiry, created_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING *
	`

	// Generate a unique ID for the event
	invitation.Id = uuid.New().String()
	invitation.Code = generateRandomNumber(6)

	if err := conn.QueryRow(
		query,
		invitation.Id,
		invitation.EventId,
		invitation.Code,
		invitation.Role,
		invitation.ExpiryDate,
		invitation.CreatedAt,
	).Scan(
		&invitation.Id,
		&invitation.Code,
		&invitation.EventId,
		&invitation.ExpiryDate,
		&invitation.Role,
		&invitation.CreatedAt,
	); err != nil {
		log.Panic(err)
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating checkpoint: %w", err)
	}

	return invitation, http.StatusCreated, nil
}

func GetInvitation(code string) (*Invitation, int, error) {
	conn := db.GetDBConnection()

	query := `
	  SELECT *
	  FROM invitations
	  WHERE code = $1
	`
	var invitation Invitation
	if err := conn.QueryRow(
		query,
		code,
	).Scan(
		&invitation.Id,
		&invitation.Code,
		&invitation.EventId,
		&invitation.ExpiryDate,
		&invitation.Role,
		&invitation.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, fmt.Errorf("invitation not found with code: %s", code)
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("error scanning row: %w", err)
	}

	return &invitation, http.StatusOK, nil
}

func UpdateInvitation(invitation *Invitation) (*Invitation, int, error) {
	conn := db.GetDBConnection()

	query := `UPDATE invitations 
	SET event_id = $1, code = $2, role = $3, expiry = $4, created_at = $5, updated_at = $6
	WHERE id = $7 RETURNING *
	`

	if err := conn.QueryRow(
		query,
		invitation.EventId,
		invitation.Code,
		invitation.Role,
		invitation.ExpiryDate,
		invitation.CreatedAt,
		invitation.Id,
	).Scan(
		&invitation.Id,
		&invitation.Code,
		&invitation.EventId,
		&invitation.ExpiryDate,
		&invitation.Role,
		&invitation.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, fmt.Errorf("invitation not found with id: %s", invitation.Id)
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("error scanning row: %w", err)
	}

	return invitation, http.StatusOK, nil
}

func DeleteInvitation(invitation_id string) (int, error) {
	conn := db.GetDBConnection()

	query := `SELECT * FROM invitations WHERE id = $1`
	del_query := "DELETE FROM invitations WHERE id = $1"

	var invitation Invitation
	if err := conn.QueryRow(
		query,
		invitation_id,
	).Scan(
		&invitation.Id,
		&invitation.Code,
		&invitation.EventId,
		&invitation.ExpiryDate,
		&invitation.Role,
		&invitation.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, fmt.Errorf("invitation not found with id: %s", invitation_id)
		}
		return http.StatusInternalServerError, fmt.Errorf("error scanning row: %w", err)
	} else {
		_, err = conn.Exec(del_query, invitation_id)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("error deleting row: %w", err)
		}
	}

	return http.StatusNoContent, nil
}

func generateRandomNumber(length int) string {
	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length)) - 1
	randomNumber := rand.Intn(max-min) + min
	return strconv.Itoa(randomNumber)
}

// func printInvitation(invitation *Invitation) {
// 	fmt.Println("invitation:")
// 	fmt.Printf(`
// 	ID: %s
// 	Event ID: %s
// 	Code: %s
// 	Role: %s
// 	Expiry: %s
// 	Created At: %s
// 	`, invitation.Id, invitation.EventId, invitation.Code, invitation.Role, invitation.ExpiryDate, invitation.CreatedAt)
// }
