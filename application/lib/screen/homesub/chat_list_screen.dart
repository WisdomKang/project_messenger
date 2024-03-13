import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:my_messenger_application/screen/screen_name.dart';

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
          ChatRoomInfo(roomId: "1", roomName: "친구"),
          ChatRoomInfo(roomId: "2", roomName: "마음"),

          ChatRoomInfo(roomId: "3", roomName: "용기"),
          ChatRoomInfo(roomId: "4", roomName: "믿음"),
          ChatRoomInfo(roomId: "5", roomName: "사랑"),
          
        ],
      )
    );
  }
}

class ChatRoomInfo extends StatelessWidget {
  String roomId;
  String roomName;


  ChatRoomInfo({super.key, required this.roomId, required this.roomName});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(roomName),
      leading: CircleAvatar(),
      onTap: (){
        print( "$roomId is tapped!");
        Navigator.of(context).pushNamed(CHATTING_SCREEN);
      },
    );
  }
}
