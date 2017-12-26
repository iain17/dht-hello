// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: pb/messaging.proto

#ifndef PROTOBUF_pb_2fmessaging_2eproto__INCLUDED
#define PROTOBUF_pb_2fmessaging_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3005000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3005000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace protobuf_pb_2fmessaging_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[4];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
void InitDefaultsDNDirectMessageResponseImpl();
void InitDefaultsDNDirectMessageResponse();
void InitDefaultsDNDirectMessageRequestImpl();
void InitDefaultsDNDirectMessageRequest();
void InitDefaultsRPCReceiveDirectMessageRequestImpl();
void InitDefaultsRPCReceiveDirectMessageRequest();
void InitDefaultsRPCDirectMessageImpl();
void InitDefaultsRPCDirectMessage();
inline void InitDefaults() {
  InitDefaultsDNDirectMessageResponse();
  InitDefaultsDNDirectMessageRequest();
  InitDefaultsRPCReceiveDirectMessageRequest();
  InitDefaultsRPCDirectMessage();
}
}  // namespace protobuf_pb_2fmessaging_2eproto
namespace pb {
class DNDirectMessageRequest;
class DNDirectMessageRequestDefaultTypeInternal;
extern DNDirectMessageRequestDefaultTypeInternal _DNDirectMessageRequest_default_instance_;
class DNDirectMessageResponse;
class DNDirectMessageResponseDefaultTypeInternal;
extern DNDirectMessageResponseDefaultTypeInternal _DNDirectMessageResponse_default_instance_;
class RPCDirectMessage;
class RPCDirectMessageDefaultTypeInternal;
extern RPCDirectMessageDefaultTypeInternal _RPCDirectMessage_default_instance_;
class RPCReceiveDirectMessageRequest;
class RPCReceiveDirectMessageRequestDefaultTypeInternal;
extern RPCReceiveDirectMessageRequestDefaultTypeInternal _RPCReceiveDirectMessageRequest_default_instance_;
}  // namespace pb
namespace pb {

// ===================================================================

class DNDirectMessageResponse : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.DNDirectMessageResponse) */ {
 public:
  DNDirectMessageResponse();
  virtual ~DNDirectMessageResponse();

  DNDirectMessageResponse(const DNDirectMessageResponse& from);

