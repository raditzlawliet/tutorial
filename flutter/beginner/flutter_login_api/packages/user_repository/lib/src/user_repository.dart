import 'dart:async';

import 'package:dio/dio.dart';
import 'package:user_repository/src/models/models.dart';

class UserRepository {
  User? _user;
  String? _lastToken;
  final String baseUrl = 'https://dummyjson.com';
  final Dio dio = Dio();

  Future<User?> getUser(String token) async {
    if (_user != null && _lastToken == token) return _user;
    
    try {
      final response = await dio.get(
        '$baseUrl/auth/me',
        options: Options(
          headers: {
            'Authorization': 'Bearer $token',
          },
        ),
      );

      if (response.statusCode == 200) {
        _user = User.fromJson(response.data);
        return _user;
      } else {
        throw Exception(
            'Failed to log in with status code: ${response.statusCode}');
      }
    } catch (e) {
      print('get user error: $e');
      throw Exception('Failed to get user');
    }
  }
}
