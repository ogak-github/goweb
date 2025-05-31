import 'dart:async';
import 'dart:developer';

import 'package:app_flutter/models/auth_data.dart';
import 'package:intl/intl.dart';
import 'package:powersync/powersync.dart';

import '../data/todo_api.dart';
import '../models/todo.dart';

class BackendConnector extends PowerSyncBackendConnector {
  final String endpoint;
  final AuthData authData;
  BackendConnector(this.endpoint, this.authData);

  @override
  Future<PowerSyncCredentials?> fetchCredentials() async {
    final format = DateFormat('EEE MMM dd HH:mm:ss \'UTC\' yyyy');
    final parsedDate = format.parseUtc(authData.expiredIn);
    try {
      return PowerSyncCredentials(
        endpoint: 'http://10.0.2.2:6100',
        token: authData.token,
        userId: authData.userData.id,
        expiresAt: parsedDate,
      );
    } catch (e) {
      log(e.toString(), name: "BackendConnector");
      return null;
    }
  }

  @override
  Future<void> uploadData(PowerSyncDatabase database) async {
    final todoApi = TodoApi();
    final transaction = await database.getNextCrudTransaction();
    if (transaction == null) {
      return;
    }

    for (var op in transaction.crud) {
      switch (op.op) {
        case UpdateType.put:
          final fullData = Map<String, dynamic>.from(op.opData!);
          final todo = Todo.fromJson(fullData);
          await todoApi.addTodo(todo);
        case UpdateType.patch:
        /*  log(op.opData.toString(), name: UpdateType.patch.json);
          final todo = Todo.fromJson(op.opData!);
          await todoApi.addTodo(todo, authData.token); */
        case UpdateType.delete:
          await todoApi.deleteTodo(op.id);
      }
    }

    await transaction.complete();
  }
}
