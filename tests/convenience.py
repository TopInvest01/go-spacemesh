import argparse
from datetime import datetime
import time

TIMESTAMP_FMT = "%Y-%m-%dT%H:%M:%S.%fZ"

PRINT_SEP = "===================================================================="


def print_nl_end(*args):
    print(*args, "\n")


def print_nl_start(*args):
    print("\n", *args)


def convert_ts_to_datetime(ts):
    return datetime.strptime(ts, TIMESTAMP_FMT)


def sleep_print_backwards(tts):
    print(f"\nsleeping for {tts} seconds\n")
    while tts != 0:
        tts -= 1
        print(f" {tts} seconds left     ", end="\r")
        time.sleep(1)


def str2bool(v):
    if isinstance(v, bool):
        return v
    if v.lower() in ('yes', 'true', 't', 'y', '1'):
        return True
    elif v.lower() in ('no', 'false', 'f', 'n', '0'):
        return False
    else:
        raise argparse.ArgumentTypeError('Boolean value expected.')