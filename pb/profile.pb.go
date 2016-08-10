// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle           string                   `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name             string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location         string                   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About            string                   `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription string                   `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website          string                   `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email            string                   `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber      string                   `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social           []*Profile_SocialAccount `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	Nsfw             bool                     `protobuf:"varint,10,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor           bool                     `protobuf:"varint,11,opt,name=vendor" json:"vendor,omitempty"`
	Moderator        bool                     `protobuf:"varint,12,opt,name=moderator" json:"moderator,omitempty"`
	PrimaryColor     string                   `protobuf:"bytes,13,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor   string                   `protobuf:"bytes,14,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor        string                   `protobuf:"bytes,15,opt,name=textColor" json:"textColor,omitempty"`
	AvatarHash       string                   `protobuf:"bytes,16,opt,name=avatarHash" json:"avatarHash,omitempty"`
	HeaderHash       string                   `protobuf:"bytes,17,opt,name=headerHash" json:"headerHash,omitempty"`
	FollowerCount    uint32                   `protobuf:"varint,18,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount   uint32                   `protobuf:"varint,19,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount     uint32                   `protobuf:"varint,20,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified     uint32                   `protobuf:"varint,21,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
}

var fileDescriptor3 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x35, 0xb6, 0xfe, 0x3b, 0x5d, 0xc6, 0x30, 0x63, 0xb2, 0x26, 0x84, 0xaa, 0x09, 0xa1,
	0x89, 0x8b, 0x5e, 0xc0, 0x13, 0xa0, 0x71, 0xc1, 0x0d, 0x08, 0x15, 0xf1, 0x00, 0x6e, 0x72, 0x42,
	0x2c, 0x39, 0x3e, 0x91, 0xed, 0xae, 0xec, 0xc5, 0xb9, 0x26, 0x3e, 0xce, 0xd2, 0xa4, 0xbb, 0xaa,
	0xbf, 0xdf, 0xf7, 0xab, 0x63, 0x27, 0x07, 0xb2, 0xc6, 0x51, 0xa9, 0x0d, 0xae, 0xdb, 0xdf, 0x40,
	0xb7, 0xff, 0x26, 0x30, 0xfb, 0x99, 0x88, 0xb8, 0x86, 0x69, 0xa5, 0x6c, 0x61, 0x50, 0x9e, 0xac,
	0x4e, 0xee, 0x16, 0x9b, 0x2e, 0x09, 0x01, 0x67, 0x56, 0xd5, 0x28, 0x5f, 0x30, 0xe5, 0xb5, 0xb8,
	0x81, 0xb9, 0xa1, 0x5c, 0x05, 0x4d, 0x56, 0x9e, 0x32, 0xef, 0xb3, 0xb8, 0x82, 0x89, 0xda, 0xd2,
	0x2e, 0xc8, 0x33, 0x2e, 0x52, 0x10, 0x1f, 0xe1, 0xd2, 0x57, 0xe4, 0xc2, 0x57, 0xf4, 0xb9, 0xd3,
	0x0d, 0xff, 0x73, 0xc2, 0xc2, 0x33, 0x2e, 0x24, 0xcc, 0xf6, 0xb8, 0xf5, 0x3a, 0xa0, 0x9c, 0xb2,
	0xf2, 0x14, 0xe3, 0xde, 0x58, 0x2b, 0x6d, 0xe4, 0x2c, 0xed, 0xcd, 0x41, 0xac, 0x60, 0xd9, 0x54,
	0x64, 0xf1, 0xc7, 0xae, 0xde, 0xa2, 0x93, 0x73, 0xee, 0x86, 0x48, 0xac, 0x61, 0xea, 0x29, 0xd7,
	0xca, 0xc8, 0xc5, 0xea, 0xf4, 0x6e, 0xf9, 0xe9, 0x7a, 0xdd, 0xdd, 0x7a, 0xfd, 0x8b, 0xf1, 0x97,
	0x3c, 0xa7, 0x9d, 0x0d, 0x9b, 0xce, 0xe2, 0x3b, 0xfb, 0x72, 0x2f, 0xa1, 0xdd, 0x6a, 0xbe, 0xe1,
	0x75, 0x7c, 0x3f, 0x0f, 0x68, 0x0b, 0x72, 0x72, 0xc9, 0xb4, 0x4b, 0xe2, 0x2d, 0x2c, 0x6a, 0x2a,
	0xd0, 0xa9, 0xd0, 0x56, 0xe7, 0x5c, 0x1d, 0x80, 0xb8, 0x85, 0xf3, 0xc6, 0xe9, 0x5a, 0xb9, 0xc7,
	0x7b, 0x32, 0xad, 0x90, 0xf1, 0xe1, 0x46, 0x4c, 0x7c, 0x80, 0x0b, 0x8f, 0x39, 0xd9, 0xa2, 0xb7,
	0x2e, 0xd8, 0x3a, 0xa2, 0xf1, 0x49, 0x01, 0xff, 0x86, 0xa4, 0xbc, 0x64, 0xe5, 0x00, 0xc4, 0x3b,
	0x00, 0xf5, 0xa0, 0x82, 0x72, 0xdf, 0x94, 0xaf, 0xe4, 0x25, 0xd7, 0x03, 0x12, 0xfb, 0x0a, 0x55,
	0x7b, 0x2e, 0xee, 0x5f, 0xa5, 0xfe, 0x40, 0xc4, 0x7b, 0xc8, 0x4a, 0x32, 0x86, 0xf6, 0xe8, 0xee,
	0xe3, 0xcb, 0x90, 0xa2, 0x55, 0xb2, 0xcd, 0x18, 0xc6, 0xb3, 0x26, 0xa0, 0xed, 0x9f, 0xa4, 0xbd,
	0x66, 0xed, 0x88, 0xc6, 0x7b, 0x1b, 0xed, 0x43, 0x6f, 0x5d, 0xb1, 0x35, 0x62, 0xec, 0x28, 0x1f,
	0xbe, 0x53, 0xa1, 0x4b, 0x8d, 0x85, 0x7c, 0xd3, 0x39, 0x03, 0x76, 0xf3, 0x1b, 0xb2, 0xd1, 0x27,
	0x8a, 0x9f, 0x26, 0x3c, 0x36, 0x4f, 0x43, 0xca, 0xeb, 0x38, 0x8e, 0x3b, 0x8f, 0x6e, 0x30, 0xa6,
	0x7d, 0x8e, 0x23, 0xd3, 0xce, 0x3a, 0x95, 0xdd, 0x9c, 0xa6, 0xb0, 0x9d, 0xf2, 0xfc, 0x7f, 0xfe,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x33, 0xed, 0x41, 0x68, 0x10, 0x03, 0x00, 0x00,
}
