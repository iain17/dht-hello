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
	while (true)
	{
		Sleep(100);
		DN_RunFrame();

		DNHealthResult* result = DN_Health();
		if (result->ready) {
			printf("DN is ready\n");
		} else {
			printf("DN is NOT ready:");
			printf(result->message.c_str());
			printf("\n");
		}

		printf("Frame\n");
	}
}