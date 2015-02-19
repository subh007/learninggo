#include <stdio.h>
#include "test.h"

extern void receiveC(char *msg);

char* myprint(char *msg) {
	receiveC(msg);
	return msg;
	}

void receiver_go(char *msg) {
	//receiveC(msg);
}
