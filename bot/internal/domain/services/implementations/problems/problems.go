package problemsService

import (
	"fmt"
	"leetcodebot/internal/domain/repositories"
	serviceInterfaces "leetcodebot/internal/domain/services/interfaces"
)

const dailyMessage = `Сегодняшний дейлик: %v
Сложность: %v
			
%v`
const morningMessage = `Доброе утро!
Сегодняшняя задача: %v
Сложность: %v

Удачного решения!
%v
`
const eveningMessage = `
День скоро кончится. Не забудь решить дейлик!

%v`

type ProblemsService struct {
	rep repositoryInterfaces.LeetcodeApi
}

func GetProblemsService(rep repositoryInterfaces.LeetcodeApi) serviceInterfaces.ProblemsService {
	return ProblemsService{rep: rep}
}

func (s ProblemsService) GetDailyInfo() (string, error) {
	problem, err := s.rep.GetDaily()
	if err != nil {
		return "", err
	}
	fmt.Println(problem)
	return fmt.Sprintf(dailyMessage, problem.Title, problem.Difficulty, problem.Link), nil
}

func (s ProblemsService) GetMorningInfo() (string, error) {
	problem, err := s.rep.GetDaily()
	if err != nil {
		return "", err
	}
	fmt.Println(problem)
	return fmt.Sprintf(morningMessage, problem.Title, problem.Difficulty, problem.Link), nil
}

func (s ProblemsService) GetEveningInfo() (string, error) {
	problem, err := s.rep.GetDaily()
	if err != nil {
		return "", err
	}
	fmt.Println(problem)
	return fmt.Sprintf(eveningMessage, problem.Link), nil
}