// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: pb/messaging.proto

#include "pb/messaging.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// This is a temporary google only hack
#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
#include "third_party/protobuf/version.h"
#endif
// @@protoc_insertion_point(includes)
namespace pb {
class DNDirectMessageResponseDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<DNDirectMessageResponse>
      _instance;
} _DNDirectMessageResponse_default_instance_;
class DNDirectMessageRequestDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<DNDirectMessageRequest>
      _instance;
} _DNDirectMessageRequest_default_instance_;
class RPCDirectMessageDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<RPCDirectMessage>
      _instance;
} _RPCDirectMessage_default_instance_;
}  // namespace pb
namespace protobuf_pb_2fmessaging_2eproto {
void InitDefaultsDNDirectMessageResponseImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  ::google::protobuf::internal::InitProtobufDefaultsForceUnique();
#else
  ::google::protobuf::internal::InitProtobufDefaults();
#endif  // GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  {
    void* ptr = &::pb::_DNDirectMessageResponse_default_instance_;
    new (ptr) ::pb::DNDirectMessageResponse();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::DNDirectMessageResponse::InitAsDefaultInstance();
}

void InitDefaultsDNDirectMessageResponse() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &InitDefaultsDNDirectMessageResponseImpl);
}

void InitDefaultsDNDirectMessageRequestImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  ::google::protobuf::internal::InitProtobufDefaultsForceUnique();
#else
  ::google::protobuf::internal::InitProtobufDefaults();
#endif  // GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  {
    void* ptr = &::pb::_DNDirectMessageRequest_default_instance_;
    new (ptr) ::pb::DNDirectMessageRequest();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::DNDirectMessageRequest::InitAsDefaultInstance();
}

void InitDefaultsDNDirectMessageRequest() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &InitDefaultsDNDirectMessageRequestImpl);
}

void InitDefaultsRPCDirectMessageImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  ::google::protobuf::internal::InitProtobufDefaultsForceUnique();
#else
  ::google::protobuf::internal::InitProtobufDefaults();
#endif  // GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  {
    void* ptr = &::pb::_RPCDirectMessage_default_instance_;
    new (ptr) ::pb::RPCDirectMessage();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::pb::RPCDirectMessage::InitAsDefaultInstance();
}

void InitDefaultsRPCDirectMessage() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &InitDefaultsRPCDirectMessageImpl);
}

::google::protobuf::Metadata file_level_metadata[3];

const ::google::protobuf::uint32 TableStruct::offsets[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::DNDirectMessageResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::DNDirectMessageResponse, delivered_),
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::DNDirectMessageRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::DNDirectMessageRequest, message_),
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCDirectMessage, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCDirectMessage, pid_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::pb::RPCDirectMessage, message_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::pb::DNDirectMessageResponse)},
  { 6, -1, sizeof(::pb::DNDirectMessageRequest)},
  { 12, -1, sizeof(::pb::RPCDirectMessage)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_DNDirectMessageResponse_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_DNDirectMessageRequest_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::pb::_RPCDirectMessage_default_instance_),
};

void protobuf_AssignDescriptors() {
  AddDescriptors();
  ::google::protobuf::MessageFactory* factory = NULL;
  AssignDescriptors(
      "pb/messaging.proto", schemas, file_default_instances, TableStruct::offsets, factory,
      file_level_metadata, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_PROTOBUF_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::internal::RegisterAllTypes(file_level_metadata, 3);
}

void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
      "\n\022pb/messaging.proto\022\002pb\",\n\027DNDirectMess"
      "ageResponse\022\021\n\tdelivered\030\001 \001(\010\")\n\026DNDire"
      "ctMessageRequest\022\017\n\007message\030\001 \001(\014\"0\n\020RPC"
      "DirectMessage\022\013\n\003pId\030\001 \001(\t\022\017\n\007message\030\002 "
      "\001(\014b\006proto3"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 171);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "pb/messaging.proto", &protobuf_RegisterTypes);
}

void AddDescriptors() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at dynamic initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;
}  // namespace protobuf_pb_2fmessaging_2eproto
namespace pb {

// ===================================================================

void DNDirectMessageResponse::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int DNDirectMessageResponse::kDeliveredFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

DNDirectMessageResponse::DNDirectMessageResponse()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageResponse();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.DNDirectMessageResponse)
}
DNDirectMessageResponse::DNDirectMessageResponse(const DNDirectMessageResponse& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  delivered_ = from.delivered_;
  // @@protoc_insertion_point(copy_constructor:pb.DNDirectMessageResponse)
}

void DNDirectMessageResponse::SharedCtor() {
  delivered_ = false;
  _cached_size_ = 0;
}

DNDirectMessageResponse::~DNDirectMessageResponse() {
  // @@protoc_insertion_point(destructor:pb.DNDirectMessageResponse)
  SharedDtor();
}

void DNDirectMessageResponse::SharedDtor() {
}

