package main

import (
	"leetcodebot/internal/adapters/handlers"
	"leetcodebot/internal/adapters/repositories"
	config "leetcodebot/internal/config"
	manageService "leetcodebot/internal/domain/services/implementations/manage"
	problemsService "leetcodebot/internal/domain/services/implementations/problems"
	serviceInterfaces "leetcodebot/internal/domain/services/interfaces"
	"log"
)

func main() {
	mongoClient := config.GetMongoClient()
	graphQLClient := config.GetGraphQLClient()
	telegramClient, err := config.GetTelegramClient()
	if err != nil {
		log.Fatal(err)
	}

	groupsRepository := repositories.GetGroupsRepository(mongoClient)
	leetcodeRepository := repositories.GetLeetcodeRepository(graphQLClient)


	manageService := manageService.GetManageService(groupsRepository)
	problemsService := problemsService.GetProblemsService(leetcodeRepository)

	services := serviceInterfaces.Services{
		ManageService: manageService,
		ProblemsService: problemsService,
	}

	handlers := handlers.GetHandlers(services, telegramClient)

	handlers.Start()

	telegramClient.Start()
}