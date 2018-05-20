<?php
/**
 * author: yanyandenuonuo
 * create: 2018.05.13
 */

$array = [3, 3];


/**
 * @param $nums
 * @param $target
 * @return array
 */
function twoSum($nums, $target) {
    for($index = 0; $index < count($nums); $index += 1) {
        $key = array_search($target - $nums[$index], $nums);
        if($key !== false && $key != $index) {
            if($key > $index) {
                return [$index, $key];
            } else {
                return [$key, $index];
            }
        }
    }
}


var_dump(twoSum($array, 6));
