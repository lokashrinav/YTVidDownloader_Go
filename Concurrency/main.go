package main

import (
	"os/exec"
	"sync"
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
	var wg sync.WaitGroup
	for _, str := range urls {
		wg.Add(1)
		go downloadVideo(str, &wg)
	}
	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println(duration)
}

func downloadVideo(new string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("yt-dlp", new)
	cmd.Run()
}