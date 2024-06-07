import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:auth_repository/auth_repository.dart';
import 'package:user_repository/user_repository.dart';

part 'auth_state.dart';

class AuthCubit extends Cubit<AuthState> {
  AuthCubit(
      {required AuthRepository authRepository,
      required UserRepository userRepository})
      : _authRepository = authRepository,
        _userRepository = userRepository,
        super(const AuthState.unknown()) {
    _authStatusSubscription =
        _authRepository.status.listen((status) => onAuthStatusChanged(status));
  }

  final AuthRepository _authRepository;
  final UserRepository _userRepository;
  late StreamSubscription<AuthStatus> _authStatusSubscription;

  @override
  Future<void> close() {
    _authStatusSubscription.cancel();
    return super.close();
  }

  void onAuthStatusChanged(AuthStatus status) async {
    switch (status) {
      case AuthStatus.unauthenticated:
        emit(const AuthState.unauthenticated());
        break;

      case AuthStatus.authenticated:
        final authToken = _authRepository.getCurrentToken();
        if (authToken == null) {
          emit(const AuthState.unauthenticated());
          break;
        }

        final user = await _tryGetUser(authToken.token);
        if (user == null) {
          emit(const AuthState.unauthenticated());
        } else {
          emit(AuthState.authenticated(user));
        }
        break;

      case AuthStatus.unknown:
        emit(const AuthState.unknown());
        break;
    }
  }

  void logout() {
    _authRepository.logout();
  }

  Future<User?> _tryGetUser(String authToken) async {
    try {
      final user = _userRepository.getUser(authToken);
      return user;
    } catch (e) {
      print('Error getting user: $e');
      return null;
    }
  }
}
