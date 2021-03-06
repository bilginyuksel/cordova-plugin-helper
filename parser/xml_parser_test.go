package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func beforeTestParseFileCreateSampleCordovaPluginXMLFile(filename string, t *testing.T) {

	content := []byte(`
    <?xml version='1.0' encoding='utf-8'?>
    <plugin id="cordova-plugin-hms-push"
            version="5.0.2"
            xmlns="http://apache.org/cordova/ns/plugins/1.0"
            xmlns:android="http://schemas.android.com/apk/res/android">
        <name>Cordova Plugin HMS Push</name>
        <description>Cordova Plugin HMS Push</description>
        <license>Apache 2.0</license>
        <keywords>android, huawei, hms, push</keywords>
    
        <engines>
            <engine name="cordova" version=">=3.0.0"/>
        </engines>
    
        <js-module name="HmsPush" src="www/HmsPush.js">
            <clobbers target="HmsPush"/>
        </js-module>
    
        <js-module name="HmsPushResultCode" src="www/HmsPushResultCode.js">
            <clobbers target="HmsPushResultCode"/>
        </js-module>
    
        <js-module name="HmsPushEvent" src="www/HmsPushEvent.js">
            <clobbers target="HmsPushEvent"/>
        </js-module>
    
        <js-module name="HmsLocalNotification" src="www/HmsLocalNotification.js">
            <clobbers target="HmsLocalNotification"/>
        </js-module>
    
    
        <js-module name="Interfaces" src="www/Interfaces.js"/>
        <js-module name="CordovaRemoteMessage" src="www/CordovaRemoteMessage.js"/>
        <js-module name="utils" src="www/utils.js"/>
    
        <platform name="android">
    
            <hook type="after_plugin_install" src="hooks/after_plugin_install.js"/>
            <hook type="before_plugin_uninstall" src="hooks/before_plugin_uninstall.js"/>
            <hook type="after_prepare" src="hooks/after_prepare.js"/>
    
            <config-file target="AndroidManifest.xml" parent="/*">
                <uses-permission android:name="android.permission.INTERNET"/>
                <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE"/>
                <!-- Below permissions are to support vibration and send scheduled local notifications -->
                <uses-permission android:name="android.permission.VIBRATE"/>
                <uses-permission android:name="android.permission.RECEIVE_BOOT_COMPLETED"/>
                <uses-permission android:name="android.permission.WAKE_LOCK"/>
                <uses-permission android:name="android.permission.SYSTEM_ALERT_WINDOW"/>
            </config-file>
    
            <config-file target="AndroidManifest.xml" parent="application">
                <receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationActionsReceiver"/>
                <!-- This receivers are for sending scheduled local notifications -->
                <receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationBootEventReceiver">
                    <intent-filter>
                        <action android:name="android.intent.action.BOOT_COMPLETED"/>
                    </intent-filter>
                </receiver>
                <receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationScheduledPublisher"
                          android:enabled="true"
                          android:exported="true">
                </receiver>
                <meta-data
                        android:name="push_kit_auto_init_enabled"
                        android:value="true"/>
            </config-file>
            <config-file target="AndroidManifest.xml" parent="application/activity">
                <intent-filter>
                    <action android:name="android.intent.action.VIEW"/>
                    <category android:name="android.intent.category.DEFAULT"/>
                    <category android:name="android.intent.category.BROWSABLE"/>
                    <data android:scheme="app"/>
                </intent-filter>
            </config-file>
    
            <config-file target="AndroidManifest.xml" parent="application">
                <service android:name="com.huawei.hms.cordova.push.remote.HmsPushMessageService" android:exported="true">
                    <intent-filter>
                        <action android:name="com.huawei.push.action.MESSAGING_EVENT"/>
                    </intent-filter>
                </service>
            </config-file>
    
            <config-file target="config.xml" parent="/*">
                <feature name="HMSPush">
                    <param name="android-package" value="com.huawei.hms.cordova.push.HMSPush"/>
                </feature>
            </config-file>
    
            <framework src="androidx.core:core:1.3.1"/>
            <framework src="com.facebook.fresco:fresco:2.2.0"/>
            <framework src="com.huawei.hms:push:5.0.2.300"/>
    
            <framework src="resources/plugin.gradle" custom="true" type="gradleReference"/>
    
            <source-file src="src/main/java/com/huawei/hms/cordova/push/HMSPush.java"
                         target-dir="src/com/huawei/hms/cordova/push"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/hmslogger/HMSLogger.java"
                         target-dir="src/com/huawei/hms/cordova/push/hmslogger"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/config/NotificationAttributes.java"
                         target-dir="src/com/huawei/hms/cordova/push/config"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/constants/Core.java"
                         target-dir="src/com/huawei/hms/cordova/push/constants"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/constants/LocalNotification.java"
                         target-dir="src/com/huawei/hms/cordova/push/constants"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/constants/NotificationConstants.java"
                         target-dir="src/com/huawei/hms/cordova/push/constants"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/constants/RemoteMessageAttributes.java"
                         target-dir="src/com/huawei/hms/cordova/push/constants"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/constants/ResultCode.java"
                         target-dir="src/com/huawei/hms/cordova/push/constants"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/listeners/HmsLocalNotificationActionPublisher.java"
                         target-dir="src/com/huawei/hms/cordova/push/listeners"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/listeners/HmsMessagePublisher.java"
                         target-dir="src/com/huawei/hms/cordova/push/listeners"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/local/BitmapDataSubscriber.java"
                         target-dir="src/com/huawei/hms/cordova/push/local"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotification.java"
                         target-dir="src/com/huawei/hms/cordova/push/local"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotificationController.java"
                         target-dir="src/com/huawei/hms/cordova/push/local"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotificationPicturesLoader.java"
                         target-dir="src/com/huawei/hms/cordova/push/local"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationActionsReceiver.java"
                         target-dir="src/com/huawei/hms/cordova/push/receiver"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationBootEventReceiver.java"
                         target-dir="src/com/huawei/hms/cordova/push/receiver"/>
            <source-file
                    src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationScheduledPublisher.java"
                    target-dir="src/com/huawei/hms/cordova/push/receiver"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/NotificationActionHandler.java"
                         target-dir="src/com/huawei/hms/cordova/push/receiver"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushInstanceId.java"
                         target-dir="src/com/huawei/hms/cordova/push/remote"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushMessageService.java"
                         target-dir="src/com/huawei/hms/cordova/push/remote"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushMessaging.java"
                         target-dir="src/com/huawei/hms/cordova/push/remote"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ApplicationUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ArrayUtil.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/BundleUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/CordovaUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/MapUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/NotificationConfigUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/RemoteMessageUtils.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/Action.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
            <source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ActionManager.java"
                         target-dir="src/com/huawei/hms/cordova/push/utils"/>
    
    
        </platform>
    </plugin>	
    `)
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		t.Error(err)
	}
}

