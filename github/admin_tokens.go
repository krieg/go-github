package github

import "fmt"

// Token represents an impersonation OAuth token used for executing
// administrative commands on a user
type Token struct {
	ID             *int       `json:"id,omitempty"`
	URL            *string    `json:"url,omitempty"`
	Scopes         []string   `json:"events,omitempty"`
	Token          *string    `json:"token,omitempty"`
	TokenLastEight *string    `json:"token_last_eight,omitempty"`
	HashedToken    *string    `json:"hashed_token,omitempty"`
	App            *App       `json:"app,omitempty"`
	Note           *string    `json:"note,omitempty"`
	NoteURL        *string    `json:"note_rul,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	Fingerprint    *string    `json:"fingerprint,omitempty"`
}

func (t Token) String() string {
	return Stringify(t)
}

// App represents an application
type App struct {
	URL      *string `json:"url,omitempty"`
	Name     *string `json:"name,omitempty"`
	ClientID *string `json:"client_id,omitempty"`
}

// CreateToken creates an impersonation token for the given user
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#create-an-impersonation-oauth-token
func (s *AdminService) CreateToken(user string, token *Token) (*Token, *Response, error) {
	u := fmt.Sprintf("admin/users/%v/authorizations", user)

	req, err := s.client.NewRequest("POST", u, token)
	if err != nil {
		return nil, nil, err
	}

	tResp := new(Token)
	resp, err := s.client.Do(req, tResp)
	if err != nil {
		return nil, resp, err
	}

	return tResp, resp, err
}

// DeleteToken deletes an impersonation token for the given user
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#delete-an-impersonation-oauth-token
func (s *AdminService) DeleteToken(user string) (*Response, error) {
	u := fmt.Sprintf("admin/users/%v/authorizations", user)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}