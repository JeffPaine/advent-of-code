#include <glog/logging.h>
#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include "absl/container/flat_hash_map.h"

int total_twos = 0;
int total_threes = 0;

int main(int argc, char* argv[]) {
  google::InitGoogleLogging(argv[0]);

  std::string filename("2018/day02/input.txt");
  std::ifstream fs(filename);
  if (!fs.is_open()) {
    LOG(ERROR) << "Could not read input file: " << filename;
    return -1;
  }

  std::string line;
  while (std::getline(fs, line)) {
    absl::flat_hash_map<char, int> ccount;

    for (auto c : line) {
      if (ccount.find(c) == ccount.end()) {
        ccount[c] = 1;
      } else {
        ccount[c] += 1;
      }
    }

    bool found_twos = false;
    bool found_threes = false;
    for (auto const& [key, value] : ccount) {
      if (value == 2) {
        found_twos = true;
      }
      if (value == 3) {
        found_threes = true;
      }
    }

    if (found_twos) {
      total_twos += 1;
    }
    if (found_threes) {
      total_threes += 1;
    }
  }

  LOG(INFO) << "total_twos: " << total_twos;
  LOG(INFO) << "total_threes: " << total_threes;

  return 0;
}
