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
