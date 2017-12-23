// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: pb/platform.proto

#ifndef PROTOBUF_pb_2fplatform_2eproto__INCLUDED
#define PROTOBUF_pb_2fplatform_2eproto__INCLUDED

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
#include "pb/matchmaking.pb.h"
#include "pb/addressbook.pb.h"
#include "pb/storage.pb.h"
#include "pb/messaging.pb.h"
#include "pb/publisher.pb.h"
#include "google/api/annotations.pb.h"
// @@protoc_insertion_point(includes)

namespace protobuf_pb_2fplatform_2eproto {
// Internal implementation detail -- do not use these members.
struct TableStruct {
  static const ::google::protobuf::internal::ParseTableField entries[];
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[];
  static const ::google::protobuf::internal::ParseTable schema[3];
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors();
void InitDefaultsRPCHealthRequestImpl();
void InitDefaultsRPCHealthRequest();
void InitDefaultsRPCHealthReplyImpl();
void InitDefaultsRPCHealthReply();
void InitDefaultsemptyImpl();
void InitDefaultsempty();
inline void InitDefaults() {
  InitDefaultsRPCHealthRequest();
  InitDefaultsRPCHealthReply();
  InitDefaultsempty();
}
}  // namespace protobuf_pb_2fplatform_2eproto
namespace pb {
class RPCHealthReply;
class RPCHealthReplyDefaultTypeInternal;
extern RPCHealthReplyDefaultTypeInternal _RPCHealthReply_default_instance_;
class RPCHealthRequest;
class RPCHealthRequestDefaultTypeInternal;
extern RPCHealthRequestDefaultTypeInternal _RPCHealthRequest_default_instance_;
class empty;
class emptyDefaultTypeInternal;
extern emptyDefaultTypeInternal _empty_default_instance_;
}  // namespace pb
namespace pb {

// ===================================================================

class RPCHealthRequest : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.RPCHealthRequest) */ {
 public:
  RPCHealthRequest();
  virtual ~RPCHealthRequest();

  RPCHealthRequest(const RPCHealthRequest& from);

