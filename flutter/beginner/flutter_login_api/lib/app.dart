import 'package:auth_repository/auth_repository.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_login_api/auth/cubit/auth_cubit.dart';
import 'package:flutter_login_api/home/screen/home_screen.dart';
import 'package:flutter_login_api/login/screen/login_screen.dart';
import 'package:flutter_login_api/splash_screen.dart';
import 'package:user_repository/user_repository.dart';

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  late final AuthRepository _authRepository;
  late final UserRepository _userRepository;

  @override
  void initState() {
    super.initState();
    _authRepository = AuthRepository();
    _userRepository = UserRepository();
  }

  @override
  void dispose() {
    _authRepository.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return RepositoryProvider.value(
      value: _authRepository,
      child: BlocProvider(
        create: (context) => AuthCubit(
          authRepository: _authRepository,
          userRepository: _userRepository,
        ),
        child: AppScreen(),
      ),
    );
  }
}

class AppScreen extends StatefulWidget {
  const AppScreen({super.key});

  @override
  State<AppScreen> createState() => _AppScreenState();
}

class _AppScreenState extends State<AppScreen> {
  final _navigatorKey = GlobalKey<NavigatorState>();

  NavigatorState get _navigator => _navigatorKey.currentState!;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      navigatorKey: _navigatorKey,
      theme: ThemeData(
        useMaterial3: true,
      ),
      builder: (context, child) {
        return BlocListener<AuthCubit, AuthState>(
          listener: (context, state) {
            print('authState.status ${state.status}');
            switch (state.status) {
              case AuthStatus.authenticated:
                _navigator.pushNamedAndRemoveUntil<void>(
                  "/home",
                  (route) => false,
                );
                break;
              default:
                _navigator.pushAndRemoveUntil<void>(
                  LoginScreen.route(),
                  (route) => false,
                );
                break;
            }
          },
          child: child,
        );
      },
      initialRoute: '/',
      onGenerateRoute: (RouteSettings settings) {
        switch (settings.name) {
          case '/':
            return SplashScreen.route();
          case '/login':
            return LoginScreen.route();
          case '/home':
            return HomeScreen.route();
          default:
            return MaterialPageRoute<void>(
              builder: (context) => const Scaffold(
                body: Center(child: Text('Page not found')),
              ),
            );
        }
      },
    );
  }
}
