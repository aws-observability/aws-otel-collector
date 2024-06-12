// From the Go documentation, it's recommended to include stdlib.h if we need
// to use C.free.
#include <stdlib.h>
#include <stdbool.h>

long create_scanner(const char* rules_as_json, const char** error, bool should_keywords_match_event_paths);
void delete_scanner(long scanner_id);

// event is a non-null terminated TODO
const char* scan(long scanner_id, const void* event, long event_size, long *retsize, long *retcap, const char** error);

void free_vec(const char* string, long len, long cap);
void free_string(const char* string);
