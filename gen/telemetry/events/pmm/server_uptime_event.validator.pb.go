// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: telemetry/events/pmm/server_uptime_event.proto

package pmmv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ServerUptimeEvent) Validate() error {
	if !(len(this.Id) == 16) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '16'`, this.Id))
	}
	if !(len(this.Version) > 4) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must have a length greater than '4'`, this.Version))
	}
	if nil == this.UpDuration {
		return github_com_mwitkow_go_proto_validators.FieldError("UpDuration", fmt.Errorf("message must exist"))
	}
	if this.UpDuration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpDuration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpDuration", err)
		}
	}
	if _, ok := DistributionMethod_name[int32(this.DistributionMethod)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("DistributionMethod", fmt.Errorf(`value '%v' must be a valid DistributionMethod field`, this.DistributionMethod))
	}
	if this.SttEnabled != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SttEnabled); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SttEnabled", err)
		}
	}
	if this.IaEnabled != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IaEnabled); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IaEnabled", err)
		}
	}
	return nil
}
func (this *ServerMetric) Validate() error {
	if !(len(this.Id) == 16) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '16'`, this.Id))
	}
	if nil == this.Time {
		return github_com_mwitkow_go_proto_validators.FieldError("Time", fmt.Errorf("message must exist"))
	}
	if this.Time != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Time); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Time", err)
		}
	}
	if !(len(this.PmmServerTelemetryId) == 16) {
		return github_com_mwitkow_go_proto_validators.FieldError("PmmServerTelemetryId", fmt.Errorf(`value '%v' must have a length equal to '16'`, this.PmmServerTelemetryId))
	}
	if !(len(this.PmmServerVersion) > 4) {
		return github_com_mwitkow_go_proto_validators.FieldError("PmmServerVersion", fmt.Errorf(`value '%v' must have a length greater than '4'`, this.PmmServerVersion))
	}
	if nil == this.UpDuration {
		return github_com_mwitkow_go_proto_validators.FieldError("UpDuration", fmt.Errorf("message must exist"))
	}
	if this.UpDuration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpDuration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpDuration", err)
		}
	}
	if _, ok := DistributionMethod_name[int32(this.DistributionMethod)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("DistributionMethod", fmt.Errorf(`value '%v' must be a valid DistributionMethod field`, this.DistributionMethod))
	}
	for _, item := range this.Metrics {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metrics", err)
			}
		}
	}
	return nil
}
func (this *ServerMetric_Metric) Validate() error {
	return nil
}
