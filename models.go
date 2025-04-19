package main

import (
	"github.com/google/uuid"
	"github.com/olavowilke/rss-api/internal/database"
	"time"
)

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	ApiKey    string    `json:"apiKey"`
}

func mapDatabaseUserToCreateUserResponse(dbUser database.User) CreateUserResponse {
	return CreateUserResponse{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type CreateFeedResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func mapDatabaseFeedToCreateFeedResponse(dbFeed database.Feed) CreateFeedResponse {
	return CreateFeedResponse{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
	}
}

type GetFeedResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func mapDatabaseFeedsToGetFeedsResponse(dbFeeds []database.Feed) []GetFeedResponse {
	feedsResponses := []GetFeedResponse{}
	for _, dbFeed := range dbFeeds {
		feedsResponses = append(feedsResponses, GetFeedResponse{
			ID:        dbFeed.ID,
			Name:      dbFeed.Name,
			Url:       dbFeed.Url,
			CreatedAt: dbFeed.CreatedAt,
			UpdatedAt: dbFeed.UpdatedAt,
		})
	}
	return feedsResponses
}

type createFeedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

func mapDatabaseFeedFollowToCreateFeedFollowResponse(feed database.FeedFollow) createFeedFollowResponse {
	return createFeedFollowResponse{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
	}
}

type GetFeedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	FeedId    uuid.UUID `json:"feedId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func mapFeedFollowsToFeedFollowsResponse(follows []database.FeedFollow) []GetFeedFollowResponse {
	feedFollowResponses := []GetFeedFollowResponse{}
	for _, dbFeedFollow := range follows {
		feedFollowResponses = append(feedFollowResponses, GetFeedFollowResponse{
			ID:        dbFeedFollow.ID,
			UserId:    dbFeedFollow.UserID,
			FeedId:    dbFeedFollow.FeedID,
			CreatedAt: dbFeedFollow.CreatedAt,
			UpdatedAt: dbFeedFollow.UpdatedAt,
		})
	}
	return feedFollowResponses
}

type GetPostsResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feedID"`
}

func mapDatabasePostsToGetPostsResponse(posts []database.GetPostsForUserRow) []GetPostsResponse {
	postsResponses := []GetPostsResponse{}
	for _, dbPost := range posts {

		postsResponses = append(postsResponses, GetPostsResponse{
			ID:          dbPost.ID,
			CreatedAt:   dbPost.CreatedAt,
			UpdatedAt:   dbPost.UpdatedAt,
			Title:       dbPost.Title,
			Description: &dbPost.Description.String,
			PublishedAt: dbPost.PublishedAt,
			Url:         dbPost.Url,
			FeedID:      dbPost.FeedID,
		})
	}
	return postsResponses
}
