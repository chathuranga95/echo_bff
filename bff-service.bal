import ballerina/http;
import ballerina/log;
import ballerina/os;

final string serviceUrl = os:getEnv("SERVICE_URL");

http:Client helloClient = check new (serviceUrl);

service / on new http:Listener(9090) {
    resource function get greeting(string subpath = "") returns json|error? {
        json resp = check helloClient->get("/" + subpath);
        log:printInfo("Response: " + resp.toJsonString());
        return resp;
    }

    resource function get diagnostic() returns json {
        json diagnostic = {
            "serviceUrl": serviceUrl,
            "diagnosticVersion": "v1.0"
        };
        log:printInfo("Details: " + diagnostic.toJsonString());
        return diagnostic;
    }
}
