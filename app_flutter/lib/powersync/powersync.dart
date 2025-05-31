import 'dart:convert';

import 'package:app_flutter/models/schema.dart';
import 'package:app_flutter/provider/auth_provider.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:path/path.dart';
import 'package:path_provider/path_provider.dart';
import 'package:powersync/powersync.dart';

import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../models/auth_data.dart';
import 'backend_connector.dart';

part 'powersync.g.dart';

/* opendatabase() async {
  final dir = await getApplicationDocumentsDirectory();
  final path = join(dir.path, 'powersync-dart.db');

  var db = PowerSyncDatabase(schema: schema, path: path);
  db.initialize();
} */

@Riverpod(keepAlive: true)
Future<PowerSyncDatabase> powerSyncInstance(Ref ref) async {
  final dir = await getApplicationDocumentsDirectory();
  final path = join(dir.path, 'powersync-dart.db');

  final sf = await SharedPreferences.getInstance();
  final authData = sf.getString('auth_data');
  final jsonDecode = json.decode(authData ?? '');
  final data = AuthData.fromJson(jsonDecode);
  if (authData == null) {
    return PowerSyncDatabase(schema: schema, path: path);
  }

  var db = PowerSyncDatabase(schema: schema, path: path);
  db.initialize();

  db.connect(connector: BackendConnector(baseUrl, data));

  ref.onDispose(db.close);

  return db;
}
