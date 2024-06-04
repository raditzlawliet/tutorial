import 'package:flutter/material.dart';
import 'package:flutter_user_list_cubit/user_list/user_list.dart';

class UserDetailPage extends StatelessWidget {
  const UserDetailPage({super.key, required this.user});
  final User user;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Padding(
              padding: const EdgeInsets.all(24.0),
              child: CircleAvatar(
                radius: 32,
                child: Text(
                  user.name.substring(0, 2).toUpperCase(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(
                child: Text(
                  user.name,
                  style: Theme.of(context).textTheme.headlineMedium,
                ),
              ),
            ),
            Center(
              child: Text(
                user.email,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
