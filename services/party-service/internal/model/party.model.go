package model

import (
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"time"
)

// Party model
type Party struct {
	PartyId      string `gorm:"primarykey;column:PartyId"`
	Slug         string
	UserID       string
	Name         string
	LocationName string
	Description  string
	Latitude     float64
	Longitude    float64
	City         string
	BannerUrl    string
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToPartyDomain(p *party.Party) *Party {
	return &Party{
		PartyId:      p.Id,
		Slug:         p.Slug,
		UserID:       p.HostId,
		Name:         p.Name,
		LocationName: p.LocationName,
		Description:  p.Description,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		City:         p.City,
		BannerUrl:    p.BannerUrl,
		StartDate:    time.Unix(p.StartDate, 0),
		EndDate:      time.Unix(p.EndDate, 0),
	}
}

func (p Party) ToPartyPb() *party.Party {
	return &party.Party{
		Id:           p.PartyId,
		Name:         p.Name,
		Description:  p.Description,
		Address:      "",
		Privacy:      0,
		BannerUrl:    p.BannerUrl,
		StartDate:    p.StartDate.Unix(),
		EndDate:      p.EndDate.Unix(),
		HostId:       p.UserID,
		LocationName: p.LocationName,
		CreateTime:   p.CreatedAt.Unix(),
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		MinimumPrice: 0,
		Slug:         p.Slug,
		City:         p.City,
	}
}
