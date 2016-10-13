package github

import "fmt"

// ListKeys lists the verified public keys for all users
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#list-all-public-keys
func (s *AdminService) ListKeys(opt *ListOptions) ([]*Key, *Response, error) {
	u := fmt.Sprintf("admin/keys")
	u, err := addOptions(u, opt)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	keys := new([]*Key)
	resp, err := s.client.Do(req, keys)
	if err != nil {
		return nil, nil, err
	}

	return *keys, resp, err
}

// DeleteKey deletes a public key
//
// https://developer.github.com/enterprise/2.7/v3/users/administration/#delete-a-public-key
func (s *AdminService) DeleteKey(id int) (*Response, error) {
	u := fmt.Sprintf("admin/keys/%v", id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