void DNDirectMessageResponse::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* DNDirectMessageResponse::descriptor() {
  ::protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const DNDirectMessageResponse& DNDirectMessageResponse::default_instance() {
  ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageResponse();
  return *internal_default_instance();
}

DNDirectMessageResponse* DNDirectMessageResponse::New(::google::protobuf::Arena* arena) const {
  DNDirectMessageResponse* n = new DNDirectMessageResponse;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void DNDirectMessageResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.DNDirectMessageResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  delivered_ = false;
  _internal_metadata_.Clear();
}

bool DNDirectMessageResponse::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.DNDirectMessageResponse)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // bool delivered = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(8u /* 8 & 0xFF */)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   bool, ::google::protobuf::internal::WireFormatLite::TYPE_BOOL>(
                 input, &delivered_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:pb.DNDirectMessageResponse)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.DNDirectMessageResponse)
  return false;
#undef DO_
}

void DNDirectMessageResponse::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.DNDirectMessageResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool delivered = 1;
  if (this->delivered() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteBool(1, this->delivered(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.DNDirectMessageResponse)
}

::google::protobuf::uint8* DNDirectMessageResponse::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.DNDirectMessageResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bool delivered = 1;
  if (this->delivered() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteBoolToArray(1, this->delivered(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.DNDirectMessageResponse)
  return target;
}

size_t DNDirectMessageResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.DNDirectMessageResponse)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // bool delivered = 1;
  if (this->delivered() != 0) {
    total_size += 1 + 1;
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void DNDirectMessageResponse::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.DNDirectMessageResponse)
  GOOGLE_DCHECK_NE(&from, this);
  const DNDirectMessageResponse* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const DNDirectMessageResponse>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.DNDirectMessageResponse)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.DNDirectMessageResponse)
    MergeFrom(*source);
  }
}

void DNDirectMessageResponse::MergeFrom(const DNDirectMessageResponse& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.DNDirectMessageResponse)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.delivered() != 0) {
    set_delivered(from.delivered());
  }
}

void DNDirectMessageResponse::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.DNDirectMessageResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void DNDirectMessageResponse::CopyFrom(const DNDirectMessageResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.DNDirectMessageResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool DNDirectMessageResponse::IsInitialized() const {
  return true;
}

void DNDirectMessageResponse::Swap(DNDirectMessageResponse* other) {
  if (other == this) return;
  InternalSwap(other);
}
void DNDirectMessageResponse::InternalSwap(DNDirectMessageResponse* other) {
  using std::swap;
  swap(delivered_, other->delivered_);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata DNDirectMessageResponse::GetMetadata() const {
  protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void DNDirectMessageRequest::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int DNDirectMessageRequest::kMessageFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

DNDirectMessageRequest::DNDirectMessageRequest()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageRequest();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.DNDirectMessageRequest)
}
DNDirectMessageRequest::DNDirectMessageRequest(const DNDirectMessageRequest& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.message().size() > 0) {
    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
  // @@protoc_insertion_point(copy_constructor:pb.DNDirectMessageRequest)
}

void DNDirectMessageRequest::SharedCtor() {
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _cached_size_ = 0;
}

DNDirectMessageRequest::~DNDirectMessageRequest() {
  // @@protoc_insertion_point(destructor:pb.DNDirectMessageRequest)
  SharedDtor();
}

void DNDirectMessageRequest::SharedDtor() {
  message_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void DNDirectMessageRequest::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* DNDirectMessageRequest::descriptor() {
  ::protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const DNDirectMessageRequest& DNDirectMessageRequest::default_instance() {
  ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageRequest();
  return *internal_default_instance();
}

DNDirectMessageRequest* DNDirectMessageRequest::New(::google::protobuf::Arena* arena) const {
  DNDirectMessageRequest* n = new DNDirectMessageRequest;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void DNDirectMessageRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.DNDirectMessageRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _internal_metadata_.Clear();
}

bool DNDirectMessageRequest::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.DNDirectMessageRequest)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // bytes message = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(10u /* 10 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_message()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:pb.DNDirectMessageRequest)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.DNDirectMessageRequest)
  return false;
#undef DO_
}

void DNDirectMessageRequest::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.DNDirectMessageRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bytes message = 1;
  if (this->message().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      1, this->message(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.DNDirectMessageRequest)
}

::google::protobuf::uint8* DNDirectMessageRequest::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.DNDirectMessageRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // bytes message = 1;
  if (this->message().size() > 0) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        1, this->message(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.DNDirectMessageRequest)
  return target;
}

size_t DNDirectMessageRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.DNDirectMessageRequest)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // bytes message = 1;
  if (this->message().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::BytesSize(
        this->message());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void DNDirectMessageRequest::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.DNDirectMessageRequest)
  GOOGLE_DCHECK_NE(&from, this);
  const DNDirectMessageRequest* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const DNDirectMessageRequest>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.DNDirectMessageRequest)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.DNDirectMessageRequest)
    MergeFrom(*source);
  }
}

