package leetcodeApi

import (
	"encoding/json"
	"fmt"
	"log"

	httpService "leetcodebot/data/httpService"
)
 
const Url = "https://leetcode.com"

const GraphQLUrl = Url + "/graphql"

const dailyQuery = `
	query questionOfToday {
		activeDailyCodingChallengeQuestion {
			date
			link
			userStatus
			question {
				difficulty
				title
			}
		}
}`

type DailyItem struct {
	ActiveDailyCodingChallengeQuestion DailyQuestion
}
	
type DailyQuestion struct {
	Link string
	Date string
	UserStatus string
	Question QuestionInfo
}
 
type QuestionInfo struct {
	Difficulty string
	Title string
}
	
func GetDaily() DailyItem {
 
	request := httpService.GraphQLRequest{
	 Query: dailyQuery,
	}
 
	response, err := httpService.SendGraphQLRequest(GraphQLUrl, request)
	if err != nil {
	 log.Fatalf("Failed to send GraphQL request: %v", err)
	}
 
	fmt.Println(string(response.Data))
 
	var result DailyItem
	if err := json.Unmarshal(response.Data, &result); err != nil {
	 log.Fatalf("Failed to unmarshal response data: %v", err)
	}
	if len(response.Errors) > 0 {
	 fmt.Printf("Response errors: %+v\n", response.Errors)
	}
	return result
}