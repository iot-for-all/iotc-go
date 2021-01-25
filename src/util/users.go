package util

import (
	"encoding/json"
)

type CustomUserCollection struct {

	// URL to get the next page of users.
	NextLink string `json:"nextLink,omitempty"`

	// The collection of users.
	// Required: true
	//Value []interface{} `json:"value"`
	Value []CustomUser `json:"value"`
}

type CustomRolesItems struct {

	// ID that specifies the role assignment for this role.
	// Required: true
	Role string `json:"role"`
}

type CustomPermission struct {

	// List of roles that specify the permissions to access the application.
	// Required: true
	// Min Items: 1
	Roles []*CustomRolesItems `json:"roles"`
}

type CustomUser struct {

	// Unique ID of the user.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Type of the user.
	// Required: true
	Type string `json:"type"`

	// List of roles that specify the permissions to access the application.
	CustomPermission

	// Email address of the user.
	// Format: email
	Email *string `json:"email,omitempty"`

	// The AAD object ID of the service principal.
	ObjectID *string `json:"objectId,omitempty"`

	// The AAD tenant ID of the service principal.
	TenantID *string `json:"tenantId,omitempty"`
}


// UnmarshalBinary interface implementation
func (m *CustomUserCollection) UnmarshalBinary(b []byte) error {
	var res CustomUserCollection

	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	*m = res
	return nil
}