import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:my_messenger_application/screen/chatlist_screen.dart';
import 'package:my_messenger_application/screen/contract_screen.dart';
import 'package:my_messenger_application/screen/setting_screen.dart';

class MainScreen extends StatefulWidget {
  const MainScreen({super.key});

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {

  int _selectedIndex = 0;

  final List<Widget> _pages = [
    const ContractScreen(),
    const ChatListScreen(),
    const SettingScreen(),
  ];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body : _pages[_selectedIndex],
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
        items: [
          BottomNavigationBarItem(icon: Icon(Icons.people), label: "연락처"),
          BottomNavigationBarItem(icon: Icon(Icons.message) , label: "대화목록"),
          BottomNavigationBarItem(icon: Icon(Icons.settings), label: "설정"),
        ],
      ),
    );
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }
}








