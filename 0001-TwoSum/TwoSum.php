<?php
/**
 * author: yanyandenuonuo
 * create: 2018.05.13
 */

$array = [1, 2, 3, 4, 5, 6, 7, 8, 9];


/**
 * @param $nums
 * @param $target
 * @return array
 */
function twoSum($nums, $target) {
    for($index = 0; $index < count($nums); $index += 1) {
        $key = array_search($nums, $target - $nums[$index]);
        if($key !== false && $key != $index) {
            return [$index, $key];
        }
    }
}


var_dump(twoSum($array, 9));
