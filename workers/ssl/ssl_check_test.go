package ssl

import (
	"testing"

	"github.com/Bulut-Bilisimciler/certman/models"
)

func TestCheckSSL_Success(t *testing.T) {
	// prepare test model.Server
	mock := &models.Server{
		IpAddress:  "172.217.17.238",
		Port:       443,
		ServerName: "www.google.com",
		Debug:      true,
		Rules:      []string{},
		OnBeforeConnection: func() error {
			t.Log("OnBeforeConnection called")
			return nil
		},
		OnAfterConnection: func() error {
			t.Log("OnAfterConnection called")
			return nil
		},
		OnCheckSuccess: func() error {
			t.Log("OnCheckSuccess called")
			return nil
		},
		OnCheckFailure: func() error {
			t.Log("OnCheckFailure called")
			return nil
		},
	}

	// call CheckSSL
	if err := CheckSSL(mock); err != nil {
		t.Error("Test CheckSSL failed, expected nil, got ", err)
	}
}

func TestCheckSSL_Fail(t *testing.T) {
	// prepare test model.Server
	mock := &models.Server{
		IpAddress:  "172.217.17.238",
		Port:       443,
		ServerName: "facebook.com",
		Debug:      true,
		Rules:      []string{},
		OnBeforeConnection: func() error {
			t.Log("OnBeforeConnection called")
			return nil
		},
		OnAfterConnection: func() error {
			t.Log("OnAfterConnection called")
			return nil
		},
		OnCheckSuccess: func() error {
			t.Log("OnCheckSuccess called")
			return nil
		},
		OnCheckFailure: func() error {
			t.Log("OnCheckFailure called")
			return nil
		},
	}

	// call CheckSSL
	if err := CheckSSL(mock); err == nil {
		t.Error("Test CheckSSL fail situation failed, expected error, got nil")
	}
}
