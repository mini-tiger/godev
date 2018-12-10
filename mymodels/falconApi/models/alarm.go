package models

import "time"

type EventCases struct {
	Id            int `orm:"pk;auto"`
	Endpoint      string
	Metric        string
	Func          string
	Cond          string
	Note          string
	MaxStep       int
	CurrentStep   int
	Priority      int
	Status        string
	Timestamp     *time.Time
	UpdateAt      *time.Time
	ClosedAt      *time.Time
	ClosedNote    string
	UserModified  int64
	TplCreator    string
	ExpressionId  int64
	StrategyId    int64
	TemplateId    int64
	ProcessNote   int64
	ProcessStatus string
}

func (this EventCases) TableName() string {
	return "event_cases"
}
