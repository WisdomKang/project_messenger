import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';


var dummyData = [
  {
    "id" : "394jf",
    "name" : "강현명",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "고양이",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강아지",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "다람쥐군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "수달군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강현명",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "고양이",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강아지",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "다람쥐군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "수달군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강현명",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "고양이",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강아지",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "다람쥐군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "수달군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강현명",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "고양이",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강아지",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "다람쥐군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "수달군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강현명",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "고양이",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "강아지",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "다람쥐군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
  {
    "id" : "394jf",
    "name" : "수달군",
    "email" : "buckgusaud@gmail.com",
    "avatar" : "/default",
  },
];


class ContractScreen extends StatelessWidget {
  const ContractScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: ListView.builder(
        itemCount: dummyData.length,
          itemBuilder: (context, index) {
            return InkWell(
              onTap: () {
                print(" tapped on ${dummyData[index]["id"]} Wow!!");
              },
              highlightColor: Colors.blue.withOpacity(0.1),
              splashColor: Colors.blue.withOpacity(0.3),
              child: ListTile(
                title: Text( dummyData[index]["name"]!),
                subtitle: Text( dummyData[index]["email"]!),
                leading: CircleAvatar(
                  backgroundImage: NetworkImage("https://media.gettyimages.com/id/1330275608/vector/female-avatar-icon.jpg?s=612x612&w=gi&k=20&c=HvX_Mulx2RIAwLmRNEHQnMQD2UUSRV-Ai1H96q-grw0="),
                ),
                selectedColor: Colors.green.withOpacity(0.3),
              ),
            );
          },
      )
    );
  }
}