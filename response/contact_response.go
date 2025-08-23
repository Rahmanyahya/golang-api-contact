package response

import (
	helper "golang-api-contact/helpers"
	"golang-api-contact/models"
)

type APIResponse struct {
	Code string `json:"code"`

	Message string `json:"message"`

	Data interface{} `json:"data"`
}

type ContactResponse struct {
	ID uint `json:"id"`

	Name string `json:"name"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Message string `json:"message"`

	CreatedAt string `json:"createdAt"`

	UpdatedAt string `json:"updatedAt"`
}

func ContactResponseFromModel(contact *models.Contact) ContactResponse {
	return ContactResponse{
		ID: contact.ID,
		Name: contact.FullName,
		Email: contact.Email,
		Phone: contact.Phone,
		Message: contact.Message,
		CreatedAt: helper.FormatTimeHuman(contact.CreatedAt),
		UpdatedAt: helper.FormatTimeHuman(contact.UpdatedAt),
	}
}