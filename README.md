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
	threadsapi "github.com/erkanzsy/threads-api"
)

func main() {
	userId, err := threadsapi.UserIDFromUsername("zuck")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId) // 314216

	profile, err := threadsapi.UserProfile("zuck", "314216")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile.Data.UserData.User.Username) // zuck

	threads, err := threadsapi.UserThreads("zuck", "314216")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(threads.Data.MediaData.Threads[0].ThreadItems[0].Post.User.Username) // zuck
}
```

## Unit Tests
<a href="https://github.com/erkanzsy/threads-api/actions">
<img src="https://img.shields.io/github/actions/workflow/status/erkanzsy/threads-api/go.yml?branch=main&label=%F0%9F%A7%AA%20&style=flat&color=75C46B">
  </a>


----

> ☢️ ** This package cannot be held responsible for any issues that may arise. The risk belongs to the user. If you perceive any risks, please refrain from using it.