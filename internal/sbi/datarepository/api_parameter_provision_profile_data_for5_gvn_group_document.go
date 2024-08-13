/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/internal/logger"
	"github.com/free5gc/udr/internal/sbi/producer"
	"github.com/free5gc/util/httpwrapper"
	"github.com/gin-gonic/gin"
)

// QueryPP5GVNGroupProfile - This resource represents the Parameter Provision Profile Data for the 5G VN Groups. It is queried by the UDM.
func QueryPP5GVNGroupProfile(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)

	rsp := producer.HandleQueryPP5GVNGroupProfile(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
