// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volume.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = config.Config{}
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VolumesServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *VolumesServiceProperties) Reset()                    { *m = VolumesServiceProperties{} }
func (m *VolumesServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*VolumesServiceProperties) ProtoMessage()               {}
func (*VolumesServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

func (m *VolumesServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeVolumesInput struct {
	Limit      int32    `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset     int32    `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	SearchWord string   `protobuf:"bytes,3,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Status     []string `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	Tags       []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	Verbose    int32    `protobuf:"varint,6,opt,name=verbose" json:"verbose,omitempty"`
	VolumeType int32    `protobuf:"varint,7,opt,name=volume_type,json=volumeType" json:"volume_type,omitempty"`
	Volumes    []string `protobuf:"bytes,8,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *DescribeVolumesInput) Reset()                    { *m = DescribeVolumesInput{} }
func (m *DescribeVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVolumesInput) ProtoMessage()               {}
func (*DescribeVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{1} }

func (m *DescribeVolumesInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeVolumesInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeVolumesInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeVolumesInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeVolumesInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeVolumesInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeVolumesInput) GetVolumeType() int32 {
	if m != nil {
		return m.VolumeType
	}
	return 0
}

func (m *DescribeVolumesInput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type DescribeVolumesOutput struct {
	Action     string    `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32     `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string    `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	TotalCount int32     `protobuf:"varint,4,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	VolumeSet  []*Volume `protobuf:"bytes,6,rep,name=volume_set,json=volumeSet" json:"volume_set,omitempty"`
}

func (m *DescribeVolumesOutput) Reset()                    { *m = DescribeVolumesOutput{} }
func (m *DescribeVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVolumesOutput) ProtoMessage()               {}
func (*DescribeVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{2} }

func (m *DescribeVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeVolumesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeVolumesOutput) GetVolumeSet() []*Volume {
	if m != nil {
		return m.VolumeSet
	}
	return nil
}

type CreateVolumesInput struct {
	Count      int32  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Size       int32  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	VolumeName string `protobuf:"bytes,3,opt,name=volume_name,json=volumeName" json:"volume_name,omitempty"`
	VolumeType int32  `protobuf:"varint,4,opt,name=volume_type,json=volumeType" json:"volume_type,omitempty"`
}

func (m *CreateVolumesInput) Reset()                    { *m = CreateVolumesInput{} }
func (m *CreateVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumesInput) ProtoMessage()               {}
func (*CreateVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{3} }

func (m *CreateVolumesInput) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CreateVolumesInput) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CreateVolumesInput) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

func (m *CreateVolumesInput) GetVolumeType() int32 {
	if m != nil {
		return m.VolumeType
	}
	return 0
}

type CreateVolumesOutput struct {
	Action  string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string   `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Volumes []string `protobuf:"bytes,5,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *CreateVolumesOutput) Reset()                    { *m = CreateVolumesOutput{} }
func (m *CreateVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumesOutput) ProtoMessage()               {}
func (*CreateVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{4} }

func (m *CreateVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateVolumesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *CreateVolumesOutput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type DeleteVolumesInput struct {
	Volumes []string `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *DeleteVolumesInput) Reset()                    { *m = DeleteVolumesInput{} }
func (m *DeleteVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteVolumesInput) ProtoMessage()               {}
func (*DeleteVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{5} }

func (m *DeleteVolumesInput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type DeleteVolumesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DeleteVolumesOutput) Reset()                    { *m = DeleteVolumesOutput{} }
func (m *DeleteVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteVolumesOutput) ProtoMessage()               {}
func (*DeleteVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{6} }

func (m *DeleteVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteVolumesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type AttachVolumesInput struct {
	Instance string   `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	Volumes  []string `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *AttachVolumesInput) Reset()                    { *m = AttachVolumesInput{} }
func (m *AttachVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*AttachVolumesInput) ProtoMessage()               {}
func (*AttachVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{7} }

func (m *AttachVolumesInput) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *AttachVolumesInput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type AttachVolumesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *AttachVolumesOutput) Reset()                    { *m = AttachVolumesOutput{} }
func (m *AttachVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachVolumesOutput) ProtoMessage()               {}
func (*AttachVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{8} }

func (m *AttachVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AttachVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AttachVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AttachVolumesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type DetachVolumesInput struct {
	Instance string   `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	Volumes  []string `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *DetachVolumesInput) Reset()                    { *m = DetachVolumesInput{} }
func (m *DetachVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*DetachVolumesInput) ProtoMessage()               {}
func (*DetachVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{9} }

func (m *DetachVolumesInput) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *DetachVolumesInput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type DetachVolumesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DetachVolumesOutput) Reset()                    { *m = DetachVolumesOutput{} }
func (m *DetachVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachVolumesOutput) ProtoMessage()               {}
func (*DetachVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{10} }

func (m *DetachVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DetachVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DetachVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DetachVolumesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ResizeVolumesInput struct {
	Size    int32    `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Volumes []string `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *ResizeVolumesInput) Reset()                    { *m = ResizeVolumesInput{} }
func (m *ResizeVolumesInput) String() string            { return proto.CompactTextString(m) }
func (*ResizeVolumesInput) ProtoMessage()               {}
func (*ResizeVolumesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{11} }

func (m *ResizeVolumesInput) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ResizeVolumesInput) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type ResizeVolumesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *ResizeVolumesOutput) Reset()                    { *m = ResizeVolumesOutput{} }
func (m *ResizeVolumesOutput) String() string            { return proto.CompactTextString(m) }
func (*ResizeVolumesOutput) ProtoMessage()               {}
func (*ResizeVolumesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{12} }

func (m *ResizeVolumesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ResizeVolumesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ResizeVolumesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResizeVolumesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ModifyVolumeAttributesInput struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Volume      string `protobuf:"bytes,2,opt,name=volume" json:"volume,omitempty"`
	VolumeName  string `protobuf:"bytes,3,opt,name=volume_name,json=volumeName" json:"volume_name,omitempty"`
}

func (m *ModifyVolumeAttributesInput) Reset()                    { *m = ModifyVolumeAttributesInput{} }
func (m *ModifyVolumeAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyVolumeAttributesInput) ProtoMessage()               {}
func (*ModifyVolumeAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{13} }

func (m *ModifyVolumeAttributesInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ModifyVolumeAttributesInput) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *ModifyVolumeAttributesInput) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

type ModifyVolumeAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifyVolumeAttributesOutput) Reset()                    { *m = ModifyVolumeAttributesOutput{} }
func (m *ModifyVolumeAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyVolumeAttributesOutput) ProtoMessage()               {}
func (*ModifyVolumeAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{14} }

func (m *ModifyVolumeAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifyVolumeAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifyVolumeAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VolumesServiceProperties)(nil), "service.VolumesServiceProperties")
	proto.RegisterType((*DescribeVolumesInput)(nil), "service.DescribeVolumesInput")
	proto.RegisterType((*DescribeVolumesOutput)(nil), "service.DescribeVolumesOutput")
	proto.RegisterType((*CreateVolumesInput)(nil), "service.CreateVolumesInput")
	proto.RegisterType((*CreateVolumesOutput)(nil), "service.CreateVolumesOutput")
	proto.RegisterType((*DeleteVolumesInput)(nil), "service.DeleteVolumesInput")
	proto.RegisterType((*DeleteVolumesOutput)(nil), "service.DeleteVolumesOutput")
	proto.RegisterType((*AttachVolumesInput)(nil), "service.AttachVolumesInput")
	proto.RegisterType((*AttachVolumesOutput)(nil), "service.AttachVolumesOutput")
	proto.RegisterType((*DetachVolumesInput)(nil), "service.DetachVolumesInput")
	proto.RegisterType((*DetachVolumesOutput)(nil), "service.DetachVolumesOutput")
	proto.RegisterType((*ResizeVolumesInput)(nil), "service.ResizeVolumesInput")
	proto.RegisterType((*ResizeVolumesOutput)(nil), "service.ResizeVolumesOutput")
	proto.RegisterType((*ModifyVolumeAttributesInput)(nil), "service.ModifyVolumeAttributesInput")
	proto.RegisterType((*ModifyVolumeAttributesOutput)(nil), "service.ModifyVolumeAttributesOutput")
}

