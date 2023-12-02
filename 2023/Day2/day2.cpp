#include <fstream>
#include <iostream>
#include <sstream>
#include <tuple>
#include <vector>

struct cubePair {
  int number;
  std::string color;
};

const int red = 12;
const int green = 13;
const int blue = 14;

std::tuple<int, int> solve(std::string &gameInput, int &gameNumber) {
  if (gameInput.empty()) {
    return std::make_tuple(0, 0);
  }

  size_t colonPos = gameInput.find(":");
  gameInput = gameInput.substr(colonPos + 1);

  std::stringstream ss(gameInput);
  std::vector<std::string> splitParts;

  std::string part;
  std::vector<cubePair> pairs;
  int minRed = 0, minGreen = 0, minBlue = 0;
  while (std::getline(ss, part, ';')) {
    std::stringstream ss(part);

    while (!ss.eof()) {
      cubePair pair;

      ss >> pair.number;

      ss >> std::ws;
      while (ss.peek() == ',') {
        ss.ignore();
        ss >> std::ws;
      }
      ss >> pair.color;

      pairs.push_back(pair);
    }

    for (const auto &pair : pairs) {
      if (pair.color[0] == 'r') {
        if (pair.number > minRed) {
          minRed = pair.number;
        }
      } else if (pair.color[0] == 'g') {
        if (pair.number > minGreen) {
          minGreen = pair.number;
        }
      } else if (pair.color[0] == 'b') {
        if (pair.number > minBlue) {
          minBlue = pair.number;
        }
      } else {
        throw std::runtime_error("Invalid color");
      }
    }
  }

  int redTotal = 0, greenTotal = 0, blueTotal = 0;
  for (const auto &pair : pairs) {
    if (pair.color[0] == 'r') {
      redTotal += pair.number;
    } else if (pair.color[0] == 'g') {
      greenTotal += pair.number;
    } else if (pair.color[0] == 'b') {
      blueTotal += pair.number;
    } else {
      throw std::runtime_error("Invalid color");
    }
  }

  int power = minRed * minGreen * minBlue;

  if (redTotal > red || greenTotal > green || blueTotal > blue) {
    return std::make_tuple(0, power);
  }

  return std::make_tuple(gameNumber, power);
}

int main() {
  std::ifstream file("input.txt");
  int count = 0;
  int score = 0;
  int score2 = 0;
  while (file.good()) {
    count++;
    std::string line;
    std::getline(file, line);

    auto [part1, part2] = solve(line, count);

    score += part1;
    score2 += part2;
  }
  std::cout << "Score: " << score << std::endl;
  std::cout << "Score2: " << score2 << std::endl;
  file.close();
}
