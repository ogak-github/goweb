// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'todo.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Todo _$TodoFromJson(Map<String, dynamic> json) => Todo(
  id: json['id'] as String?,
  title: json['title'] as String,
  content: json['content'] as String,
  createdAt: DateTime.parse(json['created_at'] as String),
  modifyAt: DateTime.parse(json['modify_at'] as String),
  userId: json['user_id'] as String,
);

Map<String, dynamic> _$TodoToJson(Todo instance) => <String, dynamic>{
  'id': instance.id,
  'title': instance.title,
  'content': instance.content,
  'created_at': instance.createdAt.toIso8601String(),
  'modify_at': instance.modifyAt.toIso8601String(),
  'user_id': instance.userId,
};