type VolumesServiceInterface interface {
	DescribeVolumes(in *DescribeVolumesInput) (out *DescribeVolumesOutput, err error)
	CreateVolumes(in *CreateVolumesInput) (out *CreateVolumesOutput, err error)
	DeleteVolumes(in *DeleteVolumesInput) (out *DeleteVolumesOutput, err error)
	AttachVolumes(in *AttachVolumesInput) (out *AttachVolumesOutput, err error)
	DetachVolumes(in *DetachVolumesInput) (out *DetachVolumesOutput, err error)
	ResizeVolumes(in *ResizeVolumesInput) (out *ResizeVolumesOutput, err error)
	ModifyVolumeAttributes(in *ModifyVolumeAttributesInput) (out *ModifyVolumeAttributesOutput, err error)
}

type VolumesService struct {
	Config     *config.Config
	Properties *VolumesServiceProperties
}

func NewVolumesService(conf *config.Config, zone string) (p *VolumesService) {
	return &VolumesService{
		Config:     conf,
		Properties: &VolumesServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Volumes(zone string) (*VolumesService, error) {
	properties := &VolumesServiceProperties{
		Zone: zone,
	}

	return &VolumesService{Config: s.Config, Properties: properties}, nil
}

func (p *VolumesService) DescribeVolumes(in *DescribeVolumesInput) (out *DescribeVolumesOutput, err error) {
	if in == nil {
		in = &DescribeVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) CreateVolumes(in *CreateVolumesInput) (out *CreateVolumesOutput, err error) {
	if in == nil {
		in = &CreateVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *CreateVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) DeleteVolumes(in *DeleteVolumesInput) (out *DeleteVolumesOutput, err error) {
	if in == nil {
		in = &DeleteVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DeleteVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) AttachVolumes(in *AttachVolumesInput) (out *AttachVolumesOutput, err error) {
	if in == nil {
		in = &AttachVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &AttachVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *AttachVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) DetachVolumes(in *DetachVolumesInput) (out *DetachVolumesOutput, err error) {
	if in == nil {
		in = &DetachVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &DetachVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DetachVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) ResizeVolumes(in *ResizeVolumesInput) (out *ResizeVolumesOutput, err error) {
	if in == nil {
		in = &ResizeVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeVolumes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ResizeVolumesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *ResizeVolumesInput) Validate() error {
	return nil
}

func (p *VolumesService) ModifyVolumeAttributes(in *ModifyVolumeAttributesInput) (out *ModifyVolumeAttributesOutput, err error) {
	if in == nil {
		in = &ModifyVolumeAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVolumeAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyVolumeAttributesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *ModifyVolumeAttributesInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor29) }

var fileDescriptor29 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xdb, 0xfc, 0x34, 0x13, 0x4a, 0xa5, 0xed, 0x8f, 0x4c, 0x5a, 0x20, 0xb2, 0x40, 0xea,
	0x29, 0x87, 0xf2, 0x04, 0xa5, 0xbd, 0xb4, 0x12, 0x50, 0xb9, 0x08, 0x8e, 0x91, 0x63, 0x4f, 0x5b,
	0x97, 0xc4, 0x1b, 0xed, 0x8e, 0x03, 0xe9, 0x89, 0x77, 0xe0, 0x59, 0x78, 0x04, 0x5e, 0x85, 0xe7,
	0x40, 0xde, 0x5d, 0x87, 0xdd, 0xd8, 0xb4, 0x17, 0x9a, 0xdb, 0xce, 0xec, 0xec, 0x37, 0xdf, 0x37,
	0xbb, 0x33, 0x36, 0x3c, 0x99, 0xf1, 0x71, 0x3e, 0xc1, 0xc1, 0x54, 0x70, 0xe2, 0xac, 0x2d, 0x51,
	0xcc, 0xd2, 0x18, 0x7b, 0x5d, 0x9a, 0x4f, 0x51, 0x6a, 0x6f, 0x30, 0x00, 0xff, 0x93, 0x8a, 0x92,
	0x97, 0x7a, 0xfb, 0x42, 0xf0, 0x29, 0x0a, 0x4a, 0x51, 0x32, 0x06, 0x8d, 0x3b, 0x9e, 0xa1, 0xef,
	0xf5, 0xbd, 0xc3, 0x4e, 0xa8, 0xd6, 0xc1, 0x6f, 0x0f, 0x76, 0x4e, 0x51, 0xc6, 0x22, 0x1d, 0xa1,
	0x39, 0x78, 0x96, 0x4d, 0x73, 0x62, 0x3b, 0xd0, 0x1c, 0xa7, 0x93, 0x94, 0x54, 0x74, 0x33, 0xd4,
	0x06, 0xdb, 0x83, 0x16, 0xbf, 0xba, 0x92, 0x48, 0xfe, 0x9a, 0x72, 0x1b, 0x8b, 0xbd, 0x84, 0xae,
	0xc4, 0x48, 0xc4, 0x37, 0xc3, 0xaf, 0x5c, 0x24, 0xfe, 0xba, 0xca, 0x00, 0xda, 0xf5, 0x99, 0x8b,
	0xa4, 0x38, 0x28, 0x29, 0xa2, 0x5c, 0xfa, 0x8d, 0xfe, 0xfa, 0x61, 0x27, 0x34, 0x56, 0xc1, 0x89,
	0xa2, 0x6b, 0xe9, 0x37, 0x95, 0x57, 0xad, 0x99, 0x0f, 0xed, 0x19, 0x8a, 0x11, 0x97, 0xe8, 0xb7,
	0x54, 0x96, 0xd2, 0x2c, 0xd2, 0xe8, 0x1a, 0x0c, 0x0b, 0xcd, 0x7e, 0x5b, 0xed, 0x82, 0x76, 0x7d,
	0x9c, 0x4f, 0x51, 0x1d, 0xd5, 0x2a, 0xfc, 0x0d, 0x85, 0x58, 0x9a, 0xc1, 0x4f, 0x0f, 0x76, 0x97,
	0x84, 0x7e, 0xc8, 0xa9, 0x50, 0xba, 0x07, 0xad, 0x28, 0xa6, 0x94, 0x67, 0xa6, 0x30, 0xc6, 0x62,
	0xcf, 0x60, 0x43, 0x20, 0x0d, 0x63, 0x9e, 0xa0, 0x51, 0xdb, 0x16, 0x48, 0x27, 0x3c, 0x51, 0x69,
	0x26, 0x28, 0x65, 0x74, 0x8d, 0x46, 0x6a, 0x69, 0x16, 0x0c, 0x89, 0x53, 0x34, 0x1e, 0xc6, 0x3c,
	0xcf, 0xc8, 0x6f, 0x68, 0x86, 0xca, 0x75, 0x52, 0x78, 0xd8, 0x00, 0x0c, 0xdf, 0x61, 0x51, 0xc5,
	0x56, 0x7f, 0xfd, 0xb0, 0x7b, 0xb4, 0x35, 0x30, 0x77, 0x39, 0xd0, 0xcc, 0xc2, 0x8e, 0x0e, 0xb9,
	0x44, 0x0a, 0xbe, 0x7b, 0xc0, 0x4e, 0x04, 0x46, 0x54, 0xb9, 0x1e, 0x9d, 0xc1, 0x5c, 0x8f, 0x32,
	0x8a, 0x6a, 0xca, 0xf4, 0xae, 0xa4, 0xab, 0xd6, 0x56, 0xcd, 0xb2, 0x68, 0x52, 0xf2, 0x35, 0x1c,
	0xde, 0x47, 0x93, 0x4a, 0x51, 0x1b, 0xcb, 0x45, 0x0d, 0x7e, 0x78, 0xb0, 0xed, 0x50, 0x78, 0x8c,
	0xc2, 0xed, 0x42, 0xeb, 0x96, 0x8f, 0x86, 0x69, 0xa2, 0x08, 0x74, 0xc2, 0xe6, 0x2d, 0x1f, 0x9d,
	0x25, 0xf6, 0x85, 0x36, 0xdd, 0x0b, 0x1d, 0x00, 0x3b, 0xc5, 0x31, 0x2e, 0xd5, 0xc5, 0x8a, 0xf7,
	0xdc, 0xf8, 0x39, 0x6c, 0x3b, 0xf1, 0xab, 0x13, 0x11, 0x9c, 0x03, 0x3b, 0x26, 0x8a, 0xe2, 0x1b,
	0x87, 0x6a, 0x0f, 0x36, 0xd2, 0x4c, 0x52, 0x94, 0xc5, 0x65, 0x4b, 0x2e, 0x6c, 0x5b, 0xc6, 0x5a,
	0x45, 0x86, 0x83, 0xb5, 0x5a, 0x19, 0xa7, 0xf8, 0xff, 0x64, 0x38, 0x58, 0x2b, 0x94, 0xf1, 0x16,
	0x58, 0x88, 0x45, 0x6b, 0x38, 0x32, 0xca, 0xd6, 0xf1, 0xac, 0xd6, 0xb9, 0x97, 0xbe, 0x83, 0xb1,
	0x42, 0xfa, 0xdf, 0x60, 0xff, 0x1d, 0x4f, 0xd2, 0xab, 0xb9, 0x4e, 0x7d, 0x4c, 0x24, 0xd2, 0x51,
	0x4e, 0xa5, 0x8e, 0x3e, 0x74, 0x13, 0x35, 0xe6, 0xa6, 0x16, 0x0f, 0xdb, 0x55, 0x90, 0xd4, 0x32,
	0x14, 0x95, 0x4e, 0x68, 0xac, 0x07, 0x07, 0x45, 0xf0, 0x05, 0x0e, 0xea, 0x33, 0x3f, 0x82, 0xfa,
	0xa3, 0x5f, 0x0d, 0x78, 0xea, 0x7e, 0xc9, 0xd8, 0x05, 0x6c, 0x2d, 0x4d, 0x70, 0xf6, 0x7c, 0x31,
	0x39, 0xeb, 0x3e, 0x62, 0xbd, 0x17, 0xff, 0xda, 0x36, 0x8c, 0xcf, 0x61, 0xd3, 0x19, 0x6c, 0x6c,
	0x7f, 0x71, 0xa0, 0x3a, 0x73, 0x7b, 0x07, 0xf5, 0x9b, 0x7f, 0xb1, 0x9c, 0xf9, 0x62, 0x61, 0x55,
	0xe7, 0x94, 0x85, 0x55, 0x37, 0x94, 0xce, 0x61, 0xd3, 0x69, 0x72, 0x0b, 0xab, 0x3a, 0x48, 0x2c,
	0xac, 0xba, 0xc9, 0xa0, 0x78, 0xd5, 0x63, 0x55, 0xbb, 0xd9, 0xe1, 0x55, 0x8b, 0xe5, 0x3c, 0x7b,
	0x0b, 0xab, 0xda, 0x52, 0x16, 0x56, 0x5d, 0xaf, 0x20, 0xec, 0xd5, 0xbf, 0x26, 0xf6, 0x6a, 0x71,
	0xee, 0x9e, 0x87, 0xde, 0x7b, 0xfd, 0x40, 0x94, 0x4e, 0x33, 0x6a, 0xa9, 0xff, 0xa2, 0x37, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x26, 0x6c, 0xb3, 0x3d, 0x09, 0x00, 0x00,
}
