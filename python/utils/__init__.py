"""Utility functions for Advent of Code solutions."""

from pathlib import Path
from typing import List


def read_input(filename: str) -> str:
    """Read the entire file and return as a string."""
    return Path(filename).read_text().strip()


def read_lines(filename: str) -> List[str]:
    """Read the file and return as a list of lines."""
    return Path(filename).read_text().strip().split('\n')


def read_ints(filename: str) -> List[int]:
    """Read the file and return as a list of integers (one per line)."""
    return [int(line) for line in read_lines(filename)]


def read_grid(filename: str) -> List[List[str]]:
    """Read the file and return as a 2D grid of characters."""
    return [list(line) for line in read_lines(filename)]
