// Package speech provides access to the Cloud Speech API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/speech/apiv1 instead.
//
// See https://cloud.google.com/speech-to-text/docs/quickstart-protocol
//
// Usage example:
//
//   import "google.golang.org/api/speech/v1"
//   ...
//   speechService, err := speech.New(oauthHttpClient)
package speech // import "google.golang.org/api/speech/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "speech:v1"
const apiName = "speech"
const apiVersion = "v1"
const basePath = "https://speech.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Operations = NewOperationsService(s)
	s.Speech = NewSpeechService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Operations *OperationsService

	Speech *SpeechService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewSpeechService(s *Service) *SpeechService {
	rs := &SpeechService{s: s}
	return rs
}

type SpeechService struct {
	s *Service
}

// LongRunningRecognizeRequest: The top-level message sent by the client
// for the `LongRunningRecognize`
// method.
type LongRunningRecognizeRequest struct {
	// Audio: *Required* The audio data to be recognized.
	Audio *RecognitionAudio `json:"audio,omitempty"`

	// Config: *Required* Provides information to the recognizer that
	// specifies how to
	// process the request.
	Config *RecognitionConfig `json:"config,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audio") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LongRunningRecognizeRequest) MarshalJSON() ([]byte, error) {
	type NoMethod LongRunningRecognizeRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RecognitionAudio: Contains audio data in the encoding specified in
// the `RecognitionConfig`.
// Either `content` or `uri` must be supplied. Supplying both or
// neither
// returns google.rpc.Code.INVALID_ARGUMENT. See
// [audio limits](https://cloud.google.com/speech/limits#content).
type RecognitionAudio struct {
	// Content: The audio data bytes encoded as specified
	// in
	// `RecognitionConfig`. Note: as with all bytes fields, protobuffers use
	// a
	// pure binary representation, whereas JSON representations use base64.
	Content string `json:"content,omitempty"`

	// Uri: URI that points to a file that contains audio data bytes as
	// specified in
	// `RecognitionConfig`. Currently, only Google Cloud Storage URIs
	// are
	// supported, which must be specified in the following
	// format:
	// `gs://bucket_name/object_name` (other URI formats
	// return
	// google.rpc.Code.INVALID_ARGUMENT). For more information, see
	// [Request URIs](https://cloud.google.com/storage/docs/reference-uris).
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RecognitionAudio) MarshalJSON() ([]byte, error) {
	type NoMethod RecognitionAudio
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RecognitionConfig: Provides information to the recognizer that
// specifies how to process the
// request.
type RecognitionConfig struct {
	// EnableWordTimeOffsets: *Optional* If `true`, the top result includes
	// a list of words and
	// the start and end time offsets (timestamps) for those words.
	// If
	// `false`, no word-level time offset information is returned. The
	// default is
	// `false`.
	EnableWordTimeOffsets bool `json:"enableWordTimeOffsets,omitempty"`

	// Encoding: Encoding of audio data sent in all `RecognitionAudio`
	// messages.
	// This field is optional for `FLAC` and `WAV` audio files and
	// required
	// for all other audio formats. For details, see AudioEncoding.
	//
	// Possible values:
	//   "ENCODING_UNSPECIFIED" - Not specified.
	//   "LINEAR16" - Uncompressed 16-bit signed little-endian samples
	// (Linear PCM).
	//   "FLAC" - `FLAC` (Free Lossless Audio
	// Codec) is the recommended encoding because it is
	// lossless--therefore recognition is not compromised--and
	// requires only about half the bandwidth of `LINEAR16`. `FLAC`
	// stream
	// encoding supports 16-bit and 24-bit samples, however, not all fields
	// in
	// `STREAMINFO` are supported.
	//   "MULAW" - 8-bit samples that compand 14-bit audio samples using
	// G.711 PCMU/mu-law.
	//   "AMR" - Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz`
	// must be 8000.
	//   "AMR_WB" - Adaptive Multi-Rate Wideband codec. `sample_rate_hertz`
	// must be 16000.
	//   "OGG_OPUS" - Opus encoded audio frames in Ogg
	// container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_her
	// tz` must be one of 8000, 12000, 16000, 24000, or 48000.
	//   "SPEEX_WITH_HEADER_BYTE" - Although the use of lossy encodings is
	// not recommended, if a very low
	// bitrate encoding is required, `OGG_OPUS` is highly preferred
	// over
	// Speex encoding. The [Speex](https://speex.org/)  encoding supported
	// by
	// Cloud Speech API has a header byte in each block, as in MIME
	// type
	// `audio/x-speex-with-header-byte`.
	// It is a variant of the RTP Speex encoding defined in
	// [RFC 5574](https://tools.ietf.org/html/rfc5574).
	// The stream is a sequence of blocks, one block per RTP packet. Each
	// block
	// starts with a byte containing the length of the block, in bytes,
	// followed
	// by one or more frames of Speex data, padded to an integral number
	// of
	// bytes (octets) as specified in RFC 5574. In other words, each RTP
	// header
	// is replaced with a single byte containing the block length. Only
	// Speex
	// wideband is supported. `sample_rate_hertz` must be 16000.
	Encoding string `json:"encoding,omitempty"`

	// LanguageCode: *Required* The language of the supplied audio as
	// a
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language
	// tag.
	// Example: "en-US".
	// See [Language
	// Support](https://cloud.google.com/speech/docs/languages)
	// for a list of the currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// MaxAlternatives: *Optional* Maximum number of recognition hypotheses
	// to be returned.
	// Specifically, the maximum number of `SpeechRecognitionAlternative`
	// messages
	// within each `SpeechRecognitionResult`.
	// The server may return fewer than `max_alternatives`.
	// Valid values are `0`-`30`. A value of `0` or `1` will return a
	// maximum of
	// one. If omitted, will return a maximum of one.
	MaxAlternatives int64 `json:"maxAlternatives,omitempty"`

	// ProfanityFilter: *Optional* If set to `true`, the server will attempt
	// to filter out
	// profanities, replacing all but the initial character in each filtered
	// word
	// with asterisks, e.g. "f***". If set to `false` or omitted,
	// profanities
	// won't be filtered out.
	ProfanityFilter bool `json:"profanityFilter,omitempty"`

	// SampleRateHertz: Sample rate in Hertz of the audio data sent in
	// all
	// `RecognitionAudio` messages. Valid values are: 8000-48000.
	// 16000 is optimal. For best results, set the sampling rate of the
	// audio
	// source to 16000 Hz. If that's not possible, use the native sample
	// rate of
	// the audio source (instead of re-sampling).
	// This field is optional for `FLAC` and `WAV` audio files and
	// required
	// for all other audio formats. For details, see AudioEncoding.
	SampleRateHertz int64 `json:"sampleRateHertz,omitempty"`

	// SpeechContexts: *Optional* A means to provide context to assist the
	// speech recognition.
	SpeechContexts []*SpeechContext `json:"speechContexts,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EnableWordTimeOffsets") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EnableWordTimeOffsets") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RecognitionConfig) MarshalJSON() ([]byte, error) {
	type NoMethod RecognitionConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RecognizeRequest: The top-level message sent by the client for the
// `Recognize` method.
type RecognizeRequest struct {
	// Audio: *Required* The audio data to be recognized.
	Audio *RecognitionAudio `json:"audio,omitempty"`

	// Config: *Required* Provides information to the recognizer that
	// specifies how to
	// process the request.
	Config *RecognitionConfig `json:"config,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audio") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RecognizeRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RecognizeRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RecognizeResponse: The only message returned to the client by the
// `Recognize` method. It
// contains the result as zero or more sequential
// `SpeechRecognitionResult`
// messages.
type RecognizeResponse struct {
	// Results: Output only. Sequential list of transcription results
	// corresponding to
	// sequential portions of audio.
	Results []*SpeechRecognitionResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Results") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Results") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RecognizeResponse) MarshalJSON() ([]byte, error) {
	type NoMethod RecognizeResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpeechContext: Provides "hints" to the speech recognizer to favor
// specific words and phrases
// in the results.
type SpeechContext struct {
	// Phrases: *Optional* A list of strings containing words and phrases
	// "hints" so that
	// the speech recognition is more likely to recognize them. This can be
	// used
	// to improve the accuracy for specific words and phrases, for example,
	// if
	// specific commands are typically spoken by the user. This can also be
	// used
	// to add additional words to the vocabulary of the recognizer.
	// See
	// [usage limits](https://cloud.google.com/speech/limits#content).
	Phrases []string `json:"phrases,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Phrases") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Phrases") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpeechContext) MarshalJSON() ([]byte, error) {
	type NoMethod SpeechContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpeechRecognitionAlternative: Alternative hypotheses (a.k.a. n-best
// list).
type SpeechRecognitionAlternative struct {
	// Confidence: Output only. The confidence estimate between 0.0 and 1.0.
	// A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. This field is set only for the top alternative of a
	// non-streaming
	// result or, of a streaming result where `is_final=true`.
	// This field is not guaranteed to be accurate and users should not rely
	// on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was
	// not set.
	Confidence float64 `json:"confidence,omitempty"`

	// Transcript: Output only. Transcript text representing the words that
	// the user spoke.
	Transcript string `json:"transcript,omitempty"`

	// Words: Output only. A list of word-specific information for each
	// recognized word.
	// Note: When enable_speaker_diarization is true, you will see all the
	// words
	// from the beginning of the audio.
	Words []*WordInfo `json:"words,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Confidence") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Confidence") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpeechRecognitionAlternative) MarshalJSON() ([]byte, error) {
	type NoMethod SpeechRecognitionAlternative
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *SpeechRecognitionAlternative) UnmarshalJSON(data []byte) error {
	type NoMethod SpeechRecognitionAlternative
	var s1 struct {
		Confidence gensupport.JSONFloat64 `json:"confidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Confidence = float64(s1.Confidence)
	return nil
}

// SpeechRecognitionResult: A speech recognition result corresponding to
// a portion of the audio.
type SpeechRecognitionResult struct {
	// Alternatives: Output only. May contain one or more recognition
	// hypotheses (up to the
	// maximum specified in `max_alternatives`).
	// These alternatives are ordered in terms of accuracy, with the top
	// (first)
	// alternative being the most probable, as ranked by the recognizer.
	Alternatives []*SpeechRecognitionAlternative `json:"alternatives,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alternatives") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alternatives") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpeechRecognitionResult) MarshalJSON() ([]byte, error) {
	type NoMethod SpeechRecognitionResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WordInfo: Word-specific information for recognized words.
type WordInfo struct {
	// EndTime: Output only. Time offset relative to the beginning of the
	// audio,
	// and corresponding to the end of the spoken word.
	// This field is only set if `enable_word_time_offsets=true` and only
	// in the top hypothesis.
	// This is an experimental feature and the accuracy of the time offset
	// can
	// vary.
	EndTime string `json:"endTime,omitempty"`

	// SpeakerTag: Output only. A distinct integer value is assigned for
	// every speaker within
	// the audio. This field specifies which one of those speakers was
	// detected to
	// have spoken this word. Value ranges from '1' to
	// diarization_speaker_count.
	// speaker_tag is set if enable_speaker_diarization = 'true' and only in
	// the
	// top alternative.
	SpeakerTag int64 `json:"speakerTag,omitempty"`

	// StartTime: Output only. Time offset relative to the beginning of the
	// audio,
	// and corresponding to the start of the spoken word.
	// This field is only set if `enable_word_time_offsets=true` and only
	// in the top hypothesis.
	// This is an experimental feature and the accuracy of the time offset
	// can
	// vary.
	StartTime string `json:"startTime,omitempty"`

	// Word: Output only. The word corresponding to this set of information.
	Word string `json:"word,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WordInfo) MarshalJSON() ([]byte, error) {
	type NoMethod WordInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "speech.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/operations/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "speech.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "speech.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/operations/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "speech.speech.longrunningrecognize":

type SpeechLongrunningrecognizeCall struct {
	s                           *Service
	longrunningrecognizerequest *LongRunningRecognizeRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// Longrunningrecognize: Performs asynchronous speech recognition:
// receive results via the
// google.longrunning.Operations interface. Returns either
// an
// `Operation.error` or an `Operation.response` which contains
// a `LongRunningRecognizeResponse` message.
func (r *SpeechService) Longrunningrecognize(longrunningrecognizerequest *LongRunningRecognizeRequest) *SpeechLongrunningrecognizeCall {
	c := &SpeechLongrunningrecognizeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.longrunningrecognizerequest = longrunningrecognizerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpeechLongrunningrecognizeCall) Fields(s ...googleapi.Field) *SpeechLongrunningrecognizeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpeechLongrunningrecognizeCall) Context(ctx context.Context) *SpeechLongrunningrecognizeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpeechLongrunningrecognizeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpeechLongrunningrecognizeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.longrunningrecognizerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/speech:longrunningrecognize")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "speech.speech.longrunningrecognize" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SpeechLongrunningrecognizeCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Performs asynchronous speech recognition: receive results via the\ngoogle.longrunning.Operations interface. Returns either an\n`Operation.error` or an `Operation.response` which contains\na `LongRunningRecognizeResponse` message.",
	//   "flatPath": "v1/speech:longrunningrecognize",
	//   "httpMethod": "POST",
	//   "id": "speech.speech.longrunningrecognize",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/speech:longrunningrecognize",
	//   "request": {
	//     "$ref": "LongRunningRecognizeRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "speech.speech.recognize":

type SpeechRecognizeCall struct {
	s                *Service
	recognizerequest *RecognizeRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Recognize: Performs synchronous speech recognition: receive results
// after all audio
// has been sent and processed.
func (r *SpeechService) Recognize(recognizerequest *RecognizeRequest) *SpeechRecognizeCall {
	c := &SpeechRecognizeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.recognizerequest = recognizerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpeechRecognizeCall) Fields(s ...googleapi.Field) *SpeechRecognizeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpeechRecognizeCall) Context(ctx context.Context) *SpeechRecognizeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpeechRecognizeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpeechRecognizeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.recognizerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/speech:recognize")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "speech.speech.recognize" call.
// Exactly one of *RecognizeResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *RecognizeResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SpeechRecognizeCall) Do(opts ...googleapi.CallOption) (*RecognizeResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RecognizeResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Performs synchronous speech recognition: receive results after all audio\nhas been sent and processed.",
	//   "flatPath": "v1/speech:recognize",
	//   "httpMethod": "POST",
	//   "id": "speech.speech.recognize",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/speech:recognize",
	//   "request": {
	//     "$ref": "RecognizeRequest"
	//   },
	//   "response": {
	//     "$ref": "RecognizeResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
