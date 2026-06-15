package api

import (
	"github.com/TwiN/gatus/v5/config"
	"github.com/TwiN/gatus/v5/config/endpoint"
	"github.com/TwiN/gatus/v5/storage/store"
	"github.com/TwiN/gatus/v5/storage/store/common/paging"
	"github.com/TwiN/logr"
	"github.com/gofiber/fiber/v2"
)

type overallStatusResponse struct {
	OverallStatus string `json:"overallStatus"`
}

func OverallStatus(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pageSize := cfg.Storage.MaximumNumberOfResults
		endpointStatuses, err := store.Get().GetAllEndpointStatuses(paging.NewEndpointStatusParams().WithResults(1, pageSize))
		if err != nil {
			logr.Errorf("[api.OverallStatus] Failed to retrieve endpoint statuses: %s", err.Error())
			return c.Status(500).SendString(err.Error())
		}

		overall := computeOverallStatus(endpointStatuses)
		return c.Status(200).JSON(overallStatusResponse{OverallStatus: overall})
	}
}

func getEndpointStatus(ep *endpoint.Status) string {
	if ep == nil {
		return "unknown"
	}
	if len(ep.Results) == 0 {
		return "unknown"
	}
	if ep.Results[len(ep.Results)-1].Success {
		return "healthy"
	}
	return "unhealthy"
}

func computeOverallStatus(endpoints []*endpoint.Status) string {
	if len(endpoints) == 0 {
		return "unknown"
	}

	statuses := make([]string, 0, len(endpoints))
	for _, ep := range endpoints {
		s := getEndpointStatus(ep)
		if s != "unknown" {
			statuses = append(statuses, s)
		}
	}

	if len(statuses) == 0 {
		return "unknown"
	}

	allHealthy := true
	allUnhealthy := true
	allDegraded := true
	anyUnhealthy := false
	anyDegraded := false

	for _, s := range statuses {
		if s != "healthy" {
			allHealthy = false
		}
		if s != "unhealthy" {
			allUnhealthy = false
		}
		if s != "degraded" {
			allDegraded = false
		}
		if s == "unhealthy" {
			anyUnhealthy = true
		}
		if s == "degraded" {
			anyDegraded = true
		}
	}

	if allHealthy {
		return "healthy"
	}
	if allUnhealthy {
		return "unhealthy"
	}
	if allDegraded {
		return "degraded"
	}
	if anyUnhealthy {
		return "unhealthy_partial"
	}
	if anyDegraded {
		return "degraded_partial"
	}

	return "unknown"
}
