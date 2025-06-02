

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:powersync/powersync.dart';

import '../models/todo.dart';
import '../powersync/powersync.dart';

class TodoQuery {
  final Future<PowerSyncDatabase> powerSyncDatabase;

  TodoQuery(this.powerSyncDatabase);

  Stream<List<Todo>> fetchTodos() async* {
    final db = await powerSyncDatabase;
    final result = db.watch("SELECT * FROM todo");
    final todos = result.map((e) => e.map((e) => Todo.fromRow(e)).toList());
    yield* todos;
  }

  Future<void> addTodo(Todo todo) async {
    final db = await powerSyncDatabase;
    await db.execute(
      "INSERT INTO todo (id, title, content, created_at, modify_at, user_id) VALUES ('${todo.id}', '${todo.title}', '${todo.content}', '${todo.createdAt}', '${todo.modifyAt}', '${todo.userId}')",
    );
  }

  Future<void> deleteTodo(String todoId) async {
    final db = await powerSyncDatabase;
    await db.execute("DELETE FROM todo WHERE id = '$todoId'");
  }
}

final todoQueryProvider = Provider<TodoQuery>((ref) {
  final db = ref.watch(powerSyncInstanceProvider.future);
  return TodoQuery(db);
});
