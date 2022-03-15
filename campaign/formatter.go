package campaign

import "strings"

type CampaignFormatter struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	ShorDescription string `json:"short_description"`
	ImageUrl        string `json:"image_url"`
	GoalAmount      int    `json:"goal_amount"`
	CurrentAmount   int    `json:"current_amount"`
	Slug            string `json:"slug"`
}

type CampaginDetailFormatter struct {
	ID              int                      `json:"id"`
	UserID          int                      `json:"user_id"`
	Name            string                   `json:"name"`
	ShorDescription string                   `json:"short_description"`
	ImageUrl        string                   `json:"image_url"`
	GoalAmount      int                      `json:"goal_amount"`
	CurrentAmount   int                      `json:"current_amount"`
	Slug            string                   `json:"slug"`
	Description     string                   `json:"description"`
	Perks           []string                 `json:"perks"`
	User            CampaginUserFormatter    `json:"user"`
	Images          []CampaignImageFormatter `json:"images"`
}

type CampaginUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	ImageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		ImageUrl = campaign.CampaignImages[0].FileName
	}
	return CampaignFormatter{
		ID:              campaign.ID,
		UserID:          campaign.UserID,
		Name:            campaign.Name,
		ShorDescription: campaign.ShorDescription,
		ImageUrl:        ImageUrl,
		GoalAmount:      campaign.GoalAmount,
		CurrentAmount:   campaign.CurrentAmount,
		Slug:            campaign.Slug,
	}
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	if len(campaigns) <= 0 {
		return []CampaignFormatter{}
	}
	var campaginsFormatter []CampaignFormatter
	for _, campaign := range campaigns {
		campaginsFormatter = append(campaginsFormatter, FormatCampaign(campaign))
	}
	return campaginsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaginDetailFormatter {
	// ini sama dengan yang atas
	ImageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		ImageUrl = campaign.CampaignImages[0].FileName
	}

	campaginDetailFormatter := CampaginDetailFormatter{}
	campaginDetailFormatter.ID = campaign.ID
	campaginDetailFormatter.UserID = campaign.UserID
	campaginDetailFormatter.Name = campaign.Name
	campaginDetailFormatter.ShorDescription = campaign.ShorDescription
	campaginDetailFormatter.ImageUrl = ImageUrl
	campaginDetailFormatter.GoalAmount = campaign.GoalAmount
	campaginDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaginDetailFormatter.Slug = campaign.Slug

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perk = strings.TrimSpace(perk)
		perks = append(perks, perk)
	}
	campaginDetailFormatter.Perks = perks

	user := campaign.User
	campaginUserFormatter := CampaginUserFormatter{}
	campaginUserFormatter.Name = user.Name
	campaginUserFormatter.ImageUrl = user.AvatarFileName
	campaginDetailFormatter.User = campaginUserFormatter

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName
		isPrimary := false
		if image.IsPrimary > 0 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}
	campaginDetailFormatter.Images = images

	return campaginDetailFormatter
}
