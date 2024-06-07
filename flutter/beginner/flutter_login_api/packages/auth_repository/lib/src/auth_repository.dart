import 'dart:async';

import 'package:dio/dio.dart';

part 'auth_token.dart';

enum AuthStatus { unknown, authenticated, unauthenticated }

class AuthRepository {
  final _controller = StreamController<AuthStatus>();
  final String baseUrl = 'https://dummyjson.com';
  final Dio dio = Dio();

  AuthToken? _currentToken;

  // To create a stream builder,
  // we create a function with the keyword ‘async*’ and have it return a stream
  Stream<AuthStatus> get status async* {
    // To emit a data, we can use the keyword ‘yield’,
    yield AuthStatus.unauthenticated;
    //to emit all data from another stream, we can use the keyword ‘yield*’
    yield* _controller.stream;
  }

  Future<void> login({
    required String username,
    required String password,
  }) async {
    try {
      final response = await dio.post(
        '$baseUrl/auth/login',
        data: {'username': username, 'password': password},
      );

      if (response.statusCode == 200) {
        _currentToken = AuthToken.fromJson(response.data);
        _controller.add(
          AuthStatus.authenticated,
        );
      } else {
        throw Exception(
            'Failed to log in with status code: ${response.statusCode}');
      }
    } catch (e) {
      print('login error: $e');
      throw Exception('Failed to log in');
    }
  }

  void logout() {
    _currentToken = null;
    _controller.add(AuthStatus.unauthenticated);
  }

  AuthToken? getCurrentToken() {
    return _currentToken;
  }

  void dispose() => _controller.close();
}
