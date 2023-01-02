package kite

import (
	"testing"
)

func TestGetUserProfile(t *testing.T) {
	t.Parallel()
	profile, err := getKite().GetUserProfile()
	if err != nil || profile.Email == "" || profile.UserID == "" {
		t.Errorf("Error while reading user profile. Error: %v", err)
	}
}

func TestGetUserMargins(t *testing.T) {
	t.Parallel()
	margins, err := getKite().GetUserMargins()
	if err != nil {
		t.Errorf("Error while reading user margins. Error: %v", err)
	}

	if !margins.Equity.Enabled || !margins.Commodity.Enabled {
		t.Errorf("Incorrect margin values.")
	}
}

func TestGetUserSegmentMargins(t *testing.T) {
	t.Parallel()
	margins, err := getKite().GetUserSegmentMargins("test")
	if err != nil {
		t.Errorf("Error while reading user margins. Error: %v", err)
	}

	if !margins.Enabled {
		t.Errorf("Incorrect segment margin values.")
	}
}
