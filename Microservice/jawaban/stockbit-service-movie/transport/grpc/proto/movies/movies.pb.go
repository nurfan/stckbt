// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movies.proto

package movies

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchRequest struct {
	Searchword           string   `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination           int64    `protobuf:"varint,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetSearchword() string {
	if m != nil {
		return m.Searchword
	}
	return ""
}

func (m *SearchRequest) GetPagination() int64 {
	if m != nil {
		return m.Pagination
	}
	return 0
}

type SearchResponse struct {
	Movie                []*Movie    `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetMovie() []*Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *SearchResponse) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type DetailRequest struct {
	ImdbID               string   `protobuf:"bytes,1,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{2}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetImdbID() string {
	if m != nil {
		return m.ImdbID
	}
	return ""
}

type DetailResponse struct {
	Data                 *DetailMovie `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{3}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetData() *DetailMovie {
	if m != nil {
		return m.Data
	}
	return nil
}

type Movie struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year                 string   `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID               string   `protobuf:"bytes,3,opt,name=ImdbID,json=imdb_id,proto3" json:"ImdbID,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster               string   `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{4}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Movie) GetImdbID() string {
	if m != nil {
		return m.ImdbID
	}
	return ""
}

func (m *Movie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Movie) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

type DetailMovie struct {
	Title                string    `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year                 string    `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated                string    `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released             string    `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime              string    `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre                string    `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director             string    `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer               string    `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors               string    `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot                 string    `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language             string    `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country              string    `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards               string    `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster               string    `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings              []*Rating `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore            string    `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating           string    `protobuf:"bytes,17,opt,name=ImdbRating,proto3" json:"ImdbRating,omitempty"`
	ImdbVotes            string    `protobuf:"bytes,18,opt,name=ImdbVotes,proto3" json:"ImdbVotes,omitempty"`
	ImdbID               string    `protobuf:"bytes,19,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type                 string    `protobuf:"bytes,20,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD                  string    `protobuf:"bytes,21,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice            string    `protobuf:"bytes,22,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production           string    `protobuf:"bytes,23,opt,name=Production,proto3" json:"Production,omitempty"`
	Website              string    `protobuf:"bytes,24,opt,name=Website,proto3" json:"Website,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DetailMovie) Reset()         { *m = DetailMovie{} }
func (m *DetailMovie) String() string { return proto.CompactTextString(m) }
func (*DetailMovie) ProtoMessage()    {}
func (*DetailMovie) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{5}
}

func (m *DetailMovie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailMovie.Unmarshal(m, b)
}
func (m *DetailMovie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailMovie.Marshal(b, m, deterministic)
}
func (m *DetailMovie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailMovie.Merge(m, src)
}
func (m *DetailMovie) XXX_Size() int {
	return xxx_messageInfo_DetailMovie.Size(m)
}
func (m *DetailMovie) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailMovie.DiscardUnknown(m)
}

var xxx_messageInfo_DetailMovie proto.InternalMessageInfo

func (m *DetailMovie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetailMovie) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *DetailMovie) GetRated() string {
	if m != nil {
		return m.Rated
	}
	return ""
}

func (m *DetailMovie) GetReleased() string {
	if m != nil {
		return m.Released
	}
	return ""
}

func (m *DetailMovie) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *DetailMovie) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *DetailMovie) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *DetailMovie) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *DetailMovie) GetActors() string {
	if m != nil {
		return m.Actors
	}
	return ""
}

func (m *DetailMovie) GetPlot() string {
	if m != nil {
		return m.Plot
	}
	return ""
}

func (m *DetailMovie) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *DetailMovie) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *DetailMovie) GetAwards() string {
	if m != nil {
		return m.Awards
	}
	return ""
}

func (m *DetailMovie) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *DetailMovie) GetRatings() []*Rating {
	if m != nil {
		return m.Ratings
	}
	return nil
}

func (m *DetailMovie) GetMetascore() string {
	if m != nil {
		return m.Metascore
	}
	return ""
}

func (m *DetailMovie) GetImdbRating() string {
	if m != nil {
		return m.ImdbRating
	}
	return ""
}

