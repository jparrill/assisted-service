package api

import (
	"context"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

type ValidationStatus string

const (
	Success ValidationStatus = "success"
	Failure ValidationStatus = "failure"
	Pending ValidationStatus = "pending"
)

// ValidationResult hold result of operator validation
type ValidationResult struct {
	// ValidationId is an id of the validation
	ValidationId string
	// Status specifies the status of the validation: success, failure or pending
	Status ValidationStatus
	// Reasons hold list of reasons of a validation failure
	Reasons []string
}

// Manifests hold generated manifests
type Manifests struct {
	// Files maps file name to file's content
	Files map[string]string
}

// Operator provides generic API of an OLM operator installation plugin
type Operator interface {
	// GetType reports type of an operator this Operator manages
	GetType() models.OperatorType
	// GetDependencies provides a list of dependencies of the Operator
	GetDependencies() []models.OperatorType
	// ValidateCluster verifies whether this operator is valid for given cluster
	ValidateCluster(ctx context.Context, cluster *common.Cluster) (ValidationResult, error)
	// ValidateHost verifies whether this operator is valid for given host
	ValidateHost(context.Context, *common.Cluster, *models.Host) (ValidationResult, error)
	// GenerateManifests generates manifests for the operator
	GenerateManifests(*common.Cluster) (*Manifests, error)
	// GetCPURequirementForWorker provides worker CPU requirements for the operator
	GetCPURequirementForWorker(context.Context, *common.Cluster) (int64, error)
	// GetCPURequirementForMaster provides master CPU requirements for the operator
	GetCPURequirementForMaster(context.Context, *common.Cluster) (int64, error)
	// GetMemoryRequirementForWorker provides worker memory requirements for the operator in MB
	GetMemoryRequirementForWorker(ctx context.Context, cluster *common.Cluster) (int64, error)
	// GetMemoryRequirementForMaster provides master memory requirements for the operator in MB
	GetMemoryRequirementForMaster(ctx context.Context, cluster *common.Cluster) (int64, error)
	// GetDisksRequirementForMaster provides a number of disks required in a master
	GetDisksRequirementForMaster(context.Context, *common.Cluster) (int64, error)
	// GetDisksRequirementForWorker provides a number of disks required in a worker
	GetDisksRequirementForWorker(ctx context.Context, cluster *common.Cluster) (int64, error)
	// GetClusterValidationID returns cluster validation ID for the Operator
	GetClusterValidationID() string
	// GetHostValidationID returns host validation ID for the Operator
	GetHostValidationID() string
}
