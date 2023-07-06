package threads_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	URL     = "https://www.threads.net"
	GRAPHQL = "https://www.threads.net/api/graphql"
)

func UserIDFromUsername(username string) (string, error) {
	url := fmt.Sprintf("%s/@%s", URL, username)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("İstek oluşturulurken bir hata oluştu: %v", err)
	}

	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	request.Header.Set("accept-language", "ko,en;q=0.9,ko-KR;q=0.8,ja;q=0.7")
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("referer", "https://www.instagram.com/")
	request.Header.Set("sec-fetch-dest", "document")
	request.Header.Set("sec-fetch-mode", "navigate")
	request.Header.Set("sec-fetch-site", "cross-site")
	request.Header.Set("sec-fetch-user", "?1")
	request.Header.Set("upgrade-insecure-requests", "1")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("İstek gönderilirken bir hata oluştu: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Yanıt okunurken bir hata oluştu: %v", err)
	}

	text := string(body)

	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\n", "")

	regex := regexp.MustCompile(`"props":{"user_id":"(\d+)"},`)
	match := regex.FindStringSubmatch(text)
	if len(match) >= 2 {
		userId := match[1]
		return userId, nil
	}

	return "", nil
}

func UserProfile(username, userId string) (UserProfileData, error) {
	formData := url.Values{
		"av":                       {"0"},
		"__user":                   {"0"},
		"__a":                      {"1"},
		"__req":                    {"1"},
		"__hs":                     {"19544.HYP:barcelona_web_pkg.2.1..0.0"},
		"dpr":                      {"1"},
		"__ccg":                    {"EXCELLENT"},
		"__rev":                    {"1007795914"},
		"__s":                      {"c1fpxh:oh98tm:os2fqi"},
		"__hsi":                    {"7252655495199472548"},
		"__dyn":                    {"7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q32360CEbo1nEhw2nVE4W0om78b87C0yE465o-cw5Mx62G3i0Bo7O2l0Fwqo31wnEfovwRwlE-U2zxe2Gew9O22362W2K0zK5o4q0GpovU1aUbodEGdwtU2ewbS1LwTwNwLw8O1pwr82gxC"},
		"__csr":                    {"j8kjt5p9e00hB4Eqw-w0Xiwrk0xE9Eixza2svazUndhEpko9xy7Ej7Saxl2U5-8m8yA4zCwxxWegQz5162a5x02UxW1g2Ex3MwM_3M25wlQ13gN0el4m2H3r16089wxwnq0w8gqd12"},
		"__comet_req":              {"29"},
		"lsd":                      {"NjppQDEgONsU_1LCzrmp6q"},
		"jazoest":                  {"21997"},
		"__spin_r":                 {"1007795914"},
		"__spin_b":                 {"trunk"},
		"__spin_t":                 {"1688640447"},
		"__jssesw":                 {"2"},
		"fb_api_caller_class":      {"RelayModern"},
		"fb_api_req_friendly_name": {"BarcelonaProfileRootQuery"},
		"variables":                {`{"userID":"` + userId + `"}`},
		"server_timestamps":        {"true"},
		"doc_id":                   {"23996318473300828"},
	}
	userProfile := UserProfileData{}

	request, err := http.NewRequest("POST", GRAPHQL, strings.NewReader(formData.Encode()))
	if err != nil {
		return userProfile, err
	}

	request.Header = getDefaultHeaders(username)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("x-fb-friendly-name", "BarcelonaProfileRootQuery")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return userProfile, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userProfile)
	if err != nil {
		return userProfile, err
	}

	return userProfile, nil

}

func UserThreads(username, userId string) (UserThreadsData, error) {
	data := url.Values{}
	data.Set("av", "0")
	data.Set("__user", "0")
	data.Set("__a", "1")
	data.Set("__req", "2")
	data.Set("__hs", "19544.HYP:barcelona_web_pkg.2.1..0.0")
	data.Set("dpr", "1")
	data.Set("__ccg", "EXCELLENT")
	data.Set("__rev", "1007795914")
	data.Set("__s", "c1fpxh:oh98tm:os2fqi")
	data.Set("__hsi", "7252655495199472548")
	data.Set("__dyn", "7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q32360CEbo1nEhw2nVE4W0om78b87C0yE465o-cw5Mx62G3i0Bo7O2l0Fwqo31wnEfovwRwlE-U2zxe2Gew9O22362W2K0zK5o4q0GpovU1aUbodEGdwtU2ewbS1LwTwNwLw8O1pwr82gxC")
	data.Set("__csr", "j8kjt5p9e00hB4Eqw-w0Xiwrk0xE9Eixza2svazUndhEpko9xy7Ej7Saxl2U5-8m8yA4zCwxxWegQz5162a5x02UxW1g2Ex3MwM_3M25wlQ13gN0el4m2H3r16089wxwnq0w8gqd12")
	data.Set("__comet_req", "29")
	data.Set("lsd", "NjppQDEgONsU_1LCzrmp6q")
	data.Set("jazoest", "21997")
	data.Set("__spin_r", "1007795914")
	data.Set("__spin_b", "trunk")
	data.Set("__spin_t", "1688640447")
	data.Set("__jssesw", "2")
	data.Set("fb_api_caller_class", "RelayModern")
	data.Set("fb_api_req_friendly_name", "BarcelonaProfileThreadsTabQuery")
	variables := fmt.Sprintf(`{"userID":"%s"}`, userId)
	data.Set("variables", variables)
	data.Set("server_timestamps", "true")
	data.Set("doc_id", "6232751443445612")

	userThreads := UserThreadsData{}

	request, err := http.NewRequest("POST", "https://www.threads.net/api/graphql", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return userThreads, err
	}

	request.Header = getDefaultHeaders(username)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("x-fb-friendly-name", "BarcelonaProfileThreadsTabQuery")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return userThreads, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&userThreads)
	if err != nil {
		return userThreads, err
	}

	return userThreads, nil
}

func getDefaultHeaders(username string) http.Header {
	headers := http.Header{}
	headers.Set("authority", "www.threads.net")
	headers.Set("accept", "*/*")
	headers.Set("accept-language", "ko")
	headers.Set("cache-control", "no-cache")
	headers.Set("origin", "https://www.threads.net")
	headers.Set("pragma", "no-cache")
	headers.Set("referer", "https://www.threads.net/@"+username)
	headers.Set("sec-ch-prefers-color-scheme", "dark")
	headers.Set("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	headers.Set("sec-ch-ua-full-version-list", "\"Not.A/Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"114.0.5735.198\", \"Google Chrome\";v=\"114.0.5735.198\"")
	headers.Set("sec-ch-ua-mobile", "?0")
	headers.Set("sec-ch-ua-platform", "\"macOS\"")
	headers.Set("sec-ch-ua-platform-version", "\"13.0.0\"")
	headers.Set("sec-fetch-dest", "empty")
	headers.Set("sec-fetch-mode", "cors")
	headers.Set("sec-fetch-site", "same-origin")
	headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	headers.Set("viewport-width", "1728")
	headers.Set("x-asbd-id", "129477")
	headers.Set("x-fb-lsd", "NjppQDEgONsU_1LCzrmp6q")
	headers.Set("x-ig-app-id", "238260118697367")

	return headers
}
