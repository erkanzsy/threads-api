package threads_api

import (
	"fmt"
	"testing"
)

func TestUserIDFromUsername_ShouldReturnNoError(t *testing.T) {
	userId, _ := UserIDFromUsername("erkno_")

	if "51180116931" != userId {
		t.Errorf(fmt.Sprintf("userId: %s", userId))
	}
}

func TestUserProfile_ShouldReturnNoError(t *testing.T) {
	userProfile, _ := UserProfile("erkno_", "51180116931")

	if "erkno_" != userProfile.Data.UserData.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", userProfile.Data.UserData.User.Username))
	}
}

func TestUserThreads_ShouldReturnNoError(t *testing.T) {
	threads, _ := UserThreads("_junhoyeo", "5438123050")

	if "_junhoyeo" != threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username {
		t.Errorf(fmt.Sprintf("username: %s", threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username))
	}
}
