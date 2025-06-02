import 'dart:convert';

import 'package:english_words/english_words.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:uuid/uuid.dart';
import '../data/todo_query.dart';
import '../models/auth_data.dart';
import '../models/todo.dart';

part 'todo_provider.g.dart';

@riverpod
class TodoData extends _$TodoData {
  @override
  Stream<List<Todo>> build() async* {
    //final todos = await _fetchTodos();
    final todosLocal = _fetchTodosFromDB();
    yield* todosLocal;
  }

  /*   Future<List<Todo>> _fetchTodos() async {
    final query = ref.watch(todoQueryProvider);
    final todos = await query.fetchTodos();
    return todos;
  } */

  Stream<List<Todo>> _fetchTodosFromDB() async* {
    final query = ref.watch(todoQueryProvider);
    final result = query.fetchTodos();
    yield* result;

    // return result.map((e) => Todo.fromRow(e)).toList();
  }

  /*   Future<void> addTodo(Todo todo) async {
    final query = ref.watch(todoQueryProvider);
    await query.addTodo(todo);
  } */

  Future<void> addRandomTodo() async {
    final sf = await SharedPreferences.getInstance();
    final authData = sf.getString('auth_data');
    if (authData == null) {
      return;
    }

    final decode = json.decode(authData);
    final user = AuthData.fromJson(decode);
    final query = ref.watch(todoQueryProvider);
    final uuid = Uuid().v4();
    final text = WordPair.random().asLowerCase;
    final content = "Content created at ${DateTime.now()}";
    await query.addTodo(
      Todo(
        id: uuid,
        title: text,
        content: content,
        createdAt: DateTime.now(),
        modifyAt: DateTime.now(),
        userId: user.userData.id,
      ),
    );
  }

  Future<void> deleteTodo(String id) async {
    final query = ref.read(todoQueryProvider);
    await query.deleteTodo(id);
  }
}
