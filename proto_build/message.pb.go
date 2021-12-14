// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: message.proto

package proto_build

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//消息类型
type MESSAGE_TYPE int32

const (
	MESSAGE_TYPE_TEXT     MESSAGE_TYPE = 0 //文本消息。
	MESSAGE_TYPE_LOCATION MESSAGE_TYPE = 1 //地理位置消息。
	MESSAGE_TYPE_FACE     MESSAGE_TYPE = 2 // 表情消息。
	MESSAGE_TYPE_SOUND    MESSAGE_TYPE = 3 //  语音消息。
	MESSAGE_TYPE_IMAGE    MESSAGE_TYPE = 4 //  图像消息。
	MESSAGE_TYPE_FILE     MESSAGE_TYPE = 5 //  文件消息。
	MESSAGE_TYPE_Video    MESSAGE_TYPE = 6 //  视频消息。
)

// Enum value maps for MESSAGE_TYPE.
var (
	MESSAGE_TYPE_name = map[int32]string{
		0: "TEXT",
		1: "LOCATION",
		2: "FACE",
		3: "SOUND",
		4: "IMAGE",
		5: "FILE",
		6: "Video",
	}
	MESSAGE_TYPE_value = map[string]int32{
		"TEXT":     0,
		"LOCATION": 1,
		"FACE":     2,
		"SOUND":    3,
		"IMAGE":    4,
		"FILE":     5,
		"Video":    6,
	}
)

func (x MESSAGE_TYPE) Enum() *MESSAGE_TYPE {
	p := new(MESSAGE_TYPE)
	*p = x
	return p
}

func (x MESSAGE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MESSAGE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MESSAGE_TYPE) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MESSAGE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MESSAGE_TYPE.Descriptor instead.
func (MESSAGE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

//图片类型
type IMAGE_TYPE int32

const (
	IMAGE_TYPE_ORIGIN    IMAGE_TYPE = 0 //原图
	IMAGE_TYPE_BIG       IMAGE_TYPE = 1 //大图
	IMAGE_TYPE_THUMBNAIL IMAGE_TYPE = 2 // 缩略图
)

// Enum value maps for IMAGE_TYPE.
var (
	IMAGE_TYPE_name = map[int32]string{
		0: "ORIGIN",
		1: "BIG",
		2: "THUMBNAIL",
	}
	IMAGE_TYPE_value = map[string]int32{
		"ORIGIN":    0,
		"BIG":       1,
		"THUMBNAIL": 2,
	}
)

func (x IMAGE_TYPE) Enum() *IMAGE_TYPE {
	p := new(IMAGE_TYPE)
	*p = x
	return p
}

func (x IMAGE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMAGE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (IMAGE_TYPE) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x IMAGE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMAGE_TYPE.Descriptor instead.
func (IMAGE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

//图片格式
type IMAGE_FORMAT int32

const (
	IMAGE_FORMAT_JPG   IMAGE_FORMAT = 0
	IMAGE_FORMAT_GIF   IMAGE_FORMAT = 1
	IMAGE_FORMAT_PNG   IMAGE_FORMAT = 2
	IMAGE_FORMAT_BMP   IMAGE_FORMAT = 3
	IMAGE_FORMAT_OTHER IMAGE_FORMAT = 255
)

// Enum value maps for IMAGE_FORMAT.
var (
	IMAGE_FORMAT_name = map[int32]string{
		0:   "JPG",
		1:   "GIF",
		2:   "PNG",
		3:   "BMP",
		255: "OTHER",
	}
	IMAGE_FORMAT_value = map[string]int32{
		"JPG":   0,
		"GIF":   1,
		"PNG":   2,
		"BMP":   3,
		"OTHER": 255,
	}
)

func (x IMAGE_FORMAT) Enum() *IMAGE_FORMAT {
	p := new(IMAGE_FORMAT)
	*p = x
	return p
}

func (x IMAGE_FORMAT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMAGE_FORMAT) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[2].Descriptor()
}

func (IMAGE_FORMAT) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[2]
}

func (x IMAGE_FORMAT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMAGE_FORMAT.Descriptor instead.
func (IMAGE_FORMAT) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

//视频格式
type VIDEO_FORMAT int32

const (
	VIDEO_FORMAT_MP4                VIDEO_FORMAT = 0
	VIDEO_FORMAT_MOV                VIDEO_FORMAT = 1
	VIDEO_FORMAT_WMV                VIDEO_FORMAT = 2
	VIDEO_FORMAT_FLV                VIDEO_FORMAT = 3
	VIDEO_FORMAT_AVI                VIDEO_FORMAT = 4
	VIDEO_FORMAT_MKV                VIDEO_FORMAT = 5
	VIDEO_FORMAT_OTHER_VIDEO_FORMAT VIDEO_FORMAT = 255
)

// Enum value maps for VIDEO_FORMAT.
var (
	VIDEO_FORMAT_name = map[int32]string{
		0:   "MP4",
		1:   "MOV",
		2:   "WMV",
		3:   "FLV",
		4:   "AVI",
		5:   "MKV",
		255: "OTHER_VIDEO_FORMAT",
	}
	VIDEO_FORMAT_value = map[string]int32{
		"MP4":                0,
		"MOV":                1,
		"WMV":                2,
		"FLV":                3,
		"AVI":                4,
		"MKV":                5,
		"OTHER_VIDEO_FORMAT": 255,
	}
)

func (x VIDEO_FORMAT) Enum() *VIDEO_FORMAT {
	p := new(VIDEO_FORMAT)
	*p = x
	return p
}

func (x VIDEO_FORMAT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VIDEO_FORMAT) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[3].Descriptor()
}

func (VIDEO_FORMAT) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[3]
}

