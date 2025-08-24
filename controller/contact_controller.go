package controller

import (
	"golang-api-contact/request"
	"golang-api-contact/response"
	"golang-api-contact/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	service services.ContactService
}

func NewContactController(service services.ContactService) *ContactController {
	return &ContactController{service}
}

func (h *ContactController) CreateContact(c *gin.Context) {
	var req request.ContactRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code: "BAD_REQUEST",
			Message: err.Error(),
			Data: nil,
		})
		return;
	}

	contact, err := h.service.CreateContact(&req);
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code: "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data: nil,
		})
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Code: "CREATED",
		Message: "Contact Already Created",
		Data: response.ContactResponseFromModel(contact),
	})
}

func (h *ContactController) GetContacts(c *gin.Context) {
	contacts, err := h.service.GetAllContact();

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code: "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	var contactResponses []response.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, response.ContactResponseFromModel(&contact))
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code: "SUCCESS",
		Message: "Contacts retrived successfully",
		Data: contactResponses,
	})
}

func (h *ContactController) GetContact(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams);

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code: "BAS_REQUEST",
			Message: "Invalid ID",
			Data: nil,
		})
		return
	}

	contact, err := h.service.GetContactById(uint(id));
	if err != nil {
		c.JSON(http.StatusNotFound, response.APIResponse{
			Code: "NOT_FOUND",
			Message: "Contact not found",
			Data: nil,	
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code: "OK",
		Message: "Contact retrived successfully",
		Data: response.ContactResponseFromModel(contact),
	})
}

func (h *ContactController) DeleteContact(c *gin.Context) {
	idParams := c.Param("id");
	id, err := strconv.Atoi(idParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code: "BAD_REQUEST",
			Message: "ID invalid",
			Data: nil,
		})
		return
	}

	contact, err := h.service.GetContactById(uint(id));
	if err != nil {
		c.JSON(http.StatusNotFound, response.APIResponse{
			Code: "NOT_FOUND",
			Message: "Contact Not Found",
			Data: nil,
		})
		return
	}

    err = h.service.DeleteContact(contact.ID);
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code: "INTERNAL_SERVER_ERROR",
			Message: "Something Went Wrong",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code: "OK",
		Message: "Success delete contact",
		Data: nil,
	})

}

func (h *ContactController) UpdateContact(c *gin.Context) {
	idParams := c.Param("id");
	id, err := strconv.Atoi(idParams);

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code: "BAS_REQUEST",
			Message: "Invalid ID",
			Data: nil,
		})
		return
	}

	var req request.ContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code: "BAD_REQUEST",
			Message: err.Error(),
			Data: nil,
		})
		return;
	}

	contact, err := h.service.GetContactById(uint(id));
	if err != nil {
		c.JSON(http.StatusNotFound, response.APIResponse{
			Code: "NOT_FOUND",
			Message: "Contact Not Found",
			Data: nil,
		})
		return;
	}


	result ,err := h.service.UpdateContact(contact.ID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code: "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data: nil,
		})
		return;
	}

	c.JSON(http.StatusAccepted, response.APIResponse{
		Code: "OK",
		Message: "Edited contact successfully",
		Data: response.ContactResponseFromModel(result),
	})

	
}