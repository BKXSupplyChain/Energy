
#include <assert.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WebServer.h>
#include <FS.h>
#include <ArduinoJson.h>

const char *credArgs[] = {"wifi_ssid", "wifi_pass", "data_url", "sensor_id"};
const int CNT_OF_ARGS = 4;

const char *AP_SSID = "ESP";
const char *AP_PASS = "3.141592";

const int WIFI_CONN_DELAY = 10000;
const int SENDING_FRAME = 1000;
const int CLIENT_CONN_DELAY = 1000;
const int JSON_STACK_SIZE = 300;

struct Credentials {
    String WiFiSSID, WiFiPASS, dataURL, sensorID; // dataURL goes with port
    Credentials(){};
    Credentials(const String &_SSID, const String &_PASS, const String &_URL, const String &_sensorID): 
        WiFiSSID(_SSID), 
        WiFiPASS(_PASS), 
        dataURL(_URL),
        sensorID(_sensorID) {
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
                return sensorID;
        }
    }

    const char *getArg(const char *_name) {
        String &_arg = getVar(_name);
        return _arg.c_str();
    }
};

HTTPClient http;
ESP8266WebServer server(80);

String getDataFromServer() {
    return http.getString(); // return only first response, TO FIX
}

int sendDataToServer(const String& payload) {
    return http.sendRequest("POST", payload);
}

bool tryToConnectWiFi(const Credentials &cred) {
    Serial.println("Trying to connect WiFi");
    if (WiFi.status() == WL_CONNECTED) {
        Serial.println("Already connected");
        return true;
    }
    delay(WIFI_CONN_DELAY);
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
    Serial.println("Open index.html");
    File file = SPIFFS.open("/index.html", "r");
    if (!file) {
        Serial.println("Can't open!");
    }
    server.streamFile(file, "text/html");
    file.close();
}

void saveCredentials(const char *filename, Credentials &cred) {
    Serial.println("SPIFFS Open");
    File credFile = SPIFFS.open(filename, "w");
    Serial.println("Credentials open!");
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
    Serial.println(cred.dataURL);
    Serial.println("File closing");
    credFile.close();
    Serial.println("File closed");
}

Credentials mainCred;

bool credentialsAreValid() {
    Serial.println("Check credentials!");
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
    Serial.println(cred.dataURL);
    if (!tryToConnectServer(cred)) {
        return false;
    }

    cred.sensorID = server.arg("sensor_id");
    Serial.println(cred.sensorID);
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
    saveCredentials("/credentials.json", mainCred);
    Serial.println("Credentials saved!");
    server.send(200, "text/html", "<h1>Successfully!</h1>");
   // WiFi.softAPdisconnect(); // reset access point of ESP
}

void setup() {
    Serial.begin(115200);
    SPIFFS.begin();
    
    WiFi.mode(WIFI_AP_STA);
    WiFi.disconnect(true);
    delay(100);
  
    if (!SPIFFS.exists("/credentials.json")) {
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
        loadCredentials("/credentials.json", mainCred);
        Serial.println("Credentials loaded");
        Serial.println("Connecting to server...");
        Serial.println(tryToConnectServer(mainCred));
    }
}

int readVoltage() {
    return analogRead(A0);
}


int total = 0; // total consumption, TODO: overflow
int prevVoltage = 0;

String getData() {
    StaticJsonBuffer<JSON_STACK_SIZE> jsonBuffer;
    JsonObject &root = jsonBuffer.createObject();
    int currVoltage = readVoltage();
    total += (currVoltage + prevVoltage) / 2;
    root["total"] = total;
    root["sensorID"] = mainCred.sensorID;
    String res;
    root.printTo(res);
    Serial.println(res);
    return res;
}

void loop() {
    delay(SENDING_FRAME);    
    Serial.println("LOOP!");
    if (mainCred.WiFiPASS != "") {
        tryToConnectWiFi(mainCred);
        sendDataToServer(getData());
    } else {
        server.handleClient();
    }
}
