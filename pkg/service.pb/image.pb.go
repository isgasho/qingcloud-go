// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImageServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ImageServiceProperties) Reset()                    { *m = ImageServiceProperties{} }
func (m *ImageServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ImageServiceProperties) ProtoMessage()               {}
func (*ImageServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *ImageServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeImagesInput struct {
	Images        []string `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
	ProcessorType string   `protobuf:"bytes,2,opt,name=processor_type,json=processorType" json:"processor_type,omitempty"`
	OsFamily      string   `protobuf:"bytes,3,opt,name=os_family,json=osFamily" json:"os_family,omitempty"`
	Visibility    string   `protobuf:"bytes,4,opt,name=visibility" json:"visibility,omitempty"`
	Provider      string   `protobuf:"bytes,5,opt,name=provider" json:"provider,omitempty"`
	Status        []string `protobuf:"bytes,6,rep,name=status" json:"status,omitempty"`
	SearchWord    string   `protobuf:"bytes,7,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Verbose       int32    `protobuf:"varint,8,opt,name=verbose" json:"verbose,omitempty"`
	Offset        int32    `protobuf:"varint,9,opt,name=offset" json:"offset,omitempty"`
	Limit         int32    `protobuf:"varint,10,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeImagesInput) Reset()                    { *m = DescribeImagesInput{} }
func (m *DescribeImagesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeImagesInput) ProtoMessage()               {}
func (*DescribeImagesInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *DescribeImagesInput) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *DescribeImagesInput) GetProcessorType() string {
	if m != nil {
		return m.ProcessorType
	}
	return ""
}

func (m *DescribeImagesInput) GetOsFamily() string {
	if m != nil {
		return m.OsFamily
	}
	return ""
}

func (m *DescribeImagesInput) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *DescribeImagesInput) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *DescribeImagesInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeImagesInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeImagesInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeImagesInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeImagesInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeImagesOutput struct {
	Action     string                               `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                                `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                               `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ImageSet   []*DescribeImagesOutput_ResponseItem `protobuf:"bytes,4,rep,name=image_set,json=imageSet" json:"image_set,omitempty"`
	TotalCount int32                                `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeImagesOutput) Reset()                    { *m = DescribeImagesOutput{} }
func (m *DescribeImagesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeImagesOutput) ProtoMessage()               {}
func (*DescribeImagesOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *DescribeImagesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeImagesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeImagesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeImagesOutput) GetImageSet() []*DescribeImagesOutput_ResponseItem {
	if m != nil {
		return m.ImageSet
	}
	return nil
}

func (m *DescribeImagesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeImagesOutput_ResponseItem struct {
	ImageId          string                      `protobuf:"bytes,1,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	ImageName        string                      `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	Description      string                      `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Size             int32                       `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	ProcessorType    string                      `protobuf:"bytes,5,opt,name=processor_type,json=processorType" json:"processor_type,omitempty"`
	Platform         string                      `protobuf:"bytes,6,opt,name=platform" json:"platform,omitempty"`
	OsFamily         string                      `protobuf:"bytes,7,opt,name=os_family,json=osFamily" json:"os_family,omitempty"`
	Visibility       string                      `protobuf:"bytes,8,opt,name=visibility" json:"visibility,omitempty"`
	Provider         string                      `protobuf:"bytes,9,opt,name=provider" json:"provider,omitempty"`
	Owner            string                      `protobuf:"bytes,10,opt,name=owner" json:"owner,omitempty"`
	RecommendedType  string                      `protobuf:"bytes,11,opt,name=recommended_type,json=recommendedType" json:"recommended_type,omitempty"`
	Status           string                      `protobuf:"bytes,12,opt,name=status" json:"status,omitempty"`
	TransitionStatus string                      `protobuf:"bytes,13,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	CreateTime       *google_protobuf1.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime       *google_protobuf1.Timestamp `protobuf:"bytes,15,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *DescribeImagesOutput_ResponseItem) Reset()         { *m = DescribeImagesOutput_ResponseItem{} }
func (m *DescribeImagesOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeImagesOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeImagesOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{2, 0}
}

func (m *DescribeImagesOutput_ResponseItem) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DescribeImagesOutput_ResponseItem) GetProcessorType() string {
	if m != nil {
		return m.ProcessorType
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetOsFamily() string {
	if m != nil {
		return m.OsFamily
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetRecommendedType() string {
	if m != nil {
		return m.RecommendedType
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetTransitionStatus() string {
	if m != nil {
		return m.TransitionStatus
	}
	return ""
}

func (m *DescribeImagesOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeImagesOutput_ResponseItem) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type CaptureInstanceInput struct {
	Instance   string `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	ImageName  string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	TargetUser string `protobuf:"bytes,3,opt,name=target_user,json=targetUser" json:"target_user,omitempty"`
}

func (m *CaptureInstanceInput) Reset()                    { *m = CaptureInstanceInput{} }
func (m *CaptureInstanceInput) String() string            { return proto.CompactTextString(m) }
func (*CaptureInstanceInput) ProtoMessage()               {}
func (*CaptureInstanceInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *CaptureInstanceInput) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *CaptureInstanceInput) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *CaptureInstanceInput) GetTargetUser() string {
	if m != nil {
		return m.TargetUser
	}
	return ""
}

type CaptureInstanceOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	ImageId string `protobuf:"bytes,5,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *CaptureInstanceOutput) Reset()                    { *m = CaptureInstanceOutput{} }
func (m *CaptureInstanceOutput) String() string            { return proto.CompactTextString(m) }
func (*CaptureInstanceOutput) ProtoMessage()               {}
func (*CaptureInstanceOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *CaptureInstanceOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CaptureInstanceOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CaptureInstanceOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CaptureInstanceOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *CaptureInstanceOutput) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

type DeleteImagesInput struct {
	Images []string `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
}

func (m *DeleteImagesInput) Reset()                    { *m = DeleteImagesInput{} }
func (m *DeleteImagesInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteImagesInput) ProtoMessage()               {}
func (*DeleteImagesInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *DeleteImagesInput) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

type DeleteImagesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DeleteImagesOutput) Reset()                    { *m = DeleteImagesOutput{} }
func (m *DeleteImagesOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteImagesOutput) ProtoMessage()               {}
func (*DeleteImagesOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *DeleteImagesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteImagesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteImagesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteImagesOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ModifyImageAttributesInput struct {
	Image       string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	ImageName   string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ModifyImageAttributesInput) Reset()                    { *m = ModifyImageAttributesInput{} }
func (m *ModifyImageAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyImageAttributesInput) ProtoMessage()               {}
func (*ModifyImageAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *ModifyImageAttributesInput) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ModifyImageAttributesInput) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ModifyImageAttributesInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ModifyImageAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifyImageAttributesOutput) Reset()                    { *m = ModifyImageAttributesOutput{} }
func (m *ModifyImageAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyImageAttributesOutput) ProtoMessage()               {}
func (*ModifyImageAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ModifyImageAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifyImageAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifyImageAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GrantImageToUsersInput struct {
	Image string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Users []string `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *GrantImageToUsersInput) Reset()                    { *m = GrantImageToUsersInput{} }
func (m *GrantImageToUsersInput) String() string            { return proto.CompactTextString(m) }
func (*GrantImageToUsersInput) ProtoMessage()               {}
func (*GrantImageToUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *GrantImageToUsersInput) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *GrantImageToUsersInput) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type GrantImageToUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GrantImageToUsersOutput) Reset()                    { *m = GrantImageToUsersOutput{} }
func (m *GrantImageToUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*GrantImageToUsersOutput) ProtoMessage()               {}
func (*GrantImageToUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *GrantImageToUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GrantImageToUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GrantImageToUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RevokeImageFromUsersInput struct {
	Image string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Users []string `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *RevokeImageFromUsersInput) Reset()                    { *m = RevokeImageFromUsersInput{} }
func (m *RevokeImageFromUsersInput) String() string            { return proto.CompactTextString(m) }
func (*RevokeImageFromUsersInput) ProtoMessage()               {}
func (*RevokeImageFromUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

func (m *RevokeImageFromUsersInput) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *RevokeImageFromUsersInput) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type RevokeImageFromUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RevokeImageFromUsersOutput) Reset()                    { *m = RevokeImageFromUsersOutput{} }
func (m *RevokeImageFromUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*RevokeImageFromUsersOutput) ProtoMessage()               {}
func (*RevokeImageFromUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

func (m *RevokeImageFromUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RevokeImageFromUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RevokeImageFromUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DescribeImageUsersInput struct {
	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	Offset  int32  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Limit   int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeImageUsersInput) Reset()                    { *m = DescribeImageUsersInput{} }
func (m *DescribeImageUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeImageUsersInput) ProtoMessage()               {}
func (*DescribeImageUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *DescribeImageUsersInput) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *DescribeImageUsersInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeImageUsersInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeImageUsersOutput struct {
	Action       string                                   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode      int32                                    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message      string                                   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ImageUserSet []*DescribeImageUsersOutput_ResponseItem `protobuf:"bytes,4,rep,name=image_user_set,json=imageUserSet" json:"image_user_set,omitempty"`
	TotalCount   int32                                    `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeImageUsersOutput) Reset()                    { *m = DescribeImageUsersOutput{} }
func (m *DescribeImageUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeImageUsersOutput) ProtoMessage()               {}
func (*DescribeImageUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{14} }

func (m *DescribeImageUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeImageUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeImageUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeImageUsersOutput) GetImageUserSet() []*DescribeImageUsersOutput_ResponseItem {
	if m != nil {
		return m.ImageUserSet
	}
	return nil
}

func (m *DescribeImageUsersOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeImageUsersOutput_ResponseItem struct {
	ImageId    string                      `protobuf:"bytes,1,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	User       *User                       `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	CreateTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
}

func (m *DescribeImageUsersOutput_ResponseItem) Reset()         { *m = DescribeImageUsersOutput_ResponseItem{} }
func (m *DescribeImageUsersOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeImageUsersOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeImageUsersOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{14, 0}
}

func (m *DescribeImageUsersOutput_ResponseItem) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *DescribeImageUsersOutput_ResponseItem) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *DescribeImageUsersOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type CloneImagesInput struct {
	Image     string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	ImageName string `protobuf:"bytes,3,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
}

func (m *CloneImagesInput) Reset()                    { *m = CloneImagesInput{} }
func (m *CloneImagesInput) String() string            { return proto.CompactTextString(m) }
func (*CloneImagesInput) ProtoMessage()               {}
func (*CloneImagesInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{15} }

func (m *CloneImagesInput) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *CloneImagesInput) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CloneImagesInput) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

type CloneImagesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CloneImagesOutput) Reset()                    { *m = CloneImagesOutput{} }
func (m *CloneImagesOutput) String() string            { return proto.CompactTextString(m) }
func (*CloneImagesOutput) ProtoMessage()               {}
func (*CloneImagesOutput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{16} }

func (m *CloneImagesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CloneImagesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CloneImagesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageServiceProperties)(nil), "service.ImageServiceProperties")
	proto.RegisterType((*DescribeImagesInput)(nil), "service.DescribeImagesInput")
	proto.RegisterType((*DescribeImagesOutput)(nil), "service.DescribeImagesOutput")
	proto.RegisterType((*DescribeImagesOutput_ResponseItem)(nil), "service.DescribeImagesOutput.ResponseItem")
	proto.RegisterType((*CaptureInstanceInput)(nil), "service.CaptureInstanceInput")
	proto.RegisterType((*CaptureInstanceOutput)(nil), "service.CaptureInstanceOutput")
	proto.RegisterType((*DeleteImagesInput)(nil), "service.DeleteImagesInput")
	proto.RegisterType((*DeleteImagesOutput)(nil), "service.DeleteImagesOutput")
	proto.RegisterType((*ModifyImageAttributesInput)(nil), "service.ModifyImageAttributesInput")
	proto.RegisterType((*ModifyImageAttributesOutput)(nil), "service.ModifyImageAttributesOutput")
	proto.RegisterType((*GrantImageToUsersInput)(nil), "service.GrantImageToUsersInput")
	proto.RegisterType((*GrantImageToUsersOutput)(nil), "service.GrantImageToUsersOutput")
	proto.RegisterType((*RevokeImageFromUsersInput)(nil), "service.RevokeImageFromUsersInput")
	proto.RegisterType((*RevokeImageFromUsersOutput)(nil), "service.RevokeImageFromUsersOutput")
	proto.RegisterType((*DescribeImageUsersInput)(nil), "service.DescribeImageUsersInput")
	proto.RegisterType((*DescribeImageUsersOutput)(nil), "service.DescribeImageUsersOutput")
	proto.RegisterType((*DescribeImageUsersOutput_ResponseItem)(nil), "service.DescribeImageUsersOutput.ResponseItem")
	proto.RegisterType((*CloneImagesInput)(nil), "service.CloneImagesInput")
	proto.RegisterType((*CloneImagesOutput)(nil), "service.CloneImagesOutput")
}

func init() { proto.RegisterFile("image.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 1126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcd, 0x72, 0xe3, 0xc4,
	0x13, 0x2f, 0x7f, 0x28, 0xb6, 0xdb, 0xd9, 0x7c, 0xcc, 0xdf, 0xc9, 0x2a, 0xca, 0x3f, 0x1b, 0xaf,
	0x16, 0xaa, 0x02, 0x4b, 0xd9, 0x55, 0xe1, 0x06, 0x27, 0x48, 0x2a, 0x29, 0x1f, 0x16, 0x16, 0x6d,
	0x80, 0xe2, 0x40, 0x09, 0x59, 0x6a, 0x3b, 0x93, 0xb5, 0x34, 0x62, 0x66, 0x9c, 0xdd, 0xec, 0x1b,
	0xc0, 0x03, 0x70, 0xe0, 0x21, 0x78, 0x30, 0x6e, 0xbc, 0x00, 0x45, 0xcd, 0x8c, 0xec, 0xc8, 0x8e,
	0xec, 0x04, 0x28, 0x73, 0x49, 0xb9, 0x7b, 0xfa, 0xe3, 0x37, 0x3d, 0xbf, 0xee, 0x56, 0xa0, 0x49,
	0xe3, 0x60, 0x88, 0x9d, 0x94, 0x33, 0xc9, 0x48, 0x4d, 0x20, 0xbf, 0xa6, 0x21, 0x3a, 0x4d, 0x79,
	0x93, 0xa2, 0x30, 0x5a, 0xe7, 0xe0, 0x47, 0x9a, 0x0c, 0xc3, 0x11, 0x1b, 0x47, 0xbe, 0x88, 0x5e,
	0xfb, 0x7c, 0x3c, 0xc2, 0xae, 0xfa, 0x93, 0x1d, 0x1f, 0x0e, 0x19, 0x1b, 0x8e, 0xb0, 0xab, 0xa5,
	0xfe, 0x78, 0xd0, 0x95, 0x34, 0x46, 0x21, 0x83, 0x38, 0x35, 0x06, 0xee, 0x47, 0xb0, 0xdb, 0x53,
	0x49, 0x5e, 0x99, 0xe0, 0x2f, 0x39, 0x4b, 0x91, 0x4b, 0x8a, 0x82, 0x10, 0xa8, 0xbe, 0x63, 0x09,
	0xda, 0xa5, 0x76, 0xe9, 0xa8, 0xe1, 0xe9, 0xdf, 0xee, 0x6f, 0x65, 0xf8, 0xdf, 0x29, 0x8a, 0x90,
	0xd3, 0x3e, 0x6a, 0x37, 0xd1, 0x4b, 0xd2, 0xb1, 0x24, 0xbb, 0xb0, 0xa6, 0xa1, 0x0a, 0xbb, 0xd4,
	0xae, 0x1c, 0x35, 0xbc, 0x4c, 0x22, 0xef, 0xc3, 0x46, 0xca, 0x59, 0x88, 0x42, 0x30, 0xee, 0x2b,
	0xd8, 0x76, 0x59, 0x47, 0x7b, 0x34, 0xd5, 0x5e, 0xdc, 0xa4, 0x48, 0xf6, 0xa1, 0xc1, 0x84, 0x3f,
	0x08, 0x62, 0x3a, 0xba, 0xb1, 0x2b, 0xda, 0xa2, 0xce, 0xc4, 0x99, 0x96, 0xc9, 0x13, 0x80, 0x6b,
	0x2a, 0x68, 0x9f, 0x8e, 0xa8, 0xbc, 0xb1, 0xab, 0xfa, 0x34, 0xa7, 0x21, 0x0e, 0xd4, 0x53, 0xce,
	0xae, 0x69, 0x84, 0xdc, 0xb6, 0x8c, 0xef, 0x44, 0x56, 0xb8, 0x84, 0x0c, 0xe4, 0x58, 0xd8, 0x6b,
	0x06, 0x97, 0x91, 0xc8, 0x21, 0x34, 0x05, 0x06, 0x3c, 0xbc, 0xf4, 0xdf, 0x30, 0x1e, 0xd9, 0x35,
	0x13, 0xd4, 0xa8, 0xbe, 0x65, 0x3c, 0x22, 0x36, 0xd4, 0xae, 0x91, 0xf7, 0x99, 0x40, 0xbb, 0xde,
	0x2e, 0x1d, 0x59, 0xde, 0x44, 0x54, 0x21, 0xd9, 0x60, 0x20, 0x50, 0xda, 0x0d, 0x7d, 0x90, 0x49,
	0xa4, 0x05, 0xd6, 0x88, 0xc6, 0x54, 0xda, 0xa0, 0xd5, 0x46, 0x70, 0x7f, 0xb7, 0xa0, 0x35, 0x5b,
	0xb0, 0x2f, 0xc7, 0x32, 0xab, 0x58, 0x10, 0x4a, 0xca, 0x92, 0xac, 0xbe, 0x99, 0x44, 0xf6, 0xa0,
	0xce, 0x51, 0xfa, 0x21, 0x8b, 0x4c, 0xad, 0x2c, 0xaf, 0xc6, 0x51, 0x9e, 0xb0, 0x08, 0x15, 0xa6,
	0x18, 0x85, 0x08, 0x86, 0x98, 0xd5, 0x68, 0x22, 0x92, 0x73, 0x68, 0xe8, 0x82, 0xfb, 0x0a, 0x56,
	0xb5, 0x5d, 0x39, 0x6a, 0x1e, 0x7f, 0xd8, 0xc9, 0xe8, 0xd2, 0x29, 0x4a, 0xdf, 0xf1, 0x50, 0xa4,
	0x2c, 0x11, 0xd8, 0x93, 0x18, 0x7b, 0x75, 0x6a, 0x18, 0x20, 0x55, 0x5d, 0x24, 0x93, 0xc1, 0xc8,
	0x0f, 0xd9, 0x38, 0x91, 0xba, 0x9c, 0x96, 0x07, 0x5a, 0x75, 0xa2, 0x34, 0xce, 0xaf, 0x55, 0x58,
	0xcf, 0xfb, 0x2a, 0xbc, 0x26, 0x35, 0x8d, 0xb2, 0x9b, 0xd4, 0xb4, 0xdc, 0x8b, 0xc8, 0x01, 0x80,
	0x39, 0x4a, 0x82, 0x78, 0xf2, 0xf0, 0x06, 0xe7, 0x17, 0x41, 0x8c, 0xa4, 0x0d, 0xcd, 0x48, 0x43,
	0x4b, 0x75, 0x19, 0xcc, 0x95, 0xf2, 0x2a, 0xc5, 0x40, 0x41, 0xdf, 0xa1, 0x7e, 0x73, 0xcb, 0xd3,
	0xbf, 0x0b, 0x18, 0x65, 0x15, 0x31, 0x4a, 0x91, 0x62, 0x14, 0xc8, 0x01, 0xe3, 0xb1, 0xbd, 0x96,
	0x91, 0x22, 0x93, 0x67, 0xd9, 0x56, 0x5b, 0xca, 0xb6, 0xfa, 0x52, 0xb6, 0x35, 0xe6, 0xd8, 0xd6,
	0x02, 0x8b, 0xbd, 0x49, 0x90, 0x6b, 0x0a, 0x34, 0x3c, 0x23, 0x90, 0x0f, 0x60, 0x8b, 0x63, 0xc8,
	0xe2, 0x18, 0x93, 0x08, 0x23, 0x83, 0xb9, 0xa9, 0x0d, 0x36, 0x73, 0x7a, 0x8d, 0xfa, 0x96, 0xae,
	0xeb, 0x86, 0x14, 0x19, 0x5d, 0x9f, 0xc3, 0xb6, 0xe4, 0x41, 0x22, 0xa8, 0x2a, 0x8b, 0x9f, 0x99,
	0x3c, 0xd2, 0x26, 0x5b, 0xb7, 0x07, 0xaf, 0x8c, 0xf1, 0xa7, 0xd0, 0x0c, 0x39, 0x06, 0x12, 0x7d,
	0xd5, 0xeb, 0xf6, 0x46, 0xbb, 0x74, 0xd4, 0x3c, 0x76, 0x3a, 0x66, 0x10, 0x74, 0x26, 0x83, 0xa0,
	0x73, 0x31, 0x19, 0x04, 0x1e, 0x18, 0x73, 0xa5, 0x50, 0xce, 0x26, 0xbc, 0x71, 0xde, 0xbc, 0xdf,
	0xd9, 0x98, 0x2b, 0x85, 0xcb, 0xa1, 0x75, 0x12, 0xa4, 0x72, 0xcc, 0xb1, 0x97, 0x08, 0x19, 0x24,
	0x21, 0x9a, 0xe9, 0xe0, 0x40, 0x9d, 0x66, 0x8a, 0x8c, 0x23, 0x53, 0xf9, 0x3e, 0x92, 0x28, 0x42,
	0x06, 0x7c, 0x88, 0xd2, 0x1f, 0x0b, 0xe4, 0x19, 0x49, 0xc0, 0xa8, 0xbe, 0x16, 0xc8, 0xdd, 0x5f,
	0x4a, 0xb0, 0x33, 0x97, 0x74, 0x15, 0x1d, 0xb6, 0x03, 0x6b, 0x57, 0xac, 0xaf, 0x48, 0x6e, 0x06,
	0x90, 0x75, 0xc5, 0xfa, 0xbd, 0x68, 0x86, 0xfd, 0xd6, 0x0c, 0xfb, 0xdd, 0xe7, 0xb0, 0x7d, 0x8a,
	0x23, 0x94, 0x0f, 0x99, 0x93, 0xee, 0x5b, 0x20, 0x79, 0xe3, 0xff, 0xee, 0x06, 0xae, 0x00, 0xe7,
	0x05, 0x8b, 0xe8, 0xe0, 0x46, 0x67, 0xfe, 0x4c, 0x4a, 0x4e, 0xfb, 0x63, 0x39, 0xc1, 0xdb, 0x02,
	0x4b, 0x23, 0xcc, 0x00, 0x18, 0xe1, 0x5f, 0x37, 0xb6, 0x7b, 0x05, 0xfb, 0x85, 0x49, 0x57, 0x70,
	0x6f, 0xf7, 0x14, 0x76, 0xcf, 0x79, 0x90, 0x48, 0x9d, 0xea, 0x82, 0x29, 0xd2, 0x2c, 0xbd, 0x5c,
	0x0b, 0x2c, 0x45, 0x35, 0x61, 0x97, 0xf5, 0x0b, 0x19, 0xc1, 0x1d, 0xc0, 0xe3, 0x3b, 0x51, 0x56,
	0x81, 0xf6, 0x1c, 0xf6, 0x3c, 0xbc, 0x66, 0xaf, 0x0d, 0x11, 0xce, 0x38, 0x8b, 0xff, 0x21, 0x60,
	0x0a, 0x4e, 0x51, 0xa0, 0x55, 0x60, 0xee, 0xc3, 0xe3, 0x99, 0x1d, 0x93, 0x43, 0xbc, 0x64, 0x3b,
	0xdc, 0xee, 0xd1, 0x72, 0xf1, 0x1e, 0xad, 0xcc, 0xec, 0xd1, 0x32, 0xd8, 0x77, 0x93, 0xac, 0xa2,
	0x4f, 0x2e, 0x60, 0xc3, 0x40, 0x56, 0x75, 0xcc, 0x2d, 0xd4, 0x4e, 0xf1, 0x42, 0xcd, 0xe1, 0x98,
	0x5d, 0xaa, 0xeb, 0x74, 0x72, 0xfc, 0xa0, 0xc5, 0xfa, 0x53, 0xe9, 0xe1, 0x8b, 0xf5, 0x29, 0x54,
	0xf5, 0x34, 0x2c, 0xeb, 0xe9, 0xfc, 0x68, 0x0a, 0x4c, 0x25, 0xf3, 0xf4, 0xd1, 0xfc, 0x12, 0xa8,
	0xfc, 0x9d, 0x25, 0xe0, 0x7e, 0x0f, 0x5b, 0x27, 0x23, 0x96, 0xcc, 0x4c, 0xae, 0x85, 0xdc, 0x33,
	0x17, 0x32, 0xe5, 0x35, 0xc2, 0xdc, 0x7c, 0xa8, 0xcc, 0xcd, 0x07, 0xf7, 0x07, 0xd8, 0xce, 0x85,
	0x5f, 0xc1, 0x1b, 0x1e, 0xff, 0x69, 0xc1, 0x7a, 0xfe, 0xab, 0x96, 0xbc, 0x80, 0x8d, 0xd9, 0xcf,
	0x20, 0xf2, 0xff, 0x05, 0xdf, 0x47, 0xfa, 0xb6, 0xce, 0xc1, 0xd2, 0xaf, 0x27, 0xf2, 0x12, 0x36,
	0xe7, 0x76, 0x0e, 0xb9, 0xf5, 0x28, 0x5a, 0x81, 0xce, 0x93, 0x45, 0xc7, 0x59, 0xc4, 0x73, 0x58,
	0xcf, 0x2f, 0x00, 0xe2, 0xe4, 0x00, 0xcc, 0x2d, 0x11, 0x67, 0xbf, 0xf0, 0x2c, 0x0b, 0xd4, 0x87,
	0x9d, 0xc2, 0xd1, 0x4a, 0x9e, 0x4d, 0xbd, 0x16, 0xcf, 0x7b, 0xe7, 0xbd, 0xe5, 0x46, 0x59, 0x8e,
	0x6f, 0x60, 0xfb, 0xce, 0x30, 0x24, 0x87, 0x53, 0xd7, 0xe2, 0x71, 0xeb, 0xb4, 0x17, 0x1b, 0x64,
	0x71, 0x7d, 0x68, 0x15, 0xcd, 0x2c, 0xe2, 0x4e, 0x3d, 0x17, 0xce, 0x46, 0xe7, 0xd9, 0x52, 0x9b,
	0x2c, 0xc1, 0x77, 0x6a, 0xcd, 0xce, 0x37, 0x2f, 0x69, 0x2f, 0xe9, 0x6c, 0x13, 0xfc, 0xe9, 0xbd,
	0xbd, 0x4f, 0x4e, 0xa1, 0x99, 0x23, 0x35, 0xd9, 0xbb, 0x7d, 0xef, 0xb9, 0x4e, 0x72, 0x9c, 0xa2,
	0x23, 0x13, 0xc5, 0x39, 0xfb, 0xf9, 0x8f, 0xea, 0xe7, 0xd0, 0xbd, 0x94, 0x32, 0x15, 0x9f, 0x74,
	0xbb, 0x11, 0x0b, 0x45, 0x67, 0xfa, 0x0f, 0x5e, 0x27, 0x64, 0x71, 0x37, 0x48, 0x69, 0x57, 0xb7,
	0x52, 0x97, 0x26, 0x11, 0xbe, 0xed, 0x5c, 0xca, 0x78, 0x44, 0x36, 0xbf, 0xa2, 0xc9, 0xf0, 0x44,
	0xdb, 0xe8, 0x70, 0xfd, 0x35, 0xdd, 0xe1, 0x1f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x89,
	0x2b, 0x77, 0x41, 0x0e, 0x00, 0x00,
}