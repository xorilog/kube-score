package score

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zegl/kube-score/config"
	ks "github.com/zegl/kube-score/domain"
	"github.com/zegl/kube-score/scorecard"
)

func TestServiceTargetsDeploymentStrategyRolling(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "service-target-deployment.yaml", "Deployment Strategy", scorecard.GradeAllOK)
}

func TestServiceNotTargetsDeploymentStrategyNotRelevant(t *testing.T) {
	t.Parallel()
	skipped := wasSkipped(t, config.Configuration{
		AllFiles: []ks.NamedReader{testFile("service-not-target-deployment.yaml")},
	}, "Deployment Strategy")
	assert.True(t, skipped)
}

func TestServiceTargetsDeploymentStrategyNotRolling(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "service-target-deployment-not-rolling.yaml", "Deployment Strategy", scorecard.GradeWarning)
}

func TestServiceTargetsDeploymentStrategyNotSet(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "service-target-deployment-strategy-not-set.yaml", "Deployment Strategy", scorecard.GradeAllOK)
}

func TestServiceTargetsDeploymentReplicasOk(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "service-target-deployment.yaml", "Deployment Replicas", scorecard.GradeAllOK)
}

func TestServiceNotTargetsDeploymentReplicasNotRelevant(t *testing.T) {
	t.Parallel()
	skipped := wasSkipped(t, config.Configuration{
		AllFiles: []ks.NamedReader{testFile("service-not-target-deployment.yaml")},
	}, "Deployment Replicas")
	assert.True(t, skipped)
}

func TestServiceTargetsDeploymentReplicasNok(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "service-target-deployment-replica-1.yaml", "Deployment Replicas", scorecard.GradeWarning)
}

func TestHPATargetsDeployment(t *testing.T) {
	t.Parallel()
	skipped := wasSkipped(t, config.Configuration{
		AllFiles: []ks.NamedReader{testFile("hpa-target-deployment.yaml")},
	}, "Deployment Replicas")
	assert.True(t, skipped)
}