func afterTestParseFileRemovePluginXML(filename string) {
	os.Remove(filename)
}

func TestParseFile_ReturnCorrectFileInformation(t *testing.T) {
	beforeTestParseFileCreateSampleCordovaPluginXMLFile("plugin.xml", t)
	xmlResult, err := ParseXML("plugin.xml")
	if err != nil {
		t.Error()
	}

	if xmlResult.ID != "cordova-plugin-hms-push" && xmlResult.Author != "" &&
		xmlResult.License != "Apache 2.0" && xmlResult.Description != "Cordova Plugin HMS Push" &&
		xmlResult.Name != "Cordova Plugin HMS Push" {
		t.Logf("Actual: %s, Expected: %s", xmlResult.ID, "cordova-plugin-hms-push")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Author, "")
		t.Logf("Actual: %s, Expected: %s", xmlResult.License, "Apache 2.0")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Description, "Cordova Plugin HMS Push")
		t.Logf("Actual: %s, Expected: %s", xmlResult.Name, "Cordova Plugin HMS Push")
		t.Error()
	}
	afterTestParseFileRemovePluginXML("plugin.xml")
}

func TestParseFile_FileNotFound(t *testing.T) {
	_, err := ParseXML("notfound.xml")
	if err == nil {
		t.Error()
	}
}

