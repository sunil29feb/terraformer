// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

type ActionOnFailure string

// Enum values for ActionOnFailure
const (
	ActionOnFailureTerminateJobFlow ActionOnFailure = "TERMINATE_JOB_FLOW"
	ActionOnFailureTerminateCluster ActionOnFailure = "TERMINATE_CLUSTER"
	ActionOnFailureCancelAndWait    ActionOnFailure = "CANCEL_AND_WAIT"
	ActionOnFailureContinue         ActionOnFailure = "CONTINUE"
)

func (enum ActionOnFailure) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActionOnFailure) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AdjustmentType string

// Enum values for AdjustmentType
const (
	AdjustmentTypeChangeInCapacity        AdjustmentType = "CHANGE_IN_CAPACITY"
	AdjustmentTypePercentChangeInCapacity AdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	AdjustmentTypeExactCapacity           AdjustmentType = "EXACT_CAPACITY"
)

func (enum AdjustmentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AdjustmentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AutoScalingPolicyState string

// Enum values for AutoScalingPolicyState
const (
	AutoScalingPolicyStatePending   AutoScalingPolicyState = "PENDING"
	AutoScalingPolicyStateAttaching AutoScalingPolicyState = "ATTACHING"
	AutoScalingPolicyStateAttached  AutoScalingPolicyState = "ATTACHED"
	AutoScalingPolicyStateDetaching AutoScalingPolicyState = "DETACHING"
	AutoScalingPolicyStateDetached  AutoScalingPolicyState = "DETACHED"
	AutoScalingPolicyStateFailed    AutoScalingPolicyState = "FAILED"
)

func (enum AutoScalingPolicyState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AutoScalingPolicyState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AutoScalingPolicyStateChangeReasonCode string

// Enum values for AutoScalingPolicyStateChangeReasonCode
const (
	AutoScalingPolicyStateChangeReasonCodeUserRequest      AutoScalingPolicyStateChangeReasonCode = "USER_REQUEST"
	AutoScalingPolicyStateChangeReasonCodeProvisionFailure AutoScalingPolicyStateChangeReasonCode = "PROVISION_FAILURE"
	AutoScalingPolicyStateChangeReasonCodeCleanupFailure   AutoScalingPolicyStateChangeReasonCode = "CLEANUP_FAILURE"
)

func (enum AutoScalingPolicyStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AutoScalingPolicyStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CancelStepsRequestStatus string

// Enum values for CancelStepsRequestStatus
const (
	CancelStepsRequestStatusSubmitted CancelStepsRequestStatus = "SUBMITTED"
	CancelStepsRequestStatusFailed    CancelStepsRequestStatus = "FAILED"
)

func (enum CancelStepsRequestStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CancelStepsRequestStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ClusterState string

// Enum values for ClusterState
const (
	ClusterStateStarting             ClusterState = "STARTING"
	ClusterStateBootstrapping        ClusterState = "BOOTSTRAPPING"
	ClusterStateRunning              ClusterState = "RUNNING"
	ClusterStateWaiting              ClusterState = "WAITING"
	ClusterStateTerminating          ClusterState = "TERMINATING"
	ClusterStateTerminated           ClusterState = "TERMINATED"
	ClusterStateTerminatedWithErrors ClusterState = "TERMINATED_WITH_ERRORS"
)

func (enum ClusterState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ClusterStateChangeReasonCode string

// Enum values for ClusterStateChangeReasonCode
const (
	ClusterStateChangeReasonCodeInternalError        ClusterStateChangeReasonCode = "INTERNAL_ERROR"
	ClusterStateChangeReasonCodeValidationError      ClusterStateChangeReasonCode = "VALIDATION_ERROR"
	ClusterStateChangeReasonCodeInstanceFailure      ClusterStateChangeReasonCode = "INSTANCE_FAILURE"
	ClusterStateChangeReasonCodeInstanceFleetTimeout ClusterStateChangeReasonCode = "INSTANCE_FLEET_TIMEOUT"
	ClusterStateChangeReasonCodeBootstrapFailure     ClusterStateChangeReasonCode = "BOOTSTRAP_FAILURE"
	ClusterStateChangeReasonCodeUserRequest          ClusterStateChangeReasonCode = "USER_REQUEST"
	ClusterStateChangeReasonCodeStepFailure          ClusterStateChangeReasonCode = "STEP_FAILURE"
	ClusterStateChangeReasonCodeAllStepsCompleted    ClusterStateChangeReasonCode = "ALL_STEPS_COMPLETED"
)

func (enum ClusterStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorGreaterThanOrEqual ComparisonOperator = "GREATER_THAN_OR_EQUAL"
	ComparisonOperatorGreaterThan        ComparisonOperator = "GREATER_THAN"
	ComparisonOperatorLessThan           ComparisonOperator = "LESS_THAN"
	ComparisonOperatorLessThanOrEqual    ComparisonOperator = "LESS_THAN_OR_EQUAL"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceCollectionType string

// Enum values for InstanceCollectionType
const (
	InstanceCollectionTypeInstanceFleet InstanceCollectionType = "INSTANCE_FLEET"
	InstanceCollectionTypeInstanceGroup InstanceCollectionType = "INSTANCE_GROUP"
)

func (enum InstanceCollectionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceCollectionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceFleetState string

// Enum values for InstanceFleetState
const (
	InstanceFleetStateProvisioning  InstanceFleetState = "PROVISIONING"
	InstanceFleetStateBootstrapping InstanceFleetState = "BOOTSTRAPPING"
	InstanceFleetStateRunning       InstanceFleetState = "RUNNING"
	InstanceFleetStateResizing      InstanceFleetState = "RESIZING"
	InstanceFleetStateSuspended     InstanceFleetState = "SUSPENDED"
	InstanceFleetStateTerminating   InstanceFleetState = "TERMINATING"
	InstanceFleetStateTerminated    InstanceFleetState = "TERMINATED"
)

func (enum InstanceFleetState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceFleetState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceFleetStateChangeReasonCode string

// Enum values for InstanceFleetStateChangeReasonCode
const (
	InstanceFleetStateChangeReasonCodeInternalError     InstanceFleetStateChangeReasonCode = "INTERNAL_ERROR"
	InstanceFleetStateChangeReasonCodeValidationError   InstanceFleetStateChangeReasonCode = "VALIDATION_ERROR"
	InstanceFleetStateChangeReasonCodeInstanceFailure   InstanceFleetStateChangeReasonCode = "INSTANCE_FAILURE"
	InstanceFleetStateChangeReasonCodeClusterTerminated InstanceFleetStateChangeReasonCode = "CLUSTER_TERMINATED"
)

func (enum InstanceFleetStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceFleetStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceFleetType string

// Enum values for InstanceFleetType
const (
	InstanceFleetTypeMaster InstanceFleetType = "MASTER"
	InstanceFleetTypeCore   InstanceFleetType = "CORE"
	InstanceFleetTypeTask   InstanceFleetType = "TASK"
)

func (enum InstanceFleetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceFleetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceGroupState string

// Enum values for InstanceGroupState
const (
	InstanceGroupStateProvisioning  InstanceGroupState = "PROVISIONING"
	InstanceGroupStateBootstrapping InstanceGroupState = "BOOTSTRAPPING"
	InstanceGroupStateRunning       InstanceGroupState = "RUNNING"
	InstanceGroupStateReconfiguring InstanceGroupState = "RECONFIGURING"
	InstanceGroupStateResizing      InstanceGroupState = "RESIZING"
	InstanceGroupStateSuspended     InstanceGroupState = "SUSPENDED"
	InstanceGroupStateTerminating   InstanceGroupState = "TERMINATING"
	InstanceGroupStateTerminated    InstanceGroupState = "TERMINATED"
	InstanceGroupStateArrested      InstanceGroupState = "ARRESTED"
	InstanceGroupStateShuttingDown  InstanceGroupState = "SHUTTING_DOWN"
	InstanceGroupStateEnded         InstanceGroupState = "ENDED"
)

func (enum InstanceGroupState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceGroupState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceGroupStateChangeReasonCode string

// Enum values for InstanceGroupStateChangeReasonCode
const (
	InstanceGroupStateChangeReasonCodeInternalError     InstanceGroupStateChangeReasonCode = "INTERNAL_ERROR"
	InstanceGroupStateChangeReasonCodeValidationError   InstanceGroupStateChangeReasonCode = "VALIDATION_ERROR"
	InstanceGroupStateChangeReasonCodeInstanceFailure   InstanceGroupStateChangeReasonCode = "INSTANCE_FAILURE"
	InstanceGroupStateChangeReasonCodeClusterTerminated InstanceGroupStateChangeReasonCode = "CLUSTER_TERMINATED"
)

func (enum InstanceGroupStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceGroupStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceGroupType string

// Enum values for InstanceGroupType
const (
	InstanceGroupTypeMaster InstanceGroupType = "MASTER"
	InstanceGroupTypeCore   InstanceGroupType = "CORE"
	InstanceGroupTypeTask   InstanceGroupType = "TASK"
)

func (enum InstanceGroupType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceGroupType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceRoleType string

// Enum values for InstanceRoleType
const (
	InstanceRoleTypeMaster InstanceRoleType = "MASTER"
	InstanceRoleTypeCore   InstanceRoleType = "CORE"
	InstanceRoleTypeTask   InstanceRoleType = "TASK"
)

func (enum InstanceRoleType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceRoleType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceState string

// Enum values for InstanceState
const (
	InstanceStateAwaitingFulfillment InstanceState = "AWAITING_FULFILLMENT"
	InstanceStateProvisioning        InstanceState = "PROVISIONING"
	InstanceStateBootstrapping       InstanceState = "BOOTSTRAPPING"
	InstanceStateRunning             InstanceState = "RUNNING"
	InstanceStateTerminated          InstanceState = "TERMINATED"
)

func (enum InstanceState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceStateChangeReasonCode string

// Enum values for InstanceStateChangeReasonCode
const (
	InstanceStateChangeReasonCodeInternalError     InstanceStateChangeReasonCode = "INTERNAL_ERROR"
	InstanceStateChangeReasonCodeValidationError   InstanceStateChangeReasonCode = "VALIDATION_ERROR"
	InstanceStateChangeReasonCodeInstanceFailure   InstanceStateChangeReasonCode = "INSTANCE_FAILURE"
	InstanceStateChangeReasonCodeBootstrapFailure  InstanceStateChangeReasonCode = "BOOTSTRAP_FAILURE"
	InstanceStateChangeReasonCodeClusterTerminated InstanceStateChangeReasonCode = "CLUSTER_TERMINATED"
)

func (enum InstanceStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of instance.
type JobFlowExecutionState string

// Enum values for JobFlowExecutionState
const (
	JobFlowExecutionStateStarting      JobFlowExecutionState = "STARTING"
	JobFlowExecutionStateBootstrapping JobFlowExecutionState = "BOOTSTRAPPING"
	JobFlowExecutionStateRunning       JobFlowExecutionState = "RUNNING"
	JobFlowExecutionStateWaiting       JobFlowExecutionState = "WAITING"
	JobFlowExecutionStateShuttingDown  JobFlowExecutionState = "SHUTTING_DOWN"
	JobFlowExecutionStateTerminated    JobFlowExecutionState = "TERMINATED"
	JobFlowExecutionStateCompleted     JobFlowExecutionState = "COMPLETED"
	JobFlowExecutionStateFailed        JobFlowExecutionState = "FAILED"
)

func (enum JobFlowExecutionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobFlowExecutionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MarketType string

// Enum values for MarketType
const (
	MarketTypeOnDemand MarketType = "ON_DEMAND"
	MarketTypeSpot     MarketType = "SPOT"
)

func (enum MarketType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MarketType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RepoUpgradeOnBoot string

// Enum values for RepoUpgradeOnBoot
const (
	RepoUpgradeOnBootSecurity RepoUpgradeOnBoot = "SECURITY"
	RepoUpgradeOnBootNone     RepoUpgradeOnBoot = "NONE"
)

func (enum RepoUpgradeOnBoot) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RepoUpgradeOnBoot) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScaleDownBehavior string

// Enum values for ScaleDownBehavior
const (
	ScaleDownBehaviorTerminateAtInstanceHour   ScaleDownBehavior = "TERMINATE_AT_INSTANCE_HOUR"
	ScaleDownBehaviorTerminateAtTaskCompletion ScaleDownBehavior = "TERMINATE_AT_TASK_COMPLETION"
)

func (enum ScaleDownBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScaleDownBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SpotProvisioningTimeoutAction string

// Enum values for SpotProvisioningTimeoutAction
const (
	SpotProvisioningTimeoutActionSwitchToOnDemand SpotProvisioningTimeoutAction = "SWITCH_TO_ON_DEMAND"
	SpotProvisioningTimeoutActionTerminateCluster SpotProvisioningTimeoutAction = "TERMINATE_CLUSTER"
)

func (enum SpotProvisioningTimeoutAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SpotProvisioningTimeoutAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Statistic string

// Enum values for Statistic
const (
	StatisticSampleCount Statistic = "SAMPLE_COUNT"
	StatisticAverage     Statistic = "AVERAGE"
	StatisticSum         Statistic = "SUM"
	StatisticMinimum     Statistic = "MINIMUM"
	StatisticMaximum     Statistic = "MAXIMUM"
)

func (enum Statistic) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Statistic) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StepCancellationOption string

// Enum values for StepCancellationOption
const (
	StepCancellationOptionSendInterrupt    StepCancellationOption = "SEND_INTERRUPT"
	StepCancellationOptionTerminateProcess StepCancellationOption = "TERMINATE_PROCESS"
)

func (enum StepCancellationOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StepCancellationOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StepExecutionState string

// Enum values for StepExecutionState
const (
	StepExecutionStatePending     StepExecutionState = "PENDING"
	StepExecutionStateRunning     StepExecutionState = "RUNNING"
	StepExecutionStateContinue    StepExecutionState = "CONTINUE"
	StepExecutionStateCompleted   StepExecutionState = "COMPLETED"
	StepExecutionStateCancelled   StepExecutionState = "CANCELLED"
	StepExecutionStateFailed      StepExecutionState = "FAILED"
	StepExecutionStateInterrupted StepExecutionState = "INTERRUPTED"
)

func (enum StepExecutionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StepExecutionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StepState string

// Enum values for StepState
const (
	StepStatePending       StepState = "PENDING"
	StepStateCancelPending StepState = "CANCEL_PENDING"
	StepStateRunning       StepState = "RUNNING"
	StepStateCompleted     StepState = "COMPLETED"
	StepStateCancelled     StepState = "CANCELLED"
	StepStateFailed        StepState = "FAILED"
	StepStateInterrupted   StepState = "INTERRUPTED"
)

func (enum StepState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StepState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StepStateChangeReasonCode string

// Enum values for StepStateChangeReasonCode
const (
	StepStateChangeReasonCodeNone StepStateChangeReasonCode = "NONE"
)

func (enum StepStateChangeReasonCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StepStateChangeReasonCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Unit string

// Enum values for Unit
const (
	UnitNone               Unit = "NONE"
	UnitSeconds            Unit = "SECONDS"
	UnitMicroSeconds       Unit = "MICRO_SECONDS"
	UnitMilliSeconds       Unit = "MILLI_SECONDS"
	UnitBytes              Unit = "BYTES"
	UnitKiloBytes          Unit = "KILO_BYTES"
	UnitMegaBytes          Unit = "MEGA_BYTES"
	UnitGigaBytes          Unit = "GIGA_BYTES"
	UnitTeraBytes          Unit = "TERA_BYTES"
	UnitBits               Unit = "BITS"
	UnitKiloBits           Unit = "KILO_BITS"
	UnitMegaBits           Unit = "MEGA_BITS"
	UnitGigaBits           Unit = "GIGA_BITS"
	UnitTeraBits           Unit = "TERA_BITS"
	UnitPercent            Unit = "PERCENT"
	UnitCount              Unit = "COUNT"
	UnitBytesPerSecond     Unit = "BYTES_PER_SECOND"
	UnitKiloBytesPerSecond Unit = "KILO_BYTES_PER_SECOND"
	UnitMegaBytesPerSecond Unit = "MEGA_BYTES_PER_SECOND"
	UnitGigaBytesPerSecond Unit = "GIGA_BYTES_PER_SECOND"
	UnitTeraBytesPerSecond Unit = "TERA_BYTES_PER_SECOND"
	UnitBitsPerSecond      Unit = "BITS_PER_SECOND"
	UnitKiloBitsPerSecond  Unit = "KILO_BITS_PER_SECOND"
	UnitMegaBitsPerSecond  Unit = "MEGA_BITS_PER_SECOND"
	UnitGigaBitsPerSecond  Unit = "GIGA_BITS_PER_SECOND"
	UnitTeraBitsPerSecond  Unit = "TERA_BITS_PER_SECOND"
	UnitCountPerSecond     Unit = "COUNT_PER_SECOND"
)

func (enum Unit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Unit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}