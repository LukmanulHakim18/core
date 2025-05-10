package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/LukmanulHakim18/core/constant"
	"github.com/LukmanulHakim18/core/feature"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/metadata"
)

const (
	MetadataUserInfo        = "user-info"        // Ex: BB12345
	MetadataAcceptLang      = "accept-language"  // Ex: en
	MetadataAppVersion      = "app-version"      // Ex: 6.2.1
	MetadataDeviceType      = "device-type"      // Ex: Redmi
	MetadataToken           = "token"            // Ex: 691d84c7585eee4a
	MetadataDeviceUuid      = "device-uuid"      // Ex: 691d84c7585eee4a
	MetadataOSVersion       = "os-version"       // Ex: 11
	MetadataOperatingSystem = "operating-system" // Ex: Android
	MetadataDeviceModel     = "device-model"     // Ex: Redmi Note 8 Pro
	MetadataManufacturer    = "manufacturer"     // Ex: Samsung
	MetadataDeviceId        = "device-id"        // Ex: 691d84c7585eee4a
	MetadataRequestId       = "request-id"       // Ex: 691d84-c7585-eee4a
	MetadataDeviceTime      = "device-time"      // Ex: 2009-11-11T06:00:00+07:00 if wib  using rfc3339
	MetadataTraceId         = "trace-id"         // Ex: uuid, auto generate if nil from client
)

var ListOfMetadataKey []string = []string{
	MetadataUserInfo,
	MetadataAcceptLang,
	MetadataAppVersion,
	MetadataDeviceType,
	MetadataToken,
	MetadataDeviceUuid,
	MetadataOSVersion,
	MetadataOperatingSystem,
	MetadataDeviceModel,
	MetadataManufacturer,
	MetadataDeviceId,
	MetadataRequestId,
	MetadataDeviceTime,
	MetadataTraceId,
}

func AllowCommonMetadata(key string) bool {
	key = strings.ToLower(key)
	return slices.Contains(ListOfMetadataKey, key)
}

type Metadata struct {
	TraceId         string
	DeviceLang      constant.DeviceLang
	AppVersion      *version.Version
	OSVersion       *version.Version
	DeviceType      string
	DeviceUUID      string
	RequestId       string
	Token           string
	DeviceTime      *time.Time
	DeviceId        string
	Manufacturer    string
	DeviceModel     string
	OperatingSystem *constant.OperatingSystem
	UserInfo        *UserInfo
}

type UserInfo struct {
	InternalID      string                 `json:"internal_id"`
	Name            string                 `json:"name"`
	Email           string                 `json:"email"`
	PhoneNumber     string                 `json:"phone_number"`
	ProfileImage    string                 `json:"profile_image"`
	ReferralCode    string                 `json:"referral_code"`
	CreatedAt       int64                  `json:"created_at"`
	VerifiedPhoneAt int64                  `json:"verified_phone_at"`
	VerifiedEmailAt int64                  `json:"verified_email_at"`
	EnabledFeature  feature.EnabledFeature `json:"enabled_features"`
}

// Deprecated: As of Aphrodite 1.7.1, this function simply calls [metadata.GetMetaDataFromContext].
func MakeMetadataFromCtx(ctx context.Context) (meta Metadata, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = fmt.Errorf("Oops! Something went wrong!")
		return
	}
	meta.setDeviceLanguage(md)

	if err = meta.setAppVersion(md); err != nil {
		return
	}
	if err = meta.setDeviceType(md); err != nil {
		return
	}

	if err = meta.setDeviceUUID(md); err != nil {
		return
	}

	if err = meta.setRequestId(md); err != nil {
		return
	}

	if err = meta.setToken(md); err != nil {
		return
	}
	return
}