func (x VIDEO_FORMAT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VIDEO_FORMAT.Descriptor instead.
func (VIDEO_FORMAT) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

//文本消息
type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *TextMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TextMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

//地理位置消息
type LocationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CoverImage string  `protobuf:"bytes,2,opt,name=cover_image,json=coverImage,proto3" json:"cover_image,omitempty"`
	Lat        float32 `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng        float32 `protobuf:"fixed32,4,opt,name=lng,proto3" json:"lng,omitempty"`
	MapLink    string  `protobuf:"bytes,5,opt,name=map_link,json=mapLink,proto3" json:"map_link,omitempty"`
	Desc       string  `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *LocationMessage) Reset() {
	*x = LocationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationMessage) ProtoMessage() {}

func (x *LocationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationMessage.ProtoReflect.Descriptor instead.
func (*LocationMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *LocationMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LocationMessage) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

func (x *LocationMessage) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LocationMessage) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *LocationMessage) GetMapLink() string {
	if x != nil {
		return x.MapLink
	}
	return ""
}

func (x *LocationMessage) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

//表情消息
type FaceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *FaceMessage) Reset() {
	*x = FaceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceMessage) ProtoMessage() {}

func (x *FaceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceMessage.ProtoReflect.Descriptor instead.
func (*FaceMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *FaceMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FaceMessage) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

//语音消息
type SoundMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Size    int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Seconds int32  `protobuf:"varint,4,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *SoundMessage) Reset() {
	*x = SoundMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundMessage) ProtoMessage() {}

func (x *SoundMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundMessage.ProtoReflect.Descriptor instead.
func (*SoundMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *SoundMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SoundMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SoundMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SoundMessage) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

//图片
type ImageMessageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   IMAGE_TYPE   `protobuf:"varint,1,opt,name=type,proto3,enum=IMAGE_TYPE" json:"type,omitempty"`       // 类型
	Format IMAGE_FORMAT `protobuf:"varint,2,opt,name=format,proto3,enum=IMAGE_FORMAT" json:"format,omitempty"` //格式
	Size   int32        `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`                       //大小
	Width  int32        `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`                     //宽度
	Height int32        `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`                   //高度
	Url    string       `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`                          //链接地址
}

func (x *ImageMessageItem) Reset() {
	*x = ImageMessageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageMessageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMessageItem) ProtoMessage() {}

func (x *ImageMessageItem) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMessageItem.ProtoReflect.Descriptor instead.
func (*ImageMessageItem) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *ImageMessageItem) GetType() IMAGE_TYPE {
	if x != nil {
		return x.Type
	}
	return IMAGE_TYPE_ORIGIN
}

func (x *ImageMessageItem) GetFormat() IMAGE_FORMAT {
	if x != nil {
		return x.Format
	}
	return IMAGE_FORMAT_JPG
}

func (x *ImageMessageItem) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageMessageItem) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageMessageItem) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageMessageItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

//图像消息
type ImageMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	List []*ImageMessageItem `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ImageMessage) Reset() {
	*x = ImageMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMessage) ProtoMessage() {}

