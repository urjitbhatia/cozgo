#include <coz.h>

void cozBegin(char* name) {
    COZ_BEGIN(name);
}

void cozEnd(char* name) {
    COZ_END(name);
}

void cozProgressNamed(char* name) {
    COZ_PROGRESS_NAMED(name);
}

void cozProgress() {
    COZ_PROGRESS;
}
