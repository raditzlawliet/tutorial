part of 'auth_repository.dart';

class AuthToken {
  final String token;

  const AuthToken({required this.token});

  factory AuthToken.fromJson(Map<String, dynamic> json) {
    return AuthToken(
      token: json['token'],
    );
  }
}
