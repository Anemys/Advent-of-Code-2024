#!/usr/bin/env python3

from os import PathLike


def parse_input(filepath: str | PathLike[str]) -> tuple[list[int], list[int]]:
    left_list, right_list = [], []

    with open(filepath, "r") as f:
        for line in f:
            ids = line.split()

            left_list.append(int(ids[0]))
            right_list.append(int(ids[1]))

    return left_list, right_list


def compute_total_distance(left_list: list[int], right_list: list[int]) -> int:
    total_distance = 0
    for left_elm, right_elm in zip(sorted(left_list), sorted(right_list)):
        total_distance += abs(left_elm - right_elm)

    return total_distance


def compute_similarity_score(left_list: list[int], right_list: list[int]) -> int:
    similarity_score = 0
    for id in left_list:
        similarity_score += id * right_list.count(id)

    return similarity_score


def main():
    left_list, right_list = parse_input("input.txt")

    total_distance = compute_total_distance(left_list, right_list)

    print(f"Total distance: {total_distance}")

    similarity_score = compute_similarity_score(left_list, right_list)

    print(f"Similarity score: {similarity_score}")


if __name__ == "__main__":
    main()
