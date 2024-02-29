import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          bottom: const TabBar(tabs: [
            Tab(icon: Icon(Icons.contact_page)),
            Tab(icon: Icon(Icons.chat) ),
            Tab(icon:Icon(Icons.settings))
          ]
          ),
        ),
        body: TabBarView(
          children: [
            Center(child: Text("연력처"),),
            Center(child: Text("채팅목록"),),
            Center(child: Text("설정"),)
          ],
        ),
      ),
    );
  }
}
