package vod

import (
	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/models"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type RedirectPlayReq struct {
	Vid        string `json:"Vid"`
	Definition string `json:"Definition"` //视频分辨率
	LogoType   string `json:"LogoType,omitempty"`
	Expires    string
}

type StartWorkflowRequest struct {
	Vid          string
	TemplateId   string
	Input        map[string]interface{}
	Priority     int
	CallbackArgs string
}

type StartWorkflowResult struct {
	RunId string
}

type StartWorkflowResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *StartWorkflowResult `json:",omitempty"`
}

type UploadMediaByUrlResult struct {
	Code    int
	Message string
}

type UploadMediaByUrlResp struct {
	base.CommonResponse
	Result UploadMediaByUrlResult
}

type VideoFormat string

const (
	MP4  VideoFormat = "mp4"
	M3U8 VideoFormat = "m3u8"
)

type UploadMediaByUrlParams struct {
	SpaceName    string
	Format       VideoFormat
	SourceUrls   []string
	CallbackArgs string
}

type FileType string

const (
	VIDEO  FileType = "video"
	IMAGE  FileType = "image"
	OBJECT FileType = "object"
)

type ApplyUploadParam struct {
	SpaceName  string
	SessionKey string
	FileType   FileType
	FileSize   int
	UploadNum  int
}

type ApplyUploadResp struct {
	base.CommonResponse
	Result ApplyUploadResult
}

type ApplyUploadResult struct {
	RequestID     string
	UploadAddress UploadAddress
}
type UploadAddress struct {
	StoreInfos    []StoreInfo
	UploadHosts   []string
	UploadHeader  map[string]string
	SessionKey    string
	AdvanceOption AdvanceOption
}

type VideoDefinition string

const (
	D1080P VideoDefinition = "1080p"
	D720P  VideoDefinition = "720p"
	D540P  VideoDefinition = "540p"
	D480P  VideoDefinition = "480p"
	D360P  VideoDefinition = "360p"
	D240P  VideoDefinition = "240p"
)

type StoreInfo struct {
	StoreUri string
	Auth     string
}

type AdvanceOption struct {
	Parallel  int
	Stream    int
	SliceSize int
}

type BaseResp struct {
	StatusMessage string
	StatusCode    int
}

type CommitUploadParam struct {
	SpaceName string
	Body      CommitUploadBody
}

type CommitUploadBody struct {
	CallbackArgs string
	SessionKey   string
	Functions    []Function
}

type Function struct {
	Name  string
	Input interface{}
}

type SnapshotInput struct {
	SnapshotTime float64
}

type EntryptionInput struct {
	Config       map[string]string
	PolicyParams map[string]string
}

type OptionInfo struct {
	Title       string
	Tags        string
	Description string
	Category    string
}

type WorkflowInput struct {
	TemplateId string
}

type CommitUploadResp struct {
	base.CommonResponse
	Result CommitUploadResult
}

type CommitUploadResult struct {
	RequestId    string
	CallbackArgs string
	Results      []UploadResult
}

type UploadResult struct {
	Vid         string
	VideoMeta   VideoMeta
	ImageMeta   ImageMeta
	ObjectMeta  ObjectMeta
	Encryption  Encryption
	SnapshotUri string
}
type VideoMeta struct {
	Uri      string
	Height   int
	Width    int
	Duration float64
	Bitrate  int
	Md5      string
	Format   string
	Size     int
}

type ImageMeta struct {
	Uri    string
	Height int
	Width  int
	Md5    string
}

type ObjectMeta struct {
	Uri string
	Md5 string
}

type Encryption struct {
	Uri       string
	SecretKey string
	Algorithm string
	Version   string
	SourceMd5 string
	Extra     map[string]string
}

type GetWeightsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           map[string]map[string]int `json:",omitempty"`
}

type DomainInfo struct {
	MainDomain   string
	BackupDomain string
}

