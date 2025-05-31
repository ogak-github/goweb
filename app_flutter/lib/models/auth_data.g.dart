// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'auth_data.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

AuthData _$AuthDataFromJson(Map<String, dynamic> json) => AuthData(
  token: json['token'] as String,
  expiredIn: json['expired_in'] as String,
  userData: UserData.fromJson(json['user_data'] as Map<String, dynamic>),
);

Map<String, dynamic> _$AuthDataToJson(AuthData instance) => <String, dynamic>{
  'token': instance.token,
  'expired_in': instance.expiredIn,
  'user_data': instance.userData,
};

UserData _$UserDataFromJson(Map<String, dynamic> json) => UserData(
  id: json['id'] as String,
  username: json['username'] as String,
  fullName: json['full_name'] as String,
  email: json['email'] as String,
);

Map<String, dynamic> _$UserDataToJson(UserData instance) => <String, dynamic>{
  'id': instance.id,
  'username': instance.username,
  'full_name': instance.fullName,
  'email': instance.email,
};
