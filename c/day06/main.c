#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define MAX_LINES 10
#define MAX_LINE_LEN 20000

char lines[MAX_LINES][MAX_LINE_LEN];
int num_lines = 0;
size_t max_len = 0;

void read_input(void) {
    FILE *fp = fopen("input.txt", "r");
    if (!fp) {
        perror("Failed to open input.txt");
        exit(1);
    }

    while (fgets(lines[num_lines], MAX_LINE_LEN, fp) && num_lines < MAX_LINES) {
        size_t len = strlen(lines[num_lines]);
        if (len > 0 && lines[num_lines][len-1] == '\n') {
            lines[num_lines][len-1] = '\0';
        }
        num_lines++;
    }
    fclose(fp);

    for (int i = 0; i < num_lines; i++) {
        size_t len = strlen(lines[i]);
        if (len > max_len) max_len = len;
    }
    for (int i = 0; i < num_lines; i++) {
        size_t len = strlen(lines[i]);
        while (len < max_len) lines[i][len++] = ' ';
        lines[i][len] = '\0';
    }
}

long long solve_part1(void) {
    char *op_line = lines[num_lines - 1];
    long long grand_total = 0;
    int col = 0;

    while (col < (int)max_len) {
        int all_spaces = 1;
        for (int row = 0; row < num_lines; row++) {
            if (lines[row][col] != ' ') { all_spaces = 0; break; }
        }
        if (all_spaces) { col++; continue; }

        int start_col = col, end_col = col;
        while (end_col < (int)max_len) {
            int is_separator = 1;
            for (int row = 0; row < num_lines; row++) {
                if (lines[row][end_col] != ' ') { is_separator = 0; break; }
            }
            if (is_separator) break;
            end_col++;
        }

        char op = ' ';
        for (int c = start_col; c < end_col; c++) {
            if (op_line[c] == '+' || op_line[c] == '*') { op = op_line[c]; break; }
        }

        long long result = 0;
        int first = 1;
        for (int row = 0; row < num_lines - 1; row++) {
            char num_str[100];
            int idx = 0;
            for (int c = start_col; c < end_col; c++) {
                if (isdigit(lines[row][c])) num_str[idx++] = lines[row][c];
            }
            num_str[idx] = '\0';
            if (idx > 0) {
                long long num = atoll(num_str);
                if (first) { result = num; first = 0; }
                else { result = (op == '+') ? result + num : result * num; }
            }
        }
        if (!first) grand_total += result;
        col = end_col;
    }
    return grand_total;
}

// Part 2: Numbers read vertically (column by column, right to left)
long long solve_part2(void) {
    char *op_line = lines[num_lines - 1];
    int num_data_rows = num_lines - 1;
    long long grand_total = 0;
    int col = 0;

    while (col < (int)max_len) {
        int all_spaces = 1;
        for (int row = 0; row < num_lines; row++) {
            if (lines[row][col] != ' ') { all_spaces = 0; break; }
        }
        if (all_spaces) { col++; continue; }

            int start_col = col, end_col = col;
        while (end_col < (int)max_len) {
            int is_separator = 1;
            for (int row = 0; row < num_lines; row++) {
                if (lines[row][end_col] != ' ') { is_separator = 0; break; }
            }
            if (is_separator) break;
            end_col++;
        }

        char op = ' ';
        for (int c = start_col; c < end_col; c++) {
            if (op_line[c] == '+' || op_line[c] == '*') { op = op_line[c]; break; }
        }

        long long result = 0;
        int first = 1;
        for (int c = end_col - 1; c >= start_col; c--) {
            char num_str[100];
            int idx = 0;
            for (int row = 0; row < num_data_rows; row++) {
                if (isdigit(lines[row][c])) num_str[idx++] = lines[row][c];
            }
            num_str[idx] = '\0';
            if (idx > 0) {
                long long num = atoll(num_str);
                if (first) { result = num; first = 0; }
                else { result = (op == '+') ? result + num : result * num; }
            }
        }
        if (!first) grand_total += result;
        col = end_col;
    }
    return grand_total;
}

int main(void) {
    read_input();
    
    printf("Part 1: %lld\n", solve_part1());
    printf("Part 2: %lld\n", solve_part2());
    
    return 0;
}
