package uiavatars

import (
	"fmt"
	"testing"
)

func TestCheckEngName(t *testing.T) {
	var fullname = "test123@gmail.com"
	initialsName := getInitials(fullname)

	fmt.Println(initialsName)

	// Test expected initials based on email format
	expectedInitials := "TE" // Adjust based on expected behavior
	if initialsName != expectedInitials {
		t.Errorf("expected %s, got %s", expectedInitials, initialsName)
	}
}

func TestCheckThaiName(t *testing.T) {
	var fullname = "อรไทย"
	initialsName := getInitials(fullname)

	fmt.Println(initialsName)

	// Test expected initials based on email format
	expectedInitials := "อร" // Adjust based on expected behavior
	if initialsName != expectedInitials {
		t.Errorf("expected %s, got %s", expectedInitials, initialsName)
	}
}
