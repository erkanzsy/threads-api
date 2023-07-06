package threads_api

import (
	"fmt"
	"testing"
)

func TestUserIDFromUsername_ShouldReturnNoError(t *testing.T) {
	userId, _ := UserIDFromUsername("zuck")

	if "314216" != userId {
		t.Errorf(fmt.Sprintf("userId: %s", userId))
	}
}

func TestUserProfile_ShouldReturnNoError(t *testing.T) {
	userProfile, _ := UserProfile("zuck", "314216")

	if "zuck" != userProfile.Data.UserData.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", userProfile.Data.UserData.User.Username))
	}
}

func TestUserThreads_ShouldReturnNoError(t *testing.T) {
	threads, _ := UserThreads("zuck", "314216")

	if "zuck" != threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username))
	}
}
