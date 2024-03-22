package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"time"
	"twitter-scraper/pkg/twitter_scraper"
)

func main() {

	scraper := twitter_scraper.New()

	scraper.WithDelay(15).WithClientTimeout(time.Second * 30).SetSearchMode(twitter_scraper.SearchLatest)
	//scraper.SetProxy("http://127.0.0.1:1088")
	//scraper.SetProxy("socks5://127.0.0.1:1085")

	var err error
	if scraper.IsLoginState() == true {
		err = scraper.Load()
	} else {
		err = scraper.Login("szd5654125@gmail.com", "qkl5641388", "qkl55916973")
	}

	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("./data.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)

	startTime, _ := time.Parse("2006-01-02", "2023-10-01")
	endTime, _ := time.Parse("2006-01-02", "2023-10-31")
	for {
		fmt.Println(startTime.AddDate(0, 0, 1).Format("2006-01-02"), startTime.Format("2006-01-02"))
		defer writer.Flush()
		for tweet := range scraper.SearchTweets(context.Background(), fmt.Sprintf("bonk until:%s since:%s", startTime.AddDate(0, 0, 1).Format("2006-01-02"), startTime.Format("2006-01-02")), 50) {
			if tweet.Error != nil {
				panic(tweet.Error)
			}
			//fmt.Println(tweet.ID)
			//writer.Write([]string{tweet.Text, strconv.Itoa(tweet.Likes), tweet.TimeParsed.Format("2006-01-02 15:04:05")})
			writer.Write([]string{tweet.ID, strconv.Itoa(tweet.Likes), tweet.Name, tweet.TimeParsed.Format("2006-01-02 15")})
			writer.Flush()
			//os.Exit(123)
		}

		if startTime.Equal(endTime) {
			break
		}
		startTime = startTime.AddDate(0, 0, 1)
		time.Sleep(3 * time.Second)
	}

}
