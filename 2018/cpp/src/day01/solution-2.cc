#include <glog/logging.h>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include "absl/container/flat_hash_map.h"
#include "absl/strings/numbers.h"

int main(int argc, char* argv[]) {
  google::InitGoogleLogging(argv[0]);

  int total = 0;
  // Map of total:count.
  absl::flat_hash_map<int, int> totals;

  while (true) {
    std::string filename("2018/day01/input.txt");
    std::ifstream fs(filename);
    if (!fs.is_open()) {
      LOG(ERROR) << "Could not read input file: " << filename;
      return -1;
    }

    std::string line;
    while (std::getline(fs, line)) {
      char sign = line[0];
      int val = 0;
      bool ok = absl::SimpleAtoi(line.substr(1, std::string::npos), &val);
      if (!ok) {
        LOG(ERROR) << "Error converting to integer";
        return -1;
      }

      if (sign == '-') {
        total -= val;
      } else if (sign == '+') {
        total += val;
      } else {
        LOG(ERROR) << "Unrecognized sign: " << sign;
        return -1;
      }

      auto pos = totals.find(total);
      if (!(pos == totals.end())) {
        LOG(INFO) << "Found repeated total: " << total;
        return 0;
      }

      totals.emplace(total, 1);
    }
  }
}
