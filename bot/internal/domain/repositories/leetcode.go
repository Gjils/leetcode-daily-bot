package repositoryInterfaces

import "leetcodebot/internal/domain/entities"

type LeetcodeApi interface {
	GetDaily() (*entities.Problem, error)

}