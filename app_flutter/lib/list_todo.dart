import 'dart:math';

import 'package:app_flutter/provider/auth_provider.dart';
import 'package:app_flutter/provider/todo_provider.dart';
import 'package:app_flutter/utils/global_data.dart';
import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

class ListTodo extends HookConsumerWidget {
  const ListTodo({super.key});
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final listTodoProvider = ref.watch(todoDataProvider);
    final users = useState<String>("");

    useEffect(() {
      Future.microtask(() async {
        final data = await getAuthData();
        users.value = data.userData.email;
      });
      return null;
    }, []);

    Color randomColorGenerator() {
      return Color.fromARGB(
        255,
        Random().nextInt(255),
        Random().nextInt(255),
        Random().nextInt(255),
      );
    }

    return Scaffold(
      appBar: AppBar(
        title: const Text("List Todo"),
        actions: [
          Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              color: randomColorGenerator(),
              borderRadius: BorderRadius.circular(8.0),
            ),
            child: Text(
              users.value,
              style: TextStyle(color: Theme.of(context).indicatorColor),
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(right: 8.0),
            child: IconButton(
              onPressed: () async {
                ref.read(authProvider.notifier).logout();
              },
              icon: const Icon(Icons.logout),
            ),
          ),
        ],
      ),

      body: listTodoProvider.when(
        data:
            (todos) =>
                todos.isEmpty
                    ? const Center(child: Text("Empty"))
                    : ListView.builder(
                      itemCount: todos.length,
                      itemBuilder: (context, index) {
                        final todo = todos[index];
                        return ListTile(
                          title: Text(todo.title),
                          //subtitle: Text(todo.content),
                          trailing: IconButton(
                            onPressed: () {
                              if (todo.id == null) {
                                return;
                              }
                              ref
                                  .read(todoDataProvider.notifier)
                                  .deleteTodo(todo.id!);
                            },
                            icon: const Icon(Icons.delete),
                          ),
                        );
                      },
                    ),
        error: (error, stackTrace) => Text(error.toString()),
        loading: () => const Center(child: CircularProgressIndicator()),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          ref.read(todoDataProvider.notifier).addRandomTodo();
        },
        child: const Icon(Icons.add),
      ),
    );
  }
}
