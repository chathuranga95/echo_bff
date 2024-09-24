import ballerina/http;
import ballerina/log;
import ballerina/os;

final string serviceUrl = os:getEnv("SERVICE_URL");
final string tokenUrl = os:getEnv("TOKEN_URL");
final string clientId = os:getEnv("CLIENT_ID");
final string clientSecret = os:getEnv("CLIENT_SECRET");

service / on new http:Listener(9090) {
    resource function get diagnostic() returns json {
        json diagnostic = {
            "serviceUrl": serviceUrl,
            "tokenUrl": tokenUrl,
            "clientId": clientId,
            "clientSecret": clientSecret,
            "diagnosticVersion": "v1.0"
        };
        log:printInfo("Details: " + diagnostic.toJsonString());
        return diagnostic;
    }
}
