# threads-api

## Installation
```bash
go get github.com/erkanzsy/threads-api
``` 

## Usage
```go
package main

import (
	"fmt"
	threads_api "github.com/erkanzsy/threads-api"
)

func main() {
	userId, _ := threads_api.UserIDFromUsername("zuck")
	fmt.Println(userId) // 314216

	profile, _ := threads_api.UserProfile("zuck", "314216")
	fmt.Println(profile.Data.UserData.User.Username) // zuck

	threads, _ := threads_api.UserThreads("zuck", "314216")
	fmt.Println(threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username) // zuck
}
```

## Unit Tests
<a href="https://github.com/erkanzsy/threads-api/actions">
<img src="https://img.shields.io/github/actions/workflow/status/erkanzsy/threads-api/go.yml?branch=main&label=%F0%9F%A7%AA%20&style=flat&color=75C46B">
  </a>