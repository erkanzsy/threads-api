package threads_api

type UserThreadsData struct {
	Data struct {
		MediaData struct {
			Threads []struct {
				ThreadItems []struct {
					Post struct {
						User struct {
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
							ID            any    `json:"id"`
							IsVerified    bool   `json:"is_verified"`
							Pk            string `json:"pk"`
						} `json:"user"`
						ImageVersions2 struct {
							Candidates []struct {
								Height   int    `json:"height"`
								URL      string `json:"url"`
								Width    int    `json:"width"`
								Typename string `json:"__typename"`
							} `json:"candidates"`
						} `json:"image_versions2"`
						OriginalWidth      int    `json:"original_width"`
						OriginalHeight     int    `json:"original_height"`
						VideoVersions      []any  `json:"video_versions"`
						CarouselMedia      any    `json:"carousel_media"`
						CarouselMediaCount any    `json:"carousel_media_count"`
						Pk                 string `json:"pk"`
						HasAudio           any    `json:"has_audio"`
						TextPostAppInfo    struct {
							LinkPreviewAttachment any `json:"link_preview_attachment"`
							ShareInfo             struct {
								QuotedPost   any `json:"quoted_post"`
								RepostedPost struct {
									Pk   string `json:"pk"`
									User struct {
										ProfilePicURL string `json:"profile_pic_url"`
										Username      string `json:"username"`
										ID            any    `json:"id"`
										IsVerified    bool   `json:"is_verified"`
										Pk            string `json:"pk"`
									} `json:"user"`
									ImageVersions2 struct {
										Candidates []any `json:"candidates"`
									} `json:"image_versions2"`
									OriginalWidth      int   `json:"original_width"`
									OriginalHeight     int   `json:"original_height"`
									VideoVersions      []any `json:"video_versions"`
									CarouselMedia      any   `json:"carousel_media"`
									CarouselMediaCount any   `json:"carousel_media_count"`
									HasAudio           any   `json:"has_audio"`
									TextPostAppInfo    struct {
										LinkPreviewAttachment any `json:"link_preview_attachment"`
										ShareInfo             struct {
											QuotedPost struct {
												TextPostAppInfo struct {
													IsPostUnavailable bool `json:"is_post_unavailable"`
													ShareInfo         struct {
														QuotedPost any `json:"quoted_post"`
													} `json:"share_info"`
													DirectReplyCount      int `json:"direct_reply_count"`
													LinkPreviewAttachment struct {
														DisplayURL string `json:"display_url"`
														FaviconURL any    `json:"favicon_url"`
														ImageURL   string `json:"image_url"`
														Title      string `json:"title"`
														URL        string `json:"url"`
													} `json:"link_preview_attachment"`
												} `json:"text_post_app_info"`
												User struct {
													IsVerified    bool   `json:"is_verified"`
													Username      string `json:"username"`
													ID            any    `json:"id"`
													ProfilePicURL string `json:"profile_pic_url"`
												} `json:"user"`
												Pk               string `json:"pk"`
												MediaOverlayInfo any    `json:"media_overlay_info"`
												Code             string `json:"code"`
												Caption          struct {
													Text string `json:"text"`
												} `json:"caption"`
												ImageVersions2 struct {
													Candidates []any `json:"candidates"`
												} `json:"image_versions2"`
												OriginalWidth      int    `json:"original_width"`
												OriginalHeight     int    `json:"original_height"`
												VideoVersions      []any  `json:"video_versions"`
												CarouselMedia      any    `json:"carousel_media"`
												CarouselMediaCount any    `json:"carousel_media_count"`
												HasAudio           any    `json:"has_audio"`
												LikeCount          int    `json:"like_count"`
												TakenAt            int    `json:"taken_at"`
												ID                 string `json:"id"`
											} `json:"quoted_post"`
										} `json:"share_info"`
										ReplyToAuthor struct {
											Username string `json:"username"`
											ID       any    `json:"id"`
										} `json:"reply_to_author"`
										IsPostUnavailable bool `json:"is_post_unavailable"`
									} `json:"text_post_app_info"`
									Caption struct {
										Text string `json:"text"`
									} `json:"caption"`
									LikeCount int    `json:"like_count"`
									TakenAt   int    `json:"taken_at"`
									Code      string `json:"code"`
									ID        string `json:"id"`
								} `json:"reposted_post"`
							} `json:"share_info"`
							ReplyToAuthor     any  `json:"reply_to_author"`
							IsPostUnavailable bool `json:"is_post_unavailable"`
						} `json:"text_post_app_info"`
						Caption          any    `json:"caption"`
						TakenAt          int    `json:"taken_at"`
						LikeCount        int    `json:"like_count"`
						Code             string `json:"code"`
						MediaOverlayInfo any    `json:"media_overlay_info"`
						ID               string `json:"id"`
					} `json:"post"`
					LineType             string `json:"line_type"`
					ViewRepliesCtaString any    `json:"view_replies_cta_string"`
					ReplyFacepileUsers   []any  `json:"reply_facepile_users"`
					ShouldShowRepliesCta bool   `json:"should_show_replies_cta"`
					Typename             string `json:"__typename"`
				} `json:"thread_items"`
				ID string `json:"id"`
			} `json:"threads"`
		} `json:"mediaData"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
}
