package leetcodeApi

import (
	"encoding/json"
	"fmt"
	"log"

	httpService "leetcodebot/data/httpService"
)
 
 type DailyItem struct {
	 ActiveDailyCodingChallengeQuestion DailyQuestion
 }
 
 type DailyQuestion struct {
	 Link string
	 Date string
	 Question QuestionInfo
 }

 type QuestionInfo struct {
	Difficulty string
	Title string
 }
 
 
 const Url = "https://leetcode.com"

 const GraphQLUrl = Url + "/graphql"

 const dailyQuery = `
	 query questionOfToday {
		 activeDailyCodingChallengeQuestion {
			 date
			 link
			 question {
				difficulty
				title
			}
		}
	}`
 func GetDaily() DailyItem {
 
	// Создание запроса
	request := httpService.GraphQLRequest{
	 Query: dailyQuery,
	}
 
	// Отправка запроса и получение ответа
	response, err := httpService.SendGraphQLRequest(GraphQLUrl, request)
	if err != nil {
	 log.Fatalf("Failed to send GraphQL request: %v", err)
	}
 
	fmt.Println(string(response.Data))
 
	// Обработка ответа
	var result DailyItem
	if err := json.Unmarshal(response.Data, &result); err != nil {
	 log.Fatalf("Failed to unmarshal response data: %v", err)
	}
 //  fmt.Println(result)
	if len(response.Errors) > 0 {
	 fmt.Printf("Response errors: %+v\n", response.Errors)
	}
	return result
 }
