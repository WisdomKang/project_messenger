import 'dart:math';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ChatRoomScreen extends StatefulWidget {
  const ChatRoomScreen({super.key});

  @override
  State<ChatRoomScreen> createState() => _ChatRoomScreenState();
}

class _ChatRoomScreenState extends State<ChatRoomScreen> {
  
  final demoText = [
    "안녕하세요 저는 강현명입니다.",
    "그러시구뇽 안녕하세요",
    "내일은 무엇을 하시나요?",
    "알빠 없잖아",
    "길고 긴 그런 글을 써보면 흠.... ㅇ게 넘치고 넘칠건데 과연 어떻게 될까요?? 궁금합니다. ㅋㅋㅋ 줄바굼을 어떻게 해야할깡??",
    "좀 도 아무거나 써봅시낟,",
    "동해물과 백두산이 마르고 닳도록 하느님이 보우하사 우리 나라 만세"
  ];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: ListView.separated(
          padding: EdgeInsets.all(5.0),
          itemBuilder: ( context, index) {
            var random = Random();
            int rNum = random.nextInt(2);
              return Container(
                  child: MessageWidget(message: demoText.elementAt(index % demoText.length ), isUsers: rNum == 0));
            },
          itemCount: 50,
          separatorBuilder: (context, index) =>  Container(height: 10,),

        ),
      ),
    );
  }
}

class MessageWidget extends StatelessWidget {
  final String message;
  final bool isUsers;

  MessageWidget({required this.message,required this.isUsers});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: isUsers ? MainAxisAlignment.start : MainAxisAlignment.end,
      children: [
        Container(
          padding: EdgeInsets.all(16.0),
          decoration: BoxDecoration(
            color: Colors.grey[200],
            borderRadius: BorderRadius.circular(8.0),
          ),
          constraints: BoxConstraints(
            maxWidth: MediaQuery.of(context).size.width * 2 / 3
          ),
          child: Text(
            message,
            style: TextStyle(fontSize: 16.0),
          ),
        ),
      ],
    );
  }
}
