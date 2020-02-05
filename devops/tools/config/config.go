package config

type GroupName string

type Build struct {
	Name       string
	Path       string
	Dockerfile string
}

// Frontend is a nextjs application from /frontend/cmd/<name>
type Frontend struct {
	Group GroupName
	Build
}

// Service is a go service from /backend/cmd/<name>
type Service struct {
	Group GroupName
	Build
}

type Group struct {
	Name      GroupName
	Frontends []Frontend
	Services  []Service
}

type Groups map[GroupName]*Group
