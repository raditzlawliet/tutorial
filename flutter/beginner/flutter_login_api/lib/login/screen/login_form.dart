import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_login_api/login/cubit/login_cubit.dart';
import 'package:flutter_login_api/login/models/models.dart';
import 'package:formz/formz.dart';

class LoginForm extends StatelessWidget {
  const LoginForm({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocListener<LoginCubit, LoginState>(
      listener: (context, state) {
        if (state.status.isFailure) {
          ScaffoldMessenger.of(context)
            ..hideCurrentSnackBar()
            ..showSnackBar(
              const SnackBar(
                content: Text("Authentication Failed"),
              ),
            );
        }
      },
      child: const Align(
        alignment: Alignment(0, -1 / 3),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            InputUsername(),
            Padding(padding: EdgeInsets.all(12)),
            InputPassword(),
            Padding(padding: EdgeInsets.all(12)),
            ButtonSubmit(),
          ],
        ),
      ),
    );
  }
}

class InputUsername extends StatelessWidget {
  const InputUsername({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LoginCubit, LoginState>(
      buildWhen: (previous, current) => previous.username != current.username,
      builder: (context, state) {
        return TextField(
          key: const Key('loginForm_usernameInput_textField'),
          onChanged: (value) =>
              context.read<LoginCubit>().usernameChanged(value),
          decoration: InputDecoration(
            labelText: 'Username',
            errorText:
                state.username.displayError == UsernameValidationError.empty
                    ? 'invalid username'
                    : null,
          ),
        );
      },
    );
  }
}

class InputPassword extends StatelessWidget {
  const InputPassword({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LoginCubit, LoginState>(
        buildWhen: (previous, current) => previous.password != current.password,
        builder: (context, state) {
          return TextField(
            key: const Key('loginForm_passwordInput_textField'),
            onChanged: (value) =>
                context.read<LoginCubit>().passwordChanged(value),
            obscureText: true,
            decoration: InputDecoration(
              labelText: "password",
              errorText: (() {
                switch (state.password.displayError) {
                  case PasswordValidationError.minimumLength:
                    return 'Password minimum 6 characters';
                  case PasswordValidationError.empty:
                    return 'Password cannot be empty';
                  default:
                    return null;
                }
              })(),
            ),
          );
        });
  }
}

class ButtonSubmit extends StatelessWidget {
  const ButtonSubmit({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<LoginCubit, LoginState>(builder: (context, state) {
      return state.status.isInProgress
          ? const CircularProgressIndicator()
          : ElevatedButton(
              key: const Key('loginForm_continue_raisedButton'),
              onPressed: () {
                context.read<LoginCubit>().loginSubmitted();
              },
              child: const Text("Login"),
            );
    });
  }
}
