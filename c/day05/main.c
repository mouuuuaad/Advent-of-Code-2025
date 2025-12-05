#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../utils/input.h"

typedef struct {
    long start;
    long end;
} Range;

int compare_ranges(const void* a, const void* b) {
    Range* r1 = (Range*)a;
    Range* r2 = (Range*)b;
    if (r1->start < r2->start) return -1;
    if (r1->start > r2->start) return 1;
    return 0;
}

int main(int argc, char* argv[]) {
    const char* filename = "input.txt";
    if (argc > 1) {
        filename = argv[1];
    }

    char* content = read_input(filename);
    if (!content) {
        fprintf(stderr, "Failed to read %s\n", filename);
        return 1;
    }

    Range ranges[1000];
    int range_count = 0;
    
    char* line = strtok(content, "\n");

    while (line != NULL) {
        if (strchr(line, '-')) {
            long start, end;
            if (sscanf(line, "%ld-%ld", &start, &end) == 2) {
                ranges[range_count].start = start;
                ranges[range_count].end = end;
                range_count++;
            }
        }
        line = strtok(NULL, "\n");
    }

    if (range_count == 0) {
        printf("0\n");
        free(content);
        return 0;
    }

    qsort(ranges, range_count, sizeof(Range), compare_ranges);

    long total_fresh = 0;
    long current_start = ranges[0].start;
    long current_end = ranges[0].end;

    for (int i = 1; i < range_count; i++) {
        if (ranges[i].start <= current_end + 1) {
             if (ranges[i].start <= current_end) {
                 if (ranges[i].end > current_end) {
                     current_end = ranges[i].end;
                 }
             } else {
                 total_fresh += (current_end - current_start + 1);
                 current_start = ranges[i].start;
                 current_end = ranges[i].end;
             }
        } else {
             total_fresh += (current_end - current_start + 1);
             current_start = ranges[i].start;
             current_end = ranges[i].end;
        }
    }
    total_fresh += (current_end - current_start + 1);

    printf("%ld\n", total_fresh);
    free(content);
    return 0;
}
