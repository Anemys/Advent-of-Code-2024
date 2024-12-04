#!/usr/bin/env python3

from os import PathLike

type Report = list[int]


def parse_input(filepath: str | PathLike[str]) -> list[Report]:
    reports = []

    with open(filepath, "r") as f:
        for line in f:
            report = list(map(int, line.split()))
            reports.append(report)

    return reports


def is_monotonic(report: Report) -> bool:
    return report == sorted(report) or report == sorted(report, reverse=True)


def has_gradual_evolution(report: Report) -> bool:
    for i, elm in enumerate(report):
        if i == 0:
            continue

        level_diff = elm - report[i - 1]
        if abs(level_diff) < 1 or abs(level_diff) > 3:
            return False

    return True


def is_safe_report(report: Report) -> bool:
    return is_monotonic(report) and has_gradual_evolution(report)


def count_safe_reports(reports: list[Report], tolerant: bool = False) -> int:
    safe_reports = 0
    for report in reports:
        if is_safe_report(report):
            safe_reports += 1
            continue

        if not tolerant:
            continue

        for i, _ in enumerate(report):
            cutted_report = report.copy()
            cutted_report.pop(i)

            if is_safe_report(cutted_report):
                safe_reports += 1
                break

    return safe_reports


def main():
    reports = parse_input("input.txt")

    safe_reports = count_safe_reports(reports)

    print(f"Safe reports: {safe_reports}")

    tolerated_reports = count_safe_reports(reports, True)

    print(f"Tolerated reports: {tolerated_reports}")


if __name__ == "__main__":
    main()
