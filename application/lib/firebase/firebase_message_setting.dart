import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:my_messenger_application/firebase_options.dart';

Future<void> initializeFirebaseMessaging() async {
  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);
  await _setFCMToken();
  return;
}

late AndroidNotificationChannel channel;
bool isFlutterLocalNotificationsInitialized = false;

late FlutterLocalNotificationsPlugin flutterLocalNotificationsPlugin;
late String? FCMToken;

Future<void> firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);
  await _setupFlutterNotifications();
  showFlutterNotificatoin(message);
}

Future<void> _setupFlutterNotifications() async {
  if (isFlutterLocalNotificationsInitialized) {
    return;
  }

  channel =const AndroidNotificationChannel(
    "my_Messendger",
    "My Messenger Notification Channel",
    description: "설명",
    importance: Importance.high,
  );

  // initialize notificationPlugin
  flutterLocalNotificationsPlugin = FlutterLocalNotificationsPlugin();

  await flutterLocalNotificationsPlugin
      .resolvePlatformSpecificImplementation<
      AndroidFlutterLocalNotificationsPlugin>()
  ?.createNotificationChannel(channel);

  await FirebaseMessaging.instance.setForegroundNotificationPresentationOptions(
    alert: true,
    badge: true,
    sound: true,
  );

  isFlutterLocalNotificationsInitialized = true;

  }

const notificationId = "myMessengerProject";

class MessageForm {
  static const String TITLE = "title";
  static const String MESSAGE = "body";
}

void showFlutterNotificatoin(RemoteMessage message) {
  RemoteNotification? notification = message.notification;

  if ( message.data != null) {
    flutterLocalNotificationsPlugin.show(
        notificationId.hashCode,
        message.data[MessageForm.TITLE],
        message.data[MessageForm.MESSAGE],
      const NotificationDetails(
        android: AndroidNotificationDetails(
            "My Channel",
            "My Channel Name",
          priority: Priority.high,
          channelDescription: "설명할꼐 없는 테스트",
          importance: Importance.max,
            // icon: 'launch_background'
        )
      ),
    );
  }else{
    print("notification is null.");
  }
}

Future<void> _setFCMToken() async {
  FirebaseMessaging firebaseMessaging = FirebaseMessaging.instance;
  FCMToken = await firebaseMessaging.getToken();

  print(" FCMToken : '$FCMToken'");
  // Token refresh setting
  firebaseMessaging.onTokenRefresh.listen((event) {

  });
}