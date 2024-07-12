package repositoryInterfaces

import "leetcodebot/internal/domain/entities"

type GroupRepository interface {
	GetAll() ([]entities.Group, error)
	Add(entities.GroupInfo) error
	Find(entities.GroupQuery) (bool, error)
	Delete(entities.GroupQuery) error
}