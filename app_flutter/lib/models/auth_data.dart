import 'package:freezed_annotation/freezed_annotation.dart';
part 'auth_data.freezed.dart';
part 'auth_data.g.dart';

@freezed
@JsonSerializable()
class AuthData with _$AuthData {
  @override
  final String token;
  @override
  @JsonKey(name: 'expired_in')
  final String expiredIn;
  @override
  @JsonKey(name: 'user_data')
  final UserData userData;

  AuthData({
    required this.token,
    required this.expiredIn,
    required this.userData,
  });

  factory AuthData.fromJson(Map<String, dynamic> json) =>
      _$AuthDataFromJson(json);

  Map<String, dynamic> toJson() => _$AuthDataToJson(this);
}

@freezed
@JsonSerializable()
class UserData with _$UserData {
  @override
  final String id;
  @override
  final String username;
  @override
  @JsonKey(name: 'full_name')
  final String fullName;
  @override
  final String email;

  UserData({
    required this.id,
    required this.username,
    required this.fullName,
    required this.email,
  });

  factory UserData.fromJson(Map<String, dynamic> json) =>
      _$UserDataFromJson(json);

  Map<String, dynamic> toJson() => _$UserDataToJson(this);
}
