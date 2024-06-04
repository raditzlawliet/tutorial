import 'package:bloc/bloc.dart';
import 'package:dio/dio.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_user_list_cubit/user_list/model/user.dart';

part "user_list_state.dart";

class UserListCubit extends Cubit<UserListState> {
  UserListCubit() : super(const UserListState.initial());

  fetchUser() async {
    try {
      emit(const UserListState.loading());
      Dio dio = Dio();

      final res = await dio.get("https://jsonplaceholder.typicode.com/users");

      if (res.statusCode == 200) {
        final List<User> users = res.data.map<User>((d) {
          return User.fromJson(d);
        }).toList();

        emit(UserListState.success(users));
      } else {
        emit(UserListState.error("error loading data: ${res.data.toString()}"));
      }
    } catch (e) {
      emit(UserListState.error("error loading data: ${e.toString()}"));
    }
  }
}