  inline DNDirectMessageResponse& operator=(const DNDirectMessageResponse& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DNDirectMessageResponse(DNDirectMessageResponse&& from) noexcept
    : DNDirectMessageResponse() {
    *this = ::std::move(from);
  }

  inline DNDirectMessageResponse& operator=(DNDirectMessageResponse&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const DNDirectMessageResponse& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DNDirectMessageResponse* internal_default_instance() {
    return reinterpret_cast<const DNDirectMessageResponse*>(
               &_DNDirectMessageResponse_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    0;

  void Swap(DNDirectMessageResponse* other);
  friend void swap(DNDirectMessageResponse& a, DNDirectMessageResponse& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DNDirectMessageResponse* New() const PROTOBUF_FINAL { return New(NULL); }

  DNDirectMessageResponse* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const DNDirectMessageResponse& from);
  void MergeFrom(const DNDirectMessageResponse& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(DNDirectMessageResponse* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // bool delivered = 1;
  void clear_delivered();
  static const int kDeliveredFieldNumber = 1;
  bool delivered() const;
  void set_delivered(bool value);

  // @@protoc_insertion_point(class_scope:pb.DNDirectMessageResponse)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  bool delivered_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fmessaging_2eproto::TableStruct;
  friend void ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageResponseImpl();
};
// -------------------------------------------------------------------

class DNDirectMessageRequest : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.DNDirectMessageRequest) */ {
 public:
  DNDirectMessageRequest();
  virtual ~DNDirectMessageRequest();

  DNDirectMessageRequest(const DNDirectMessageRequest& from);

  inline DNDirectMessageRequest& operator=(const DNDirectMessageRequest& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DNDirectMessageRequest(DNDirectMessageRequest&& from) noexcept
    : DNDirectMessageRequest() {
    *this = ::std::move(from);
  }

  inline DNDirectMessageRequest& operator=(DNDirectMessageRequest&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const DNDirectMessageRequest& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DNDirectMessageRequest* internal_default_instance() {
    return reinterpret_cast<const DNDirectMessageRequest*>(
               &_DNDirectMessageRequest_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    1;

  void Swap(DNDirectMessageRequest* other);
  friend void swap(DNDirectMessageRequest& a, DNDirectMessageRequest& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DNDirectMessageRequest* New() const PROTOBUF_FINAL { return New(NULL); }

  DNDirectMessageRequest* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const DNDirectMessageRequest& from);
  void MergeFrom(const DNDirectMessageRequest& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(DNDirectMessageRequest* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // bytes message = 2;
  void clear_message();
  static const int kMessageFieldNumber = 2;
  const ::std::string& message() const;
  void set_message(const ::std::string& value);
  #if LANG_CXX11
  void set_message(::std::string&& value);
  #endif
  void set_message(const char* value);
  void set_message(const void* value, size_t size);
  ::std::string* mutable_message();
  ::std::string* release_message();
  void set_allocated_message(::std::string* message);

  // uint32 channel = 1;
  void clear_channel();
  static const int kChannelFieldNumber = 1;
  ::google::protobuf::uint32 channel() const;
  void set_channel(::google::protobuf::uint32 value);

  // @@protoc_insertion_point(class_scope:pb.DNDirectMessageRequest)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr message_;
  ::google::protobuf::uint32 channel_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fmessaging_2eproto::TableStruct;
  friend void ::protobuf_pb_2fmessaging_2eproto::InitDefaultsDNDirectMessageRequestImpl();
};
// -------------------------------------------------------------------

class RPCReceiveDirectMessageRequest : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.RPCReceiveDirectMessageRequest) */ {
 public:
  RPCReceiveDirectMessageRequest();
  virtual ~RPCReceiveDirectMessageRequest();

  RPCReceiveDirectMessageRequest(const RPCReceiveDirectMessageRequest& from);

  inline RPCReceiveDirectMessageRequest& operator=(const RPCReceiveDirectMessageRequest& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  RPCReceiveDirectMessageRequest(RPCReceiveDirectMessageRequest&& from) noexcept
    : RPCReceiveDirectMessageRequest() {
    *this = ::std::move(from);
  }

  inline RPCReceiveDirectMessageRequest& operator=(RPCReceiveDirectMessageRequest&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const RPCReceiveDirectMessageRequest& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const RPCReceiveDirectMessageRequest* internal_default_instance() {
    return reinterpret_cast<const RPCReceiveDirectMessageRequest*>(
               &_RPCReceiveDirectMessageRequest_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    2;

  void Swap(RPCReceiveDirectMessageRequest* other);
  friend void swap(RPCReceiveDirectMessageRequest& a, RPCReceiveDirectMessageRequest& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline RPCReceiveDirectMessageRequest* New() const PROTOBUF_FINAL { return New(NULL); }

  RPCReceiveDirectMessageRequest* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const RPCReceiveDirectMessageRequest& from);
  void MergeFrom(const RPCReceiveDirectMessageRequest& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(RPCReceiveDirectMessageRequest* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // uint32 channel = 1;
  void clear_channel();
  static const int kChannelFieldNumber = 1;
  ::google::protobuf::uint32 channel() const;
  void set_channel(::google::protobuf::uint32 value);

  // @@protoc_insertion_point(class_scope:pb.RPCReceiveDirectMessageRequest)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::uint32 channel_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fmessaging_2eproto::TableStruct;
  friend void ::protobuf_pb_2fmessaging_2eproto::InitDefaultsRPCReceiveDirectMessageRequestImpl();
};
// -------------------------------------------------------------------

class RPCDirectMessage : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.RPCDirectMessage) */ {
 public:
  RPCDirectMessage();
  virtual ~RPCDirectMessage();

  RPCDirectMessage(const RPCDirectMessage& from);

  inline RPCDirectMessage& operator=(const RPCDirectMessage& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  RPCDirectMessage(RPCDirectMessage&& from) noexcept
    : RPCDirectMessage() {
    *this = ::std::move(from);
  }

  inline RPCDirectMessage& operator=(RPCDirectMessage&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const RPCDirectMessage& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const RPCDirectMessage* internal_default_instance() {
    return reinterpret_cast<const RPCDirectMessage*>(
               &_RPCDirectMessage_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    3;

  void Swap(RPCDirectMessage* other);
  friend void swap(RPCDirectMessage& a, RPCDirectMessage& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline RPCDirectMessage* New() const PROTOBUF_FINAL { return New(NULL); }

  RPCDirectMessage* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const RPCDirectMessage& from);
  void MergeFrom(const RPCDirectMessage& from);
  void Clear() PROTOBUF_FINAL;
  bool IsInitialized() const PROTOBUF_FINAL;

  size_t ByteSizeLong() const PROTOBUF_FINAL;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) PROTOBUF_FINAL;
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const PROTOBUF_FINAL;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* target) const PROTOBUF_FINAL;
  int GetCachedSize() const PROTOBUF_FINAL { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const PROTOBUF_FINAL;
  void InternalSwap(RPCDirectMessage* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return NULL;
  }
  inline void* MaybeArenaPtr() const {
    return NULL;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const PROTOBUF_FINAL;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string pId = 2;
  void clear_pid();
  static const int kPIdFieldNumber = 2;
  const ::std::string& pid() const;
  void set_pid(const ::std::string& value);
  #if LANG_CXX11
  void set_pid(::std::string&& value);
  #endif
  void set_pid(const char* value);
  void set_pid(const char* value, size_t size);
  ::std::string* mutable_pid();
  ::std::string* release_pid();
  void set_allocated_pid(::std::string* pid);

  // bytes message = 3;
  void clear_message();
  static const int kMessageFieldNumber = 3;
  const ::std::string& message() const;
  void set_message(const ::std::string& value);
  #if LANG_CXX11
  void set_message(::std::string&& value);
  #endif
  void set_message(const char* value);
  void set_message(const void* value, size_t size);
  ::std::string* mutable_message();
  ::std::string* release_message();
  void set_allocated_message(::std::string* message);

  // uint32 channel = 1;
  void clear_channel();
  static const int kChannelFieldNumber = 1;
  ::google::protobuf::uint32 channel() const;
  void set_channel(::google::protobuf::uint32 value);

  // @@protoc_insertion_point(class_scope:pb.RPCDirectMessage)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr pid_;
  ::google::protobuf::internal::ArenaStringPtr message_;
  ::google::protobuf::uint32 channel_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fmessaging_2eproto::TableStruct;
  friend void ::protobuf_pb_2fmessaging_2eproto::InitDefaultsRPCDirectMessageImpl();
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// DNDirectMessageResponse

// bool delivered = 1;
inline void DNDirectMessageResponse::clear_delivered() {
  delivered_ = false;
}
inline bool DNDirectMessageResponse::delivered() const {
  // @@protoc_insertion_point(field_get:pb.DNDirectMessageResponse.delivered)
  return delivered_;
}
inline void DNDirectMessageResponse::set_delivered(bool value) {
  
  delivered_ = value;
  // @@protoc_insertion_point(field_set:pb.DNDirectMessageResponse.delivered)
}

// -------------------------------------------------------------------

// DNDirectMessageRequest

// uint32 channel = 1;
inline void DNDirectMessageRequest::clear_channel() {
  channel_ = 0u;
}
inline ::google::protobuf::uint32 DNDirectMessageRequest::channel() const {
  // @@protoc_insertion_point(field_get:pb.DNDirectMessageRequest.channel)
  return channel_;
}
inline void DNDirectMessageRequest::set_channel(::google::protobuf::uint32 value) {
  
  channel_ = value;
  // @@protoc_insertion_point(field_set:pb.DNDirectMessageRequest.channel)
}

// bytes message = 2;
inline void DNDirectMessageRequest::clear_message() {
  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& DNDirectMessageRequest::message() const {
  // @@protoc_insertion_point(field_get:pb.DNDirectMessageRequest.message)
  return message_.GetNoArena();
}
inline void DNDirectMessageRequest::set_message(const ::std::string& value) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:pb.DNDirectMessageRequest.message)
}
#if LANG_CXX11
inline void DNDirectMessageRequest::set_message(::std::string&& value) {
  
  message_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:pb.DNDirectMessageRequest.message)
}
#endif
inline void DNDirectMessageRequest::set_message(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:pb.DNDirectMessageRequest.message)
}
inline void DNDirectMessageRequest::set_message(const void* value, size_t size) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:pb.DNDirectMessageRequest.message)
}
inline ::std::string* DNDirectMessageRequest::mutable_message() {
  
  // @@protoc_insertion_point(field_mutable:pb.DNDirectMessageRequest.message)
  return message_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* DNDirectMessageRequest::release_message() {
  // @@protoc_insertion_point(field_release:pb.DNDirectMessageRequest.message)
  
  return message_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void DNDirectMessageRequest::set_allocated_message(::std::string* message) {
  if (message != NULL) {
    
  } else {
    
  }
  message_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), message);
  // @@protoc_insertion_point(field_set_allocated:pb.DNDirectMessageRequest.message)
}

// -------------------------------------------------------------------

// RPCReceiveDirectMessageRequest

// uint32 channel = 1;
inline void RPCReceiveDirectMessageRequest::clear_channel() {
  channel_ = 0u;
}
inline ::google::protobuf::uint32 RPCReceiveDirectMessageRequest::channel() const {
  // @@protoc_insertion_point(field_get:pb.RPCReceiveDirectMessageRequest.channel)
  return channel_;
}
inline void RPCReceiveDirectMessageRequest::set_channel(::google::protobuf::uint32 value) {
  
  channel_ = value;
  // @@protoc_insertion_point(field_set:pb.RPCReceiveDirectMessageRequest.channel)
}

// -------------------------------------------------------------------

// RPCDirectMessage

// uint32 channel = 1;
inline void RPCDirectMessage::clear_channel() {
  channel_ = 0u;
}
inline ::google::protobuf::uint32 RPCDirectMessage::channel() const {
  // @@protoc_insertion_point(field_get:pb.RPCDirectMessage.channel)
  return channel_;
}
inline void RPCDirectMessage::set_channel(::google::protobuf::uint32 value) {
  
  channel_ = value;
  // @@protoc_insertion_point(field_set:pb.RPCDirectMessage.channel)
}

// string pId = 2;
inline void RPCDirectMessage::clear_pid() {
  pid_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& RPCDirectMessage::pid() const {
  // @@protoc_insertion_point(field_get:pb.RPCDirectMessage.pId)
  return pid_.GetNoArena();
}
inline void RPCDirectMessage::set_pid(const ::std::string& value) {
  
  pid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:pb.RPCDirectMessage.pId)
}
#if LANG_CXX11
inline void RPCDirectMessage::set_pid(::std::string&& value) {
  
  pid_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:pb.RPCDirectMessage.pId)
}
#endif
inline void RPCDirectMessage::set_pid(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  pid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:pb.RPCDirectMessage.pId)
}
inline void RPCDirectMessage::set_pid(const char* value, size_t size) {
  
  pid_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:pb.RPCDirectMessage.pId)
}
inline ::std::string* RPCDirectMessage::mutable_pid() {
  
  // @@protoc_insertion_point(field_mutable:pb.RPCDirectMessage.pId)
  return pid_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* RPCDirectMessage::release_pid() {
  // @@protoc_insertion_point(field_release:pb.RPCDirectMessage.pId)
  
  return pid_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void RPCDirectMessage::set_allocated_pid(::std::string* pid) {
  if (pid != NULL) {
    
  } else {
    
  }
  pid_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), pid);
  // @@protoc_insertion_point(field_set_allocated:pb.RPCDirectMessage.pId)
}

// bytes message = 3;
inline void RPCDirectMessage::clear_message() {
  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& RPCDirectMessage::message() const {
  // @@protoc_insertion_point(field_get:pb.RPCDirectMessage.message)
  return message_.GetNoArena();
}
inline void RPCDirectMessage::set_message(const ::std::string& value) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:pb.RPCDirectMessage.message)
}
#if LANG_CXX11
inline void RPCDirectMessage::set_message(::std::string&& value) {
  
  message_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:pb.RPCDirectMessage.message)
}
#endif
inline void RPCDirectMessage::set_message(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:pb.RPCDirectMessage.message)
}
inline void RPCDirectMessage::set_message(const void* value, size_t size) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:pb.RPCDirectMessage.message)
}
inline ::std::string* RPCDirectMessage::mutable_message() {
  
  // @@protoc_insertion_point(field_mutable:pb.RPCDirectMessage.message)
  return message_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* RPCDirectMessage::release_message() {
  // @@protoc_insertion_point(field_release:pb.RPCDirectMessage.message)
  
  return message_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void RPCDirectMessage::set_allocated_message(::std::string* message) {
  if (message != NULL) {
    
  } else {
    
  }
  message_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), message);
  // @@protoc_insertion_point(field_set_allocated:pb.RPCDirectMessage.message)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace pb

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_pb_2fmessaging_2eproto__INCLUDED
