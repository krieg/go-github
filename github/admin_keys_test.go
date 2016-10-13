package github

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAdminService_ListKeys(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"page": "2"})
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 2}
	keys, _, err := client.Admin.ListKeys(opt)
	if err != nil {
		t.Errorf("Admin.ListKeys returned error: %v", err)
	}

	want := []*Key{{ID: Int(1)}}
	if !reflect.DeepEqual(keys, want) {
		t.Errorf("Admin.ListKeys returned %+v, want %+v", keys, want)
	}
}

func TestAdminService_DeleteKey(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/keys/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Admin.DeleteKey(1)
	if err != nil {
		t.Errorf("Admin.DeleteKey returned error: %v", err)
	}
}