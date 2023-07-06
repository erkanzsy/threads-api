package threads_api

type UserProfileData struct {
	Data struct {
		UserData struct {
			User struct {
				IsPrivate            bool   `json:"is_private"`
				ProfilePicURL        string `json:"profile_pic_url"`
				Username             string `json:"username"`
				HdProfilePicVersions []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"hd_profile_pic_versions"`
				IsVerified                  bool   `json:"is_verified"`
				Biography                   string `json:"biography"`
				BiographyWithEntities       any    `json:"biography_with_entities"`
				FollowerCount               int    `json:"follower_count"`
				ProfileContextFacepileUsers any    `json:"profile_context_facepile_users"`
				BioLinks                    []struct {
					URL string `json:"url"`
				} `json:"bio_links"`
				Pk       string `json:"pk"`
				FullName string `json:"full_name"`
				ID       any    `json:"id"`
			} `json:"user"`
		} `json:"userData"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
}
