package main

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/olavowilke/rss-api/internal/database"
	"log"
	"strings"
	"sync"
	"time"
)

func startScraping(
	db *database.Queries,
	concurrentFeeds int,
	requestInterval time.Duration) {
	log.Printf("Scraping RSS Feeds with %d concurrent feeds every %v\n", concurrentFeeds, requestInterval)

	ticker := time.NewTicker(requestInterval)
	for ; ; <-ticker.C {
		// Get all feeds from the database
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrentFeeds))
		if err != nil {
			log.Println("Error getting feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.SetFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error setting feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Error parsing published date:", err)
			publishedAt = time.Now()
		}

		_, dbErr := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: description,
			PublishedAt: publishedAt,
			Url:         item.Link,
			FeedID:      feed.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if dbErr != nil {
			if strings.Contains(dbErr.Error(), "duplicate key") {
				continue
			}
			log.Println("Error creating post:", dbErr)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
