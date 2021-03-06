package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/deissh/go-utils"
	"gorm.io/gorm"
)

type User struct {
	UserShort

	HasSupported bool `json:"has_supported"`

	Statistics        UserStatistics      `json:"statistics" gorm:"foreignkey:user_id;references:id"`
	MonthlyPlaycounts []MonthlyPlaycounts `json:"monthly_playcounts" gorm:"foreignkey:user_id;references:id"`
	UserAchievements  []UserAchievements  `json:"user_achievements" gorm:"foreignkey:user_id;references:id"`
	RankHistory       RankHistory         `json:"rank_history" gorm:"foreignkey:user_id;references:id"`
	// todo: this
	ReplaysWatchedCounts []interface{} `json:"replays_watched_counts" gorm:"-"`

	Discord                          utils.NullString `json:"discord"`
	Skype                            utils.NullString `json:"skype"`
	Title                            utils.NullString `json:"title"`
	TitleURL                         utils.NullString `json:"title_url"`
	Twitter                          utils.NullString `json:"twitter"`
	Website                          utils.NullString `json:"website"`
	Interests                        utils.NullString `json:"interests"`
	Kudosu                           Kudosu           `json:"kudosu"`
	Location                         utils.NullString `json:"location"`
	MaxBlocks                        int              `json:"max_blocks"`
	MaxFriends                       int              `json:"max_friends"`
	Occupation                       string           `json:"occupation"`
	Playmode                         string           `json:"playmode"`
	Playstyle                        []string         `json:"playstyle"`
	PostCount                        int              `json:"post_count"`
	ProfileOrder                     []string         `json:"profile_order"`
	AccountHistory                   []interface{}    `json:"account_history"`
	ActiveTournamentBanner           []interface{}    `json:"active_tournament_banner"`
	Badges                           []interface{}    `json:"badges"`
	BeatmapPlaycountsCount           int              `json:"beatmap_playcounts_count"`
	FavouriteBeatmapsetCount         int              `json:"favourite_beatmapset_count"`
	FollowerCount                    int              `json:"follower_count"`
	GraveyardBeatmapsetCount         int              `json:"graveyard_beatmapset_count"`
	LovedBeatmapsetCount             int              `json:"loved_beatmapset_count"`
	Page                             Page             `json:"page"`
	PreviousUsernames                []string         `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int              `json:"ranked_and_approved_beatmapset_count"`
	ScoresBestCount                  int              `json:"scores_best_count"`
	ScoresFirstCount                 int              `json:"scores_first_count"`
	ScoresRecentCount                int              `json:"scores_recent_count"`
	UnrankedBeatmapsetCount          int              `json:"unranked_beatmapset_count"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if err := u.UserShort.AfterFind(tx); err != nil {
		return err
	}
	return
}

// Country with code
type Country struct {
	Id   uint   `json:"-"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (Country) TableName() string {
	return "countries"
}

// Cover file url
type Cover struct {
	CustomURL string `json:"custom_url" gorm:"-"`
	URL       string `json:"url"`
	ID        string `json:"id" gorm:"-"`
}

// Kudosu value in user profile
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}

func (c Kudosu) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Kudosu) Scan(value interface{}) error {
	result := Kudosu{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// Page customization
type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

func (c Page) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Page) Scan(value interface{}) error {
	result := Page{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}
