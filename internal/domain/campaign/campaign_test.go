package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"teste@gmail.com", "email@gmail.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	// Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

	// without lib
	// if campaign.ID != "1" {
	// 	t.Error("expected 1")
	// } else if campaign.Name != name {
	// 	t.Error("expected correct name", campaign.Name)
	// } else if campaign.Content != content {
	// 	t.Error("expected correct content", campaign.Content)
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Error("expected correct qtd contacts", len(campaign.Contacts))
	// }
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
