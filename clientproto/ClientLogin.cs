// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: client-login.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
/// <summary>Holder for reflection information generated from client-login.proto</summary>
public static partial class ClientLoginReflection {

  #region Descriptor
  /// <summary>File descriptor for client-login.proto</summary>
  public static pbr::FileDescriptor Descriptor {
    get { return descriptor; }
  }
  private static pbr::FileDescriptor descriptor;

  static ClientLoginReflection() {
    byte[] descriptorData = global::System.Convert.FromBase64String(
        string.Concat(
          "ChJjbGllbnQtbG9naW4ucHJvdG8aEWNsaWVudC1nYXRlLnByb3RvIi0KB0NM",
          "TG9nb24SDAoETmFtZRgBIAEoCRIUCgxQYXNzd29yZEhhc2gYAiABKAkiLQoH",
          "Q0xMb2dpbhIMCgROYW1lGAEgASgJEhQKDFBhc3N3b3JkSGFzaBgCIAEoCSJS",
          "ChBDTENoYW5nZVBhc3N3b3JkEgwKBE5hbWUYASABKAkSFwoPT2xkUGFzc3dv",
          "cmRIYXNoGAIgASgJEhcKD05ld1Bhc3N3b3JkSGFzaBgDIAEoCSJsCg1DbGll",
          "bnRUb0xvZ2luEhcKBUxvZ29uGAEgASgLMgguQ0xMb2dvbhIXCgVMb2dpbhgC",
          "IAEoCzIILkNMTG9naW4SKQoOQ2hhbmdlUGFzc3dvcmQYAyABKAsyES5DTENo",
          "YW5nZVBhc3N3b3JkImUKDUxvZ2luVG9DbGllbnQSDwoHU3VjY2VzcxgBIAEo",
          "CBIhCgtTZXNzaW9uQ2VydBgCIAEoCzIMLlNlc3Npb25DZXJ0Eg4KBkdhdGVJ",
          "UBgDIAEoDBIQCghHYXRlUG9ydBgEIAEoBUIYWhZJTlNlcnZlci9zcmMvcHJv",
          "dG8vbXNnYgZwcm90bzM="));
    descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
        new pbr::FileDescriptor[] { global::ClientGateReflection.Descriptor, },
        new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
          new pbr::GeneratedClrTypeInfo(typeof(global::CLLogon), global::CLLogon.Parser, new[]{ "Name", "PasswordHash" }, null, null, null, null),
          new pbr::GeneratedClrTypeInfo(typeof(global::CLLogin), global::CLLogin.Parser, new[]{ "Name", "PasswordHash" }, null, null, null, null),
          new pbr::GeneratedClrTypeInfo(typeof(global::CLChangePassword), global::CLChangePassword.Parser, new[]{ "Name", "OldPasswordHash", "NewPasswordHash" }, null, null, null, null),
          new pbr::GeneratedClrTypeInfo(typeof(global::ClientToLogin), global::ClientToLogin.Parser, new[]{ "Logon", "Login", "ChangePassword" }, null, null, null, null),
          new pbr::GeneratedClrTypeInfo(typeof(global::LoginToClient), global::LoginToClient.Parser, new[]{ "Success", "SessionCert", "GateIP", "GatePort" }, null, null, null, null)
        }));
  }
  #endregion

}
#region Messages
public sealed partial class CLLogon : pb::IMessage<CLLogon> {
  private static readonly pb::MessageParser<CLLogon> _parser = new pb::MessageParser<CLLogon>(() => new CLLogon());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<CLLogon> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::ClientLoginReflection.Descriptor.MessageTypes[0]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogon() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogon(CLLogon other) : this() {
    name_ = other.name_;
    passwordHash_ = other.passwordHash_;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogon Clone() {
    return new CLLogon(this);
  }

  /// <summary>Field number for the "Name" field.</summary>
  public const int NameFieldNumber = 1;
  private string name_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Name {
    get { return name_; }
    set {
      name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "PasswordHash" field.</summary>
  public const int PasswordHashFieldNumber = 2;
  private string passwordHash_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string PasswordHash {
    get { return passwordHash_; }
    set {
      passwordHash_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as CLLogon);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(CLLogon other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (Name != other.Name) return false;
    if (PasswordHash != other.PasswordHash) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (Name.Length != 0) hash ^= Name.GetHashCode();
    if (PasswordHash.Length != 0) hash ^= PasswordHash.GetHashCode();
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
    if (Name.Length != 0) {
      output.WriteRawTag(10);
      output.WriteString(Name);
    }
    if (PasswordHash.Length != 0) {
      output.WriteRawTag(18);
      output.WriteString(PasswordHash);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (Name.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
    }
    if (PasswordHash.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(PasswordHash);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(CLLogon other) {
    if (other == null) {
      return;
    }
    if (other.Name.Length != 0) {
      Name = other.Name;
    }
    if (other.PasswordHash.Length != 0) {
      PasswordHash = other.PasswordHash;
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
          Name = input.ReadString();
          break;
        }
        case 18: {
          PasswordHash = input.ReadString();
          break;
        }
      }
    }
  }

}

public sealed partial class CLLogin : pb::IMessage<CLLogin> {
  private static readonly pb::MessageParser<CLLogin> _parser = new pb::MessageParser<CLLogin>(() => new CLLogin());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<CLLogin> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::ClientLoginReflection.Descriptor.MessageTypes[1]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogin() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogin(CLLogin other) : this() {
    name_ = other.name_;
    passwordHash_ = other.passwordHash_;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLLogin Clone() {
    return new CLLogin(this);
  }

  /// <summary>Field number for the "Name" field.</summary>
  public const int NameFieldNumber = 1;
  private string name_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Name {
    get { return name_; }
    set {
      name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "PasswordHash" field.</summary>
  public const int PasswordHashFieldNumber = 2;
  private string passwordHash_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string PasswordHash {
    get { return passwordHash_; }
    set {
      passwordHash_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as CLLogin);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(CLLogin other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (Name != other.Name) return false;
    if (PasswordHash != other.PasswordHash) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (Name.Length != 0) hash ^= Name.GetHashCode();
    if (PasswordHash.Length != 0) hash ^= PasswordHash.GetHashCode();
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
    if (Name.Length != 0) {
      output.WriteRawTag(10);
      output.WriteString(Name);
    }
    if (PasswordHash.Length != 0) {
      output.WriteRawTag(18);
      output.WriteString(PasswordHash);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (Name.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
    }
    if (PasswordHash.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(PasswordHash);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(CLLogin other) {
    if (other == null) {
      return;
    }
    if (other.Name.Length != 0) {
      Name = other.Name;
    }
    if (other.PasswordHash.Length != 0) {
      PasswordHash = other.PasswordHash;
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
          Name = input.ReadString();
          break;
        }
        case 18: {
          PasswordHash = input.ReadString();
          break;
        }
      }
    }
  }

}

public sealed partial class CLChangePassword : pb::IMessage<CLChangePassword> {
  private static readonly pb::MessageParser<CLChangePassword> _parser = new pb::MessageParser<CLChangePassword>(() => new CLChangePassword());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<CLChangePassword> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::ClientLoginReflection.Descriptor.MessageTypes[2]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLChangePassword() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLChangePassword(CLChangePassword other) : this() {
    name_ = other.name_;
    oldPasswordHash_ = other.oldPasswordHash_;
    newPasswordHash_ = other.newPasswordHash_;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public CLChangePassword Clone() {
    return new CLChangePassword(this);
  }

  /// <summary>Field number for the "Name" field.</summary>
  public const int NameFieldNumber = 1;
  private string name_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string Name {
    get { return name_; }
    set {
      name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "OldPasswordHash" field.</summary>
  public const int OldPasswordHashFieldNumber = 2;
  private string oldPasswordHash_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string OldPasswordHash {
    get { return oldPasswordHash_; }
    set {
      oldPasswordHash_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "NewPasswordHash" field.</summary>
  public const int NewPasswordHashFieldNumber = 3;
  private string newPasswordHash_ = "";
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public string NewPasswordHash {
    get { return newPasswordHash_; }
    set {
      newPasswordHash_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as CLChangePassword);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(CLChangePassword other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (Name != other.Name) return false;
    if (OldPasswordHash != other.OldPasswordHash) return false;
    if (NewPasswordHash != other.NewPasswordHash) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (Name.Length != 0) hash ^= Name.GetHashCode();
    if (OldPasswordHash.Length != 0) hash ^= OldPasswordHash.GetHashCode();
    if (NewPasswordHash.Length != 0) hash ^= NewPasswordHash.GetHashCode();
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
    if (Name.Length != 0) {
      output.WriteRawTag(10);
      output.WriteString(Name);
    }
    if (OldPasswordHash.Length != 0) {
      output.WriteRawTag(18);
      output.WriteString(OldPasswordHash);
    }
    if (NewPasswordHash.Length != 0) {
      output.WriteRawTag(26);
      output.WriteString(NewPasswordHash);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (Name.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
    }
    if (OldPasswordHash.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(OldPasswordHash);
    }
    if (NewPasswordHash.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeStringSize(NewPasswordHash);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(CLChangePassword other) {
    if (other == null) {
      return;
    }
    if (other.Name.Length != 0) {
      Name = other.Name;
    }
    if (other.OldPasswordHash.Length != 0) {
      OldPasswordHash = other.OldPasswordHash;
    }
    if (other.NewPasswordHash.Length != 0) {
      NewPasswordHash = other.NewPasswordHash;
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
          Name = input.ReadString();
          break;
        }
        case 18: {
          OldPasswordHash = input.ReadString();
          break;
        }
        case 26: {
          NewPasswordHash = input.ReadString();
          break;
        }
      }
    }
  }

}

public sealed partial class ClientToLogin : pb::IMessage<ClientToLogin> {
  private static readonly pb::MessageParser<ClientToLogin> _parser = new pb::MessageParser<ClientToLogin>(() => new ClientToLogin());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<ClientToLogin> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::ClientLoginReflection.Descriptor.MessageTypes[3]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public ClientToLogin() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public ClientToLogin(ClientToLogin other) : this() {
    logon_ = other.logon_ != null ? other.logon_.Clone() : null;
    login_ = other.login_ != null ? other.login_.Clone() : null;
    changePassword_ = other.changePassword_ != null ? other.changePassword_.Clone() : null;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public ClientToLogin Clone() {
    return new ClientToLogin(this);
  }

  /// <summary>Field number for the "Logon" field.</summary>
  public const int LogonFieldNumber = 1;
  private global::CLLogon logon_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public global::CLLogon Logon {
    get { return logon_; }
    set {
      logon_ = value;
    }
  }

  /// <summary>Field number for the "Login" field.</summary>
  public const int LoginFieldNumber = 2;
  private global::CLLogin login_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public global::CLLogin Login {
    get { return login_; }
    set {
      login_ = value;
    }
  }

  /// <summary>Field number for the "ChangePassword" field.</summary>
  public const int ChangePasswordFieldNumber = 3;
  private global::CLChangePassword changePassword_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public global::CLChangePassword ChangePassword {
    get { return changePassword_; }
    set {
      changePassword_ = value;
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as ClientToLogin);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(ClientToLogin other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (!object.Equals(Logon, other.Logon)) return false;
    if (!object.Equals(Login, other.Login)) return false;
    if (!object.Equals(ChangePassword, other.ChangePassword)) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (logon_ != null) hash ^= Logon.GetHashCode();
    if (login_ != null) hash ^= Login.GetHashCode();
    if (changePassword_ != null) hash ^= ChangePassword.GetHashCode();
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
    if (logon_ != null) {
      output.WriteRawTag(10);
      output.WriteMessage(Logon);
    }
    if (login_ != null) {
      output.WriteRawTag(18);
      output.WriteMessage(Login);
    }
    if (changePassword_ != null) {
      output.WriteRawTag(26);
      output.WriteMessage(ChangePassword);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (logon_ != null) {
      size += 1 + pb::CodedOutputStream.ComputeMessageSize(Logon);
    }
    if (login_ != null) {
      size += 1 + pb::CodedOutputStream.ComputeMessageSize(Login);
    }
    if (changePassword_ != null) {
      size += 1 + pb::CodedOutputStream.ComputeMessageSize(ChangePassword);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(ClientToLogin other) {
    if (other == null) {
      return;
    }
    if (other.logon_ != null) {
      if (logon_ == null) {
        Logon = new global::CLLogon();
      }
      Logon.MergeFrom(other.Logon);
    }
    if (other.login_ != null) {
      if (login_ == null) {
        Login = new global::CLLogin();
      }
      Login.MergeFrom(other.Login);
    }
    if (other.changePassword_ != null) {
      if (changePassword_ == null) {
        ChangePassword = new global::CLChangePassword();
      }
      ChangePassword.MergeFrom(other.ChangePassword);
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
          if (logon_ == null) {
            Logon = new global::CLLogon();
          }
          input.ReadMessage(Logon);
          break;
        }
        case 18: {
          if (login_ == null) {
            Login = new global::CLLogin();
          }
          input.ReadMessage(Login);
          break;
        }
        case 26: {
          if (changePassword_ == null) {
            ChangePassword = new global::CLChangePassword();
          }
          input.ReadMessage(ChangePassword);
          break;
        }
      }
    }
  }

}

public sealed partial class LoginToClient : pb::IMessage<LoginToClient> {
  private static readonly pb::MessageParser<LoginToClient> _parser = new pb::MessageParser<LoginToClient>(() => new LoginToClient());
  private pb::UnknownFieldSet _unknownFields;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pb::MessageParser<LoginToClient> Parser { get { return _parser; } }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public static pbr::MessageDescriptor Descriptor {
    get { return global::ClientLoginReflection.Descriptor.MessageTypes[4]; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  pbr::MessageDescriptor pb::IMessage.Descriptor {
    get { return Descriptor; }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public LoginToClient() {
    OnConstruction();
  }

  partial void OnConstruction();

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public LoginToClient(LoginToClient other) : this() {
    success_ = other.success_;
    sessionCert_ = other.sessionCert_ != null ? other.sessionCert_.Clone() : null;
    gateIP_ = other.gateIP_;
    gatePort_ = other.gatePort_;
    _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public LoginToClient Clone() {
    return new LoginToClient(this);
  }

  /// <summary>Field number for the "Success" field.</summary>
  public const int SuccessFieldNumber = 1;
  private bool success_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Success {
    get { return success_; }
    set {
      success_ = value;
    }
  }

  /// <summary>Field number for the "SessionCert" field.</summary>
  public const int SessionCertFieldNumber = 2;
  private global::SessionCert sessionCert_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public global::SessionCert SessionCert {
    get { return sessionCert_; }
    set {
      sessionCert_ = value;
    }
  }

  /// <summary>Field number for the "GateIP" field.</summary>
  public const int GateIPFieldNumber = 3;
  private pb::ByteString gateIP_ = pb::ByteString.Empty;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public pb::ByteString GateIP {
    get { return gateIP_; }
    set {
      gateIP_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
    }
  }

  /// <summary>Field number for the "GatePort" field.</summary>
  public const int GatePortFieldNumber = 4;
  private int gatePort_;
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int GatePort {
    get { return gatePort_; }
    set {
      gatePort_ = value;
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override bool Equals(object other) {
    return Equals(other as LoginToClient);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public bool Equals(LoginToClient other) {
    if (ReferenceEquals(other, null)) {
      return false;
    }
    if (ReferenceEquals(other, this)) {
      return true;
    }
    if (Success != other.Success) return false;
    if (!object.Equals(SessionCert, other.SessionCert)) return false;
    if (GateIP != other.GateIP) return false;
    if (GatePort != other.GatePort) return false;
    return Equals(_unknownFields, other._unknownFields);
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public override int GetHashCode() {
    int hash = 1;
    if (Success != false) hash ^= Success.GetHashCode();
    if (sessionCert_ != null) hash ^= SessionCert.GetHashCode();
    if (GateIP.Length != 0) hash ^= GateIP.GetHashCode();
    if (GatePort != 0) hash ^= GatePort.GetHashCode();
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
    if (Success != false) {
      output.WriteRawTag(8);
      output.WriteBool(Success);
    }
    if (sessionCert_ != null) {
      output.WriteRawTag(18);
      output.WriteMessage(SessionCert);
    }
    if (GateIP.Length != 0) {
      output.WriteRawTag(26);
      output.WriteBytes(GateIP);
    }
    if (GatePort != 0) {
      output.WriteRawTag(32);
      output.WriteInt32(GatePort);
    }
    if (_unknownFields != null) {
      _unknownFields.WriteTo(output);
    }
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public int CalculateSize() {
    int size = 0;
    if (Success != false) {
      size += 1 + 1;
    }
    if (sessionCert_ != null) {
      size += 1 + pb::CodedOutputStream.ComputeMessageSize(SessionCert);
    }
    if (GateIP.Length != 0) {
      size += 1 + pb::CodedOutputStream.ComputeBytesSize(GateIP);
    }
    if (GatePort != 0) {
      size += 1 + pb::CodedOutputStream.ComputeInt32Size(GatePort);
    }
    if (_unknownFields != null) {
      size += _unknownFields.CalculateSize();
    }
    return size;
  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
  public void MergeFrom(LoginToClient other) {
    if (other == null) {
      return;
    }
    if (other.Success != false) {
      Success = other.Success;
    }
    if (other.sessionCert_ != null) {
      if (sessionCert_ == null) {
        SessionCert = new global::SessionCert();
      }
      SessionCert.MergeFrom(other.SessionCert);
    }
    if (other.GateIP.Length != 0) {
      GateIP = other.GateIP;
    }
    if (other.GatePort != 0) {
      GatePort = other.GatePort;
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
        case 8: {
          Success = input.ReadBool();
          break;
        }
        case 18: {
          if (sessionCert_ == null) {
            SessionCert = new global::SessionCert();
          }
          input.ReadMessage(SessionCert);
          break;
        }
        case 26: {
          GateIP = input.ReadBytes();
          break;
        }
        case 32: {
          GatePort = input.ReadInt32();
          break;
        }
      }
    }
  }

}

#endregion


#endregion Designer generated code
