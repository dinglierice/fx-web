package workflow

import (
	"context"
)

type WorkflowService struct {
}

func (s *WorkflowService) StartWorkflow(ctx context.Context, request *StartWorkflowRequest) (*StartWorkflowResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WorkflowService) GetWorkflowStatus(ctx context.Context, request *GetWorkflowStatusRequest) (*GetWorkflowStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WorkflowService) mustEmbedUnimplementedWorkflowServiceServer() {
	//TODO implement me
	panic("implement me")
}
