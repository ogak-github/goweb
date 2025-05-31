import 'package:powersync/powersync.dart';

const schema = Schema(([
  Table('todo', [
    Column.text('title'),
    Column.text('content'),
    Column.text('created_at'),
    Column.text('modify_at'),
    Column.text('user_id'),
  ]),
]));
