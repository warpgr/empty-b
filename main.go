package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var configCounter int = 1

// var configDB
var configsDB map[string]interface{}
var statisticsDB map[string]*map[ExchangeName]map[string]PairStatistic
var processingDB map[string]*map[ExchangeName]map[string]PairProcessingStatus

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getProcessing(c *gin.Context) {
	userName := c.GetHeader("user")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty user name.",
		})
		return
	}

	var userData *map[ExchangeName]map[string]PairProcessingStatus
	if userData, ok := processingDB[userName]; !ok || userData == nil {
		msg := fmt.Sprintf("User %s not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": msg,
		})
		return
	}

	exchangeName := c.GetHeader("exchange")
	exchangeNameExists := true
	if exchangeName == "" {
		exchangeNameExists = false
	}
	pairName := c.GetHeader("pair")
	pairNameExists := true
	if pairName == "" {
		pairNameExists = false
	}

	if !pairNameExists && !exchangeNameExists {
		c.JSON(http.StatusOK, *userData)
		return
	} else if !exchangeNameExists && pairNameExists {
		var pairStatuses []PairProcessingStatus
		for _, exchangeStatuses := range *userData {
			if pairStatus, ok := exchangeStatuses[pairName]; ok {
				pairStatuses = append(pairStatuses, pairStatus)
			}
		}
		c.JSON(http.StatusOK, pairStatuses)
		return
	} else if exchangeNameExists && !pairNameExists {
		exchange := ExchangeName(exchangeName)
		if exchangeStatuses, ok := (*userData)[exchange]; ok {
			c.JSON(http.StatusOK, exchangeStatuses)
			return
		}
		msg := fmt.Sprintf("Statuses for %s exchange not found.", exchange)
		c.JSON(http.StatusNotFound, gin.H{
			"message": msg,
		})
		fmt.Print(msg)
		return
	} else {
		pairStatus, ok := (*userData)[ExchangeName(exchangeName)][pairName]
		if !ok {
			msg := fmt.Sprintf("Data with %s pair on %s exchange not found.", exchangeName, pairName)
			c.JSON(http.StatusNotFound, gin.H{
				"message": msg,
			})
			fmt.Print(msg)
			return
		}
		c.JSON(http.StatusOK, pairStatus)
	}
}

func getStatistics(c *gin.Context) {
	userName := c.Request.Header.Get("user")
	if userName == "" {
		fmt.Print("Empty user filed")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty user name.",
		})
		return
	}

	userStatistics, ok := statisticsDB[userName]
	if !ok || userStatistics == nil {
		msg := fmt.Sprintf("User %s statistic not found.", userName)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		fmt.Print(msg)
		return
	}

	var exchangeName ExchangeName
	exchangeParamExists := true
	var exNameStr string
	if exNameStr = c.GetHeader(ExchangeKWD); exNameStr != "" {
		exchangeName = ExchangeName(exNameStr)
	} else {
		exchangeParamExists = false
	}

	pairParamExists := true
	pair := c.GetHeader(PairKWD)
	if pair == "" {
		pairParamExists = false
	}

	if exchangeParamExists {
		exchangeOrders, ok := (*userStatistics)[exchangeName]
		if !ok {
			msg := fmt.Sprintf("Statistics on %s exchange not found.", exchangeName)
			c.JSON(http.StatusNotFound, gin.H{
				"message": msg,
			})
			return
		}
		if pairParamExists {
			if pairOrders, ok := exchangeOrders[pair]; ok {
				c.JSON(http.StatusOK, pairOrders)
				return
			} else {
				msg := fmt.Sprintf("Orders with %s pair on %s exchange not found.", pair, exchangeName)
				c.JSON(http.StatusNotFound, gin.H{
					"message": msg,
				})
				return
			}
		} else {
			c.JSON(http.StatusOK, exchangeOrders)
			return
		}
	} else if pairParamExists {
		var pairOrders []PairStatistic
		for _, exchangeStatistic := range *userStatistics {
			if pairOrder, ok := exchangeStatistic[pair]; ok {
				pairOrders = append(pairOrders, pairOrder)
			}
		}
		c.JSON(http.StatusOK, pairOrders)
		return
	}

	// Otherwise
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Exchange and pair parameters not sent.",
	})
	return
}

func setConfigurations(c *gin.Context) {
	var config interface{}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty body.",
		})
		return
	}
	configsDB[fmt.Sprintf("config_%d", configCounter)] = config
	configCounter++
	c.JSON(http.StatusOK, gin.H{
		"message": "Configurations was set.",
	})
	return
}

func main() {
	configFilePath := "emulated_db/configs.json"
	processingFilePath := "emulated_db/in_process.json"
	statisticsFilePath := "emulated_db/statistics.json"

	serializedConfigs, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Can't load configurations from '%s' file. Error: %s.", configFilePath, err)
		os.Exit(1)
	}
	if err := json.Unmarshal(serializedConfigs, &configsDB); err != nil {
		fmt.Printf("Can't unmarshal config file. Error %s.", err)
		os.Exit(1)
	}

	serializedProcessing, err := os.ReadFile(processingFilePath)
	if err != nil {
		fmt.Printf("Can't load processing db from %s file. Error: %s.", processingFilePath, err)
		os.Exit(1)
	}
	if err = json.Unmarshal(serializedProcessing, &processingDB); err != nil {
		fmt.Printf("Can't unmarshal processing db. Error: %s.", err)
		os.Exit(1)
	}

	serializedStatisticsFile, err := os.ReadFile(statisticsFilePath)
	if err != nil {
		fmt.Printf("Can't load statistics db from '%s' file. Error: %s.", statisticsFilePath, err)
		os.Exit(1)
	}
	if err = json.Unmarshal(serializedStatisticsFile, &statisticsDB); err != nil {
		fmt.Printf("Can't unmarshal statistics db. Error %s.", err)
		os.Exit(1)
	}

	router := gin.Default()
	router.GET("/ping", handlePing)
	router.GET("/statistic", getStatistics)
	router.GET("/processing", getProcessing)
	router.POST("/config", setConfigurations)

	router.Run("localhost:8080")

}
