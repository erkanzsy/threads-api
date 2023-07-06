package threads_api

import (
	"fmt"
	"testing"
)

var (
	USERNAME = "zuck"
	USERID   = "314216"
)

func TestUserIDFromUsername_ShouldReturnNoError(t *testing.T) {
	userId, _ := UserIDFromUsername(USERNAME)

	if USERID != userId {
		t.Errorf(fmt.Sprintf("userId: %s", userId))
	}
}

func TestUserProfile_ShouldReturnNoError(t *testing.T) {
	userProfile, _ := UserProfile(USERNAME, USERID)

	if USERNAME != userProfile.Data.UserData.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", userProfile.Data.UserData.User.Username))
	}
}

func TestUserThreads_ShouldReturnNoError(t *testing.T) {
	threads, _ := UserThreads(USERNAME, USERID)

	if USERNAME != threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username))
	}
}
