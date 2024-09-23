// From the Go documentation, it's recommended to include stdlib.h if we need
// to use C.free.
#include <stdlib.h>
#include <stdbool.h>

long create_regex_rule(const char* json_config);

long create_scanner(long rule_list, const char** error, bool should_keywords_match_event_paths);
void delete_scanner(long scanner_id);

// event is a non-null terminated TODO
const char* scan(long scanner_id, const void* event, long event_size, long *retsize, long *retcap, const char** error);

void free_vec(const char* string, long len, long cap);
void free_string(const char* string);

void free_any_rule(long rule_ptr);

long create_rule_list();
void append_rule_to_list(long rule_ptr, long list_ptr);
void free_rule_list(long list_ptr);
