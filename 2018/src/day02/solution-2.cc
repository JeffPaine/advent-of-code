#include <glog/logging.h>
#include <fstream>
#include <iostream>
#include <queue>
#include <sstream>
#include <string>
#include "absl/strings/str_cat.h"
#include "absl/strings/str_join.h"

int main(int argc, char* argv[]) {
  google::InitGoogleLogging(argv[0]);

  std::string filename("2018/day02/input.txt");
  std::ifstream fs(filename);
  if (!fs.is_open()) {
    LOG(ERROR) << "Could not read input file: " << filename;
    return -1;
  }

  std::vector<std::string> entries;

  std::string line;
  while (std::getline(fs, line)) {
    entries.emplace_back(line);
  }

  while (!entries.empty()) {
    std::string last = entries.back();
    entries.pop_back();

    for (auto const& entry : entries) {
      if (entry.size() != last.size()) {
        continue;
      }

      int non_matching = 0;
      for (size_t i = 0; i < entry.size(); i++) {
        if (entry[i] != last[i]) {
          non_matching += 1;
        }
      }

      if (non_matching == 1) {
        LOG(INFO) << "last:  " << last;
        LOG(INFO) << "entry: " << entry;
        std::vector<std::string> matching_chars;
        for (size_t i = 0; i < entry.size(); i++) {
          if (entry[i] == last[i]) {
            matching_chars.emplace_back(std::string(1, entry[i]));
          }
        }
        LOG(INFO) << "matching_chars: " << absl::StrJoin(matching_chars, "");
      }
    }
  }

  return 0;
}
