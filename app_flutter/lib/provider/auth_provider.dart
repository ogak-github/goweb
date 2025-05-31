import 'dart:convert';
import 'dart:developer';

import 'package:app_flutter/models/auth_data.dart';
import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:shared_preferences/shared_preferences.dart';

part 'auth_provider.g.dart';

final baseUrl = 'http://10.0.2.2:9001';

@riverpod
class Auth extends _$Auth {
  @override
  Future<AuthData?> build() async {
    final sf = await SharedPreferences.getInstance();
    final authStr = sf.getString('auth_data');
    if (authStr == null) {
      return null;
    }
    final decode = json.decode(authStr);
    final authData = AuthData.fromJson(decode);
    return authData;
  }

  void authLogin(String username, String password) async {
    final dio = Dio();
    state = const AsyncLoading();
    final result = await dio.post(
      '$baseUrl/api/login',
      data: {'username': username, 'password': password},
    );
    final response = result.data['data'];
    if (response != null) {
      final authData = AuthData.fromJson(json.decode(json.encode(response)));
      final sf = await SharedPreferences.getInstance();
      sf.setString('auth_data', json.encode(authData.toJson()));
      state = AsyncData(authData);
      ref.invalidateSelf();
      return;
    }
  }

  void logout() async {
    log('logout', name: 'logout called');
    final sf = await SharedPreferences.getInstance();
    log(sf.getString('auth_data').toString(), name: 'auth_data');
    sf.remove('auth_data');
    state = const AsyncData(null);
    ref.invalidateSelf();
  }
}
