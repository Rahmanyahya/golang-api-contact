package services

import (
	"golang-api-contact/models"
	"golang-api-contact/repositories"
	"golang-api-contact/request"

	"github.com/go-playground/validator/v10"
)

type ContactService interface {

	CreateContact(req *request.ContactRequest) (*models.Contact, error)

	GetAllContact() ([]models.Contact, error)

	GetContactById(id uint) (*models.Contact, error)

	UpdateContact(id uint, req *request.ContactRequest) (*models.Contact, error)

	DeleteContact(id uint) error

}

type contactService struct {
	repository repositories.ContactRepository
	validate *validator.Validate
}

func NewContactService(repository repositories.ContactRepository) ContactService {
	return &contactService{
		repository: repository,
		validate: validator.New(),
	}
}

func (s *contactService) CreateContact(req *request.ContactRequest) (*models.Contact, error) {

	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	contact := models.Contact{
		FullName: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Message: req.Message,
	}

	err := s.repository.Create(&contact)
	return &contact, err

}

func (s *contactService) GetAllContact() ([]models.Contact, error) {
	return s.repository.FindAll()
}

func (s *contactService) GetContactById(id uint) (*models.Contact, error) {
	return s.repository.FindById(id) 
}

func (s *contactService) UpdateContact(id uint, req *request.ContactRequest) (*models.Contact, error) {

	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	contact, err := s.repository.FindById(id);

	if err != nil {
		return nil, err
	}

	contact.FullName = req.Name
	contact.Email = req.Email
	contact.Message = req.Message
	contact.Phone = req.Phone

	err = s.repository.Update(contact)

	return contact, err
}

func (s *contactService) DeleteContact(id uint) error {
	
	contact, err := s.repository.FindById(id);

	if err != nil {
		return err
	}

	return s.repository.Delete(contact)
}
