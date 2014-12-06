package steam

import (
	"testing"
)

const (
	accessKey = ""
)

func TestGetServerInfo(t *testing.T) {
	info, err := GetServerInfo()
	if err != nil {
		t.Fatal(err)
	}
	if len(info.ServerTimeString) == 0 {
		t.Error("ServerInfo.ServerTimeString is empty")
	}
}

func TestGetSupportedAPIList(t *testing.T) {
	list, err := GetSupportedAPIList(accessKey)
	if err != nil {
		t.Fatal(err)
	}
	if list.Interfaces == nil || len(list.Interfaces) == 0 {
		t.Error("apilist.interfaces is empty")
	}
}

func TestGetAppList(t *testing.T) {
	list, err := GetAppList()
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Apps) == 0 {
		t.Error("applist.apps is empty")
	}
}

func TestGetAppDetails(t *testing.T) {
	ids := []int{}
	details, err := GetAppDetails(ids...)
	if err != nil {
		t.Fatal(err)
	}
	if len(details) != len(ids) {
		t.Errorf("Requested %d AppDetails, got %d", len(details), len(ids))
	}
}
