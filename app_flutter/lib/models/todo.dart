// ignore_for_file: invalid_annotation_target

import 'package:freezed_annotation/freezed_annotation.dart';

part 'todo.freezed.dart';
part 'todo.g.dart';

@freezed
@JsonSerializable()
class Todo with _$Todo {
  @override
  final String? id;
  @override
  final String title;
  @override
  final String content;
  @override
  @JsonKey(name: "created_at")
  final DateTime createdAt;
  @override
  @JsonKey(name: "modify_at")
  final DateTime modifyAt;
  @override
  @JsonKey(name: "user_id")
  final String userId;

  Todo({
    this.id,
    required this.title,
    required this.content,
    required this.createdAt,
    required this.modifyAt,
    required this.userId,
  });

  factory Todo.fromJson(Map<String, dynamic> json) => _$TodoFromJson(json);

  Map<String, dynamic> toJson() => _$TodoToJson(this);

  factory Todo.fromRow(Map<String, dynamic> row) {
    return Todo(
      id: row['id'],
      title: row['title'],
      content: row['content'],
      createdAt: DateTime.parse(row['created_at']),
      modifyAt: DateTime.parse(row['modify_at']),
      userId: row['user_id'],
    );
  }

  /// Create insert todo without id
  factory Todo.create(Todo todo) {
    return Todo(
      title: todo.title,
      content: todo.content,
      createdAt: todo.createdAt,
      modifyAt: todo.modifyAt,
      userId: todo.userId,
    );
  }

  /// Create insert query sql without id
  String createQuery() {
    return "INSERT INTO todo (title, content, created_at, modify_at, user_id) VALUES ('$title', '$content', '$createdAt', '$modifyAt', '$userId')";
  }
}
