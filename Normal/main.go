package main

import (
	"os/exec"
	"time"
	"fmt"
)

var urls = [] string {
	"https://www.youtube.com/watch?v=MvsAesQ-4zA",
	"https://www.youtube.com/watch?v=TK4N5W22Gts",
	"https://www.youtube.com/watch?v=5DEdR5lqnDE",
}

func main() {
	startTime := time.Now()
	for _, str := range urls {
		downloadVideo(str)
	}
	duration := time.Since(startTime)
	fmt.Println(duration)
}

func downloadVideo(new string) {
	cmd := exec.Command("yt-dlp", new)
	cmd.Run()
}