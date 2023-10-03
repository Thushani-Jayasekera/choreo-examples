import ballerina/http;

listener http:Listener probeEP = new (9091);

service /probes on probeEP {
    resource function get healthz() returns boolean {
        return true;
    }
    resource function get health() returns boolean {
        return true;
    }
}