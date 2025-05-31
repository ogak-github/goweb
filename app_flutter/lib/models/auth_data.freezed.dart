// dart format width=80
// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'auth_data.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

// dart format off
T _$identity<T>(T value) => value;

/// @nodoc
mixin _$AuthData {

 String get token; String get expiredIn; UserData get userData;
/// Create a copy of AuthData
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$AuthDataCopyWith<AuthData> get copyWith => _$AuthDataCopyWithImpl<AuthData>(this as AuthData, _$identity);



@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is AuthData&&(identical(other.token, token) || other.token == token)&&(identical(other.expiredIn, expiredIn) || other.expiredIn == expiredIn)&&(identical(other.userData, userData) || other.userData == userData));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,token,expiredIn,userData);

@override
String toString() {
  return 'AuthData(token: $token, expiredIn: $expiredIn, userData: $userData)';
}


}

/// @nodoc
abstract mixin class $AuthDataCopyWith<$Res>  {
  factory $AuthDataCopyWith(AuthData value, $Res Function(AuthData) _then) = _$AuthDataCopyWithImpl;
@useResult
$Res call({
 String token, String expiredIn, UserData userData
});




}
/// @nodoc
class _$AuthDataCopyWithImpl<$Res>
    implements $AuthDataCopyWith<$Res> {
  _$AuthDataCopyWithImpl(this._self, this._then);

  final AuthData _self;
  final $Res Function(AuthData) _then;

/// Create a copy of AuthData
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? token = null,Object? expiredIn = null,Object? userData = null,}) {
  return _then(AuthData(
token: null == token ? _self.token : token // ignore: cast_nullable_to_non_nullable
as String,expiredIn: null == expiredIn ? _self.expiredIn : expiredIn // ignore: cast_nullable_to_non_nullable
as String,userData: null == userData ? _self.userData : userData // ignore: cast_nullable_to_non_nullable
as UserData,
  ));
}

}



/// @nodoc
mixin _$UserData {

 String get id; String get username; String get fullName; String get email;
/// Create a copy of UserData
/// with the given fields replaced by the non-null parameter values.
@JsonKey(includeFromJson: false, includeToJson: false)
@pragma('vm:prefer-inline')
$UserDataCopyWith<UserData> get copyWith => _$UserDataCopyWithImpl<UserData>(this as UserData, _$identity);



@override
bool operator ==(Object other) {
  return identical(this, other) || (other.runtimeType == runtimeType&&other is UserData&&(identical(other.id, id) || other.id == id)&&(identical(other.username, username) || other.username == username)&&(identical(other.fullName, fullName) || other.fullName == fullName)&&(identical(other.email, email) || other.email == email));
}

@JsonKey(includeFromJson: false, includeToJson: false)
@override
int get hashCode => Object.hash(runtimeType,id,username,fullName,email);

@override
String toString() {
  return 'UserData(id: $id, username: $username, fullName: $fullName, email: $email)';
}


}

/// @nodoc
abstract mixin class $UserDataCopyWith<$Res>  {
  factory $UserDataCopyWith(UserData value, $Res Function(UserData) _then) = _$UserDataCopyWithImpl;
@useResult
$Res call({
 String id, String username, String fullName, String email
});




}
/// @nodoc
class _$UserDataCopyWithImpl<$Res>
    implements $UserDataCopyWith<$Res> {
  _$UserDataCopyWithImpl(this._self, this._then);

  final UserData _self;
  final $Res Function(UserData) _then;

/// Create a copy of UserData
/// with the given fields replaced by the non-null parameter values.
@pragma('vm:prefer-inline') @override $Res call({Object? id = null,Object? username = null,Object? fullName = null,Object? email = null,}) {
  return _then(UserData(
id: null == id ? _self.id : id // ignore: cast_nullable_to_non_nullable
as String,username: null == username ? _self.username : username // ignore: cast_nullable_to_non_nullable
as String,fullName: null == fullName ? _self.fullName : fullName // ignore: cast_nullable_to_non_nullable
as String,email: null == email ? _self.email : email // ignore: cast_nullable_to_non_nullable
as String,
  ));
}

}


// dart format on
