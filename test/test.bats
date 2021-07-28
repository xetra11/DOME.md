#!/usr/bin/env bats

@test "should print 'done'" {
  result="$(../do test.md)"
  [[ "$result" == *"Parsing README.md"* ]]
}