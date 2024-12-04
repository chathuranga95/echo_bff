import ballerina/http;
import ballerina/log;
import ballerina/os;

final string serviceUrl = os:getEnv("SERVICE_URL");
final string tokenUrl = os:getEnv("TOKEN_URL");
final string clientId = os:getEnv("CLIENT_ID");
final string clientSecret = os:getEnv("CLIENT_SECRET");
final string serviceUrl2 = os:getEnv("SERVICE_URL_2");
final string tokenUrl2 = os:getEnv("TOKEN_URL_2");
final string clientId2 = os:getEnv("CLIENT_ID_2");
final string clientSecret2 = os:getEnv("CLIENT_SECRET_2");

http:Client helloClient = check new (serviceUrl,
    auth = {
        tokenUrl: tokenUrl == "" ? "https://sts.preview-dv.choreo.dev/oauth2/token" : tokenUrl,
        clientId: clientId,
        clientSecret: clientSecret
    }
);

service / on new http:Listener(9090) {
    resource function get greeting(string subpath = "") returns json|error? {
        json resp = check helloClient->get("/" + subpath);
        log:printInfo("Response: " + resp.toJsonString());
        return resp;
    }

    resource function get diagnostic() returns json {
        json diagnostic = {
            "serviceUrl": serviceUrl,
            "tokenUrl": tokenUrl,
            "clientId": clientId,
            "clientSecret": clientSecret,
            "serviceUrl2": serviceUrl2,
            "tokenUrl2": tokenUrl2,
            "clientId2": clientId2,
            "clientSecret2": clientSecret2,
            "diagnosticVersion": "v1.0"
        };
        log:printInfo("Details: " + diagnostic.toJsonString());
        return diagnostic;
    }
}
