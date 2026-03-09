#!/usr/bin/env python3
import sys
import subprocess
from pathlib import Path
import datetime

# 统一基准测试脚本：
# - 调用 tests/gen_testdata.py 生成测试数据
# - 运行 Huffman / Arithmetic / Range / RLE 的基准测试
# - 将各自输出写入 reports/ 目录下的文本文件

ROOT = Path(__file__).resolve().parent.parent
TESTS_DIR = ROOT / "tests"
TEST_DATA_DIR = TESTS_DIR / "data"
REPORTS_DIR = ROOT / "reports"

PY = sys.executable or "python3"


def run_capture(cmd, cwd: Path, report_path: Path, title: str):
    print(f"[run_all_bench] running {title}: {' '.join(map(str, cmd))} (cwd={cwd})")
    proc = subprocess.run(cmd, cwd=cwd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
    REPORTS_DIR.mkdir(parents=True, exist_ok=True)
    header = f"# {title}\n# command: {' '.join(map(str, cmd))}\n# cwd: {cwd}\n# returncode: {proc.returncode}\n\n"
    report_path.write_text(header + proc.stdout, encoding="utf-8", errors="ignore")
    if proc.returncode != 0:
        print(f"[run_all_bench] WARNING: {title} exited with code {proc.returncode}, see {report_path}")
    else:
        print(f"[run_all_bench] {title} OK, report -> {report_path}")


def ensure_testdata():
    print("[run_all_bench] generating test data...")
    subprocess.check_call([PY, "tests/gen_testdata.py"], cwd=ROOT)
    target = TEST_DATA_DIR / "random_10MiB.bin"
    if not target.is_file():
        raise SystemExit(f"test data not found: {target}")
    return target


def main():
    REPORTS_DIR.mkdir(parents=True, exist_ok=True)

    input_file = ensure_testdata()
    ts = datetime.datetime.now().strftime("%Y%m%d-%H%M%S")

    # Huffman 跨语言 benchmark
    huffman_bench = ROOT / "huffman" / "benchmark" / "bench.py"
    if huffman_bench.is_file():
        run_capture(
            [PY, "bench.py", str(input_file)],
            cwd=huffman_bench.parent,
            report_path=REPORTS_DIR / f"huffman_report_{ts}.txt",
            title="Huffman benchmark",
        )

    # Arithmetic 跨语言 benchmark
    arithmetic_bench = ROOT / "arithmetic" / "benchmark" / "bench.py"
    if arithmetic_bench.is_file():
        run_capture(
            [PY, "bench.py", str(input_file)],
            cwd=arithmetic_bench.parent,
            report_path=REPORTS_DIR / f"arithmetic_report_{ts}.txt",
            title="Arithmetic benchmark",
        )

    # Range coder 跨语言 benchmark
    range_bench = ROOT / "range" / "benchmark" / "bench.py"
    if range_bench.is_file():
        run_capture(
            [PY, "bench.py", str(input_file)],
            cwd=range_bench.parent,
            report_path=REPORTS_DIR / f"range_report_{ts}.txt",
            title="Range coder benchmark",
        )

    # RLE 跨语言 benchmark
    rle_bench = ROOT / "rle" / "benchmark" / "bench.py"
    if rle_bench.is_file():
        run_capture(
            [PY, "bench.py", str(input_file)],
            cwd=rle_bench.parent,
            report_path=REPORTS_DIR / f"rle_report_{ts}.txt",
            title="RLE benchmark",
        )

    print("[run_all_bench] all done. Reports are in:")
    for p in sorted(REPORTS_DIR.glob("*_report_*.txt")):
        print("  ", p)


if __name__ == "__main__":
    main()
