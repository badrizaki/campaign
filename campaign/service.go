package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
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
