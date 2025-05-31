import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../models/auth_data.dart';

Future<AuthData> getAuthData() async {
  final sf = await SharedPreferences.getInstance();
  final authStr = sf.getString('auth_data');
  if (authStr == null) {
    throw Exception('No auth data found');
  }
  final decode = json.decode(authStr);
  final authData = AuthData.fromJson(decode);
  return authData;
}