func (x *ImageMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMessage.ProtoReflect.Descriptor instead.
func (*ImageMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *ImageMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ImageMessage) GetList() []*ImageMessageItem {
	if x != nil {
		return x.List
	}
	return nil
}

//文件消息
type FileMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Size int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` //大小
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`  //文件名
	Url  string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`    //链接地址
}

func (x *FileMessage) Reset() {
	*x = FileMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMessage) ProtoMessage() {}

func (x *FileMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMessage.ProtoReflect.Descriptor instead.
func (*FileMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *FileMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

//视频消息
type VideoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Size        int32        `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Second      int32        `protobuf:"varint,3,opt,name=second,proto3" json:"second,omitempty"`
	Url         string       `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Format      string       `protobuf:"bytes,5,opt,name=format,proto3" json:"format,omitempty"`
	ThumbUrl    string       `protobuf:"bytes,6,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	ThumbSize   string       `protobuf:"bytes,7,opt,name=thumb_size,json=thumbSize,proto3" json:"thumb_size,omitempty"`
	ThumbWidth  string       `protobuf:"bytes,8,opt,name=thumb_width,json=thumbWidth,proto3" json:"thumb_width,omitempty"`
	ThumbHeight string       `protobuf:"bytes,9,opt,name=thumb_height,json=thumbHeight,proto3" json:"thumb_height,omitempty"`
	ThumbFormat IMAGE_FORMAT `protobuf:"varint,10,opt,name=thumb_format,json=thumbFormat,proto3,enum=IMAGE_FORMAT" json:"thumb_format,omitempty"`
}

func (x *VideoMessage) Reset() {
	*x = VideoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoMessage) ProtoMessage() {}

func (x *VideoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoMessage.ProtoReflect.Descriptor instead.
func (*VideoMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *VideoMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VideoMessage) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

func (x *VideoMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoMessage) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *VideoMessage) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *VideoMessage) GetThumbSize() string {
	if x != nil {
		return x.ThumbSize
	}
	return ""
}

func (x *VideoMessage) GetThumbWidth() string {
	if x != nil {
		return x.ThumbWidth
	}
	return ""
}

func (x *VideoMessage) GetThumbHeight() string {
	if x != nil {
		return x.ThumbHeight
	}
	return ""
}

func (x *VideoMessage) GetThumbFormat() IMAGE_FORMAT {
	if x != nil {
		return x.ThumbFormat
	}
	return IMAGE_FORMAT_JPG
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x37, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x22, 0x35, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x5e, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x45, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x57, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xa6, 0x02, 0x0a, 0x0c, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x30, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2a, 0x5b, 0x0a, 0x0c, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x43,
	0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c,
	0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x06, 0x2a, 0x30,
	0x0a, 0x0a, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x47, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e, 0x41, 0x49, 0x4c, 0x10, 0x02,
	0x2a, 0x3e, 0x0a, 0x0c, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54,
	0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x49, 0x46,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x4d, 0x50, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0xff, 0x01,
	0x2a, 0x5d, 0x0a, 0x0c, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x50, 0x34, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4f, 0x56,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x4d, 0x56, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x46,
	0x4c, 0x56, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x56, 0x49, 0x10, 0x04, 0x12, 0x07, 0x0a,
	0x03, 0x4d, 0x4b, 0x56, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x12, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f,
	0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0xff, 0x01, 0x42,
	0x1a, 0x5a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []interface{}{
	(MESSAGE_TYPE)(0),        // 0: MESSAGE_TYPE
	(IMAGE_TYPE)(0),          // 1: IMAGE_TYPE
	(IMAGE_FORMAT)(0),        // 2: IMAGE_FORMAT
	(VIDEO_FORMAT)(0),        // 3: VIDEO_FORMAT
	(*TextMessage)(nil),      // 4: TextMessage
	(*LocationMessage)(nil),  // 5: LocationMessage
	(*FaceMessage)(nil),      // 6: FaceMessage
	(*SoundMessage)(nil),     // 7: SoundMessage
	(*ImageMessageItem)(nil), // 8: ImageMessageItem
	(*ImageMessage)(nil),     // 9: ImageMessage
	(*FileMessage)(nil),      // 10: FileMessage
	(*VideoMessage)(nil),     // 11: VideoMessage
}
var file_message_proto_depIdxs = []int32{
	1, // 0: ImageMessageItem.type:type_name -> IMAGE_TYPE
	2, // 1: ImageMessageItem.format:type_name -> IMAGE_FORMAT
	8, // 2: ImageMessage.list:type_name -> ImageMessageItem
	2, // 3: VideoMessage.thumb_format:type_name -> IMAGE_FORMAT
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoundMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageMessageItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
