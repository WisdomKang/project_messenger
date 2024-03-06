import 'package:flutter/material.dart';

class LoginScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('로그인'),
      ),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: LoginForm(),
      ),
    );
  }
}

class LoginForm extends StatefulWidget {
  @override
  _LoginFormState createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        TextField(
          controller: _emailController,
          decoration: InputDecoration(labelText: '이메일'),
        ),
        SizedBox(height: 16.0),
        TextField(
          controller: _passwordController,
          decoration: InputDecoration(labelText: '비밀번호'),
          obscureText: true,
        ),
        SizedBox(height: 24.0),
        TextButton(
          onPressed: () {
            // 로그인 버튼을 눌렀을 때의 동작 구현
            String email = _emailController.text;
            String password = _passwordController.text;
            // 로그인 로직 추가
            print('로그인 시도 - 이메일: $email, 비밀번호: $password');

            Navigator.of(context).pushNamed("/main_screen");
          },
          style: TextButton.styleFrom(
            backgroundColor: Colors.blue,
          ),
          child: Text('로그인'),
        ),
      ],
    );
  }
}
