import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ChatListScreen extends StatefulWidget {
  const ChatListScreen({super.key});

  @override
  State<ChatListScreen> createState() => _ChatListScreenState();
}

class _ChatListScreenState extends State<ChatListScreen> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: ListView(
        padding: EdgeInsets.all(8.0),
        children: [
          Row(
            children: [
              Text("Notification"),
              Switch(
                value: true,
                activeColor: Colors.lime,
                onChanged: (value) {

                },
              )
            ],
          ),
        ],
      )
    );
  }
}
