import 'package:flutter/material.dart';

class ContractListScreen extends StatefulWidget {
  const ContractListScreen({super.key});

  @override
  State<ContractListScreen> createState() => _ContractListScreenState();
}

var demoContractList = [
  {
    "name": "막시묵스",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "궁구구루",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "토르",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "강현명",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "로추로",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "뺴상",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "율류",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
  {
    "name": "테스터",
    "imageUrl": "test",
    "contractId": "1fk3",
  },
];

class _ContractListScreenState extends State<ContractListScreen> {
  @override
  Widget build(BuildContext context) {
    print("build is called!");
    return Center(
      child: ListView.separated(
        itemBuilder: (context, index) => ContractRow(
            contractId: demoContractList.elementAt(index)["contractId"]!,
            contractName: demoContractList.elementAt(index)["name"]!,
            imageUrl: demoContractList.elementAt(index)["imageUrl"]!,
            onTappedItem: onTappedItem,
        ),
        itemCount: demoContractList.length,
        separatorBuilder: (context, index) => Container(
          height: 5.0,
        ),
        padding: EdgeInsets.all(8.0),
      ),
    );
  }

  void onTappedItem(){
    print("Taaped!");
    setState(() {
      demoContractList.shuffle();

      // demoContractList = [
      //   {"name": "혼자", "contractId": "imageUrl", "imageUrl": "imageUrl"}
      // ];
    });
  }



  @override
  void didUpdateWidget(ContractListScreen oldWidget) {
    super.didUpdateWidget(oldWidget);
    print("didUpdateWidget is called!");
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    print("didChangeDependencies is called!");
  }
}

class ContractRow extends StatelessWidget {
  final String? imageUrl;
  final String contractName;
  final String contractId;

  final void Function() onTappedItem;

  const ContractRow(
      {super.key,
      required this.contractId,
      required this.contractName,
      required this.imageUrl,
      required this.onTappedItem,
      });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: Image(
        image: AssetImage("assets/RamZ-600.png"),
        height: 50,
      ),
      title: Text(contractName),
      onTap: onTappedItem,
    );
  }
}
