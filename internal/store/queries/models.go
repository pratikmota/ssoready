// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package queries

import (
	"time"

	"github.com/google/uuid"
)

type ApiKey struct {
	ID                uuid.UUID
	AppOrganizationID uuid.UUID
	SecretValue       string
}

type AppOrganization struct {
	ID                 uuid.UUID
	GoogleHostedDomain *string
}

type AppSession struct {
	ID         uuid.UUID
	AppUserID  uuid.UUID
	CreateTime time.Time
	ExpireTime time.Time
	Token      string
}

type AppUser struct {
	ID                uuid.UUID
	AppOrganizationID uuid.UUID
	DisplayName       string
	Email             *string
}

type Environment struct {
	ID                uuid.UUID
	RedirectUrl       *string
	AppOrganizationID uuid.UUID
	DisplayName       *string
}

type Organization struct {
	ID            uuid.UUID
	EnvironmentID uuid.UUID
	ExternalID    *string
}

type SamlConnection struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	IdpRedirectUrl     *string
	IdpX509Certificate []byte
	IdpEntityID        *string
}

type SamlSession struct {
	ID                   uuid.UUID
	SamlConnectionID     uuid.UUID
	SecretAccessToken    *uuid.UUID
	SubjectID            *string
	SubjectIdpAttributes []byte
}

type SchemaMigration struct {
	Version int64
	Dirty   bool
}
