package main

import (
	"context"
	"strconv"

	githubactions "github.com/sethvargo/go-githubactions"
	action "github.com/wwilkinson/aws-ecs-deploy-task-definition/pkg/deploy_task_definition"
)

type Config struct {
	TaskDefinition            string
	Cluster                   string
	WaitForServiceStability   bool
	CodeDeployAppSpec         string
	CodeDeployApplication     string
	CodeDeployDeploymentGroup string
}

func newFromInputs(action *githubactions.Action) (*Config, error) {
	waitForServiceStability, err := strconv.ParseBool(action.GetInput("wait-for-service-stability"))

	if err != nil {
		return nil, err
	}

	c := Config{
		TaskDefinition:            action.GetInput("task-definition"),
		Cluster:                   action.GetInput("cluster"),
		WaitForServiceStability:   waitForServiceStability,
		CodeDeployAppSpec:         action.GetInput("code-deploy-app-spec"),
		CodeDeployApplication:     action.GetInput("code-deploy-application"),
		CodeDeployDeploymentGroup: action.GetInput("code-deploy-deployment-group"),
	}
	return &c, nil
}

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := newFromInputs( action)

	if err != nil {
		return err
	}

	return action.Run(ctx, cfg)
}

func main() {
	action, err := run()
	if err != nil {
		action.Fatalf("%v", err)
	}
}
