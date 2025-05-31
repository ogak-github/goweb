import 'dart:developer';

import 'package:app_flutter/provider/auth_provider.dart';
import 'package:app_flutter/utils/global_data.dart';
import 'package:dio/dio.dart';

import '../models/todo.dart';

class TodoApi {
  final Dio _dio = Dio();

  TodoApi();

  Future<void> addTodo(Todo todo) async {
    final auth = await getAuthData();
    final result = await _dio.post(
      '$baseUrl/api/todo/create',
      options: Options(headers: {'Authorization': 'Bearer ${auth.token}'}),
      data: {'title': todo.title, 'content': todo.content},
    );
    if (result.statusCode == 201) {
      log(result.data.toString(), name: 'addTodo');
      return;
    }
    if (result.statusCode == 400) {
      throw Exception(result.data['message']);
    }
    if (result.statusCode == 401) {
      throw Exception(result.data['message']);
    }
  }

  Future<void> deleteTodo(String todoId) async {
    final auth = await getAuthData();
    final result = await _dio.delete(
      '$baseUrl/api/todo/delete/$todoId',
      options: Options(headers: {'Authorization': 'Bearer ${auth.token}'}),
    );
    if (result.statusCode == 200) {
      log(result.data.toString(), name: 'deleteTodo');
      return;
    }
    if (result.statusCode == 400) {
      throw Exception(result.data['message']);
    }
    if (result.statusCode == 401) {
      throw Exception(result.data['message']);
    }
  }
}
