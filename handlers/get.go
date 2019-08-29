package handlers

import (
	"net/http"

	"github.com/byuoitav/common/status"
	"github.com/byuoitav/epson-projector-ms/helpers"
	"github.com/labstack/echo"
)

// GetPower .
func GetPower(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetPower(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetBlanked .
func GetBlanked(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetBlanked(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetInput .
func GetInput(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetInput(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetVolume .
func GetVolume(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetVolume(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetMuted .
func GetMuted(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetMuted(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

/*
// GetHardwareInfo .
func GetHardwareInfo(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetHardwareInfo(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}

// GetActiveSignal .
func GetActiveSignal(ectx echo.Context) error {
	address := ectx.Param("address")

	resp, err := helpers.GetActiveSignal(address)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	return ectx.JSON(http.StatusOK, resp)
}
*/