func GetMetaDataFromContext(ctx context.Context) Metadata {
	var (
		res = Metadata{}
		tmp []string
	)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return res
	}

	// Accept-language
	langArr := md.Get(MetadataAcceptLang)
	if langArr != nil {
		lang := strings.ToUpper(langArr[0])
		if lang == "ID" {
			res.DeviceLang = constant.DEVICE_LANG_ID
		} else {
			res.DeviceLang = constant.DEVICE_LANG_EN
		}
	} else {
		res.DeviceLang = constant.DEVICE_LANG_EN // Default language
	}

	// App-Version
	tmp = md.Get(MetadataAppVersion)
	if len(tmp) > 0 {
		res.AppVersion, _ = version.NewSemver(tmp[0])
	}

	// Token
	tmp = md.Get(MetadataToken)
	if len(tmp) > 0 {
		res.Token = tmp[0]
	}

	// Device-Type
	tmp = md.Get(MetadataDeviceType)
	if len(tmp) > 0 {
		res.DeviceType = tmp[0]
	}

	// Device-Uuid
	tmp = md.Get(MetadataDeviceUuid)
	if len(tmp) > 0 {
		res.DeviceUUID = tmp[0]
	}

	// Os-Version
	tmp = md.Get(MetadataOSVersion)
	if len(tmp) > 0 {
		res.OSVersion, _ = version.NewSemver(tmp[0])
	}

	// Operating-System
	tmp = md.Get(MetadataOperatingSystem)
	if len(tmp) > 0 {
		res.OperatingSystem, _ = constant.OSFromStr(tmp[0])
	}

	// Device-Model
	tmp = md.Get(MetadataDeviceModel)
	if len(tmp) > 0 {
		res.DeviceModel = tmp[0]
	}

	// Manufacturer
	tmp = md.Get(MetadataManufacturer)
	if len(tmp) > 0 {
		res.Manufacturer = tmp[0]
	}

	// Device-Id
	tmp = md.Get(MetadataDeviceId)
	if len(tmp) > 0 {
		res.DeviceId = tmp[0]
	}

	// Request-Id
	tmp = md.Get(MetadataRequestId)
	if len(tmp) > 0 {
		res.RequestId = tmp[0]
	}

	// Device-Time
	tmp = md.Get(MetadataDeviceTime)
	if len(tmp) > 0 {
		deviceTime := tmp[0]
		layout := time.RFC3339
		parsedTime, err := time.Parse(layout, deviceTime)
		if err != nil {
			log.Printf("[GetMetaDataFromContext] error when parse device-time string to time.Time. error: %s device-time: %s\n", err.Error(), deviceTime)
			parsedTime = time.Now()
		}
		res.DeviceTime = &parsedTime

	} else {
		layout := time.RFC3339
		now := time.Now().Format(layout)
		parsedTime, err := time.Parse(layout, now)
		if err == nil {
			res.DeviceTime = &parsedTime
		}
	}

	// User-Info
	tmp = md.Get(MetadataUserInfo)
	if len(tmp) > 0 {
		userInfo := UserInfo{}
		err := json.Unmarshal([]byte(tmp[0]), &userInfo)
		if err == nil {
			res.UserInfo = &userInfo
		}
	}

	// Trace-Id
	tmp = md.Get(MetadataTraceId)
	if len(tmp) > 0 {
		res.TraceId = tmp[0]
	} else {
		res.TraceId = uuid.NewString()
	}

	return res
}

