package main

import (
	"net/http"
	"time"
	"fmt"
	"sync"
)

func main() {
	URLs := []string{
		"www.wikipedia.org",
		"www.instagram.com",
		"www.facebook.com",
		"www.twitter.com",
		"www.asvfgdasda.gu",
	}

	urlMap := make(map[string]string)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, v := range URLs{
		wg.Add(1)
		go func(a string){
			defer wg.Done()
			mu.Lock()
			urlMap[a] = linkTest(a)
			mu.Unlock()
		}(v)	
	}

	wg.Wait()
	fmt.Println(urlMap)
}


func linkTest(link string) string {
	client := http.Client{
		Timeout : 3*time.Second,
	}

	_, err := client.Get("http://"+link)
	if err != nil {
		resultString := "bad"
		return resultString
	}
	resultString := "good"
	return resultString
}