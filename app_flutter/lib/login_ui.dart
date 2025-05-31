import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

import 'provider/auth_provider.dart';

class LoginUi extends HookConsumerWidget {
  const LoginUi({super.key});
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final usernameCtrl = useTextEditingController();
    final passwordCtrl = useTextEditingController();
    final auth = ref.watch(authProvider);
    return Scaffold(
      appBar: AppBar(title: const Text("App Flutter")),
      body: Container(
        padding: EdgeInsets.all(16),
        height: MediaQuery.of(context).size.height,
        child: Form(
          child: Column(
            children: [
              TextFormField(
                controller: usernameCtrl,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Username',
                ),
              ),
              SizedBox(height: 8),
              TextFormField(
                controller: passwordCtrl,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Password',
                ),
                obscureText: true,
              ),
              SizedBox(height: 16),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed:
                      auth.isLoading
                          ? null
                          : () {
                            ref
                                .read(authProvider.notifier)
                                .authLogin(
                                  usernameCtrl.text.trim(),
                                  passwordCtrl.text.trim(),
                                );
                          },
                  child:
                      auth.isLoading
                          ? Text("Processing...")
                          : const Text("Login"),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
