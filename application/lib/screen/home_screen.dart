import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:my_messenger_application/screen/homesub/chat_list_screen.dart';
import 'package:my_messenger_application/screen/homesub/contract_screen.dart';
import 'package:my_messenger_application/screen/homesub/setting_screen.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {

  int _pageIndex = 0;

  final List<Widget> _pageWidgetList = <Widget>[
    const ContractListScreen(),
    const ChatListScreen(),
    const SettingScreen()
  ];

  final List<String> _pageNameList = <String>[
    "Contracts",
    "Chats",
    "Setting",
  ];

  void _onItemTapped(int index){
    setState(() {
      _pageIndex = index;
    });
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(_pageNameList.elementAt(_pageIndex)),
          backgroundColor: Colors.greenAccent,
          toolbarHeight: 40.0,
        ),
        bottomNavigationBar: NavigationBar(
          onDestinationSelected: (int index){
            setState(() {
              _pageIndex = index;
            });
          },

          destinations: <Widget>[
            NavigationDestination(
              selectedIcon:  Icon(Icons.contact_page_outlined),
                icon: Icon(Icons.contact_page) ,
                label: _pageNameList[0] ,
            ),
            NavigationDestination(icon: Icon(Icons.chat) , label: _pageNameList[1] , ),
            NavigationDestination(icon:Icon(Icons.settings) , label: _pageNameList[2], ),
          ],
          selectedIndex: _pageIndex,
          indicatorColor: Colors.greenAccent,
        ),
        body: _pageWidgetList.elementAt(_pageIndex)
      );
  }
}
