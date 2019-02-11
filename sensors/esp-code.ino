#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WebServer.h>
#include <FS.h>
#include <ArduinoJson.h>
#include <assert.h>

const char *credArgs[] = {"wifi_ssid", "wifi_pass", "data_url", "data_port", "token"};
const int CNT_OF_ARGS = 5;

const char *AP_SSID = "ESP";
const char *AP_PASS = "3.141592";

const int WIFI_CONN_DELAY = 10000;
const int SENDING_FRAME = 1000;
const int CLIENT_CONN_DELAY = 1000;
const int JSON_STACK_SIZE = 8192;

struct Credentials {
    String WiFiSSID, WiFiPASS, dataURL, dataPort, token;
    Credentials(){};
    Credentials(const String &_SSID, const String &_PASS, const String &_URL, const String &_port, const String &_token): 
        WiFiSSID(_SSID), 
        WiFiPASS(_PASS), 
        dataURL(_URL),
        token(_token),
        dataPort(_port) {
    };

    String &getVar(const char *_name) {
        int pos = -1;
        for (int i = 0; i < CNT_OF_ARGS; ++i) {
            if (strcmp(credArgs[i], _name) == 0) {
                pos = i;
                break;
            }
        }
        assert(pos != -1);
        switch (pos) {
            case 0:
                return WiFiSSID;
            case 1:
                return WiFiPASS;
            case 2:
                return dataURL;
            case 3:
                return dataPort;
            case 4:
                return token;
        }
    }

    const char *getArg(const char *_name) {
        String &_arg = getVar(_name);
        return _arg.c_str();
    }

    int getPort() const {
        return strtol(dataPort.c_str(), NULL, 10);
    }
};

WiFiClient client;
HTTPClient http;
ESP8266WebServer server(80);

String getResponseFromServer() {
    if (client.available()) {
        Serial.println("Available");
        String line = client.readStringUntil('\n');
        Serial.println(line);
    } else {
        Serial.println("Not available");
    }
}

void sendDataToServer(const String& data) {
    client.println("POST /posts HTTP/1.1");
    client.println("Host: jsonplaceholder.typicode.com");
    client.println("Cache-Control: no-cache");
    client.println("Content-Type: application/x-www-form-urlencoded");
    client.print("Content-Length: ");
    client.println(data.length());
    client.println();
    client.println(data);
}

bool tryToConnectWiFi(const Credentials &cred) {
    Serial.println("Trying to connect WiFi");
    if (WiFi.status() == WL_CONNECTED) {
        Serial.println("Already connected");
        return true;
    }
    delay(WIFI_CONN_DELAY);
    Serial.println(cred.WiFiSSID);
    Serial.println(cred.WiFiPASS);
    WiFi.begin(cred.WiFiSSID.c_str(), cred.WiFiPASS.c_str());
    if (WiFi.waitForConnectResult() == WL_CONNECTED) {
        Serial.println("Success!");
        return true;
    }
    Serial.println("Unable to connect!");
    return false;
}

bool tryToConnectServer(WiFiClient& cl, const Credentials &cred) {
    delay(CLIENT_CONN_DELAY);
    cl.connect(cred.dataURL.c_str(), cred.getPort());
    return cl.connected() == true; // return true if client is not connected and there is still unread data
}

void handleBad() {
    server.send(200, "text/html", "<h1>Not well</h1>");
}

const char *getAP_PASS() {
    return AP_PASS;
}

const char *getAP_SSID() {
    return AP_SSID;
}

void handleRegistration() {
    File file = SPIFFS.open("/index.html", "r");
    server.streamFile(file, "text/html");
    file.close();
}

void saveCredentials(const char *filename, Credentials &cred) {
    File credFile = SPIFFS.open(filename, "w");
    StaticJsonBuffer<JSON_STACK_SIZE> jsonBuffer;
    JsonObject &root = jsonBuffer.createObject();
    for (int i = 0; i < CNT_OF_ARGS; ++i) {
        root[credArgs[i]] = cred.getArg(credArgs[i]);
    }
    root.printTo(Serial); 
    if (root.printTo(credFile) == 0) {
        Serial.println("Unable write to file!"); 
    } else {
        Serial.println("Saved!");
    }
    credFile.close();
}

