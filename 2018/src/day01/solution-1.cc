#include <glog/logging.h>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include "absl/strings/numbers.h"

int main(int argc, char* argv[]) {
  google::InitGoogleLogging(argv[0]);

  std::string filename("2018/day01/input.txt");
  std::ifstream fs(filename);
  if (!fs.is_open()) {
    LOG(ERROR) << "Could not read input file: " << filename;
    return -1;
  }

  int total = 0;
  std::string line;
  while (std::getline(fs, line)) {
    // LOG(INFO) << "line: " << line;

    char sign = line[0];
    // LOG(INFO) << "sign: " << sign;

    int val = 0;
    bool ok = absl::SimpleAtoi(line.substr(1, std::string::npos), &val);
    if (!ok) {
      LOG(ERROR) << "Error converting to integer";
      return -1;
    }
    // LOG(INFO) << "val: " << val;

    if (sign == '-') {
      total -= val;
    }
    if (sign == '+') {
      total += val;
    }
  }

  LOG(INFO) << "total: " << total;

  return 0;
}