  inline RPCHealthRequest& operator=(const RPCHealthRequest& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  RPCHealthRequest(RPCHealthRequest&& from) noexcept
    : RPCHealthRequest() {
    *this = ::std::move(from);
  }

  inline RPCHealthRequest& operator=(RPCHealthRequest&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const RPCHealthRequest& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const RPCHealthRequest* internal_default_instance() {
    return reinterpret_cast<const RPCHealthRequest*>(
               &_RPCHealthRequest_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    0;

  void Swap(RPCHealthRequest* other);
  friend void swap(RPCHealthRequest& a, RPCHealthRequest& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline RPCHealthRequest* New() const PROTOBUF_FINAL { return New(NULL); }

  RPCHealthRequest* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const RPCHealthRequest& from);
  void MergeFrom(const RPCHealthRequest& from);
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
  void InternalSwap(RPCHealthRequest* other);
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

  // @@protoc_insertion_point(class_scope:pb.RPCHealthRequest)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fplatform_2eproto::TableStruct;
  friend void ::protobuf_pb_2fplatform_2eproto::InitDefaultsRPCHealthRequestImpl();
};
// -------------------------------------------------------------------

class RPCHealthReply : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.RPCHealthReply) */ {
 public:
  RPCHealthReply();
  virtual ~RPCHealthReply();

  RPCHealthReply(const RPCHealthReply& from);

  inline RPCHealthReply& operator=(const RPCHealthReply& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  RPCHealthReply(RPCHealthReply&& from) noexcept
    : RPCHealthReply() {
    *this = ::std::move(from);
  }

  inline RPCHealthReply& operator=(RPCHealthReply&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const RPCHealthReply& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const RPCHealthReply* internal_default_instance() {
    return reinterpret_cast<const RPCHealthReply*>(
               &_RPCHealthReply_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    1;

  void Swap(RPCHealthReply* other);
  friend void swap(RPCHealthReply& a, RPCHealthReply& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline RPCHealthReply* New() const PROTOBUF_FINAL { return New(NULL); }

  RPCHealthReply* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const RPCHealthReply& from);
  void MergeFrom(const RPCHealthReply& from);
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
  void InternalSwap(RPCHealthReply* other);
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

  // string message = 2;
  void clear_message();
  static const int kMessageFieldNumber = 2;
  const ::std::string& message() const;
  void set_message(const ::std::string& value);
  #if LANG_CXX11
  void set_message(::std::string&& value);
  #endif
  void set_message(const char* value);
  void set_message(const char* value, size_t size);
  ::std::string* mutable_message();
  ::std::string* release_message();
  void set_allocated_message(::std::string* message);

  // bool ready = 1;
  void clear_ready();
  static const int kReadyFieldNumber = 1;
  bool ready() const;
  void set_ready(bool value);

  // @@protoc_insertion_point(class_scope:pb.RPCHealthReply)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr message_;
  bool ready_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fplatform_2eproto::TableStruct;
  friend void ::protobuf_pb_2fplatform_2eproto::InitDefaultsRPCHealthReplyImpl();
};
// -------------------------------------------------------------------

class empty : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:pb.empty) */ {
 public:
  empty();
  virtual ~empty();

  empty(const empty& from);

  inline empty& operator=(const empty& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  empty(empty&& from) noexcept
    : empty() {
    *this = ::std::move(from);
  }

  inline empty& operator=(empty&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor();
  static const empty& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const empty* internal_default_instance() {
    return reinterpret_cast<const empty*>(
               &_empty_default_instance_);
  }
  static PROTOBUF_CONSTEXPR int const kIndexInFileMessages =
    2;

  void Swap(empty* other);
  friend void swap(empty& a, empty& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline empty* New() const PROTOBUF_FINAL { return New(NULL); }

  empty* New(::google::protobuf::Arena* arena) const PROTOBUF_FINAL;
  void CopyFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void MergeFrom(const ::google::protobuf::Message& from) PROTOBUF_FINAL;
  void CopyFrom(const empty& from);
  void MergeFrom(const empty& from);
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
  void InternalSwap(empty* other);
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

  // @@protoc_insertion_point(class_scope:pb.empty)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  mutable int _cached_size_;
  friend struct ::protobuf_pb_2fplatform_2eproto::TableStruct;
  friend void ::protobuf_pb_2fplatform_2eproto::InitDefaultsemptyImpl();
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// RPCHealthRequest

// -------------------------------------------------------------------

// RPCHealthReply

// bool ready = 1;
inline void RPCHealthReply::clear_ready() {
  ready_ = false;
}
inline bool RPCHealthReply::ready() const {
  // @@protoc_insertion_point(field_get:pb.RPCHealthReply.ready)
  return ready_;
}
inline void RPCHealthReply::set_ready(bool value) {
  
  ready_ = value;
  // @@protoc_insertion_point(field_set:pb.RPCHealthReply.ready)
}

// string message = 2;
inline void RPCHealthReply::clear_message() {
  message_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& RPCHealthReply::message() const {
  // @@protoc_insertion_point(field_get:pb.RPCHealthReply.message)
  return message_.GetNoArena();
}
inline void RPCHealthReply::set_message(const ::std::string& value) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:pb.RPCHealthReply.message)
}
#if LANG_CXX11
inline void RPCHealthReply::set_message(::std::string&& value) {
  
  message_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:pb.RPCHealthReply.message)
}
#endif
inline void RPCHealthReply::set_message(const char* value) {
  GOOGLE_DCHECK(value != NULL);
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:pb.RPCHealthReply.message)
}
inline void RPCHealthReply::set_message(const char* value, size_t size) {
  
  message_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:pb.RPCHealthReply.message)
}
inline ::std::string* RPCHealthReply::mutable_message() {
  
  // @@protoc_insertion_point(field_mutable:pb.RPCHealthReply.message)
  return message_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* RPCHealthReply::release_message() {
  // @@protoc_insertion_point(field_release:pb.RPCHealthReply.message)
  
  return message_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void RPCHealthReply::set_allocated_message(::std::string* message) {
  if (message != NULL) {
    
  } else {
    
  }
  message_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), message);
  // @@protoc_insertion_point(field_set_allocated:pb.RPCHealthReply.message)
}

// -------------------------------------------------------------------

// empty

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace pb

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_pb_2fplatform_2eproto__INCLUDED
