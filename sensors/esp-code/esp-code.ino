#include <NTPClient.h>
#include <assert.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WebServer.h>
#include <FS.h>
#include <time.h>
#include <ArduinoJson.h>

const char *credArgs[] = {"wifi_ssid", "wifi_pass", "data_url", "data_port", "user_token", "neighbor_token"};
const int CNT_OF_ARGS = 6;

const char *AP_SSID = "ESP";
const char *AP_PASS = "3.141592";

const int WIFI_CONN_DELAY = 10000;
const int SENDING_FRAME = 1000;
const int CLIENT_CONN_DELAY = 1000;
const int JSON_STACK_SIZE = 8192;

struct Credentials {
    String WiFiSSID, WiFiPASS, dataURL, dataPort, userToken, neighborToken;
    Credentials(){};
    Credentials(const String &_SSID, const String &_PASS, const String &_URL, const String &_port, const String &_userToken, const String &_neighborToken): 
        WiFiSSID(_SSID), 
        WiFiPASS(_PASS), 
        dataURL(_URL),
        dataPort(_port),
        userToken(_userToken),
        neighborToken(_neighborToken) {
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
                return userToken;
            case 5:
                return neighborToken;
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

HTTPClient http;
ESP8266WebServer server(80);

String getDataFromServer() {
    return http.getString();
}

int sendDataToServer(const String& data) {
    http.addHeader("Content-Type", "text/plain");
    return http.POST(data.c_str()); // response code
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

bool tryToConnectServer(const Credentials &cred) {
    delay(CLIENT_CONN_DELAY);
    return http.begin(cred.dataURL);
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
    if (!tryToConnectServer(cred)) {
        return false;
    }

    // Check response from server by sending token and receiving response, TODO
    
    cred.userToken = server.arg("user_token");
    cred.neighborToken = server.arg("neighbor_token");
    Serial.println(cred.userToken);
    Serial.println(cred.neighborToken);

    sendDataToServer(cred.userToken + ":" + cred.neighborToken);
    if (getDataFromServer() != "OK") {
        Serial.println("Bad data");
    }
    mainCred = cred;
    Serial.println("Data is valid!");
    return true;
}

void handleLogin() {
    if (!credentialsAreValid()) {
        // Give info about error, TODO
        Serial.println("Credentials not valid!");
        return;
    }
    saveCredentials("/credentials.txt", mainCred);
    server.send(200, "text/html", "<h1>Successfully!</h1>");
   // WiFi.softAPdisconnect(); // reset access point of ESP
}

unsigned int getTimestamp() {
    time_t result = time(nullptr);
    Serial.println(asctime(localtime(&result)));
    return (unsigned int)result;
}

void setup() {
    Serial.begin(115200);
    SPIFFS.begin();
    Serial.println(getTimestamp());
 
    WiFi.mode(WIFI_AP_STA);
    WiFi.disconnect(true);
    delay(100);
  
    if (!SPIFFS.exists("/credentials.txt")) {
        Serial.println("Credentials file not found. Starting registration...");
        delay(100);
        Serial.println("Starting access point...");

        WiFi.softAP(AP_SSID, AP_PASS);
        delay(500);
        IPAddress myIP = WiFi.softAPIP();
        
        Serial.print("Access point IP: ");
        Serial.println(myIP);
        
        server.on("/", handleRegistration);
        server.on("/login", handleLogin);
        server.begin();
        
        Serial.println("HTTP server started...");
        
        // Need to set up, TODO
    } else {
        Serial.println("Credentials exists!");
        loadCredentials("/credentials.txt", mainCred);
    }
}

int readVoltage() {
    return 100;
}

void loop() {
    if (mainCred.WiFiPASS != "") {
        tryToConnectWiFi(mainCred);
        sendDataToServer(String(readVoltage()) + ":" + String(getTimestamp()));
        Serial.println(getDataFromServer());
    } else {
        server.handleClient();
    }
}