func TestParseFile_ExtensionError(t *testing.T) {
	_, err := ParseXML("file.yaml")
	if err == nil {
		t.Error()
	}
}

func TestCreateXML_CreateNewFile(t *testing.T) {
	plugin := &Plugin{ID: "id-test", Version: "version-test", Xmlns: "xmlns-test",
		XmlnsAndroid: "android-test", Name: "name-test", Description: "description-test", License: "license-test"}
	CreateXML(plugin, "test.xml")
	xmlFile, _ := ParseXML("test.xml")
	fmt.Println(xmlFile)
	var (
		isIDEqual           bool
		isVersionEqual      bool
		isXmlnsEqual        bool
		isXmlnsAndroidEqual bool
		isNameEqual         bool
		isDescriptionEqual  bool
		isLicenseEqual      bool
	)
	isIDEqual = plugin.ID == xmlFile.ID
	isVersionEqual = plugin.Version == xmlFile.Version
	isXmlnsEqual = plugin.Xmlns == xmlFile.Xmlns
	isXmlnsAndroidEqual = plugin.XmlnsAndroid == xmlFile.XmlnsAndroid
	isNameEqual = plugin.Name == xmlFile.Name
	isDescriptionEqual = plugin.Description == xmlFile.Description
	isLicenseEqual = plugin.License == xmlFile.License

	if !isIDEqual || !isVersionEqual || !isXmlnsEqual || !isXmlnsAndroidEqual ||
		!isNameEqual || !isDescriptionEqual || !isLicenseEqual {
		t.Error()
	}
	afterTestParseFileRemovePluginXML("test.xml")
}

func TestCreateXMLFilenameNotXML_CreateFileWithXMLExtension(t *testing.T) {

	isFileExists := func(filename string) bool {
		_, err := os.Stat(filename)
		return !os.IsNotExist(err)
	}

	if isFileExists("test.xml") {
		t.Error()
	}

	plugin := &Plugin{}
	CreateXML(plugin, "test")

	if !isFileExists("test.xml") {
		t.Error()
	}

	afterTestParseFileRemovePluginXML("test.xml")
}

func TestNewSourceFrom_AddSourceFile(t *testing.T) {
	plg := Plugin{Platform: &Platform{}}
	var javaFiles []string
	for i := 1; i < 30; i++ {
		file := fmt.Sprintf("src/test%d.java", i)
		javaFiles = append(javaFiles, file)
	}
	plg.Platform.NewSourceFrom(javaFiles)
	if len(javaFiles) != len(plg.Platform.SourceFiles) {
		t.Error()
	}
	for i := 0; i < len(javaFiles); i++ {
		dir, _ := filepath.Split(javaFiles[i])
		isSourceFileEqual := javaFiles[i] == plg.Platform.SourceFiles[i].Src && dir == plg.Platform.SourceFiles[i].TargetDir
		if !isSourceFileEqual {
			t.Error()
		}
	}
}

func TestNewJsModulesFrom_AddJsFile(t *testing.T) {
	plg := Plugin{Platform: &Platform{}}
	var jsFiles []string
	for i := 1; i < 30; i++ {
		file := fmt.Sprintf("www/test%d.js", i)
		jsFiles = append(jsFiles, file)
	}
	plg.NewJsModulesFrom(jsFiles)
	if len(jsFiles) != len(plg.JsModule) {
		t.Error()
	}
	for i := 0; i < len(jsFiles); i++ {
		_, name := filepath.Split(jsFiles[i])
		name = strings.TrimSuffix(name, filepath.Ext(jsFiles[i]))
		isJsFileEqual := jsFiles[i] == plg.JsModule[i].Src && name == plg.JsModule[i].Name
		if !isJsFileEqual {
			t.Error()
		}
	}
}