func (m *DetailMovie) GetImdbVotes() string {
	if m != nil {
		return m.ImdbVotes
	}
	return ""
}

func (m *DetailMovie) GetImdbID() string {
	if m != nil {
		return m.ImdbID
	}
	return ""
}

func (m *DetailMovie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DetailMovie) GetDVD() string {
	if m != nil {
		return m.DVD
	}
	return ""
}

func (m *DetailMovie) GetBoxOffice() string {
	if m != nil {
		return m.BoxOffice
	}
	return ""
}

func (m *DetailMovie) GetProduction() string {
	if m != nil {
		return m.Production
	}
	return ""
}

func (m *DetailMovie) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

type Rating struct {
	Source               string   `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rating) Reset()         { *m = Rating{} }
func (m *Rating) String() string { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()    {}
func (*Rating) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{6}
}

func (m *Rating) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rating.Unmarshal(m, b)
}
func (m *Rating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rating.Marshal(b, m, deterministic)
}
func (m *Rating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rating.Merge(m, src)
}
func (m *Rating) XXX_Size() int {
	return xxx_messageInfo_Rating.Size(m)
}
func (m *Rating) XXX_DiscardUnknown() {
	xxx_messageInfo_Rating.DiscardUnknown(m)
}

var xxx_messageInfo_Rating proto.InternalMessageInfo

func (m *Rating) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Rating) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Pagination struct {
	PageSize             int64    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	CurrentPage          int64    `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	TotalPage            int64    `protobuf:"varint,3,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	TotalResult          int64    `protobuf:"varint,4,opt,name=total_result,json=totalResult,proto3" json:"total_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{7}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pagination) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *Pagination) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *Pagination) GetTotalResult() int64 {
	if m != nil {
		return m.TotalResult
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "movies.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "movies.SearchResponse")
	proto.RegisterType((*DetailRequest)(nil), "movies.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "movies.DetailResponse")
	proto.RegisterType((*Movie)(nil), "movies.Movie")
	proto.RegisterType((*DetailMovie)(nil), "movies.DetailMovie")
	proto.RegisterType((*Rating)(nil), "movies.Rating")
	proto.RegisterType((*Pagination)(nil), "movies.Pagination")
}

func init() { proto.RegisterFile("movies.proto", fileDescriptor_546d72ade507cae9) }

var fileDescriptor_546d72ade507cae9 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0x1b, 0x31,
	0x10, 0x55, 0x9a, 0x64, 0x43, 0x26, 0x1f, 0xa5, 0x06, 0x82, 0x45, 0x3f, 0x04, 0xe9, 0x81, 0x9c,
	0x38, 0xa4, 0x52, 0x51, 0x8f, 0x6d, 0x23, 0x55, 0x48, 0x45, 0x44, 0x0b, 0x02, 0xf5, 0x84, 0x9c,
	0xdd, 0x21, 0xb5, 0xb4, 0x59, 0xa7, 0xb6, 0x17, 0x0a, 0x3f, 0xa2, 0x7f, 0xb8, 0x97, 0xca, 0x1e,
	0x6f, 0xb2, 0x69, 0x4f, 0x3d, 0xad, 0xdf, 0xf3, 0xf8, 0xcd, 0xb3, 0x77, 0x66, 0xa0, 0xbb, 0x50,
	0xf7, 0x12, 0xcd, 0xc9, 0x52, 0x2b, 0xab, 0x58, 0x44, 0x68, 0x78, 0x01, 0xbd, 0x4b, 0x14, 0x3a,
	0xf9, 0x1e, 0xe3, 0x8f, 0x02, 0x8d, 0x65, 0x6f, 0x00, 0x8c, 0x27, 0x1e, 0x94, 0x4e, 0x79, 0xed,
	0xb0, 0x36, 0x6a, 0xc7, 0x15, 0xc6, 0xed, 0x2f, 0xc5, 0x5c, 0xe6, 0xc2, 0x4a, 0x95, 0xf3, 0x67,
	0x87, 0xb5, 0x51, 0x3d, 0xae, 0x30, 0x43, 0x09, 0xfd, 0x52, 0xd0, 0x2c, 0x55, 0x6e, 0x90, 0xbd,
	0x85, 0xa6, 0x4f, 0xc6, 0x6b, 0x87, 0xf5, 0x51, 0x67, 0xdc, 0x3b, 0x09, 0x46, 0xce, 0xdd, 0x27,
	0xa6, 0x3d, 0x36, 0xfe, 0x47, 0xb6, 0x33, 0x66, 0x65, 0xe4, 0x74, 0xb5, 0xb3, 0x91, 0xea, 0x18,
	0x7a, 0x13, 0xb4, 0x42, 0x66, 0xa5, 0xf7, 0x01, 0x44, 0x67, 0x8b, 0x74, 0x76, 0x36, 0x09, 0xbe,
	0x03, 0x1a, 0x7e, 0x80, 0x7e, 0x19, 0x18, 0x3c, 0x1d, 0x43, 0x23, 0x15, 0x56, 0xf8, 0xb8, 0xce,
	0x78, 0xa7, 0x4c, 0x44, 0x51, 0x64, 0xcc, 0x07, 0x0c, 0xef, 0xa1, 0xe9, 0x21, 0xdb, 0x85, 0xe6,
	0x95, 0xb4, 0x19, 0x06, 0x69, 0x02, 0x8c, 0x41, 0xe3, 0x1b, 0x0a, 0xed, 0x0d, 0xb7, 0x63, 0xbf,
	0x66, 0xfb, 0x2b, 0x17, 0x75, 0xcf, 0xb6, 0xe4, 0x22, 0x9d, 0xdd, 0xca, 0xd4, 0x05, 0x5f, 0x3d,
	0x2e, 0x91, 0x37, 0x28, 0xd8, 0xad, 0x9d, 0xe5, 0xa9, 0x32, 0x16, 0x35, 0x6f, 0x92, 0x65, 0x42,
	0xc3, 0xdf, 0x0d, 0xe8, 0x54, 0xdc, 0xfc, 0x47, 0xfa, 0x5d, 0x68, 0xc6, 0xc2, 0x62, 0x1a, 0xb2,
	0x13, 0x60, 0x07, 0xb0, 0x15, 0x63, 0x86, 0xc2, 0x60, 0x1a, 0xf2, 0xaf, 0x30, 0xe3, 0xd0, 0x8a,
	0x8b, 0xdc, 0xca, 0x05, 0x06, 0x13, 0x25, 0x74, 0x5a, 0x5f, 0x30, 0xd7, 0xc8, 0x23, 0xd2, 0xf2,
	0xc0, 0x69, 0x4d, 0xa4, 0xc6, 0xc4, 0x2a, 0xcd, 0x5b, 0xa4, 0x55, 0x62, 0x77, 0x9f, 0x1b, 0x2d,
	0xdd, 0x7d, 0xb6, 0xe8, 0x3e, 0x84, 0x1c, 0xff, 0xd1, 0x05, 0x18, 0xde, 0x26, 0x9e, 0x90, 0xbb,
	0xc1, 0x34, 0x53, 0x96, 0x03, 0xdd, 0xc0, 0xad, 0x9d, 0xfe, 0x57, 0x91, 0xcf, 0x0b, 0x31, 0x47,
	0xde, 0x21, 0xfd, 0x12, 0x3b, 0xaf, 0x9f, 0x55, 0x91, 0x5b, 0xfd, 0xc8, 0xbb, 0xe4, 0x35, 0x40,
	0x9f, 0xe1, 0x41, 0xe8, 0xd4, 0xf0, 0x5e, 0xc8, 0xe0, 0x51, 0xe5, 0x85, 0xfb, 0xd5, 0x17, 0x66,
	0x23, 0x68, 0xc5, 0xc2, 0xca, 0x7c, 0x6e, 0xf8, 0x73, 0x5f, 0x98, 0xfd, 0xb2, 0x0a, 0x88, 0x8e,
	0xcb, 0x6d, 0xf6, 0x0a, 0xda, 0xe7, 0x68, 0x85, 0x49, 0x94, 0x46, 0xbe, 0xed, 0x45, 0xd6, 0x84,
	0x6b, 0x08, 0xf7, 0xbb, 0x29, 0x98, 0xbf, 0xa0, 0x86, 0x59, 0x33, 0xee, 0xb4, 0x43, 0xd7, 0xca,
	0xa2, 0xe1, 0x8c, 0x4e, 0xaf, 0x88, 0x4a, 0xc9, 0xee, 0x54, 0x4b, 0x76, 0x55, 0x2b, 0xbb, 0x95,
	0x5a, 0xd9, 0x86, 0xfa, 0xe4, 0x7a, 0xc2, 0xf7, 0x3c, 0xe5, 0x96, 0x4e, 0xfb, 0x93, 0xfa, 0x79,
	0x71, 0x77, 0x27, 0x13, 0xe4, 0x03, 0xd2, 0x5e, 0x11, 0xce, 0xd9, 0x54, 0xab, 0xb4, 0x48, 0x7c,
	0x4f, 0xed, 0x93, 0xb3, 0x35, 0xe3, 0xde, 0xf2, 0x06, 0x67, 0x46, 0x5a, 0xe4, 0x9c, 0xde, 0x32,
	0xc0, 0xe1, 0x7b, 0x88, 0x82, 0xfb, 0x01, 0x44, 0x97, 0xaa, 0xd0, 0x49, 0x59, 0x78, 0x01, 0xb9,
	0xca, 0xb8, 0x16, 0x59, 0x81, 0xa1, 0xf4, 0x08, 0x0c, 0x7f, 0xd5, 0x00, 0xd6, 0xcd, 0xca, 0x5e,
	0x42, 0x7b, 0x29, 0xe6, 0x78, 0x6b, 0xe4, 0x13, 0x9d, 0xaf, 0xc7, 0x5b, 0x8e, 0xb8, 0x94, 0x4f,
	0xc8, 0x8e, 0xa0, 0x9b, 0x14, 0x5a, 0x63, 0x6e, 0x6f, 0x1d, 0x17, 0x46, 0x49, 0x27, 0x70, 0x53,
	0xf7, 0xb3, 0x5f, 0x03, 0x58, 0x65, 0x45, 0x46, 0x01, 0x75, 0x1f, 0xd0, 0xf6, 0x8c, 0xdf, 0x3e,
	0x82, 0x2e, 0x6d, 0x6b, 0x34, 0x45, 0x66, 0x7d, 0x5d, 0xd7, 0xe3, 0x8e, 0xe7, 0x62, 0x4f, 0x8d,
	0x9f, 0x20, 0xf2, 0xfd, 0x63, 0xd8, 0x29, 0x44, 0x34, 0x97, 0xd8, 0x5e, 0xf9, 0x9f, 0x37, 0x06,
	0xdf, 0xc1, 0xe0, 0x6f, 0x3a, 0x8c, 0x8a, 0x53, 0x88, 0xa8, 0x11, 0xd7, 0x07, 0x37, 0xa6, 0xce,
	0xfa, 0xe0, 0xe6, 0x8c, 0x99, 0x45, 0x7e, 0xd2, 0xbe, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x51,
	0xe7, 0x03, 0xc5, 0x79, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoviesClient is the client API for Movies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoviesClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type moviesClient struct {
	cc *grpc.ClientConn
}

func NewMoviesClient(cc *grpc.ClientConn) MoviesClient {
	return &moviesClient{cc}
}

func (c *moviesClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/movies.Movies/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, "/movies.Movies/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServer is the server API for Movies service.
type MoviesServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
}

// UnimplementedMoviesServer can be embedded to have forward compatible implementations.
type UnimplementedMoviesServer struct {
}

func (*UnimplementedMoviesServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedMoviesServer) Detail(ctx context.Context, req *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}

func RegisterMoviesServer(s *grpc.Server, srv MoviesServer) {
	s.RegisterService(&_Movies_serviceDesc, srv)
}

func _Movies_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movies.Movies/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movies.Movies/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Movies_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movies.Movies",
	HandlerType: (*MoviesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Movies_Search_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Movies_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movies.proto",
}