void DNDirectMessageRequest::MergeFrom(const DNDirectMessageRequest& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.DNDirectMessageRequest)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.message().size() > 0) {

    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
}

void DNDirectMessageRequest::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.DNDirectMessageRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void DNDirectMessageRequest::CopyFrom(const DNDirectMessageRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.DNDirectMessageRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool DNDirectMessageRequest::IsInitialized() const {
  return true;
}

void DNDirectMessageRequest::Swap(DNDirectMessageRequest* other) {
  if (other == this) return;
  InternalSwap(other);
}
void DNDirectMessageRequest::InternalSwap(DNDirectMessageRequest* other) {
  using std::swap;
  message_.Swap(&other->message_);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata DNDirectMessageRequest::GetMetadata() const {
  protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void RPCDirectMessage::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int RPCDirectMessage::kPIdFieldNumber;
const int RPCDirectMessage::kMessageFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

RPCDirectMessage::RPCDirectMessage()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    ::protobuf_pb_2fmessaging_2eproto::InitDefaultsRPCDirectMessage();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:pb.RPCDirectMessage)
}
RPCDirectMessage::RPCDirectMessage(const RPCDirectMessage& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  pid_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.pid().size() > 0) {
    pid_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.pid_);
  }
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.message().size() > 0) {
    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
  // @@protoc_insertion_point(copy_constructor:pb.RPCDirectMessage)
}

void RPCDirectMessage::SharedCtor() {
  pid_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  message_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _cached_size_ = 0;
}

RPCDirectMessage::~RPCDirectMessage() {
  // @@protoc_insertion_point(destructor:pb.RPCDirectMessage)
  SharedDtor();
}

void RPCDirectMessage::SharedDtor() {
  pid_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  message_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void RPCDirectMessage::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* RPCDirectMessage::descriptor() {
  ::protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const RPCDirectMessage& RPCDirectMessage::default_instance() {
  ::protobuf_pb_2fmessaging_2eproto::InitDefaultsRPCDirectMessage();
  return *internal_default_instance();
}

RPCDirectMessage* RPCDirectMessage::New(::google::protobuf::Arena* arena) const {
  RPCDirectMessage* n = new RPCDirectMessage;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void RPCDirectMessage::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.RPCDirectMessage)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  pid_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _internal_metadata_.Clear();
}

bool RPCDirectMessage::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:pb.RPCDirectMessage)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // string pId = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(10u /* 10 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_pid()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->pid().data(), static_cast<int>(this->pid().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "pb.RPCDirectMessage.pId"));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // bytes message = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(18u /* 18 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_message()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:pb.RPCDirectMessage)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:pb.RPCDirectMessage)
  return false;
#undef DO_
}

void RPCDirectMessage::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:pb.RPCDirectMessage)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string pId = 1;
  if (this->pid().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->pid().data(), static_cast<int>(this->pid().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCDirectMessage.pId");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      1, this->pid(), output);
  }

  // bytes message = 2;
  if (this->message().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      2, this->message(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:pb.RPCDirectMessage)
}

::google::protobuf::uint8* RPCDirectMessage::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:pb.RPCDirectMessage)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string pId = 1;
  if (this->pid().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->pid().data(), static_cast<int>(this->pid().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "pb.RPCDirectMessage.pId");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        1, this->pid(), target);
  }

  // bytes message = 2;
  if (this->message().size() > 0) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        2, this->message(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.RPCDirectMessage)
  return target;
}

size_t RPCDirectMessage::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.RPCDirectMessage)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // string pId = 1;
  if (this->pid().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->pid());
  }

  // bytes message = 2;
  if (this->message().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::BytesSize(
        this->message());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void RPCDirectMessage::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:pb.RPCDirectMessage)
  GOOGLE_DCHECK_NE(&from, this);
  const RPCDirectMessage* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const RPCDirectMessage>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:pb.RPCDirectMessage)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:pb.RPCDirectMessage)
    MergeFrom(*source);
  }
}

void RPCDirectMessage::MergeFrom(const RPCDirectMessage& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.RPCDirectMessage)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.pid().size() > 0) {

    pid_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.pid_);
  }
  if (from.message().size() > 0) {

    message_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.message_);
  }
}

void RPCDirectMessage::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:pb.RPCDirectMessage)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void RPCDirectMessage::CopyFrom(const RPCDirectMessage& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.RPCDirectMessage)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool RPCDirectMessage::IsInitialized() const {
  return true;
}

void RPCDirectMessage::Swap(RPCDirectMessage* other) {
  if (other == this) return;
  InternalSwap(other);
}
void RPCDirectMessage::InternalSwap(RPCDirectMessage* other) {
  using std::swap;
  pid_.Swap(&other->pid_);
  message_.Swap(&other->message_);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata RPCDirectMessage::GetMetadata() const {
  protobuf_pb_2fmessaging_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_pb_2fmessaging_2eproto::file_level_metadata[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace pb

// @@protoc_insertion_point(global_scope)
