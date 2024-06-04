import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_user_list_cubit/user_list/user_list.dart';
import 'package:flutter_user_list_cubit/user_list/view/user_detail_page.dart';

part "user_list_view.dart";

class UserListPage extends StatelessWidget {
  const UserListPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => UserListCubit(),
      child: const UserListView(),
    );
  }
}
