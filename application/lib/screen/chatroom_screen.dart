import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ChatRoomScreen extends StatefulWidget {
  const ChatRoomScreen({super.key});

  @override
  State<ChatRoomScreen> createState() => _ChatRoomScreenState();
}

class _ChatRoomScreenState extends State<ChatRoomScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      bottomSheet: TextField(),
      body: Center(
        child: ListView.builder(
          padding: EdgeInsets.all(5.0),

          itemBuilder: ( context, index) {
              return Container(
                height: 50.0,
                color: Colors.greenAccent,
                padding: EdgeInsets.all(2.0),
                child: Row(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  mainAxisSize: MainAxisSize.max,
                  children: [
                    Placeholder(child: Text("Avartor Image"),),
                    Placeholder(child: Text("Message contents"),)
                  ],
                ),
              );
            },
          itemCount: 50,
        ),
      ),
    );
  }
}
