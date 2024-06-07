import 'package:auth_repository/auth_repository.dart';
import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_login_api/login/models/models.dart';
import 'package:formz/formz.dart';

part 'login_state.dart';

class LoginCubit extends Cubit<LoginState> {
  LoginCubit({
    required AuthRepository authRepository,
  })  : _authRepository = authRepository,
        super(const LoginState());

  final AuthRepository _authRepository;

  void usernameChanged(String value) {
    final username = Username.dirty(value);
    emit(state.copyWith(
      username: username,
      isValid: Formz.validate([username, state.password]),
      status: FormzSubmissionStatus.initial,
    ));
  }

  void passwordChanged(String value) {
    final password = Password.dirty(value);
    emit(state.copyWith(
      password: password,
      isValid: Formz.validate([state.username, password]),
      status: FormzSubmissionStatus.initial,
    ));
  }

  Future<void> loginSubmitted() async {
    if (!state.isValid) return;
    emit(state.copyWith(status: FormzSubmissionStatus.inProgress));
    try {
      await _authRepository.login(
        username: state.username.value,
        password: state.password.value,
      );
      emit(state.copyWith(status: FormzSubmissionStatus.success));
    } catch (e) {
      print('login submitted error $e');
      emit(state.copyWith(status: FormzSubmissionStatus.failure));
    }
  }
}
