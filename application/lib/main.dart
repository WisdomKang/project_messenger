import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter/material.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:my_messenger_application/firebase/firebase_message_setting.dart';
import 'package:my_messenger_application/screen/chatroom_screen.dart';
import 'package:my_messenger_application/screen/home_screen.dart';
import 'package:my_messenger_application/screen/login_screen.dart';
import 'package:my_messenger_application/screen/screen_name.dart';

Future<void> main() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();

  // Splash screen setting function
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);
  FlutterNativeSplash.remove();

  // await initializeFirebaseMessaging();
  // FirebaseMessaging.onBackgroundMessage(firebaseMessagingBackgroundHandler);

  runApp(MaterialApp(
    routes: {
      LOGIN_SCREEN: (context) => LoginScreen(),
      MAIN_SCREEN: (context) => HomeScreen(),
      CHATTING_SCREEN : (BuildContext context) => ChatRoomScreen(),
    },
    initialRoute: MAIN_SCREEN,
  ));
}