func GetMetaDataFromContextWithDeviceTimeQueryUnescape(ctx context.Context) Metadata {
	var (
		res = Metadata{}
		tmp []string
	)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return res
	}

	// Accept-language
	langArr := md.Get(MetadataAcceptLang)
	if langArr != nil {
		lang := strings.ToUpper(langArr[0])
		if lang == "ID" {
			res.DeviceLang = constant.DEVICE_LANG_ID
		} else {
			res.DeviceLang = constant.DEVICE_LANG_EN
		}
	} else {
		res.DeviceLang = constant.DEVICE_LANG_EN // Default language
	}

	// App-Version
	tmp = md.Get(MetadataAppVersion)
	if len(tmp) > 0 {
		res.AppVersion, _ = version.NewSemver(tmp[0])
	}

	// Token
	tmp = md.Get(MetadataToken)
	if len(tmp) > 0 {
		res.Token = tmp[0]
	}

	// Device-Type
	tmp = md.Get(MetadataDeviceType)
	if len(tmp) > 0 {
		res.DeviceType = tmp[0]
	}

	// Device-Uuid
	tmp = md.Get(MetadataDeviceUuid)
	if len(tmp) > 0 {
		res.DeviceUUID = tmp[0]
	}

	// Os-Version
	tmp = md.Get(MetadataOSVersion)
	if len(tmp) > 0 {
		res.OSVersion, _ = version.NewSemver(tmp[0])
	}

	// Operating-System
	tmp = md.Get(MetadataOperatingSystem)
	if len(tmp) > 0 {
		res.OperatingSystem, _ = constant.OSFromStr(tmp[0])
	}

	// Device-Model
	tmp = md.Get(MetadataDeviceModel)
	if len(tmp) > 0 {
		res.DeviceModel = tmp[0]
	}

	// Manufacturer
	tmp = md.Get(MetadataManufacturer)
	if len(tmp) > 0 {
		res.Manufacturer = tmp[0]
	}

	// Device-Id
	tmp = md.Get(MetadataDeviceId)
	if len(tmp) > 0 {
		res.DeviceId = tmp[0]
	}

	// Request-Id
	tmp = md.Get(MetadataRequestId)
	if len(tmp) > 0 {
		res.RequestId = tmp[0]
	}

	// Device-Time
	tmp = md.Get(MetadataDeviceTime)
	if len(tmp) > 0 {
		deviceTime := tmp[0]

		// URL-decode the time string first
		deviceTimeAfterQueryUnescape, errQueryUnescape := url.QueryUnescape(deviceTime)
		if errQueryUnescape != nil {
			log.Printf("[GetMetaDataFromContextWithDeviceTimeQueryUnescape] error when url.QueryUnescape. error: %s device-time: %s\n", errQueryUnescape.Error(), deviceTime)
		} else {
			deviceTime = deviceTimeAfterQueryUnescape
		}

		layout := time.RFC3339
		parsedTime, err := time.Parse(layout, deviceTime)
		if err == nil {
			res.DeviceTime = &parsedTime
		} else {
			log.Printf("[GetMetaDataFromContextWithDeviceTimeQueryUnescape] error when parse device-time string to time.Time. error: %s device-time: %s\n", err.Error(), deviceTime)
			now := time.Now().Format(layout)
			parsedTime, err := time.Parse(layout, now)
			if err == nil {
				res.DeviceTime = &parsedTime
			}
		}
	} else {
		layout := time.RFC3339
		now := time.Now().Format(layout)
		parsedTime, err := time.Parse(layout, now)
		if err == nil {
			res.DeviceTime = &parsedTime
		}
	}

	// User-Info
	tmp = md.Get(MetadataUserInfo)
	if len(tmp) > 0 {
		userInfo := UserInfo{}
		err := json.Unmarshal([]byte(tmp[0]), &userInfo)
		if err == nil {
			res.UserInfo = &userInfo
		}
	}

	// Trace-Id
	tmp = md.Get(MetadataTraceId)
	if len(tmp) > 0 {
		res.TraceId = tmp[0]
	} else {
		res.TraceId = uuid.NewString()
	}

	return res
}

func CopyMetaDataToNewContext(newCtx context.Context, oldCtx context.Context) context.Context {
	if md, ok := metadata.FromIncomingContext(oldCtx); ok {
		return metadata.NewIncomingContext(newCtx, md)
	}
	return oldCtx
}

func (m *Metadata) setDeviceLanguage(md metadata.MD) {
	// Get language requested from metadata(header in http), default is english
	langArr := md.Get(MetadataAcceptLang)
	if langArr != nil {
		lang := strings.ToUpper(langArr[0])
		if lang == "ID" {
			m.DeviceLang = constant.DEVICE_LANG_ID
		} else {
			m.DeviceLang = constant.DEVICE_LANG_EN
		}
	} else {
		m.DeviceLang = constant.DEVICE_LANG_EN // Default language
	}
}

func (m *Metadata) setAppVersion(md metadata.MD) (err error) {

	// Get language requested, default is english
	appVersionSlice := md.Get(MetadataAppVersion)
	if appVersionSlice == nil {
		return fmt.Errorf("undefined version")
	}
	if m.AppVersion, err = version.NewSemver(appVersionSlice[0]); err != nil {
		return err
	}
	return nil
}

