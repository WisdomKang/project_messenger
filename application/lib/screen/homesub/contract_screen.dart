import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ContractListScreen extends StatefulWidget {
  const ContractListScreen ({super.key});

  @override
  State<ContractListScreen> createState() => _ContractListScreenState();
}

class _ContractListScreenState extends State<ContractListScreen> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: ListView.builder(
        itemCount: 10,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text("test $index"),
            onTap: () {
              Navigator.pushNamed(context, "/chatting");
            },
          );
        },
      ),
    );
  }
}
