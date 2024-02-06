package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelServerInfo struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyID       primitive.ObjectID `json:"company_id,omitempty" bson:"company_id,omitempty"`
	Host            string             `json:"host" bson:"host"`
	Port            int                `json:"port" bson:"port"`
	Protocol        string             `json:"protocol" bson:"protocol"`
	IsPublic        bool               `json:"isPublic" bson:"isPublic"`
	Status          string             `json:"status" bson:"status"`
	StartTime       time.Time          `json:"startTime" bson:"startTime"`
	TestTime        time.Time          `json:"testTime" bson:"testTime"`
	EngineVersion   string             `json:"engineVersion" bson:"engineVersion"`
	CriteriaVersion string             `json:"criteriaVersion" bson:"criteriaVersion"`
	Endpoints       []Endpoint         `json:"endpoints" bson:"endpoints"`
}

type Endpoint struct {
	IPAddress         string `json:"ipAddress" bson:"ipAddress"`
	ServerName        string `json:"serverName" bson:"serverName"`
	StatusMessage     string `json:"statusMessage" bson:"statusMessage"`
	Grade             string `json:"grade" bson:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored" bson:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings" bson:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional" bson:"isExceptional"`
	Progress          int    `json:"progress" bson:"progress"`
	Duration          int    `json:"duration" bson:"duration"`
	Delegation        int    `json:"delegation" bson:"delegation"`
}

func MapSSLLabsResultToServerInfo(companyID primitive.ObjectID, result map[string]interface{}) *HotelServerInfo {
	serverInfo := &HotelServerInfo{
		CompanyID:       companyID,
		Host:            result["host"].(string),
		Port:            int(result["port"].(float64)),
		Protocol:        result["protocol"].(string),
		IsPublic:        result["isPublic"].(bool),
		Status:          result["status"].(string),
		StartTime:       time.Unix(0, int64(result["startTime"].(float64))*int64(time.Millisecond)),
		TestTime:        time.Unix(0, int64(result["testTime"].(float64))*int64(time.Millisecond)),
		EngineVersion:   result["engineVersion"].(string),
		CriteriaVersion: result["criteriaVersion"].(string),
		Endpoints:       make([]Endpoint, 0),
	}

	endpoints := result["endpoints"].([]interface{})
	for _, ep := range endpoints {
		endpoint := ep.(map[string]interface{})
		serverInfo.Endpoints = append(serverInfo.Endpoints, Endpoint{
			IPAddress:         endpoint["ipAddress"].(string),
			ServerName:        endpoint["serverName"].(string),
			StatusMessage:     endpoint["statusMessage"].(string),
			Grade:             endpoint["grade"].(string),
			GradeTrustIgnored: endpoint["gradeTrustIgnored"].(string),
			HasWarnings:       endpoint["hasWarnings"].(bool),
			IsExceptional:     endpoint["isExceptional"].(bool),
			Progress:          int(endpoint["progress"].(float64)),
			Duration:          int(endpoint["duration"].(float64)),
			Delegation:        int(endpoint["delegation"].(float64)),
		})
	}

	return serverInfo
}