void loadCredentials(const char *filename, Credentials &cred) {
    File credFile = SPIFFS.open(filename, "r");
    StaticJsonBuffer<JSON_STACK_SIZE> jsonBuffer;
    JsonObject &root = jsonBuffer.parseObject(credFile);
    if (!root.success()) {
        Serial.println("Can't read credentials!");
        credFile.close();
        return;
    }
    Serial.println("Successfully read!");
    for (int i = 0; i < CNT_OF_ARGS; ++i) {
        String &var = cred.getVar(credArgs[i]);
        var = root[credArgs[i]].as<String>();
    }
    Serial.println(cred.WiFiSSID);
    Serial.println(cred.WiFiPASS);
    credFile.close();
}

Credentials mainCred;

bool credentialsAreValid() {
    for (int i = 0; i < CNT_OF_ARGS; ++i) {
        if (!server.hasArg(credArgs[i]) || server.arg(credArgs[i]) == NULL) {
            return false;
        }
    }
    Credentials cred;
    cred.WiFiSSID = server.arg("wifi_ssid");
    cred.WiFiPASS = server.arg("wifi_pass");
    if (!tryToConnectWiFi(cred)) {
        return false;
    }
    cred.dataURL = server.arg("data_url");
    cred.dataPort = server.arg("data_port");
    Serial.println(cred.dataURL);
    Serial.println(cred.dataPort);
    if (!tryToConnectServer(client, cred)) {
        return false;
    }

    // Check response from server by sending token and receiving response, TODO
    
    cred.token = server.arg("token");
    sendDataToServer(cred.token);
    mainCred = cred;
    return true;
}

void handleLogin() {
    Serial.println("HANDLE LOGIN!");
    if (!credentialsAreValid()) {
        // Give info about error, TODO
        Serial.println("Credentials not valid!");
        return;
    }
    saveCredentials("/credentials.txt", mainCred);
    server.send(200, "text/html", "<h1>Successfully!</h1>");
   // WiFi.softAPdisconnect(); // reset access point of ESP
}

void setup() {
    Serial.begin(115200);
    SPIFFS.begin();
 
    WiFi.mode(WIFI_AP_STA);
    WiFi.disconnect(true);
    delay(100);

//    WiFi.begin("", "");
//    if (WiFi.waitForConnectResult() == WL_CONNECTED) {
//        Serial.println("Success!");
//    } else {
//        Serial.println("Unable to connect!");
//    }
//    
//    delay(100);
//    client.connect("207.154.227.181", 1000);
//    
//    if (!client.connected()) {
//        Serial.println("bad");
//    }
//    delay(100);
//    sendDataToServer("abcde");
//    int ans = client.read();
//    Serial.println(ans);
  
    if (!SPIFFS.exists("/credentials.txt")) {
        Serial.println("Credentials file not found. Starting registration...");
        delay(100);
        Serial.println("Starting access point...");

        WiFi.softAP(getAP_SSID(), getAP_PASS());
        delay(500);
        IPAddress myIP = WiFi.softAPIP();
        
        Serial.print("Access point IP: ");
        Serial.println(myIP);
        
        server.on("/", handleRegistration);
        server.on("/login", handleLogin);
        server.on("/bad", handleBad);
        server.begin();
        
        Serial.println("HTTP server started...");
        
        // Need to set up, TODO
        return;
    } else {
        Serial.println("Credentials exists!");
        loadCredentials("/credentials.txt", mainCred);
    }
}

void loop() {
    if (mainCred.WiFiPASS != "") {
        tryToConnectWiFi(mainCred);
        if (!client.connected()) {
            Serial.println("Disconnected from server!");
            if (tryToConnectServer(client, mainCred)) {
                Serial.println("Connected to server...");
            }
        } else {
            Serial.println(client.availableForWrite());
            client.write("Hello, world!");
        }
    } else {
        server.handleClient();
    }
    /*
    server.handleClient();
    delay(SENDING_FRAME);
    tryToConnectWiFi();
    if (!client.connected()) {
        Serial.println("Disconnected from server!");
        if (tryToConnectServer(client)) {
            Serial.println("Connected to server...");
        }
        // need to pass credentials TODO
    } else {
        Serial.println(client.availableForWrite());
        client.flush();
        client.write("Hello, world!");
    }
    */
}