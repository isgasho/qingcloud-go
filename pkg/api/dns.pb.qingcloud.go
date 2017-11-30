// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: dns.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"
import "github.com/chai2010/qingcloud-go/pkg/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}
var _ = data.Operation{}

type DNSAliasServiceInterface interface {
	DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error)
	AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error)
	DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error)
	GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error)
}

type DNSAliasService struct {
	Config           *config.Config
	Properties       *DNSAliasServiceProperties
	LastResponseBody string
}

func NewDNSAliasService(conf *config.Config, zone string) (p *DNSAliasService) {
	return &DNSAliasService{
		Config:     conf,
		Properties: &DNSAliasServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) DNSAlias(zone string) (*DNSAliasService, error) {
	properties := &DNSAliasServiceProperties{
		Zone: proto.String(zone),
	}

	return &DNSAliasService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *DNSAliasService) DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error) {
	if in == nil {
		in = &DescribeDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeDNSAliases",
		RequestMethod: "GET",
	}

	x := &DescribeDNSAliasesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DNSAliasService) AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error) {
	if in == nil {
		in = &AssociateDNSAliasInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateDNSAlias",
		RequestMethod: "GET",
	}

	x := &AssociateDNSAliasOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DNSAliasService) DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error) {
	if in == nil {
		in = &DissociateDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateDNSAliases",
		RequestMethod: "GET",
	}

	x := &DissociateDNSAliasesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DNSAliasService) GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error) {
	if in == nil {
		in = &GetDNSLabelInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetDNSLabel",
		RequestMethod: "GET",
	}

	x := &GetDNSLabelOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}
