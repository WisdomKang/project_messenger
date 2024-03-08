import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class SettingScreen extends StatefulWidget {
  const SettingScreen({super.key});

  @override
  State<SettingScreen> createState() => _SettingScreenState();
}

class _SettingScreenState extends State<SettingScreen> {

  bool isNotification = false;

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: EdgeInsets.all(15.0),
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text("Notification"),
            Switch(
              value: isNotification,
              activeColor: Colors.lime,
              onChanged: (value) {
                setState(() {
                  isNotification = !isNotification;
                });
              },
            )
          ],
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text("Contracts Update"),
            ElevatedButton(onPressed: (){}, child: Icon(Icons.refresh))
          ],
        ),
        Row(
          children: [
            Text("Server Address"),
          ],
        ),
      ],
    );
  }
}
