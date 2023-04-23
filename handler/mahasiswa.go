package handler

import (
	"fmt"
	"mahasiswa/helper"
	"mahasiswa/modules/mahasiswa"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mahasiswaHandler struct {
	service mahasiswa.Service
}

func NewMahasiswaHandler(service mahasiswa.Service) *mahasiswaHandler {
	return &mahasiswaHandler{service}
}

func (h *mahasiswaHandler) GetAllMahasiswa(c *gin.Context) {
	mahasiswas, err := h.service.GetAllMahasiswa()

	if err != nil {
		response := helper.APIResponse("Failed to get mahasiswa", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of mahasiswa", 200, "success", mahasiswa.FormatMahasiswas(mahasiswas))
	c.JSON(http.StatusOK, response)
}

func (h *mahasiswaHandler) GetDetailMahasiswa(c *gin.Context) {
	var input mahasiswa.GetMahasiswaDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mahasiswaDetail, err := h.service.GetMahasiswaByID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mahasiswaDetail.ID == 0 {
		response := helper.APIResponse("Failed to get detail mahasiswa", 404, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Detail of mahasiswa", 200, "success", mahasiswa.FormatMahasiswaDetail(mahasiswaDetail))

	c.JSON(http.StatusOK, response)
}

func (h *mahasiswaHandler) StoreMahasiswa(c *gin.Context) {
	var input mahasiswa.CreateMahasiswaInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create mahasiswa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMahasiswa, err := h.service.CreateMahasiswa(input)

	if err != nil {
		response := helper.APIResponse("Failed to create mahasiswa", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create mahasiswa", 201, "success", mahasiswa.FormatMahasiswaDetail(newMahasiswa))
	c.JSON(http.StatusCreated, response)
}

func (h *mahasiswaHandler) UpdateMahasiswa(c *gin.Context) {
	var inputID mahasiswa.GetMahasiswaDetailInput
	var inputData mahasiswa.CreateMahasiswaInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(&inputData)
	fmt.Println(err)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mahasiswa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check mahasiswa exist
	mahasiswaDetail, err := h.service.GetMahasiswaByID(inputID)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mahasiswaDetail.ID == 0 {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	updatedMahasiswa, err := h.service.UpdateMahasiswa(inputID, inputData)

	if err != nil {
		response := helper.APIResponse("Failed to update mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update mahasiswa", 200, "success", mahasiswa.FormatMahasiswaDetail(updatedMahasiswa))

	c.JSON(http.StatusOK, response)
}

func (h *mahasiswaHandler) DeleteMahasiswa(c *gin.Context) {
	var input mahasiswa.GetMahasiswaDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check mahasiswa exist
	mahasiswaDetail, err := h.service.GetMahasiswaByID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mahasiswaDetail.ID == 0 {
		response := helper.APIResponse("Failed to get detail mahasiswa", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// delete mahasiswa
	deletedMahasiswa, err := h.service.DeleteMahasiswa(input)

	if err != nil {
		response := helper.APIResponse("Failed to delete mahasiswa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete mahasiswa", 200, "success", mahasiswa.FormatMahasiswa(deletedMahasiswa))

	c.JSON(http.StatusOK, response)
}
