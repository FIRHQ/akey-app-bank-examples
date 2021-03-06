// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: global.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
/// <summary>Holder for reflection information generated from global.proto</summary>
public static partial class GlobalReflection {

  #region Descriptor
  /// <summary>File descriptor for global.proto</summary>
  public static pbr::FileDescriptor Descriptor {
    get { return descriptor; }
  }
  private static pbr::FileDescriptor descriptor;

  static GlobalReflection() {
    byte[] descriptorData = global::System.Convert.FromBase64String(
        string.Concat(
          "CgxnbG9iYWwucHJvdG8aG2dvb2dsZS9wcm90b2J1Zi9lbXB0eS5wcm90byJa",
          "CgtWZXJzaW9uSW5mbxIPCgd2ZXJzaW9uGAEgASgJEg8KB3NlcnZpY2UYAiAB",
          "KAkSEwoLZGVzY3JpcHRpb24YAyABKAkSFAoMYWNjZXNzX3JpZ2h0GAQgASgJ",
          "QloKHmNvbS5ha2V5Lm9wZW4uYmFzZS5ncnBjLmdsb2JhbEILR2xvYmFsUHJv",
          "dG9QAVoiYWtleS1hcHAtYmFuay1leGFtcGxlcy9ncnBjL2dsb2JhbKICBEFL",
          "ZXliBnByb3RvMw=="));
    descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
        new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.EmptyReflection.Descriptor, },
        new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
          new pbr::GeneratedClrTypeInfo(typeof(global::VersionInfo), global::VersionInfo.Parser, new[]{ "Version", "Service", "Description", "AccessRight" }, null, null, null)
        }));
  }
  #endregion

}
#region Messages
public sealed partial class VersionInfo : pb::IMessage<VersionInfo> {
  private static readonly pb::MessageParser<VersionInfo> _parser = new pb::MessageParser<VersionInfo>(() => new VersionInfo());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<VersionInfo> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::GlobalReflection.Descriptor.MessageTypes[0]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public VersionInfo() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public VersionInfo(VersionInfo other) : this() {
    version_ = other.version_;
    service_ = other.service_;
    description_ = other.description_;
    accessRight_ = other.accessRight_;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public VersionInfo Clone() {
    return new VersionInfo(this);
  }

  /// <summary>Field number for the "version" field.</summary>
  public const int VersionFieldNumber = 1;
  private string version_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Version {
    get { return version_; }
    set {
      version_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "service" field.</summary>
  public const int ServiceFieldNumber = 2;
  private string service_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Service {
    get { return service_; }
    set {
      service_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "description" field.</summary>
  public const int DescriptionFieldNumber = 3;
  private string description_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Description {
    get { return description_; }
    set {
      description_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "access_right" field.</summary>
  public const int AccessRightFieldNumber = 4;
  private string accessRight_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string AccessRight {
    get { return accessRight_; }
    set {
      accessRight_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as VersionInfo);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(VersionInfo other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (Version != other.Version) return false;
    if (Service != other.Service) return false;
    if (Description != other.Description) return false;
    if (AccessRight != other.AccessRight) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (Version.Length != 0) hash ^= Version.GetHashCode();
    if (Service.Length != 0) hash ^= Service.GetHashCode();
    if (Description.Length != 0) hash ^= Description.GetHashCode();
    if (AccessRight.Length != 0) hash ^= AccessRight.GetHashCode();
    if (_unknownFields != null) {
      hash ^= _unknownFields.GetHashCode();
    }
    return hash;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override string ToString() {
    return pb::JsonFormatter.ToDiagnosticString(this);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void WriteTo(pb::CodedOutputStream output) {
    if (Version.Length != 0) {
      output.WriteRawTag(10);
      output.WriteString(Version);
    }
    if (Service.Length != 0) {
      output.WriteRawTag(18);
      output.WriteString(Service);
    }
    if (Description.Length != 0) {
      output.WriteRawTag(26);
      output.WriteString(Description);
    }
    if (AccessRight.Length != 0) {
      output.WriteRawTag(34);
      output.WriteString(AccessRight);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (Version.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Version);
    }
    if (Service.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Service);
    }
    if (Description.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Description);
    }
    if (AccessRight.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(AccessRight);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(VersionInfo other) {
    if (other == null) {
      return;
    }
    if (other.Version.Length != 0) {
      Version = other.Version;
    }
    if (other.Service.Length != 0) {
      Service = other.Service;
    }
    if (other.Description.Length != 0) {
      Description = other.Description;
    }
    if (other.AccessRight.Length != 0) {
      AccessRight = other.AccessRight;
    }
    _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(pb::CodedInputStream input) {
    uint tag;
    while ((tag = input.ReadTag()) != 0) {
      switch(tag) {
        default:
          _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
          break;
        case 10: {
          Version = input.ReadString();
          break;
        }
        case 18: {
          Service = input.ReadString();
          break;
        }
        case 26: {
          Description = input.ReadString();
          break;
        }
        case 34: {
          AccessRight = input.ReadString();
          break;
        }
      }
    }
  }

}

#endregion


#endregion Designer generated code
