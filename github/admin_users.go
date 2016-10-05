package github

import "fmt"

// AdminService handles communication with the administration
// related methods of the GitHub Enterprise API
//
// Enterprise API docs: https://developer.github.com/enterprise/2.7/v3/users/administration/
type AdminService service

/*
// User represents a GitHub user that administrative commands are executed upon
type User struct {
	Login      *string `json:"login,omitempty"`
	ID         *int    `json:"id,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitempty"`
	HTMLURL    *string `json:"html_url,omitempty"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	Type       *string `json:"type,omitempty"`
	SiteAdmin  *bool   `json:"site_admin,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}
*/

// Create creates a new GitHub user -- if external authN is used, ensure the login name
// matches in the external system
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#create-a-new-user
func (s *AdminService) Create(user *User) (*User, *Response, error) {
	u := "admin/users"

	req, err := s.client.NewRequest("POST", u, user)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

// Rename modifies the login for a user -- if external authN is used, ensure the login name
// matches in the external system
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#rename-an-existing-user
func (s *AdminService) Rename(newLogin string, user *User) (*User, *Response, error) {
	u := fmt.Sprintf("/admin/users/%v", newLogin)

	req, err := s.client.NewRequest("PATCH", u, user)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)

	return uResp, resp, err
}

// Promote promotes an ordinary user to a site administrator
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#promote-an-ordinary-user-to-a-site-administrator
func (s *AdminService) Promote(user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/site_admin", user)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Length", "0")

	return s.client.Do(req, nil)
}

// Demote demotes a site administrator to an ordinary user
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#demote-a-site-administrator-to-an-ordinary-user
func (s *AdminService) Demote(user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/site_admin", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Suspend suspends a user from the enterprise instance
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#suspend-a-user
func (s *AdminService) Suspend(user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/suspended", user)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Length", "0")

	return s.client.Do(req, nil)
}

// Unsuspend reinstates a user to the enterprise instance
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#unsuspend-a-user
func (s *AdminService) Unsuspend(user string) (*Response, error) {
	u := fmt.Sprintf("users/%v/suspended", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Delete deletes a user including all of their repositories, gists, applications and
// personal settings
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#delete-a-user
func (s *AdminService) Delete(user string) (*Response, error) {
	u := fmt.Sprintf("admin/users/%v", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}