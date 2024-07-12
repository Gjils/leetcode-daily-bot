package serviceInterfaces

type ProblemsService interface {
	GetDailyInfo() (string, error)
	GetMorningInfo() (string, error)
	GetEveningInfo() (string, error)
}