#! /bin/bash

git clone $(curl -s https://api.github.com/repos/01-edu/the-final-cl-test/git/refs/heads/main | grep '"sha"' | head -n 1 | awk '{print $2}' | tail -c 41) https://github.com/tensorflow/tensorflow.git