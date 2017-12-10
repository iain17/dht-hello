// example.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "libdn.h"
#include <iostream>

void DN_LogCB(const char* message)
{
	printf("[NP] %s", message);
}

int main()
{
	printf("DN_Init()\n");
	DN_Init(DN_LogCB);
	bool status = false;
	while (!status) {
		status = DN_Connect("10.1.1.34", 3036);
	}
	
	DNSessionInfo info;
	info.name = "Big tests";
	info.type = 1337;
	info.port = 8080;
	info.address = 123;
	//info->details.insert("cool", "certainly");
	auto request = DN_UpsertSession(&info);
	request->SetCallback([](DNAsync<DNUpsertSessionResult>* async) {
		DNUpsertSessionResult* answer = async->GetResult();
		printf("session id: %d\n", answer->sessionId);
	}, NULL);

	while (true) {
		DN_RunFrame();
		auto request = DN_GetSessionIds(1337, "", "");
		if (request->Wait(7500) != nullptr) {
			std::vector<DNSID>* answer = request->GetResult();
			printf("Received %d session ids\n", answer->size());
			for (DNSID sessionId : *answer) {
				printf("Session: %d\n", sessionId);
			}
			printf("-----\n", answer->size());
		}
		printf("Frame\n");
		Sleep(100);
	}
}