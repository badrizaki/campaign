package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	var err error
	if userId == 0 {
		campaigns, err = s.repository.FetchAll()
	} else {
		campaigns, err = s.repository.FindByUserID(userId)
	}
	return campaigns, err
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	return s.repository.FindByID(input.ID)
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %s", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	return newCampaign, err
}
