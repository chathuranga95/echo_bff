import ballerina/http;
import ballerina/log;
import ballerina/os;

final string serviceUrl = os:getEnv("SERVICE_URL");
final string tokenUrl = os:getEnv("TOKEN_URL");
final string clientId = os:getEnv("CLIENT_ID");
final string clientSecret = os:getEnv("CLIENT_SECRET");

http:Client helloClient = check new (serviceUrl,
    auth = {
        tokenUrl: tokenUrl,
        clientId: clientId,
        clientSecret: clientSecret
    }
);

service / on new http:Listener(9090) {
    resource function get greeting() returns json|error? {
        json resp = check helloClient->get("/");
        log:printInfo("Response: " + resp.toJsonString());
        return resp;
    }

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
