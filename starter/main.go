package main

import (
	"context"
	"log"

	"github.com/arjunmalhotra1/hellotemporal-1/helloworkflow"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client ", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "workflow-options-id-starter-main.go",
		TaskQueue: "hello-world-queue-1",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworkflow.Workflow, "string-name-to-workflow")
	if err != nil {
		log.Fatalln("Unable to execute workflow.", err)
	}
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result.")
	}
	log.Println("Workflow result:", result)
}
