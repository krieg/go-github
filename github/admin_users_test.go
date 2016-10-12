package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAdminService_Create_validUser(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Login: String("l"), Email: String("e")}

	mux.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	user, _, err := client.Admin.Create(input)
	if err != nil {
		t.Errorf("Admin.Create returned error: %v", err)
	}

	want := &User{ID: Int(1)}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Admin.Create returned %+v, want %+v", user, want)
	}
}

func TestAdminService_Rename(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Login: String("l")}

	mux.HandleFunc("/admin/users/u", func(w http.ResponseWriter, r *http.Request) {
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	user, _, err := client.Admin.Rename("u", input)
	if err != nil {
		t.Errorf("Admin.Rename returned error: %v", err)
	}

	want := &User{ID: Int(1)}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Admin.Rename returned %+v, want %+v", user, want)
	}
}

func TestAdminService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/users/u", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Admin.Delete("u")
	if err != nil {
		t.Errorf("Admin.Delete returned error: %v", err)
	}
}