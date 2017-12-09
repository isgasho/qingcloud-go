// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: dns.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type DNSAliasServiceInterface interface {
	DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error)
	AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error)
	DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error)
	GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error)
}

type DNSAliasService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewDNSAliasService(server *ServerInfo) (p *DNSAliasService) {
	return &DNSAliasService{
		ServerInfo: server,
	}
}

func (p *DNSAliasService) DescribeDNSAliases(input *DescribeDNSAliasesInput) (output *DescribeDNSAliasesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeDNSAliasesOutput)

	err = client.CallMethod(nil, "DescribeDNSAliases", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *DNSAliasService) AssociateDNSAlias(input *AssociateDNSAliasInput) (output *AssociateDNSAliasOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(AssociateDNSAliasOutput)

	err = client.CallMethod(nil, "AssociateDNSAlias", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *DNSAliasService) DissociateDNSAliases(input *DissociateDNSAliasesInput) (output *DissociateDNSAliasesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DissociateDNSAliasesOutput)

	err = client.CallMethod(nil, "DissociateDNSAliases", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *DNSAliasService) GetDNSLabel(input *GetDNSLabelInput) (output *GetDNSLabelOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(GetDNSLabelOutput)

	err = client.CallMethod(nil, "GetDNSLabel", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