type ImgUrl struct {
	MainUrl   string
	BackupUrl string
}

//UpdateVideoInfo
type UpdateVideoInfoRequest struct {
	models.UpdateVideoInfoRequest
}

func NewUpdateVideoInfoRequest() *UpdateVideoInfoRequest {
	return &UpdateVideoInfoRequest{}
}

func (r *UpdateVideoInfoRequest) IsSetPosterUri() bool {
	return r.PosterUri != nil
}

func (r *UpdateVideoInfoRequest) IsSetTitle() bool {
	return r.Title != nil
}

func (r *UpdateVideoInfoRequest) IsSetDescription() bool {
	return r.Description != nil
}

func (r *UpdateVideoInfoRequest) IsSetTags() bool {
	return r.Tags != nil
}

func (r *UpdateVideoInfoRequest) SetVid(vid string) {
	r.Vid = vid
}

func (r *UpdateVideoInfoRequest) SetPosterUri(posterUri string) {
	r.PosterUri = &wrappers.StringValue{Value: posterUri}
}

func (r *UpdateVideoInfoRequest) SetTitle(title string) {
	r.Title = &wrappers.StringValue{Value: title}
}

func (r *UpdateVideoInfoRequest) SetDescription(desc string) {
	r.Description = &wrappers.StringValue{Value: desc}
}

func (r *UpdateVideoInfoRequest) SetTags(tags string) {
	r.Tags = &wrappers.StringValue{Value: tags}
}

func (r *UpdateVideoInfoRequest) GetVid() string {
	return r.Vid
}

func (r *UpdateVideoInfoRequest) GetPosterUri() string {
	return r.PosterUri.GetValue()
}

func (r *UpdateVideoInfoRequest) GetTitle() string {
	return r.Title.GetValue()
}

func (r *UpdateVideoInfoRequest) GetDescription() string {
	return r.Description.GetValue()
}

func (r *UpdateVideoInfoRequest) GetTags() string {
	return r.Tags.GetValue()
}

type UpdateVideoInfoResponse struct {
	ResponseMetadata *base.ResponseMetadata
}

// UpdateVideoPublishStatus
func NewUpdateVideoPublishStatusRequest() *UpdateVideoPublishStatusRequest {
	return &UpdateVideoPublishStatusRequest{}
}

type UpdateVideoPublishStatusRequest struct {
	models.UpdateVideoPublishStatusRequest
}

func (r *UpdateVideoPublishStatusRequest) SetVid(vid string) {
	r.Vid = vid
}

func (r *UpdateVideoPublishStatusRequest) SetStatus(status string) {
	r.Status = status
}

func (r *UpdateVideoPublishStatusRequest) GetVid() string {
	return r.Vid
}

func (r *UpdateVideoPublishStatusRequest) GetStatus() string {
	return r.Status
}

type UpdateVideoPublishStatusResponse struct {
	ResponseMetadata *base.ResponseMetadata
}

//视频详情
type GetVideoInfosRequest struct {
	models.GetVideoInfosRequest
}

func NewGetVideoInfosRequest() *GetVideoInfosRequest {
	return &GetVideoInfosRequest{}
}

func (r *GetVideoInfosRequest) SetVids(vids string) {
	r.Vids = vids
}

func (r *GetVideoInfosRequest) GetVids() string {
	return r.Vids
}

type GetVideoInfosResponse struct {
	models.GetVideoInfosResponse
}

//候选封面
type GetRecommendedPostersRequest struct {
	models.GetRecommendedPostersRequest
}

func NewGetRecommendedPostersRequest() *GetRecommendedPostersRequest {
	return &GetRecommendedPostersRequest{}
}

func (r *GetRecommendedPostersRequest) SetVids(vids string) {
	r.Vids = vids
}

func (r *GetRecommendedPostersRequest) GetVids() string {
	return r.Vids
}

type GetRecommendedPostersResponse struct {
	models.GetRecPostersResponse
}