func (m *Metadata) setDeviceType(md metadata.MD) (err error) {

	// Get language requested, default is english
	deviceType := md.Get(MetadataDeviceType)
	if deviceType == nil {
		return fmt.Errorf("undefined device type")
	}
	m.DeviceType = deviceType[0]
	return
}

func (m *Metadata) setDeviceUUID(md metadata.MD) (err error) {
	// Get language requested, default is english
	deviceUuid := md.Get(MetadataDeviceUuid)
	if deviceUuid == nil {
		return fmt.Errorf("undefined device uuid")
	}
	m.DeviceUUID = deviceUuid[0]
	return
}

func (m *Metadata) setRequestId(md metadata.MD) (err error) {
	// Get language requested, default is english
	requestId := md.Get(MetadataRequestId)
	if requestId == nil {
		return fmt.Errorf("undefined request id")
	}
	m.RequestId = requestId[0]
	return
}

func (m *Metadata) setToken(md metadata.MD) error {
	token := md.Get(MetadataToken)
	if token == nil {
		return fmt.Errorf("undefined token")
	}
	m.Token = token[0]
	return nil
}

func GetDeviceLanguageFromCtx(ctx context.Context) constant.DeviceLang {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return constant.DEVICE_LANG_EN
	}
	// Get language requested from metadata(header in http), default is english
	langArr := md.Get(MetadataAcceptLang)
	if langArr != nil {
		lang := strings.ToUpper(langArr[0])
		if lang == "ID" {
			return constant.DEVICE_LANG_ID
		} else {
			return constant.DEVICE_LANG_EN
		}
	} else {
		return constant.DEVICE_LANG_EN // Default language
	}
}

func GetDeviceVersionFromCtx(ctx context.Context) (ver *version.Version, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("error parsing")
	}
	// Get language requested, default is english
	appVersionSlice := md.Get(MetadataAppVersion)
	if appVersionSlice == nil {
		return nil, fmt.Errorf("undefined version")
	}
	deviceVersion, err := version.NewSemver(appVersionSlice[0])
	if err != nil {
		return nil, err
	}
	return deviceVersion, nil
}

func GetDeviceTypeFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("error parsing")
	}
	// Get language requested, default is english
	deviceType := md.Get(MetadataDeviceType)
	if deviceType == nil {
		return "", fmt.Errorf("device type not found")
	}

	return deviceType[0], nil
}

func GetDeviceUuid(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("error parsing")
	}
	// Get language requested, default is english
	deviceUuid := md.Get(MetadataDeviceUuid)
	if deviceUuid == nil {
		return "", fmt.Errorf("device uuid not found")
	}

	return deviceUuid[0], nil
}

func GetRequestId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("error parsing")
	}
	// Get language requested, default is english
	requestId := md.Get(MetadataRequestId)
	if requestId == nil {
		return "", fmt.Errorf("request id not found")
	}

	return requestId[0], nil
}

func GetTokenFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("token not found")
	}
	// Get language requested, default is english
	requestId := md.Get(MetadataToken)
	if requestId == nil {
		return "", fmt.Errorf("user id not found")
	}
	return requestId[0], nil
}

// Get trace-id from context
func GetTraceIdFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("trace-id not found")
	}
	data := md.Get(MetadataTraceId)
	if data == nil {
		return "", fmt.Errorf("trace-id not found")
	}
	return data[0], nil
}

// using lower case key
func GetMetadataByKey(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	// Get language requested, default is english
	mdSlice := md.Get(key)
	if len(mdSlice) == 0 {
		return ""
	}
	return mdSlice[0]
}

func InitiateTraceId(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}

	traceId := md.Get(MetadataTraceId)
	if len(traceId) == 0 {
		newTraceId := uuid.NewString()
		md.Set(MetadataTraceId, newTraceId)
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	ctx = context.WithValue(ctx, MetadataTraceId, md.Get(MetadataTraceId)[0])

	return ctx
}
