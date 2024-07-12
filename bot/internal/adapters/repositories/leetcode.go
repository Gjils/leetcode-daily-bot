package repositories

import (
	"encoding/json"
	"fmt"

	"leetcodebot/internal/domain/repositories"
	"leetcodebot/internal/domain/entities"
	"leetcodebot/internal/domain/interfaces"
)
 
const url = "https://leetcode.com"

const graphQLUrl = url + "/graphql"

const dailyQuery = `
	query questionOfToday {
		activeDailyCodingChallengeQuestion {
			link
			question {
				difficulty
				title
			}
		}
}`

type LeetcodeRepository struct {
	GraphQLClient interfaces.GraphQLClient
}

func GetLeetcodeRepository(cli interfaces.GraphQLClient) repositoryInterfaces.LeetcodeApi {
	return LeetcodeRepository{GraphQLClient: cli}
}

type DailyResponse struct {
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
	
func (api LeetcodeRepository) GetDaily() (*entities.Problem, error) {
 
	request := interfaces.GraphQLRequest{
	 Query: dailyQuery,
	}
 
	response, err := api.GraphQLClient.SendGraphQLRequest(graphQLUrl, request)
	if err != nil {
		return nil, err
	}
 
	fmt.Println(string(response.Data))
 
	var responseObj DailyResponse
	if err := json.Unmarshal(response.Data, &responseObj); err != nil {
	 return nil, err
	}
	if len(response.Errors) > 0 {
	 fmt.Printf("Response errors: %+v\n", response.Errors)
	}

	result := entities.Problem{
		Link: url + responseObj.ActiveDailyCodingChallengeQuestion.Link,
		Title: responseObj.ActiveDailyCodingChallengeQuestion.Question.Title,
		Difficulty: responseObj.ActiveDailyCodingChallengeQuestion.Question.Difficulty,
	}

	return &result, nil
}